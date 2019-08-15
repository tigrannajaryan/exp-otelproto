package traceprotobuf

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// GeneratorB allows to generate a ExportRequest.
type GeneratorB struct {
	random     *rand.Rand
	tracesSent uint64
	spansSent  uint64
}

func NewGeneratorB() *GeneratorB {
	return &GeneratorB{
		random: rand.New(rand.NewSource(99)),
	}
}

func (g *GeneratorB) genRandByteString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = byte(g.random.Intn(128))
	}
	return string(b)
}

func (g *GeneratorB) GenerateBatch(spansPerBatch int, attrsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batch := &ExportRequestB{NodeSpans: []*NodeSpansB{{}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &SpanB{
			TraceId:           generateTraceID(traceID),
			SpanId:            generateSpanID(spanID),
			Name:              "load-GeneratorB-span",
			Kind:              SpanB_CLIENT,
			StartTimeUnixnano: timeToTimestamp(startTime),
			EndTimeUnixnano:   timeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = map[string]*AttributeValue{}

			if attrsPerSpan >= 2 {
				span.Attributes["load_GeneratorB.span_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}}
				span.Attributes["load_GeneratorB.trace_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}}
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

		batch.NodeSpans[0].Spans = append(batch.NodeSpans[0].Spans, span)
	}
	return batch
}
