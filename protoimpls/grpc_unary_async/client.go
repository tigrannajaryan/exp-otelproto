package grpc_unary_async

import (
	"context"
	"log"

	"google.golang.org/grpc"

	otlp "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client otlp.TraceServiceClient
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
	c.client = otlp.NewTraceServiceClient(conn)
	c.sem = make(chan bool, CONCURRENCY)
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	c.sem <- true
	go func() {
		defer func() { <-c.sem }()
		request := batch.(*otlp.ExportTraceServiceRequest)
		_, err := c.client.Export(context.Background(), request)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (c *Client) Shutdown() {
	for i := 0; i < CONCURRENCY; i++ {
		c.sem <- true
	}
}
