package grpc_stream_lb_srv

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/duration"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

type GrpcServer struct {
	experimental.UnimplementedStreamExporterServer
	srv       *Server
	onReceive func(batch core.ExportRequest, spanCount int)
}

func (s *GrpcServer) ExportTraces(stream experimental.StreamExporter_ExportTracesServer) error {
	lastStreamOpen := time.Now()
	var requestsProcessed uint = 0
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

		requestsProcessed++

		// Check if time to re-establish the stream.
		if requestsProcessed > s.srv.StreamReopenRequestCount ||
			time.Since(lastStreamOpen) > s.srv.StreamReopenPeriod {
			// Close the stream. The client will reopen it.
			break
		}
	}

	st, err := status.New(codes.Unavailable, "Server is unavailable").
		WithDetails(&errdetails.RetryInfo{RetryDelay: &duration.Duration{Seconds: 0}})
	if err != nil {
		log.Fatal(err)
	}

	return st.Err()
}

type Server struct {
	s                        *grpc.Server
	StreamReopenPeriod       time.Duration
	StreamReopenRequestCount uint
}

func (srv *Server) Listen(
	endpoint string, onReceive func(batch core.ExportRequest, spanCount int),
) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.s = grpc.NewServer()
	experimental.RegisterStreamExporterServer(srv.s, &GrpcServer{srv: srv, onReceive: onReceive})
	if err := srv.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (srv *Server) Stop() {
	srv.s.Stop()
}
