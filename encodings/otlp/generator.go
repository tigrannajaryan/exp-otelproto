package otlp

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	otlpmetriccol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	otlpcommon "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	otlpmetric "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
	otlpresource "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"

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

func GenResource() *otlpresource.Resource {
	return &otlpresource.Resource{
		Attributes: []*otlpcommon.AttributeKeyValue{
			{Key: "StartTimeUnixnano", IntValue: 12345678},
			{Key: "Pid", IntValue: 1234},
			{Key: "HostName", StringValue: "fakehost"},
			{Key: "ServiceName", StringValue: "generator"},
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	il := &otlptrace.InstrumentationLibrarySpans{}
	batch := &otlptracecol.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource:                    GenResource(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{il},
			},
		},
	}

	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &otlptrace.Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              otlptrace.Span_CLIENT,
			StartTimeUnixNano: core.TimeToTimestamp(startTime),
			EndTimeUnixNano:   core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*otlpcommon.AttributeKeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&otlpcommon.AttributeKeyValue{
						Key:      "load_generator.span_seq_num",
						Type:     otlpcommon.AttributeKeyValue_INT,
						IntValue: int64(spanID),
					})
				span.Attributes = append(span.Attributes,
					&otlpcommon.AttributeKeyValue{
						Key:      "load_generator.trace_seq_num",
						Type:     otlpcommon.AttributeKeyValue_INT,
						IntValue: int64(traceID),
					})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes = append(span.Attributes,
					&otlpcommon.AttributeKeyValue{
						Key:         attrName,
						Type:        otlpcommon.AttributeKeyValue_STRING,
						StringValue: g.genRandByteString(g.random.Intn(20) + 1),
					})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.Events = append(span.Events, &otlptrace.Span_Event{
					TimeUnixNano: core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
					// TimeStartDeltaNano: (time.Duration(i) * time.Millisecond).Nanoseconds(),
					Attributes: []*otlpcommon.AttributeKeyValue{
						{Key: "te", Type: otlpcommon.AttributeKeyValue_INT, IntValue: int64(spanID)},
					},
				})
			}
		}

		il.Spans = append(il.Spans, span)
	}
	return batch
}

func (g *Generator) GenerateLogBatch(logsPerBatch int, attrsPerLog int) core.ExportRequest {
	/*
		traceID := atomic.AddUint64(&g.tracesSent, 1)

		batch := &ExportLogsServiceRequest{ResourceLogs: []*ResourceLogs{{Resource: GenResource()}}}
		for i := 0; i < logsPerBatch; i++ {
			startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

			spanID := atomic.AddUint64(&g.spansSent, 1)

			// Create a log.
			log := &Log{
				TraceId:      core.GenerateTraceID(traceID),
				SpanId:       core.GenerateSpanID(spanID),
				TimeUnixnano: core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
				EventType:    "auto_generated_event",
				Body: &AttributeValue{
					Type:        AttributeValueType_STRING,
					StringValue: fmt.Sprintf("Log message %d of %d, traceid=%q, spanid=%q", i, logsPerBatch, traceID, spanID),
				},
			}

			if attrsPerLog >= 0 {
				// Append attributes.
				log.Attributes = []*AttributeKeyValue{}

				if attrsPerLog >= 2 {
					log.Attributes = append(log.Attributes,
						&AttributeKeyValue{Key: "load_generator.span_seq_num", Type: AttributeKeyValue_INT, IntValue: int64(spanID)})
					log.Attributes = append(log.Attributes,
						&AttributeKeyValue{Key: "load_generator.trace_seq_num", Type: AttributeKeyValue_INT, IntValue: int64(traceID)})
				}

				for j := len(log.Attributes); j < attrsPerLog; j++ {
					attrName := g.genRandByteString(g.random.Intn(50) + 1)
					log.Attributes = append(log.Attributes,
						&AttributeKeyValue{Key: attrName, Type: AttributeKeyValue_STRING, StringValue: g.genRandByteString(g.random.Intn(20) + 1)})
				}
			}

			batch.ResourceLogs[0].Logs = append(batch.ResourceLogs[0].Logs, log)
		}
		return batch
	*/
	return nil
}

