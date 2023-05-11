package ws_stream_sync

import (
	"log"
	"net/http"

	"github.com/tigrannajaryan/exp-otelproto/encodings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type Server struct {
}

var upgrader = websocket.Upgrader{} // use default options

func telemetryReceiver(w http.ResponseWriter, r *http.Request, onReceive func(batch core.ExportRequest, spanCount int)) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, bytes, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			break
		}

		request := encodings.Decode(bytes)

		//if request.GetExport().Id == 0 {
		//	log.Fatal("Received 0 Id")
		//}

		onReceive(request, len(request.ResourceSpans[0].ScopeSpans[0].Spans))

		response := &otlptracecol.ExportTraceServiceResponse{}
		responseBytes, err := proto.Marshal(response)
		if err != nil {
			log.Fatal("cannot encode:", err)
			break
		}

		err = c.WriteMessage(mt, responseBytes)
		if err != nil {
			log.Fatal("write:", err)
			break
		}
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
