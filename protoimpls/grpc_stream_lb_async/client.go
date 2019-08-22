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
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client                      traceprotobuf.StreamExporterClient
	stream                      traceprotobuf.StreamExporter_ExportClient
	lastStreamOpen              time.Time
	pendingAckMap               map[uint64]*list.Element
	pendingAckMutex             sync.Mutex
	pendingAckList              *list.List
	StreamReopenPeriod          time.Duration
	StreamReopenRequestCount    uint32
	nextId                      uint64
	requestsSentSinceStreamOpen uint32
}

type pendingRequest struct {
	request  *traceprotobuf.ExportRequest
	deadline time.Time
}

func (c *Client) Connect(server string) error {
	c.pendingAckMap = make(map[uint64]*list.Element)
	c.pendingAckList = list.New()

	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewStreamExporterClient(conn)

	go c.processTimeouts()

	// Establish stream to server.
	return c.openStream()
}

func (c *Client) openStream() error {
	var err error
	c.stream, err = c.client.Export(context.Background())
	if err != nil {
		log.Fatalf("cannot open stream: %v", err)
	}
	c.lastStreamOpen = time.Now()
	atomic.StoreUint32(&c.requestsSentSinceStreamOpen, 0)

	go c.readStream(c.stream)

	return nil
}

func (c *Client) readStream(stream traceprotobuf.StreamExporter_ExportClient) {
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

func (c *Client) processTimeouts() {
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

func (c *Client) Export(batch core.ExportRequest) {
	// Send the batch via stream.
	request := batch.(*traceprotobuf.ExportRequest)
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
	if requestsSent > c.StreamReopenRequestCount || time.Since(c.lastStreamOpen) > c.StreamReopenPeriod {
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

func (c *Client) resendPending() {
	var requests []*traceprotobuf.ExportRequest
	c.pendingAckMutex.Lock()
	for _, request := range c.pendingAckMap {
		requests = append(requests, request.Value.(*traceprotobuf.ExportRequest))
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
