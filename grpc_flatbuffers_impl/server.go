package grpc_flatbuffers_impl

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/traceflatbuffers"
)

type GrpcServer struct {
	onReceive func(batch core.SpanBatch)
}

func (s *GrpcServer) SendBatch(ctx context.Context, batch *traceflatbuffers.BatchRequest) (*traceflatbuffers.BatchResponse, error) {
	spanBatch := traceflatbuffers.GetRootAsSpanBatch(batch.EncodedSpans, 0)

	log.Printf("Received %d spans", spanBatch.SpansLength())
	s.onReceive(spanBatch)
	return &traceflatbuffers.BatchResponse{}, nil
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
	traceflatbuffers.RegisterTracerServer(s, &GrpcServer{onReceive})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
