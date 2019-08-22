package grpc_unary_async

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

type GrpcServer struct {
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) Hello(context.Context, *traceprotobuf.HelloRequest) (*traceprotobuf.HelloResponse, error) {
	return &traceprotobuf.HelloResponse{}, nil
}

func (s *GrpcServer) Export(ctx context.Context, batch *traceprotobuf.ExportRequest) (*traceprotobuf.ExportResponse, error) {
	if batch.Id == 0 {
		log.Fatal("Received 0 Id")
	}

	s.onReceive(batch, len(batch.NodeSpans[0].Spans))
	return &traceprotobuf.ExportResponse{Id: batch.Id}, nil
}

type Server struct {
	s *grpc.Server
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest, spanCount int)) error {
	// log.Println("Starting GRPC Server...")

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	traceprotobuf.RegisterUnaryExporterServer(srv.s, &GrpcServer{onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
