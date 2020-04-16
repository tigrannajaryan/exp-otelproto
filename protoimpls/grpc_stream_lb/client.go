package grpc_stream_lb

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client                  experimental.StreamExporterClient
	stream                  experimental.StreamExporter_ExportTracesClient
	lastStreamOpen          time.Time
	ReopenAfterEveryRequest bool
	StreamReopenPeriod      time.Duration
	nextId                  uint64
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = experimental.NewStreamExporterClient(conn)

	// Establish stream to server.
	return c.openStream()
}

func (c *Client) openStream() error {
	var err error
	c.stream, err = c.client.ExportTraces(context.Background())
	if err != nil {
		log.Fatalf("Cannot open stream: %v", err)
	}
	c.lastStreamOpen = time.Now()
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	// Send the batch via stream.
	request := batch.(*experimental.TraceExportRequest)
	request.Id = atomic.AddUint64(&c.nextId, 1)
	c.stream.Send(request)

	// Wait for response from server. This is full synchronous operation,
	// we do not send batches concurrently.
	_, err := c.stream.Recv()

	if err != nil {
		log.Fatal("Error from server when expecting batch response")
	}

	if c.ReopenAfterEveryRequest || time.Since(c.lastStreamOpen) > c.StreamReopenPeriod {
		// Close and reopen the stream.
		c.lastStreamOpen = time.Now()
		err = c.stream.CloseSend()
		if err != nil {
			log.Fatal("Error closing stream")
		}
		if err = c.openStream(); err != nil {
			log.Fatal("Error opening stream")
		}
	}
}

func (c *Client) Shutdown() {
}
