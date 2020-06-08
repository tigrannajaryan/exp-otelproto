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

package sapmprotocol

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/golang/protobuf/proto"

	splunksapm "github.com/signalfx/sapm-proto/gen"
)

type poolObj struct {
	gr   *gzip.Reader
	jeff *bytes.Buffer
	tmp  []byte
}

var (
	// ErrBadContentType indicates an incompatible content type was received
	ErrBadContentType = errors.New("bad content type")

	pool = &sync.Pool{
		New: func() interface{} {
			// create a new gzip reader with a bytes reader and array of bytes containing only the gzip header
			gr, _ := gzip.NewReader(bytes.NewReader([]byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 255, 255, 1, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0}))
			jeff := &bytes.Buffer{}
			tmp := make([]byte, 32*1024)
			return &poolObj{
				gr:   gr,
				jeff: jeff,
				tmp:  tmp,
			}
		},
	}
)

// ParseTraceV2Request processes an http request request into SAPM
func ParseTraceV2Request(req *http.Request) (*splunksapm.PostSpansRequest, error) {
	// content type MUST be application/x-protobuf
	if req.Header.Get(ContentTypeHeaderName) != ContentTypeHeaderValue {
		return nil, ErrBadContentType
	}

	var err error
	var reader io.Reader

	obj := pool.Get().(*poolObj)
	defer pool.Put(obj)
	obj.jeff.Reset()

	// content encoding SHOULD be gzip
	if req.Header.Get(ContentEncodingHeaderName) == GZipEncodingHeaderValue {
		// get the gzip reader
		// reset the reader with the request body
		err = obj.gr.Reset(req.Body)
		if err != nil {
			return nil, err
		}
		reader = obj.gr
	} else {
		reader = req.Body
	}

	_, err = io.CopyBuffer(obj.jeff, reader, obj.tmp)
	if err != nil {
		return nil, err
	}

	var sapm = &splunksapm.PostSpansRequest{}

	// unmarshal request body
	err = proto.Unmarshal(obj.jeff.Bytes(), sapm)
	if err != nil {
		return nil, err
	}

	return sapm, err
}
