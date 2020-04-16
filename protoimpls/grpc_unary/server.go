package grpc_unary

import (
	"context"
	"log"
	"net"

	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type GrpcServer struct {
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) Export(ctx context.Context, batch *otlptracecol.ExportTraceServiceRequest) (*otlptracecol.ExportTraceServiceResponse, error) {
	//if batch.Id == 0 {
	//	log.Fatal("Received 0 Id")
	//}

	s.onReceive(batch, len(batch.ResourceSpans[0].InstrumentationLibrarySpans[0].Spans))
	return &otlptracecol.ExportTraceServiceResponse{}, nil
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
	otlptracecol.RegisterTraceServiceServer(srv.s, &GrpcServer{onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
