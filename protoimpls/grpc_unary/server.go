package grpc_unary

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

type GrpcServer struct {
	onReceive func(batch core.SpanBatch)
}

func (s *GrpcServer) SendBatch(ctx context.Context, batch *traceprotobuf.SpanBatch) (*traceprotobuf.BatchResponse, error) {
	// log.Printf("Received %d spans", len(batch.Spans))
	s.onReceive(batch)
	return &traceprotobuf.BatchResponse{Id: batch.Id}, nil
}

type Server struct {
	s *grpc.Server
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.SpanBatch)) error {
	// log.Println("Starting GRPC Server...")

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	traceprotobuf.RegisterUnaryTracerServer(srv.s, &GrpcServer{onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
