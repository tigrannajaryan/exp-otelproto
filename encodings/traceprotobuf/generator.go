package traceprotobuf

import (
	"encoding/binary"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a SpanBatch.
type Generator struct {
	tracesSent uint64
	spansSent  uint64
}

func genRandByteString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = byte(rand.Intn(128))
	}
	return string(b)
}

func (g *Generator) GenerateBatch(spansPerBatch int, attrsPerSpan int) core.SpanBatch {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batch := &SpanBatch{}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           generateTraceID(traceID),
			SpanId:            generateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: timeToTimestamp(startTime),
			EndTimeUnixnano:   timeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = map[string]*AttributeValue{}

			if attrsPerSpan >= 2 {
				span.Attributes["load_generator.span_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}}
				span.Attributes["load_generator.trace_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}}
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := genRandByteString(rand.Intn(50) + 1)
				span.Attributes[attrName] = &AttributeValue{
					Value: &AttributeValue_StringValue{
						StringValue: genRandByteString(rand.Intn(100) + 1),
					},
				}
			}

		}

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

func timeToTimestamp(t time.Time) int64 {
	return t.UnixNano()
}
