package ws_stream_async

import (
	"log"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/tigrannajaryan/exp-otelproto/encodings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	conn            *websocket.Conn
	pendingAck      map[uint64]core.ExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
	Compression     traceprotobuf.CompressionMethod
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	c.pendingAck = make(map[uint64]core.ExportRequest)

	u := url.URL{Scheme: "ws", Host: server, Path: "/telemetry"}

	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	go c.readStream()

	return nil
}

func (c *Client) readStream() {
	// defer close(done)
	for {
		_, bytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}
		var response traceprotobuf.ExportResponse
		err = proto.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatal("cannnot decode:", err)
			break
		}

		c.pendingAckMutex.Lock()
		_, ok := c.pendingAck[response.Id]
		if !ok {
			c.pendingAckMutex.Unlock()
			log.Fatalf("Received ack on batch ID that does not exist: %v", response.Id)
		}
		delete(c.pendingAck, response.Id)
		c.pendingAckMutex.Unlock()
	}
}

func (c *Client) Export(batch core.ExportRequest) {
	request := batch.(*traceprotobuf.ExportRequest)
	request.Id = atomic.AddUint64(&c.nextId, 1)

	body := &traceprotobuf.RequestBody{Body: &traceprotobuf.RequestBody_Export{request}}
	bytes := encodings.Encode(body, c.Compression)

	err := c.conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Fatal("write:", err)
	}

	// Add the ID to pendingAck map
	c.pendingAckMutex.Lock()
	c.pendingAck[request.Id] = batch
	c.pendingAckMutex.Unlock()
}
