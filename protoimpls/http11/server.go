package http11

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type Server struct {
}

func telemetryReceiver(w http.ResponseWriter, r *http.Request, onReceive func(batch core.ExportRequest, spanCount int)) {
	//log.Printf("Incoming WS connection.")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("read:", err)
	}

	//request := encodings.Decode(bytes)
	var request otlptracecol.ExportTraceServiceRequest
	err = proto.Unmarshal(bytes, &request)
	if err != nil {
		log.Fatal("Unmarshal:", err)
	}

	//Id := request.GetExport().Id
	//if Id == 0 {
	//	log.Fatal("Received 0 Id")
	//}

	onReceive(request, len(request.ResourceSpans[0].InstrumentationLibrarySpans[0].Spans))

	response := &otlptracecol.ExportTraceServiceResponse{}
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
	//http.HandleFunc(
	//	"/telemetry",
	//	func(w http.ResponseWriter, r *http.Request) { telemetryReceiver(w, r, onReceive) },
	//)

	m := http.NewServeMux()
	m.HandleFunc(
		"/telemetry",
		func(w http.ResponseWriter, r *http.Request) { telemetryReceiver(w, r, onReceive) },
	)

	s := &http.Server{
		Handler:      m,
		Addr:         endpoint,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(s.ListenAndServe())
	// log.Fatal(http.ListenAndServe(endpoint, nil))
	return nil
}

func (srv *Server) Stop() {
}
