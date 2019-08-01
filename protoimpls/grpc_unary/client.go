package grpc_unary

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client traceprotobuf.UnaryTracerClient
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewUnaryTracerClient(conn)
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	c.client.Export(context.Background(), batch.(*traceprotobuf.ExportRequest))
}
