package grpc_stream_lb_srv

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
	client          traceprotobuf.StreamExporterClient
	stream          traceprotobuf.StreamExporter_ExportClient
	pendingAck      map[uint64]*traceprotobuf.ExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
}

func (c *Client) Connect(server string) error {
	c.pendingAck = make(map[uint64]*traceprotobuf.ExportRequest)

	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewStreamExporterClient(conn)

	// Establish stream to server.
	return c.openStream()
}

func (c *Client) openStream() error {
	response, err := c.client.Hello(context.Background(), &traceprotobuf.HelloRequest{})
	if response == nil || err != nil {
		log.Fatalf("No response on Hello: %v", err)
	}

	c.stream, err = c.client.Export(context.Background())
	if err != nil {
		log.Fatalf("cannot open stream: %v", err)
	}

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

	// Add the ID to pendingAck map
	c.pendingAckMutex.Lock()
	c.pendingAck[request.Id] = request
	c.pendingAckMutex.Unlock()

	if err := c.stream.Send(request); err != nil {
		if err == io.EOF {
			// Server closed the stream or disconnected. Try reopening the stream once.
			time.Sleep(1 * time.Second)
			if err = c.openStream(); err != nil {
				log.Fatal("Error opening stream")
			}
			c.resendPending()
		} else {
			log.Fatalf("cannot send request: %v", err)
		}
	}
}

func (c *Client) resendPending() {
	var requests []*traceprotobuf.ExportRequest
	c.pendingAckMutex.Lock()
	for _, request := range c.pendingAck {
		requests = append(requests, request)
	}
	c.pendingAckMutex.Unlock()

	for _, request := range requests {
		if err := c.stream.Send(request); err != nil {
			log.Fatalf("cannot send request: %v", err)
		}
	}
}
