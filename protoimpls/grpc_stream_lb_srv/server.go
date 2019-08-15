package grpc_stream_lb_srv

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

type GrpcServer struct {
	srv       *Server
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) Hello(context.Context, *traceprotobuf.HelloRequest) (*traceprotobuf.HelloResponse, error) {
	return &traceprotobuf.HelloResponse{}, nil
}

func (s *GrpcServer) Export(stream traceprotobuf.StreamExporter_ExportServer) error {
	lastStreamOpen := time.Now()
	var requestsProcessed uint = 0
	for {
		// Wait for batch from client.
		batch, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}

		if batch.Id == 0 {
			log.Fatal("Received 0 Id")
		}

		// Process received batch.
		s.onReceive(batch, len(batch.NodeSpans[0].Spans))

		// Send response to client.
		stream.Send(&traceprotobuf.ExportResponse{Id: batch.Id})

		requestsProcessed++

		// Check if time to re-establish the stream.
		if requestsProcessed > s.srv.StreamReopenRequestCount ||
			time.Since(lastStreamOpen) > s.srv.StreamReopenPeriod {
			// Close the stream. The client will reopen it.
			break
		}
	}
	return nil
}

type Server struct {
	s                        *grpc.Server
	StreamReopenPeriod       time.Duration
	StreamReopenRequestCount uint
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest, spanCount int)) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	traceprotobuf.RegisterStreamExporterServer(srv.s, &GrpcServer{srv, onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
