package grpc_unary_async

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
	sem    chan bool
}

const CONCURRENCY = 10

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = otlp.NewUnaryExporterClient(conn)
	c.sem = make(chan bool, CONCURRENCY)
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	c.sem <- true
	go func() {
		defer func() { <-c.sem }()
		request := batch.(*otlp.TraceExportRequest)
		request.Id = atomic.AddUint64(&c.nextId, 1)
		response, err := c.client.ExportTraces(context.Background(), request)
		if err != nil {
			log.Fatal(err)
		}
		if response.Id != request.Id {
			log.Fatal("ack id mismatch")
		}
	}()
}

func (c *Client) Shutdown() {
	for i := 0; i < CONCURRENCY; i++ {
		c.sem <- true
	}
}
