package sapmenc

import (
	"math/rand"
	"sync/atomic"
	"time"

	jaegerpb "github.com/jaegertracing/jaeger/model"

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

func GenResource() *jaegerpb.Process {
	return &jaegerpb.Process{
		Tags: []jaegerpb.KeyValue{
			{Key: "StartTimeUnixnano", VInt64: 12345678, VType: jaegerpb.ValueType_INT64},
			{Key: "Pid", VInt64: 1234, VType: jaegerpb.ValueType_INT64},
			{Key: "HostName", VStr: "fakehost", VType: jaegerpb.ValueType_STRING},
			{Key: "ServiceName", VStr: "generator", VType: jaegerpb.ValueType_STRING},
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	batch := &jaegerpb.Batch{Process: GenResource()}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &jaegerpb.Span{
			TraceID:       jaegerpb.TraceID{core.GenerateTraceIDLow(traceID), core.GenerateTraceIDHigh(traceID)},
			SpanID:        jaegerpb.SpanID(spanID),
			OperationName: "load-generator-span",
			StartTime:     startTime,
			Duration:      time.Duration(i) * time.Millisecond,
		}

		span.Tags = []jaegerpb.KeyValue{
			{Key: "span.kind", VType: jaegerpb.ValueType_STRING, VStr: "client"},
		}

		if attrsPerSpan >= 0 {
			// Append attributes.

			if attrsPerSpan >= 2 {
				span.Tags = append(span.Tags,
					jaegerpb.KeyValue{Key: "load_generator.span_seq_num", VType: jaegerpb.ValueType_INT64, VInt64: int64(spanID)})
				span.Tags = append(span.Tags,
					jaegerpb.KeyValue{Key: "load_generator.trace_seq_num", VType: jaegerpb.ValueType_INT64, VInt64: int64(traceID)})
			}

			for j := 2; j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(20) + 1)
				span.Tags = append(span.Tags,
					jaegerpb.KeyValue{Key: attrName, VType: jaegerpb.ValueType_STRING,
						VStr: g.genRandByteString(g.random.Intn(20) + 1)})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.Logs = append(span.Logs, jaegerpb.Log{
					Timestamp: startTime.Add(time.Duration(i) * time.Millisecond),
					Fields: []jaegerpb.KeyValue{
						{Key: "te", VType: jaegerpb.ValueType_INT64, VInt64: int64(spanID)},
					},
				})
			}
		}

		batch.Spans = append(batch.Spans, span)
	}
	return batch
}
func (g *Generator) GenerateMetricBatch(
	metricsPerBatch int,
	valuesPerTimeseries int,
	int64 bool,
	histogram bool,
	summary bool,
) core.ExportRequest {

	return nil
}
