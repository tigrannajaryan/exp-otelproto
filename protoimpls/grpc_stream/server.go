package grpc_stream

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

type GrpcServer struct {
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

		if batch.Id == 0 {
			log.Fatal("Received 0 Id")
		}

		// Process received batch.
		s.onReceive(batch, len(batch.ResourceSpans[0].InstrumentationLibrarySpans[0].Spans))

		// Send response to client.
		stream.Send(&experimental.ExportResponse{Id: batch.Id})
	}
}

type Server struct {
	s *grpc.Server
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest, spanCount int)) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	experimental.RegisterStreamExporterServer(srv.s, &GrpcServer{onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
