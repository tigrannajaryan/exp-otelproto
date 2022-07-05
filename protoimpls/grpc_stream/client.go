package grpc_stream

import (
	"context"
	"log"
	"sync/atomic"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client experimental.StreamExporterClient
	stream experimental.StreamExporter_ExportTracesClient
	nextId uint64
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = experimental.NewStreamExporterClient(conn)

	// Establish stream to server.
	c.stream, err = c.client.ExportTraces(context.Background())
	if err != nil {
		log.Fatalf("cannot open stream: %v", err)
	}

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
}

func (c *Client) Shutdown() {
}
