package http11

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"

	"github.com/golang/protobuf/proto"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type Client struct {
	Compression   experimental.CompressionMethod
	clientStreams []*worker
	Concurrency   int
	requestsCh    chan *otlptracecol.ExportTraceServiceRequest
	//semaphor   chan int
	nextStream int64
}

// clientStream can connect to a server and send a batch of spans.
type worker struct {
	//conn            *websocket.Conn
	//pendingAck      map[uint64]core.ExportRequest
	//pendingAckMutex sync.Mutex
	nextId      uint64
	Compression experimental.CompressionMethod
	requestsCh  chan *otlptracecol.ExportTraceServiceRequest
	url         string
	httpClient  *http.Client
}

func (c *Client) Connect(server string) error {
	//c.semaphor = make(chan int, c.Concurrency)
	c.requestsCh = make(chan *otlptracecol.ExportTraceServiceRequest, 10*c.Concurrency)
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
	//c.requestsCh = make(chan *experimental.TraceExportRequest, 0)
	c.Compression = client.Compression
	//c.sentCh = client.sentCh
	// c.pendingAckList = list.New()
	// go c.processTimeouts()

	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic("defaultRoundTripper not an *http.Transport")
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	c.httpClient = &http.Client{Transport: &defaultTransport}
	go c.processSendRequests()
	return &c
}

func (c *Client) Export(batch core.ExportRequest) {
	if c.Concurrency == 1 {
		c.clientStreams[0].sendRequest(batch.(*otlptracecol.ExportTraceServiceRequest))
		return
	}

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
	request := batch.(*otlptracecol.ExportTraceServiceRequest)
	//if request.Id != 0 {
	//	log.Fatal("Request is still processing but got overwritten")
	//}

	//Id := atomic.AddUint64(&c.nextId, 1)
	//request.Id = Id

	//body := &experimental.RequestBody{
	//	RequestType: experimental.RequestType_TraceExport,
	//	Export:      request,
	//}
	//b := encodings.Encode(body, c.Compression)
	//request.Id = 0

	b, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("write:", err)
	}

	buf := bytes.NewBuffer(b)
	resp, err := c.httpClient.Post(c.url, "application/x-protobuf", buf)
	if err != nil {
		log.Fatal("write:", err)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("write:", err)
	}

	resp.Body.Close()

	var response otlptracecol.ExportTraceServiceResponse
	err = proto.Unmarshal(b, &response)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}
}

func (c *worker) Shutdown() {
}
