package grpc_unary

import (
	"context"
	"log"
	"sync/atomic"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client otlp.UnaryExporterClient
	nextId uint64
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = otlp.NewUnaryExporterClient(conn)
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	request := batch.(*otlp.TraceExportRequest)
	request.Id = atomic.AddUint64(&c.nextId, 1)
	c.client.Export(context.Background(), request)
}

func (c *Client) Shutdown() {
}
