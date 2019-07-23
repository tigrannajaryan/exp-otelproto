package grpc_protobuf_impl

import (
	"context"
	"log"
	"net"

	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type GrpcServer struct {
	onReceive func(batch core.SpanBatch)
}

func (s *GrpcServer) SendBatch(ctx context.Context, batch *traceprotobuf.SpanBatch) (*traceprotobuf.BatchResponse, error) {
	log.Printf("Received %d spans", len(batch.Spans))
	s.onReceive(batch)
	return &traceprotobuf.BatchResponse{}, nil
}

type Server struct {
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.SpanBatch)) error {
	log.Println("Starting GRPC Server...")

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	traceprotobuf.RegisterTracerServer(s, &GrpcServer{onReceive})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
