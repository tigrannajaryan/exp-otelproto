package ws_async_worker

import (
	"log"
	"net/http"

	"github.com/tigrannajaryan/exp-otelproto/encodings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

type Server struct {
}

var upgrader = websocket.Upgrader{} // use default options

type ConnectionHandler struct {
	conn             *websocket.Conn
	requestsToDecode chan RequestToDecode
	onReceive        func(batch core.ExportRequest, spanCount int)
	responsesToSend  chan []byte
}

type RequestToDecode struct {
	bytes []byte
}

func (ch *ConnectionHandler) requestDecoder() {
	for toDecode := range ch.requestsToDecode {
		request := encodings.Decode(toDecode.bytes)

		//Id := request.GetExport().Id
		//if Id == 0 {
		//	log.Fatal("Received 0 Id")
		//}
		//if Id != lastId+1 {
		//	log.Fatalf("Received out of order request ID=%d instead of expected ID=%d", Id, lastId+1)
		//}
		//lastId = Id

		ch.onReceive(request, len(request.ResourceSpans[0].InstrumentationLibrarySpans[0].Spans))

		response := &otlptracecol.ExportTraceServiceResponse{
			//ResponseType: experimental.RequestType_TraceExport,
			//Export:       &experimental.ExportResponse{Id: Id},
		}
		responseBytes, err := proto.Marshal(response)
		if err != nil {
			log.Fatal("cannot encode:", err)
			break
		}

		ch.responsesToSend <- responseBytes
	}
}

func (ch *ConnectionHandler) responseSender() {
	for responseBytes := range ch.responsesToSend {
		err := ch.conn.WriteMessage(websocket.BinaryMessage, responseBytes)
		if err != nil {
			log.Fatal("write:", err)
			break
		}
	}
}

func telemetryReceiver(w http.ResponseWriter, r *http.Request, onReceive func(batch core.ExportRequest, spanCount int)) {
	//log.Printf("Incoming WS connection.")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade:", err)
		return
	}
	defer conn.Close()

	decoderCount := 1 // runtime.GOMAXPROCS(0)

	ch := ConnectionHandler{
		conn:             conn,
		onReceive:        onReceive,
		requestsToDecode: make(chan RequestToDecode, decoderCount),
		responsesToSend:  make(chan []byte, decoderCount),
	}

	go ch.responseSender()
	defer close(ch.responsesToSend)

	for i := 0; i < decoderCount; i++ {
		go ch.requestDecoder()
	}
	defer close(ch.requestsToDecode)

	//lastId := uint64(0)
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			break
		}

		ch.requestsToDecode <- RequestToDecode{bytes: bytes}
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
