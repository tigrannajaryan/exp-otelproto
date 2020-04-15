package http11

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/tigrannajaryan/exp-otelproto/encodings"

	"github.com/golang/protobuf/proto"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

type Client struct {
	Compression   otlp.CompressionMethod
	clientStreams []*worker
	Concurrency   int
	requestsCh    chan *otlp.TraceExportRequest
	//semaphor   chan int
	nextStream int64
}

// clientStream can connect to a server and send a batch of spans.
type worker struct {
	//conn            *websocket.Conn
	//pendingAck      map[uint64]core.ExportRequest
	//pendingAckMutex sync.Mutex
	nextId      uint64
	Compression otlp.CompressionMethod
	requestsCh  chan *otlp.TraceExportRequest
	url         string
}

func (c *Client) Connect(server string) error {
	//c.semaphor = make(chan int, c.Concurrency)
	c.requestsCh = make(chan *otlp.TraceExportRequest, 10*c.Concurrency)
	c.clientStreams = make([]*worker, c.Concurrency)

	for i := 0; i < c.Concurrency; i++ {
		c.clientStreams[i] = newWorker(c)
		err := c.clientStreams[i].Connect(server)
		if err != nil {
			return err
		}
	}
	return nil
}

func newWorker(client *Client) *worker {
	c := worker{}
	// c.client = client
	c.requestsCh = client.requestsCh
	//c.requestsCh = make(chan *otlp.TraceExportRequest, 0)
	c.Compression = client.Compression
	//c.sentCh = client.sentCh
	// c.pendingAckList = list.New()
	// go c.processTimeouts()
	go c.processSendRequests()
	return &c
}

func (c *Client) Export(batch core.ExportRequest) {
	if c.Concurrency == 1 {
		c.clientStreams[0].sendRequest(batch.(*otlp.TraceExportRequest))
		return
	}

	// Make sure we have only up to c.Concurrency Export calls in progress
	// concurrently. It means no single stream has concurrent sendRequests
	// in progress, so sendRequest does not need to be safe for concurrent call.

	//si := atomic.AddInt64(&c.nextStream, 1)
	//c.semaphor <- 1
	//c.clientStreams[si%int64(c.Concurrency)].requestsCh <- batch.(*otlp.TraceExportRequest)
	//<-c.semaphor

	c.requestsCh <- batch.(*otlp.TraceExportRequest)
}

func (c *Client) Shutdown() {
	for _, cs := range c.clientStreams {
		cs.Shutdown()
	}
}

func (c *worker) Connect(server string) error {
	url := url.URL{Scheme: "http", Host: server, Path: "/telemetry"}
	c.url = url.String()
	return nil
}

func (c *worker) processSendRequests() {
	for request := range c.requestsCh {
		c.sendRequest(request)
	}
}

func (c *worker) sendRequest(batch core.ExportRequest) {
	request := batch.(*otlp.TraceExportRequest)
	if request.Id != 0 {
		log.Fatal("Request is still processing but got overwritten")
	}

	Id := atomic.AddUint64(&c.nextId, 1)
	request.Id = Id

	body := &otlp.RequestBody{Body: &otlp.RequestBody_Export{request}}
	b := encodings.Encode(body, c.Compression)
	request.Id = 0

	buf := bytes.NewBuffer(b)
	resp, err := http.Post(c.url, "application/x-prtobuf", buf)
	if err != nil {
		log.Fatal("write:", err)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("write:", err)
	}

	var response otlp.Response
	err = proto.Unmarshal(b, &response)
	if err != nil {
		log.Fatal("cannnot decode:", err)
	}
}

func (c *worker) Shutdown() {
}
