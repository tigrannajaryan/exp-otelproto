package otlp

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

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

func genResource() *Resource {
	return &Resource{
		Labels: []*AttributeKeyValue{
			{Key: "StartTimeUnixnano", Int64Value: 12345678},
			{Key: "Pid", Int64Value: 1234},
			{Key: "HostName", StringValue: "fakehost"},
			{Key: "ServiceName", StringValue: "generator"},
		},
	}
}

func (g *Generator) GenerateBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	batch := &TraceExportRequest{ResourceSpans: []*ResourceSpans{{Resource: genResource()}}}
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

func (g *Generator) GenerateMetricBatch(metricsPerBatch int) core.ExportRequest {
	batch := &MetricExportRequest{ResourceMetrics: []*ResourceMetrics{{Resource: genResource()}}}
	for i := 0; i < metricsPerBatch; i++ {
		startTime := time.Now()

		labelKeys := []string{
			"label1",
			"label2",
		}

		descr := &MetricDescriptor{
			Name:        "metric" + strconv.Itoa(i),
			Description: "some description: " + strconv.Itoa(i),
			Type:        MetricDescriptor_GAUGE_INT64,
			LabelKeys:   labelKeys,
		}

		var timeseries []*TimeSeries
		for j := 0; j < 5; j++ {
			var points []*Point

			pointTs := startTime.Add(time.Duration(time.Millisecond))

			for k := 0; k < 5; k++ {
				point := Point{
					Timestamp: timeToTimestamp(pointTs),
					Value:     &Point_Int64Value{Int64Value: int64(i * j * k)},
				}
				points = append(points, &point)
			}

			ts := TimeSeries{
				LabelValues: []*LabelValue{
					{Value: "val1"},
					{Value: "val2"},
				},
				Points: points,
			}
			timeseries = append(timeseries, &ts)
		}

		metric := &Metric{
			MetricDescriptor: descr,
			Timeseries:       timeseries,
		}

		batch.ResourceMetrics[0].Metrics = append(batch.ResourceMetrics[0].Metrics, metric)
	}
	return batch
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}
