package ws_stream_async

import (
	"log"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/tigrannajaryan/exp-otelproto/encodings"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type Client struct {
	Compression   experimental.CompressionMethod
	clientStreams []*clientStream
	Concurrency   int
	requestsCh    chan *experimental.TraceExportRequest
	//semaphor   chan int
	nextStream int64
}

// clientStream can connect to a server and send a batch of spans.
type clientStream struct {
	conn            *websocket.Conn
	pendingAck      map[uint64]core.ExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
	Compression     experimental.CompressionMethod
	requestsCh      chan *experimental.TraceExportRequest
}

func (c *Client) Connect(server string) error {
	//c.semaphor = make(chan int, c.Concurrency)
	c.requestsCh = make(chan *experimental.TraceExportRequest, 10*c.Concurrency)
	c.clientStreams = make([]*clientStream, c.Concurrency)

	for i := 0; i < c.Concurrency; i++ {
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
	//c.sentCh = client.sentCh
	// c.pendingAckList = list.New()
	// go c.processTimeouts()
	go c.processSendRequests()
	return &c
}

func (c *Client) Export(batch core.ExportRequest) {
	if c.Concurrency == 1 {
		c.clientStreams[0].sendRequest(batch.(*experimental.TraceExportRequest))
		return
	}

	// Make sure we have only up to c.Concurrency Export calls in progress
	// concurrently. It means no single stream has concurrent sendRequests
	// in progress, so sendRequest does not need to be safe for concurrent call.

	//si := atomic.AddInt64(&c.nextStream, 1)
	//c.semaphor <- 1
	//c.clientStreams[si%int64(c.Concurrency)].requestsCh <- batch.(*experimental.TraceExportRequest)
	//<-c.semaphor

	c.requestsCh <- batch.(*experimental.TraceExportRequest)
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
	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	go c.readStream()

	return nil
}

func (c *clientStream) processSendRequests() {
	for request := range c.requestsCh {
		c.sendRequest(request)
		//c.sentCh <- true
	}
}

func (c *clientStream) readStream() {
	// defer close(done)
	lastId := uint64(0)
	for {
		_, bytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}
		var response experimental.Response
		err = proto.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatal("cannnot decode:", err)
			break
		}

		Id := response.GetExport().Id
		if Id != lastId+1 {
			log.Fatalf("Received out of order response ID=%d", Id)
		}
		lastId = Id

		//c.pendingAckMutex.Lock()
		//_, ok := c.pendingAck[Id]
		//if !ok {
		//	c.pendingAckMutex.Unlock()
		//	log.Fatalf("Received ack on batch ID that does not exist: %v", Id)
		//}
		//delete(c.pendingAck, Id)
		//c.pendingAckMutex.Unlock()
	}
}

func (c *clientStream) sendRequest(batch core.ExportRequest) {
	request := batch.(*experimental.TraceExportRequest)
	if request.Id != 0 {
		log.Fatal("Request is still processing but got overwritten")
	}

	Id := atomic.AddUint64(&c.nextId, 1)
	request.Id = Id

	body := &experimental.RequestBody{Body: &experimental.RequestBody_Export{request}}
	bytes := encodings.Encode(body, c.Compression)
	request.Id = 0

	//// Add the ID to pendingAck map
	//c.pendingAckMutex.Lock()
	//c.pendingAck[Id] = batch
	//c.pendingAckMutex.Unlock()

	err := c.conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Fatal("write:", err)
	}
}

func (c *clientStream) Shutdown() {
}
