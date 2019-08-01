package grpc_stream_lb_async

import (
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
	client          traceprotobuf.StreamTracerClient
	stream          traceprotobuf.StreamTracer_ExportClient
	lastStreamOpen  time.Time
	pendingAck      map[uint64]core.ExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
}

// How often to reopen the stream to help LB's rebalance traffic.
var streamReopenPeriod = 30 * time.Second

func (c *Client) Connect(server string) error {
	c.pendingAck = make(map[uint64]core.ExportRequest)

	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewStreamTracerClient(conn)

	// Establish stream to server.
	return c.openStream()
}

func (c *Client) openStream() error {
	var err error
	c.stream, err = c.client.Export(context.Background())
	if err != nil {
		return err
	}
	c.lastStreamOpen = time.Now()

	go c.readStream(c.stream)

	return nil
}

func (c *Client) readStream(stream traceprotobuf.StreamTracer_ExportClient) {
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
	// Send the batch via stream.
	request := batch.(*traceprotobuf.ExportRequest)
	request.Id = atomic.AddUint64(&c.nextId, 1)

	c.stream.Send(request)

	// Add the ID to pendingAck map
	c.pendingAckMutex.Lock()
	c.pendingAck[request.Id] = batch
	c.pendingAckMutex.Unlock()

	// Check if time to re-establish the stream.
	if time.Since(c.lastStreamOpen) > streamReopenPeriod {
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
