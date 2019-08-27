package grpc_stream_lb_async

import (
	"container/list"
	"context"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client                   otlp.StreamExporterClient
	clientStreams            []*clientStream
	Concurrency              int
	requestsCh               chan *otlp.TraceExportRequest
	sentCh                   chan bool
	nextStream               int
	nextStreamMux            sync.Mutex
	StreamReopenPeriod       time.Duration
	StreamReopenRequestCount uint32
}

type clientStream struct {
	client                      *Client
	stream                      otlp.StreamExporter_ExportTracesClient
	lastStreamOpen              time.Time
	pendingAckMap               map[uint64]*list.Element
	pendingAckMutex             sync.Mutex
	pendingAckList              *list.List
	nextId                      uint64
	requestsSentSinceStreamOpen uint32
	requestsCh                  chan *otlp.TraceExportRequest
	sentCh                      chan bool
}

func newClientStream(client *Client) *clientStream {
	c := clientStream{}
	c.client = client
	c.requestsCh = client.requestsCh
	c.sentCh = client.sentCh
	c.pendingAckMap = make(map[uint64]*list.Element)
	c.pendingAckList = list.New()
	go c.processTimeouts()
	if err := c.openStream(); err != nil {
		log.Fatal(err)
	}
	go c.processSendRequests()
	return &c
}

type pendingRequest struct {
	request  *otlp.TraceExportRequest
	deadline time.Time
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = otlp.NewStreamExporterClient(conn)

	if c.Concurrency < 1 {
		c.Concurrency = 1
	}
	c.requestsCh = make(chan *otlp.TraceExportRequest, c.Concurrency)
	c.sentCh = make(chan bool, c.Concurrency)

	for i := 0; i < c.Concurrency; i++ {
		c.clientStreams = append(c.clientStreams, newClientStream(c))
	}
	return nil
}

func (c *clientStream) openStream() error {
	var err error
	c.stream, err = c.client.client.ExportTraces(context.Background())
	if err != nil {
		log.Fatalf("cannot open stream: %v", err)
	}
	c.lastStreamOpen = time.Now()
	atomic.StoreUint32(&c.requestsSentSinceStreamOpen, 0)

	go c.readStream(c.stream)

	return nil
}

func (c *clientStream) readStream(stream otlp.StreamExporter_ExportTracesClient) {
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// read done.
			return
		}
		if err != nil {
			// log.Printf("Failed to read from stream: %v", err)
			return
		}

		c.pendingAckMutex.Lock()
		elem, ok := c.pendingAckMap[response.Id]
		if !ok {
			c.pendingAckMutex.Unlock()
			log.Fatalf("Received ack on batch ID that does not exist: %v", response.Id)
		}
		delete(c.pendingAckMap, response.Id)
		c.pendingAckList.Remove(elem)
		c.pendingAckMutex.Unlock()
	}
}

func (c *clientStream) processTimeouts() {
	ch := time.Tick(1 * time.Second)
	for now := range ch {
		for {
			c.pendingAckMutex.Lock()
			elem := c.pendingAckList.Back()
			if elem == nil {
				c.pendingAckMutex.Unlock()
				break
			}
			pr := elem.Value.(pendingRequest)
			if pr.deadline.Before(now) {
				c.pendingAckList.Remove(elem)
				delete(c.pendingAckMap, pr.request.Id)
			} else {
				c.pendingAckMutex.Unlock()
				break
			}
			c.pendingAckMutex.Unlock()
			log.Printf("Request %v timed out", pr.request.Id)
		}
	}
}

func (c *clientStream) processSendRequests() {
	for request := range c.requestsCh {
		c.sendRequest(request)
		//c.sentCh <- true
	}
}

func (c *clientStream) sendRequest(
	request *otlp.TraceExportRequest,
) {
	// Send the batch via stream.
	request.Id = atomic.AddUint64(&c.nextId, 1)

	pr := pendingRequest{request: request, deadline: time.Now().Add(30 * time.Second)}

	// Add the ID to pendingAckMap map
	c.pendingAckMutex.Lock()
	elem := c.pendingAckList.PushFront(pr)
	c.pendingAckMap[request.Id] = elem
	c.pendingAckMutex.Unlock()

	if err := c.stream.Send(request); err != nil {
		if err == io.EOF {
			// Server closed the stream or disconnected. Try reopening the stream once.
			log.Print("Reopening stream")
			time.Sleep(1 * time.Second)
			if err = c.openStream(); err != nil {
				log.Fatal("Error opening stream")
			}
			c.resendPending()
		} else {
			log.Fatalf("cannot send request: %v", err)
		}
	}

	requestsSent := atomic.AddUint32(&c.requestsSentSinceStreamOpen, 1)

	// Check if time to re-establish the stream.
	if requestsSent > c.client.StreamReopenRequestCount ||
		time.Since(c.lastStreamOpen) > c.client.StreamReopenPeriod {
		// Close and reopen the stream.
		c.lastStreamOpen = time.Now()
		err := c.stream.CloseSend()
		if err != nil {
			log.Fatal("Error closing stream")
		}
		if err = c.openStream(); err != nil {
			log.Fatal("Error opening stream")
		}
	}
}

func (c *Client) Export(batch core.ExportRequest) {
	if c.Concurrency == 1 {
		c.clientStreams[0].sendRequest(batch.(*otlp.TraceExportRequest))
		return
	}

	// Make sure we have only up to c.Concurrency Export calls in progress
	// concurrently. It means no single stream has concurrent sendRequests
	// in progress, so sendRequest does not need to be safe for concurrent call.
	c.requestsCh <- batch.(*otlp.TraceExportRequest)

	// Wait until it is sent
	// <-c.sentCh

	//go func() {
	//	defer func() { <-c.concurrencySem }()
	//
	//	// Find a stream to send via. Use list of stream circularly.
	//	c.nextStreamMux.Lock()
	//	c.nextStream++
	//	if c.nextStream >= c.Concurrency {
	//		c.nextStream = 0
	//	}
	//	streamIndex := c.nextStream
	//	c.nextStreamMux.Unlock()
	//
	//	c.clientStreams[streamIndex].sendRequest(batch.(*attrlist.TraceExportRequest))
	//}()
}

func (c *clientStream) resendPending() {
	var requests []*otlp.TraceExportRequest
	c.pendingAckMutex.Lock()
	for _, request := range c.pendingAckMap {
		requests = append(requests, request.Value.(*otlp.TraceExportRequest))
	}
	c.pendingAckMutex.Unlock()

	log.Printf("Resending %d requests", len(requests))
	for _, request := range requests {
		if err := c.stream.Send(request); err != nil {
			log.Fatalf("cannot send request: %v", err)
		}
	}
}

func (c *Client) Shutdown() {

}