func GenInt64Timeseries(startTime time.Time, offset int, valuesPerTimeseries int) []*otlpmetric.Int64DataPoint {
	var timeseries []*otlpmetric.Int64DataPoint
	for j := 0; j < 5; j++ {
		var points []*otlpmetric.Int64DataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))

			point := otlpmetric.Int64DataPoint{
				TimeUnixNano: pointTs,
				Value:        int64(offset * j * k),
			}

			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}

			points = append(points, &point)
		}

		timeseries = append(timeseries, points...)
	}

	return timeseries
}

func genInt64Gauge(startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int) *otlpmetric.Metric {
	descr := GenMetricDescriptor(i)

	metric1 := &otlpmetric.Metric{
		MetricDescriptor: descr,
		Int64DataPoints:  GenInt64Timeseries(startTime, i, valuesPerTimeseries),
	}

	return metric1
}

func GenMetricDescriptor(i int) *otlpmetric.MetricDescriptor {
	descr := &otlpmetric.MetricDescriptor{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Type:        otlpmetric.MetricDescriptor_GAUGE_INT64,
		Labels: []*otlpcommon.StringKeyValue{
			{
				Key:   "label1",
				Value: "val1",
			},
			{
				Key:   "label2",
				Value: "val2",
			},
		},
	}
	return descr
}

func genHistogram(startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int) *otlpmetric.Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i)
	descr.Type = otlpmetric.MetricDescriptor_GAUGE_HISTOGRAM

	var timeseries2 []*otlpmetric.HistogramDataPoint
	for j := 0; j < 1; j++ {
		var points []*otlpmetric.HistogramDataPoint

		//prevPointTs := int64(0)
		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			//diffTs := pointTs - prevPointTs
			//prevPointTs = pointTs
			val := float64(i * j * k)
			point := otlpmetric.HistogramDataPoint{
				TimeUnixNano: pointTs,
				Count:        1,
				Sum:          val,
				Buckets: []*otlpmetric.HistogramDataPoint_Bucket{
					{
						Count: 12,
						Exemplar: &otlpmetric.HistogramDataPoint_Bucket_Exemplar{
							Value:        val,
							TimeUnixNano: pointTs,
						},
					},
					{
						Count: 345,
					},
				},
				ExplicitBounds: []float64{0, 1000000},
			}
			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}
			points = append(points, &point)
		}

		timeseries2 = append(timeseries2, points...)
	}

	metric2 := &otlpmetric.Metric{
		MetricDescriptor:    descr,
		HistogramDataPoints: timeseries2,
	}

	return metric2
}

func genSummary(startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int) *otlpmetric.Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i)
	descr.Type = otlpmetric.MetricDescriptor_SUMMARY

	var timeseries2 []*otlpmetric.SummaryDataPoint
	for j := 0; j < 1; j++ {
		var points []*otlpmetric.SummaryDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			val := float64(i * j * k)
			point := otlpmetric.SummaryDataPoint{
				TimeUnixNano: pointTs,
				Count:        1,
				Sum:          val,
				PercentileValues: []*otlpmetric.SummaryDataPoint_ValueAtPercentile{
					{
						Percentile: 99,
						Value:      val / 10,
					},
				},
			}
			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}
			points = append(points, &point)
		}

		timeseries2 = append(timeseries2, points...)
	}

	metric2 := &otlpmetric.Metric{
		MetricDescriptor:  descr,
		SummaryDataPoints: timeseries2,
	}

	return metric2
}

func (g *Generator) GenerateMetricBatch(
	metricsPerBatch int,
	valuesPerTimeseries int,
	int64 bool,
	histogram bool,
	summary bool,
) core.ExportRequest {

	il := &otlpmetric.InstrumentationLibraryMetrics{}
	batch := &otlpmetriccol.ExportMetricsServiceRequest{
		ResourceMetrics: []*otlpmetric.ResourceMetrics{
			{
				Resource:                      GenResource(),
				InstrumentationLibraryMetrics: []*otlpmetric.InstrumentationLibraryMetrics{il},
			},
		},
	}

	for i := 0; i < metricsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		labelKeys := []string{
			"label1",
			"label2",
		}

		if int64 {
			il.Metrics = append(il.Metrics, genInt64Gauge(startTime, i, labelKeys, valuesPerTimeseries))
		}
		if histogram {
			il.Metrics = append(il.Metrics, genHistogram(startTime, i, labelKeys, valuesPerTimeseries))
		}
		if summary {
			il.Metrics = append(il.Metrics, genSummary(startTime, i, labelKeys, valuesPerTimeseries))
		}
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
