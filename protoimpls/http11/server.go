package http11

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

type Server struct {
}

func telemetryReceiver(w http.ResponseWriter, r *http.Request, onReceive func(batch core.ExportRequest, spanCount int)) {
	//log.Printf("Incoming WS connection.")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("read:", err)
	}

	request := encodings.Decode(bytes)

	Id := request.GetExport().Id
	if Id == 0 {
		log.Fatal("Received 0 Id")
	}

	onReceive(request, len(request.GetExport().ResourceSpans[0].Spans))

	response := &otlp.Response{
		Body: &otlp.Response_Export{
			Export: &otlp.ExportResponse{Id: Id},
		},
	}
	responseBytes, err := proto.Marshal(response)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	_, err = w.Write(responseBytes)
	if err != nil {
		log.Fatal("write:", err)
	}
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest, spanCount int)) error {
	http.HandleFunc(
		"/telemetry",
		func(w http.ResponseWriter, r *http.Request) { telemetryReceiver(w, r, onReceive) },
	)
	log.Fatal(http.ListenAndServe(endpoint, nil))
	return nil
}

func (srv *Server) Stop() {
}
