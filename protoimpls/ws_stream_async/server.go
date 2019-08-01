package ws_stream_async

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/gorilla/websocket"
	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

type Server struct {
}

var upgrader = websocket.Upgrader{} // use default options

func telemetryReceiver(w http.ResponseWriter, r *http.Request, onReceive func(batch core.SpanBatch)) {
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
		// log.Printf("recv: %s", message)

		var spanBatch traceprotobuf.SpanBatch
		err = proto.Unmarshal(bytes, &spanBatch)
		if err != nil {
			log.Fatal("cannnot decode:", err)
			break
		}

		onReceive(spanBatch)

		responseBytes, err := proto.Marshal(&traceprotobuf.BatchResponse{Id: spanBatch.Id})
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

func (srv *Server) Listen(endpoint string, onReceive func(batch core.SpanBatch)) error {
	http.HandleFunc(
		"/telemetry",
		func(w http.ResponseWriter, r *http.Request) { telemetryReceiver(w, r, onReceive) },
	)
	log.Fatal(http.ListenAndServe(endpoint, nil))
	return nil
}

func (srv *Server) Stop() {
}
