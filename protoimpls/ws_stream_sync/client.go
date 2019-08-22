package ws_stream_sync

import (
	"log"
	"net/url"
	"sync/atomic"

	"github.com/tigrannajaryan/exp-otelproto/encodings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	conn        *websocket.Conn
	nextId      uint64
	Compression traceprotobuf.CompressionMethod
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	u := url.URL{Scheme: "ws", Host: server, Path: "/telemetry"}

	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	return nil
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

	_, bytes, err = c.conn.ReadMessage()
	if err != nil {
		log.Fatal("read:", err)
		return
	}
	var response traceprotobuf.Response
	err = proto.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatal("cannnot decode:", err)
	}
	if response.GetExport().Id != request.Id {
		log.Fatal("received ack on unexpected ID")
	}
}

func (c *Client) Shutdown() {
}
