package grpcimpl

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/tracerprotobuf"
)

type GrpcServer struct {
	onReceive func(batch core.SpanBatch)
}

func (s *GrpcServer) SendBatch(ctx context.Context, batch *tracerprotobuf.SpanBatch) (*tracerprotobuf.BatchResponse, error) {
	s.onReceive(batch)
	return &tracerprotobuf.BatchResponse{}, nil
}

type Server struct {
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.SpanBatch)) {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tracerprotobuf.RegisterTracerServer(s, &GrpcServer{onReceive})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
