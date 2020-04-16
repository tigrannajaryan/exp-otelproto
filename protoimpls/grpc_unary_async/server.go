package grpc_unary_async

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	otlp "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	"github.com/tigrannajaryan/exp-otelproto/core"
)

type GrpcServer struct {
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) Export(ctx context.Context, batch *otlp.ExportTraceServiceRequest) (*otlp.ExportTraceServiceResponse, error) {
	s.onReceive(batch, len(batch.ResourceSpans[0].InstrumentationLibrarySpans[0].Spans))
	return &otlp.ExportTraceServiceResponse{}, nil
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
	otlp.RegisterTraceServiceServer(srv.s, &GrpcServer{onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
