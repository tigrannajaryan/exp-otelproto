package grpc_stream_lb_async

import (
	"io"
	"log"
	"net"

	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type GrpcServer struct {
	experimental.UnimplementedStreamExporterServer
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) ExportTraces(stream experimental.StreamExporter_ExportTracesServer) error {
	for {
		// Wait for batch from client.
		batch, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}

		//if batch.Id == 0 {
		//	log.Fatal("Received 0 Id")
		//}

		// Process received batch.
		s.onReceive(batch, len(batch.ResourceSpans[0].ScopeSpans[0].Spans))

		// Send response to client.
		stream.Send(&experimental.ExportResponse{})
	}
}

type Server struct {
	s *grpc.Server
}

func (srv *Server) Listen(
	endpoint string, onReceive func(batch core.ExportRequest, spanCount int),
) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	experimental.RegisterStreamExporterServer(srv.s, &GrpcServer{onReceive: onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
