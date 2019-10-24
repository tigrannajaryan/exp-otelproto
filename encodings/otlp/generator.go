package otlp

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a ExportRequest.
type Generator struct {
	random     *rand.Rand
	tracesSent uint64
	spansSent  uint64
}

func NewGenerator() *Generator {
	return &Generator{
		random: rand.New(rand.NewSource(99)),
	}
}

func (g *Generator) genRandByteString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = byte(g.random.Intn(10) + 33)
	}
	return string(b)
}

func (g *Generator) GenerateBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	resource := Resource{
		Labels: []*AttributeKeyValue{
			{Key: "StartTimeUnixnano", Int64Value: 12345678},
			{Key: "Pid", Int64Value: 1234},
			{Key: "HostName", StringValue: "fakehost"},
			{Key: "ServiceName", StringValue: "generator"},
		},
	}

	batch := &TraceExportRequest{ResourceSpans: []*ResourceSpans{{Resource: &resource}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			ParentSpanId:      []byte{},
			Name:              "load-generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: core.TimeToTimestamp(startTime),
			EndTimeUnixnano:   core.TimeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*AttributeKeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.span_seq_num", Type: AttributeKeyValue_INT64, Int64Value: int64(spanID)})
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.trace_seq_num", Type: AttributeKeyValue_INT64, Int64Value: int64(traceID)})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: attrName, Type: AttributeKeyValue_STRING, StringValue: g.genRandByteString(g.random.Intn(20) + 1)})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.TimedEvents = append(span.TimedEvents, &Span_TimedEvent{
					TimeUnixnano: core.TimeToTimestamp(startTime),
					Attributes: []*AttributeKeyValue{
						{Key: "te", Type: AttributeKeyValue_INT64, Int64Value: int64(spanID)},
					},
				})
			}
		}

		batch.ResourceSpans[0].Spans = append(batch.ResourceSpans[0].Spans, span)
	}
	return batch
}
