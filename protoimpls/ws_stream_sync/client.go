package ws_stream_sync

import (
	"log"
	"net/url"
	"sync/atomic"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/encodings/wsframing"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	conn   *websocket.Conn
	nextId uint64
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

	bytes := wsframing.Encode(request)
	err := c.conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Fatal("write:", err)
	}

	_, bytes, err = c.conn.ReadMessage()
	if err != nil {
		log.Fatal("read:", err)
		return
	}
	var response traceprotobuf.ExportResponse
	err = proto.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatal("cannnot decode:", err)
	}
}
