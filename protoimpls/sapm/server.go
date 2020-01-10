package sapm

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	splunksapm "github.com/signalfx/sapm-proto/gen"
	"github.com/signalfx/sapm-proto/sapmprotocol"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

var (
	// ErrAlreadyStarted indicates an error on starting an already-started
	// receiver/processor/exporter.
	ErrAlreadyStarted = errors.New("already started")

	// ErrAlreadyStopped indicates an error on stoping an already-stopped
	// receiver/processor/exporter.
	ErrAlreadyStopped = errors.New("already stopped")

	// ErrNilNextConsumer indicates an error on nil next consumer.
	ErrNilNextConsumer = errors.New("nil nextConsumer")
)

var gzipWriterPool = &sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

type Server struct {
	mu        sync.Mutex
	startOnce sync.Once
	stopOnce  sync.Once
	server    *http.Server

	// defaultResponse is a placeholder. For now this receiver returns an empty sapm response.
	// This defaultResponse is an optimization so we don't have to proto.Marshal the response
	// for every request. At some point this may be removed when there is actual content to return.
	defaultResponse []byte

	onReceive func(batch core.ExportRequest, spanCount int)
}

func (sr *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest, spanCount int)) error {

	sr.onReceive = onReceive

	// build the response message
	var err error
	sr.defaultResponse, err = proto.Marshal(&splunksapm.PostSpansResponse{})
	if err != nil {
		return fmt.Errorf("failed to marshal default response body for receiver: %v", err)
	}

	sr.mu.Lock()
	defer sr.mu.Unlock()

	err = ErrAlreadyStarted
	sr.startOnce.Do(func() {
		var ln net.Listener

		// set up the listener
		ln, err = net.Listen("tcp", endpoint)
		if err != nil {
			err = fmt.Errorf("failed to bind to address %s: %v", endpoint, err)
			return
		}

		// use gorilla mux to create a router/handler
		nr := mux.NewRouter()
		nr.HandleFunc(sapmprotocol.TraceEndpointV2, sr.HTTPHandlerFunc)

		// create a server with the handler
		sr.server = &http.Server{Handler: nr}

		// run the server on a routine
		go func() {
			log.Fatal(sr.server.Serve(ln))
		}()
	})
	return err
}

func (sr *Server) Stop() {
}

// handleRequest parses an http request containing sapm and passes the trace data to the next consumer
func (sr *Server) handleRequest(req *http.Request) error {
	sapm, err := sapmprotocol.ParseTraceV2Request(req)
	// errors processing the request should return http.StatusBadRequest
	if err != nil {
		return err
	}

	// process sapm batches
	for _, request := range sapm.Batches {
		// pass the trace data to the next consumer
		sr.onReceive(request, len(request.Spans))
	}

	return nil
}

// HTTPHandlerFunction returns an http.HandlerFunc that handles SAPM requests
func (sr *Server) HTTPHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	// handle the request payload
	err := sr.handleRequest(req)
	if err != nil {
		// TODO account for this error (throttled logging or metrics)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// respBytes are bytes to write to the http.Response

	// build the response message
	// NOTE currently the response is an empty struct.  As an optimization this receiver will pass a
	// byte array that was generated in the receiver's constructor.  If this receiver needs to return
	// more than an empty struct, then the sapm.PostSpansResponse{} struct will need to be marshalled
	// and on error a http.StatusInternalServerError should be written to the http.ResponseWriter and
	// this function should immediately return.
	var respBytes = sr.defaultResponse
	rw.Header().Set(sapmprotocol.ContentTypeHeaderName, sapmprotocol.ContentTypeHeaderValue)

	// write the response if client does not accept gzip encoding
	if req.Header.Get(sapmprotocol.AcceptEncodingHeaderName) != sapmprotocol.GZipEncodingHeaderValue {
		// write the response bytes
		rw.Write(respBytes)
		return
	}

	// gzip the response

	// get the gzip writer
	writer := gzipWriterPool.Get().(*gzip.Writer)
	defer gzipWriterPool.Put(writer)

	var gzipBuffer bytes.Buffer

	// reset the writer with the gzip buffer
	writer.Reset(&gzipBuffer)

	// gzip the responseBytes
	_, err = writer.Write(respBytes)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// close the gzip writer and write gzip footer
	err = writer.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write the successfully gzipped payload
	rw.Header().Set(sapmprotocol.ContentEncodingHeaderName, sapmprotocol.GZipEncodingHeaderValue)
	rw.Write(gzipBuffer.Bytes())
}

// StopTraceRetention stops the the sapmReceiver's server
func (sr *Server) Shutdown() error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	var err = ErrAlreadyStopped
	sr.stopOnce.Do(func() {
		if sr.server != nil {
			err = sr.server.Close()
			sr.server = nil
		}
	})

	return err
}
