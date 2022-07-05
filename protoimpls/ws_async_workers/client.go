package ws_async_worker

import (
	"log"
	"net/url"
	"sync"

	"github.com/tigrannajaryan/exp-otelproto/encodings"
	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

type Client struct {
	Compression   experimental.CompressionMethod
	clientStreams []*clientStream
	Concurrency   int
	requestsCh    chan *otlptracecol.ExportTraceServiceRequest
	//semaphor   chan int
	nextStream int64
}

// clientStream can connect to a server and send a batch of spans.
type clientStream struct {
	conn            *websocket.Conn
	encoderCount    int
	pendingAck      map[uint64]core.ExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
	Compression     experimental.CompressionMethod
	requestsCh      chan *otlptracecol.ExportTraceServiceRequest
	bytesCh         chan []byte
}

func (c *Client) Connect(server string) error {
	//c.semaphor = make(chan int, c.Concurrency)
	c.requestsCh = make(chan *otlptracecol.ExportTraceServiceRequest, 10*c.Concurrency)
	c.clientStreams = make([]*clientStream, 1)

	for i := 0; i < 1; i++ {
		c.clientStreams[i] = newClientStream(c)
		err := c.clientStreams[i].Connect(server)
		if err != nil {
			return err
		}
	}
	return nil
}

func newClientStream(client *Client) *clientStream {
	c := clientStream{}
	// c.client = client
	c.requestsCh = client.requestsCh
	//c.requestsCh = make(chan *experimental.TraceExportRequest, 0)
	c.Compression = client.Compression
	c.bytesCh = make(chan []byte, 10*client.Concurrency)
	c.encoderCount = 1
	//c.sentCh = client.sentCh
	// c.pendingAckList = list.New()
	// go c.processTimeouts()
	for i := 0; i < c.encoderCount; i++ {
		go c.processEncodeRequests()
	}

	go c.processReadyToSendBytes()

	return &c
}

func (c *Client) Export(batch core.ExportRequest) {
	//if c.Concurrency == 1 {
	//	c.clientStreams[0].encodeRequest(batch.(*otlptracecol.ExportTraceServiceRequest))
	//	return
	//}

	// Make sure we have only up to c.Concurrency Export calls in progress
	// concurrently. It means no single stream has concurrent sendRequests
	// in progress, so sendRequest does not need to be safe for concurrent call.

	//si := atomic.AddInt64(&c.nextStream, 1)
	//c.semaphor <- 1
	//c.clientStreams[si%int64(c.Concurrency)].requestsCh <- batch.(*experimental.TraceExportRequest)
	//<-c.semaphor

	c.requestsCh <- batch.(*otlptracecol.ExportTraceServiceRequest)
}

func (c *Client) Shutdown() {
	for _, cs := range c.clientStreams {
		cs.Shutdown()
	}
}

func (c *clientStream) Connect(server string) error {
	// Set up a connection to the server.
	c.pendingAck = make(map[uint64]core.ExportRequest)

	u := url.URL{Scheme: "ws", Host: server, Path: "/telemetry"}

	var err error
	dialer := *websocket.DefaultDialer
	dialer.WriteBufferSize = 256 * 1024
	c.conn, _, err = dialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	go c.readStream()

	return nil
}

func (c *clientStream) processEncodeRequests() {
	for request := range c.requestsCh {
		c.encodeRequest(request)
	}
}

func (c *clientStream) processReadyToSendBytes() {
	for bytes := range c.bytesCh {
		c.sendBytes(bytes)
	}
}

func (c *clientStream) readStream() {
	for {
		_, bytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}
		var response otlptracecol.ExportTraceServiceResponse
		err = proto.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatal("cannnot decode:", err)
			break
		}
	}
}

func (c *clientStream) encodeRequest(batch core.ExportRequest) {
	request := batch.(*otlptracecol.ExportTraceServiceRequest)
	bytes := encodings.Encode(request, c.Compression)
	c.bytesCh <- bytes
}

func (c *clientStream) sendBytes(bytes []byte) {
	err := c.conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Fatal("write:", err)
	}
}

func (c *clientStream) Shutdown() {
}
