// Copyright 2019 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	jaegerpb "github.com/jaegertracing/jaeger/model"
	"go.opencensus.io/trace"

	sapmpb "github.com/signalfx/sapm-proto/gen"
)

// worker is not safe to be called from multiple goroutines. Each caller must use locks to avoid races
// and data corruption. In case a caller needs to export multiple requests at the same time, it should
// use one worker per request.
type worker struct {
	client             *http.Client
	accessToken        string
	endpoint           string
	gzipWriter         *gzip.Writer
	disableCompression bool
}

func newWorker(client *http.Client, endpoint string, accessToken string, disableCompression bool) *worker {
	w := &worker{
		client:             client,
		accessToken:        accessToken,
		endpoint:           endpoint,
		disableCompression: disableCompression,
		gzipWriter:         gzip.NewWriter(nil),
	}
	return w
}

func (w *worker) export(ctx context.Context, batches []*jaegerpb.Batch) *ErrSend {
	ctx, span := trace.StartSpan(ctx, "export")
	defer span.End()

	var spansCount int
	for _, batch := range batches {
		spansCount += len(batch.Spans)
	}

	span.AddAttributes(trace.Int64Attribute("spans", int64(spansCount)))
	span.AddAttributes(trace.Int64Attribute("batches", int64(len(batches))))
	if spansCount == 0 {
		return nil
	}

	sr, err := w.prepare(ctx, batches, spansCount)
	if err != nil {
		recordEncodingFailure(ctx, sr)
		span.SetStatus(trace.Status{
			Code: trace.StatusCodeInvalidArgument,
		})
		return &ErrSend{Err: err}
	}

	serr := w.send(ctx, sr)
	if serr == nil {
		recordSuccess(ctx, sr)
		return nil
	}
	recordSendFailure(ctx, sr)
	span.SetStatus(trace.Status{
		Code: OCStatusCodeFromHTTP(int32(serr.StatusCode)),
	})

	if serr.Permanent {
		recordDrops(ctx, sr)
	}
	return serr
}

func (w *worker) send(ctx context.Context, r *sendRequest) *ErrSend {
	_, span := trace.StartSpan(ctx, "export")
	defer span.End()

	req, err := http.NewRequest("POST", w.endpoint, bytes.NewBuffer(r.message))
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: err.Error(),
		})
		return &ErrSend{Err: err, Permanent: true}
	}
	req.Header.Add(headerContentType, headerValueXProtobuf)

	if !w.disableCompression {
		req.Header.Add(headerContentEncoding, headerValueGZIP)
	}

	if w.accessToken != "" {
		req.Header.Add(headerAccessToken, w.accessToken)
	}

	resp, err := w.client.Do(req)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInternal,
			Message: err.Error(),
		})
		return &ErrSend{Err: err}
	}
	io.CopyN(ioutil.Discard, resp.Body, maxHTTPBodyReadBytes)
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}

	// Drop the batch if server thinks it is malformed in some way or client is not authorized
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized {
		msg := fmt.Sprintf("server responded with: %d", resp.StatusCode)
		span.SetStatus(trace.Status{
			Code:    OCStatusCodeFromHTTP(int32(resp.StatusCode)),
			Message: msg,
		})
		return &ErrSend{
			Err:        fmt.Errorf("dropping request: %s", msg),
			StatusCode: http.StatusBadRequest,
			Permanent:  true,
		}
	}

	// Check if server is overwhelmed and requested to pause sending for a while.
	// Pause from sending more data till the specified number of seconds in the Retry-After header.
	// Fallback to defaultRateLimitingBackoffSeconds if the header is not present
	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := defaultRateLimitingBackoffSeconds
		if val := resp.Header.Get(headerRetryAfter); val != "" {
			if seconds, err := strconv.Atoi(val); err == nil {
				retryAfter = seconds
			}
		}
		span.SetStatus(trace.Status{Code: trace.StatusCodeResourceExhausted})
		return &ErrSend{
			Err:               errors.New("server responded with 429"),
			StatusCode:        resp.StatusCode,
			RetryDelaySeconds: retryAfter,
		}
	}

	// TODO: handle 301, 307, 308
	// redirects are not handled right now but should be to confirm with the spec.

	span.SetStatus(trace.Status{Code: OCStatusCodeFromHTTP(int32(resp.StatusCode))})
	return &ErrSend{
		Err:        fmt.Errorf("error exporting spans. server responded with status %d", resp.StatusCode),
		StatusCode: resp.StatusCode,
	}
}

// prepare takes a jaeger batches, converts them to a SAPM PostSpansRequest, compresses it and returns a request ready
// to be sent.
func (w *worker) prepare(ctx context.Context, batches []*jaegerpb.Batch, spansCount int) (*sendRequest, error) {
	_, span := trace.StartSpan(ctx, "export")
	defer span.End()

	psr := &sapmpb.PostSpansRequest{
		Batches: batches,
	}

	encoded, err := proto.Marshal(psr)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: "failed to marshal request",
		})
		return nil, err
	}

	if w.disableCompression {
		return &sendRequest{
			message: encoded,
			batches: int64(len(batches)),
			spans:   int64(spansCount),
		}, nil
	}

	buf := bytes.NewBuffer([]byte{})
	w.gzipWriter.Reset(buf)

	_, err = w.gzipWriter.Write(encoded)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: "failed to gzip request",
		})
		return nil, err
	}

	if err := w.gzipWriter.Close(); err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: "failed to gzip request",
		})
		return nil, err
	}
	sr := &sendRequest{
		message: buf.Bytes(),
		batches: int64(len(batches)),
		spans:   int64(spansCount),
	}
	return sr, nil
}
