package traceprotobufb

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
		b[i] = byte(g.random.Intn(128))
	}
	return string(b)
}

func (g *Generator) GenerateBatch(spansPerBatch int, attrsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	resource := Resource{
		Identifier: &ProcessIdentifier{
			StartTimeUnixnano: 12345678,
			Pid:               1234,
			HostName:          "fakehost",
		},
	}

	batch := &ExportRequest{SpanBatch: []*SpanBatch{{Resource: &resource}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateTraceID(spanID),
			Name:              "load-Generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: core.TimeToTimestamp(startTime),
			EndTimeUnixnano:   core.TimeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = map[string]*AttributeValue{}

			if attrsPerSpan >= 2 {
				span.Attributes["load_Generator.span_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}}
				span.Attributes["load_Generator.trace_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}}
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes[attrName] = &AttributeValue{
					Value: &AttributeValue_StringValue{
						StringValue: g.genRandByteString(g.random.Intn(100) + 1),
					},
				}
			}

		}

		batch.SpanBatch[0].Spans = append(batch.SpanBatch[0].Spans, span)
	}
	return batch
}
