package encodings

import (
	"log"
	"runtime"
	"testing"

	"github.com/golang/protobuf/proto"
	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
	v1 "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
)

/*
func TestPreparedTrace(t *testing.T) {

	g := otlp.NewGenerator()
	request := g.GenerateSpanBatch(100, 5, 0).(*otlp.TraceExportRequest)

	requestBytes, err := proto.Marshal(request)
	if err != nil {
		t.Fatal()
	}

	var preparedRequest otlp.TraceExportRequestPrepared
	if err := proto.Unmarshal(requestBytes, &preparedRequest); err != nil {
		t.Fatal()
	}

	for i, spans := range preparedRequest.ResourceSpans {

		var resource otlp.Resource
		if err := proto.Unmarshal(spans.Resource, &resource); err != nil {
			t.Fatal()
		}

		if !proto.Equal(&resource, request.ResourceSpans[i].Resource) {
			t.Fatal()
		}

		for j, span := range spans.Spans {
			for k, attrBytes := range span.Attributes {
				var attr otlp.AttributeKeyValue
				if err := proto.Unmarshal(attrBytes, &attr); err != nil {
					t.Fatal()
				}

				if !proto.Equal(&attr, request.ResourceSpans[i].Spans[j].Attributes[k]) {
					t.Fatal()
				}
			}
		}
	}

	preparedRequestBytes, err := proto.Marshal(&preparedRequest)
	if err != nil {
		t.Fatal()
	}

	if c := bytes.Compare(requestBytes, preparedRequestBytes); c != 0 {
		t.Fatal()
	}
}
*/

func encodeUnpreparedTraces(spanCount int) proto.Message {
	g := experimental.NewGenerator()
	request := g.GenerateSpanBatch(spanCount, 5, 0).(*v1.ExportTraceServiceRequest)
	return request
}

func BenchmarkEncodeTraces(b *testing.B) {
	tests := []struct {
		name    string
		encoder func(spanCount int) proto.Message
	}{
		{
			name:    "Unprepared",
			encoder: encodeUnpreparedTraces,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batch := test.encoder(100)
			runtime.GC()
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				var err error
				bytes, err := proto.Marshal(batch)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}

func BenchmarkDecodeEncodeTraces(b *testing.B) {
	tests := []struct {
		name            string
		emptyMsgCreator func() proto.Message
	}{
		{
			name:            "Full",
			emptyMsgCreator: func() proto.Message { return &v1.TraceExportRequest{} },
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batch := encodeUnpreparedTraces(100)
			runtime.GC()

			var err error
			bytes, err := proto.Marshal(batch)
			if err != nil || len(bytes) == 0 {
				log.Fatal("Cannot encode batch")
			}

			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				msg := test.emptyMsgCreator()
				err = proto.Unmarshal(bytes, msg)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot decode batch")
				}

				bytes, err := proto.Marshal(msg)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}
