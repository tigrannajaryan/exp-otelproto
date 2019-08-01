package grpc_oc

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
)

type GrpcServer struct {
	onReceive func(batch core.ExportRequest)
	SendAck   bool
}

func (s *GrpcServer) Export(stream octraceprotobuf.OCStreamTracer_ExportServer) error {
	for {
		// Wait for batch from client.
		batch, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}

		// Process received batch.
		s.onReceive(batch)

		if s.SendAck {
			// Send response to client.
			stream.Send(&octraceprotobuf.ExportResponse{})
		}
	}
}

type Server struct {
	s       *grpc.Server
	SendAck bool
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest)) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	octraceprotobuf.RegisterOCStreamTracerServer(srv.s,
		&GrpcServer{onReceive, srv.SendAck})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
