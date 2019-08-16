package grpc_flatbuffers_impl

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceflatbuffers"
)

type GrpcServer struct {
	onReceive func(batch core.ExportRequest)
}

func (s *GrpcServer) SendBatch(ctx context.Context, batch *traceflatbuffers.BatchRequest) (*traceflatbuffers.BatchResponse, error) {
	//spanBatch := traceflatbuffers.GetRootAsSpanBatch(batch.EncodedSpans, 0)
	//
	//log.Printf("Received %d spans", spanBatch.SpansLength())
	//
	//for i := 0; i < spanBatch.SpansLength(); i++ {
	//	var span traceflatbuffers.Span
	//	if !spanBatch.Spans(&span, i) {
	//		log.Fatalf("Cannot decode span %d", i)
	//	}
	//	startTime := time.Unix(0, span.StartTime())
	//	endTime := time.Unix(0, span.EndTime())
	//	log.Printf("Span name=%s, traceid=%v, spanid=%v, start time=%s, end time=%s",
	//		span.Name(), span.TraceIdLo(), span.SpanId(), startTime.String(), endTime.String())
	//
	//	for j := 0; j < span.AttributesLength(); j++ {
	//		var attribute traceflatbuffers.Attribute
	//		if !span.Attributes(&attribute, j) {
	//			log.Fatalf("Cannot decode attributes of span %d", i)
	//		}
	//
	//		key := string(attribute.Key())
	//		log.Printf("Attribute %d, key=%v", j, key)
	//
	//		unionTable := new(flatbuffers.Table)
	//		if attribute.Value(unionTable) {
	//			unionType := attribute.ValueType()
	//			if unionType == traceflatbuffers.AttributeValueInt64Value {
	//				attrVal := new(traceflatbuffers.Int64Value)
	//				attrVal.Init(unionTable.Bytes, unionTable.Pos)
	//				log.Printf("Value of key=%d", attrVal.Int64Value())
	//			}
	//		}
	//
	//	}
	//}
	//
	//s.onReceive(spanBatch)
	//return &traceflatbuffers.BatchResponse{}, nil
	return nil, nil
}

type Server struct {
}

func (srv *Server) Listen(endpoint string, onReceive func(batch core.ExportRequest)) error {
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
