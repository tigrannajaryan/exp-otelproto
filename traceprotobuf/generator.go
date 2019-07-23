package traceprotobuf

import (
	"encoding/binary"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a SpanBatch.
type Generator struct {
	tracesSent uint64
	spansSent  uint64
}

func (g *Generator) GenerateBatch() core.SpanBatch {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batch := &SpanBatch{}
	for i := 0; i < 100; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId: generateTraceID(traceID),
			SpanId:  generateSpanID(spanID),
			Name:    &TruncatableString{Value: "load-generator-span"},
			Kind:    Span_CLIENT,
			Attributes: &Span_Attributes{
				AttributeMap: map[string]*AttributeValue{
					"load_generator.span_seq_num":  &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}},
					"load_generator.trace_seq_num": &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}},
				},
			},
			StartTime: timeToTimestamp(startTime),
			EndTime:   timeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		// Append attributes.
		//for k, v := range g.options.Attributes {
		//	span.Attributes[k] = v
		//}

		batch.Spans = append(batch.Spans, span)
	}
	return batch
}

func generateTraceID(id uint64) []byte {
	var traceID [16]byte
	binary.PutUvarint(traceID[:], id)
	return traceID[:]
}

func generateSpanID(id uint64) []byte {
	var spanID [8]byte
	binary.PutUvarint(spanID[:], id)
	return spanID[:]
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}
