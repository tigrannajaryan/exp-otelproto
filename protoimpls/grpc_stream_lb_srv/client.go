package grpc_stream_lb_srv

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client          experimental.StreamExporterClient
	stream          experimental.StreamExporter_ExportTracesClient
	pendingAck      map[uint64]*experimental.TraceExportRequest
	pendingAckMutex sync.Mutex
	nextId          uint64
}

func (c *Client) Connect(server string) error {
	c.pendingAck = make(map[uint64]*experimental.TraceExportRequest)

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
		log.Fatalf("cannot open stream: %v", err)
	}

	go c.readStream(c.stream)

	return nil
}

func (c *Client) readStream(stream experimental.StreamExporter_ExportTracesClient) {
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// read done.
			return
		}
		if err != nil {
			st := status.Convert(err)
			for _, detail := range st.Details() {
				switch t := detail.(type) {
				case *errdetails.RetryInfo:
					if t.RetryDelay.Seconds > 0 {
						// TODO: Need to wait before retrying.
						fmt.Printf("Request was rejected by the server. Retry after %v sec\n",
							t.RetryDelay.Seconds)
					}
				}
			}
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
	request := batch.(*experimental.TraceExportRequest)
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
	var requests []*experimental.TraceExportRequest
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

func (c *Client) Shutdown() {
}
