package traceprotobuf

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

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	node := Node{
		Identifier: &ProcessIdentifier{
			StartTimeUnixnano: 12345678,
			Pid:               1234,
			HostName:          "fakehost",
		},
		ServiceInfo: &ServiceInfo{
			Name: "generator",
		},
	}

	batch := &ExportRequest{NodeSpans: []*NodeSpans{{Node: &node}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: core.TimeToTimestamp(startTime),
			EndTimeUnixnano:   core.TimeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = map[string]*AttributeValue{}

			if attrsPerSpan >= 2 {
				span.Attributes["load_generator.span_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}}
				span.Attributes["load_generator.trace_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}}
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes[attrName] = &AttributeValue{
					Value: &AttributeValue_StringValue{
						StringValue: g.genRandByteString(g.random.Intn(20) + 1),
					},
				}
			}

		}

		if timedEventsPerSpan > 0 {
			span.TimeEvents = &Span_TimeEvents{}
			for i := 0; i < timedEventsPerSpan; i++ {
				span.TimeEvents.TimeEvent = append(span.TimeEvents.TimeEvent, &Span_TimeEvent{
					TimeUnixnano: core.TimeToTimestamp(startTime),
					Value: &Span_TimeEvent_Annotation_{
						Annotation: &Span_TimeEvent_Annotation{
							Attributes: map[string]*AttributeValue{
								"te": {Value: &AttributeValue_IntValue{IntValue: int64(spanID)}},
							},
						},
					},
				})
			}
		}

		batch.NodeSpans[0].Spans = append(batch.NodeSpans[0].Spans, span)
	}
	return batch
}

func (g *Generator) GenerateMetricBatch(metricsPerBatch int) core.ExportRequest {
	return nil
}
