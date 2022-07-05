package otlp

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	otlplogcol "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/logs/v1"
	otlpmetriccol "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/metrics/v1"
	otlptracecol "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
	otlpcommon "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/common/v1"
	otlplogs "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/logs/v1"
	otlpmetric "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/metrics/v1"
	otlpresource "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/resource/v1"
	otlptrace "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/trace/v1"

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
		Attributes: []*otlpcommon.KeyValue{
			{
				Key:   "StartTimeUnixnano",
				Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: 12345678}},
			},
			{
				Key:   "Pid",
				Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: 1234}},
			},
			{
				Key:   "HostName",
				Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "fakehost"}},
			},
			{
				Key:   "ServiceName",
				Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "generator"}},
			},
		},
	}
}

func (g *Generator) GenerateSpanBatch(
	spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int,
) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	il := &otlptrace.ScopeSpans{
		Scope: &otlpcommon.InstrumentationScope{Name: "io.opentelemetry"},
	}
	batch := &otlptracecol.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource:   GenResource(),
				ScopeSpans: []*otlptrace.ScopeSpans{il},
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
			Kind:              otlptrace.Span_SPAN_KIND_CLIENT,
			StartTimeUnixNano: core.TimeToTimestamp(startTime),
			EndTimeUnixNano:   core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*otlpcommon.KeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(
					span.Attributes,
					&otlpcommon.KeyValue{
						Key:   "load_generator.span_seq_num",
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: int64(spanID)}},
					},
				)
				span.Attributes = append(
					span.Attributes,
					&otlpcommon.KeyValue{
						Key:   "load_generator.trace_seq_num",
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: int64(traceID)}},
					},
				)
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := core.GenRandAttrName(g.random)
				span.Attributes = append(
					span.Attributes,
					&otlpcommon.KeyValue{
						Key:   attrName,
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: g.genRandByteString(g.random.Intn(20) + 1)}},
					},
				)
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.Events = append(
					span.Events, &otlptrace.Span_Event{
						TimeUnixNano: core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
						// TimeStartDeltaNano: (time.Duration(i) * time.Millisecond).Nanoseconds(),
						Attributes: []*otlpcommon.KeyValue{
							{
								Key:   "te",
								Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: int64(spanID)}},
							},
						},
					},
				)
			}
		}

		il.Spans = append(il.Spans, span)
	}
	return batch
}

func (g *Generator) GenerateLogBatch(logsPerBatch int, attrsPerLog int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	il := &otlplogs.ScopeLogs{
		Scope: &otlpcommon.InstrumentationScope{Name: "io.opentelemetry"},
	}
	batch := &otlplogcol.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				Resource:  GenResource(),
				ScopeLogs: []*otlplogs.ScopeLogs{il},
			},
		},
	}

	logs := []*otlplogs.LogRecord{}
	for i := 0; i < logsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a log.
		log := &otlplogs.LogRecord{
			TimeUnixNano:   core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
			TraceId:        core.GenerateTraceID(traceID),
			SpanId:         core.GenerateSpanID(spanID),
			SeverityNumber: otlplogs.SeverityNumber_SEVERITY_NUMBER_INFO,
			SeverityText:   "info",
			Body:           &otlpcommon.AnyValue{
				//Type: ValueType_KVLIST,
				//ListValues: &ValueList{
				//	ListValues: []*KeyValue{
				//		{
				//			Key:         "bodykey",
				//			Type:        ValueType_STRING,
				//			StringValue: fmt.Sprintf("Log message %d of %d, traceid=%q, spanid=%q", i, logsPerBatch, traceID, spanID),
				//		},
				//	},
				//},
			},
		}

		if attrsPerLog >= 0 {
			// Append attributes.
			log.Attributes = []*otlpcommon.KeyValue{}

			if attrsPerLog >= 2 {
				log.Attributes = append(
					log.Attributes,
					&otlpcommon.KeyValue{
						Key:   "load_generator.span_seq_num",
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: int64(spanID)}},
					},
				)
				log.Attributes = append(
					log.Attributes,
					&otlpcommon.KeyValue{
						Key:   "load_generator.trace_seq_num",
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: int64(traceID)}},
					},
				)
			}

			for j := len(log.Attributes); j < attrsPerLog; j++ {
				attrName := core.GenRandAttrName(g.random)
				log.Attributes = append(
					log.Attributes,
					&otlpcommon.KeyValue{
						Key:   attrName,
						Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: g.genRandByteString(g.random.Intn(20) + 1)}},
					},
				)
			}

			log.Attributes = append(
				log.Attributes,
				&otlpcommon.KeyValue{
					Key:   "event_type",
					Value: &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "auto_generated_event"}},
				},
			)

		}

		logs = append(logs, log)
	}

	il.LogRecords = logs

	return batch
}

func GenInt64Timeseries(
	startTime time.Time, offset int, valuesPerTimeseries int,
) []*otlpmetric.NumberDataPoint {
	var timeseries []*otlpmetric.NumberDataPoint
	for j := 0; j < 5; j++ {
		var points []*otlpmetric.NumberDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))

			point := otlpmetric.NumberDataPoint{
				TimeUnixNano: pointTs,
				Value: &otlpmetric.NumberDataPoint_AsInt{
					AsInt: int64(offset * j * k),
				},
				Attributes: []*otlpcommon.KeyValue{
					{
						Key: "label1",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{
								StringValue: "val1",
							},
						},
					},
					{
						Key: "label2",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{"val2"},
						},
					},
				},
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

func genInt64Gauge(startTime time.Time, i int, valuesPerTimeseries int) *otlpmetric.Metric {
	metric1 := GenMetricDescriptor(i)

	metric1.Data = &otlpmetric.Metric_Gauge{
		Gauge: &otlpmetric.Gauge{
			DataPoints: GenInt64Timeseries(startTime, i, valuesPerTimeseries),
		},
	}

	return metric1
}

func GenMetricDescriptor(i int) *otlpmetric.Metric {
	descr := &otlpmetric.Metric{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
	}
	return descr
}

func genHistogram(startTime time.Time, i int, valuesPerTimeseries int) *otlpmetric.Metric {
	// Add Histogram
	metric2 := GenMetricDescriptor(i)

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
				Sum:          &val,
				BucketCounts: []uint64{12, 345},
				Exemplars: []*otlpmetric.Exemplar{

					&otlpmetric.Exemplar{
						Value: &otlpmetric.Exemplar_AsDouble{
							AsDouble: val,
						},
						TimeUnixNano: pointTs,
					},
				},

				ExplicitBounds: []float64{0, 1000000},

				Attributes: []*otlpcommon.KeyValue{
					{
						Key: "label1",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{"val1"},
						},
					},
					{
						Key: "label2",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{"val2"},
						},
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

	metric2.Data = &otlpmetric.Metric_Histogram{
		Histogram: &otlpmetric.Histogram{
			DataPoints: timeseries2,
		},
	}

	return metric2
}

func genSummary(startTime time.Time, i int, valuesPerTimeseries int) *otlpmetric.Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i)

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

				QuantileValues: []*otlpmetric.SummaryDataPoint_ValueAtQuantile{
					{
						Quantile: 99,
						Value:    val / 10,
					},
				},

				Attributes: []*otlpcommon.KeyValue{
					{
						Key: "label1",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{"val1"},
						},
					},
					{
						Key: "label2",
						Value: &otlpcommon.AnyValue{
							Value: &otlpcommon.AnyValue_StringValue{"val2"},
						},
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

	descr.Data = &otlpmetric.Metric_Summary{
		Summary: &otlpmetric.Summary{
			DataPoints: timeseries2,
		},
	}

	return descr
}

func (g *Generator) GenerateMetricBatch(
	metricsPerBatch int,
	valuesPerTimeseries int,
	int64 bool,
	histogram bool,
	summary bool,
) core.ExportRequest {

	il := &otlpmetric.ScopeMetrics{}
	batch := &otlpmetriccol.ExportMetricsServiceRequest{
		ResourceMetrics: []*otlpmetric.ResourceMetrics{
			{
				Resource:     GenResource(),
				ScopeMetrics: []*otlpmetric.ScopeMetrics{il},
			},
		},
	}

	for i := 0; i < metricsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		if int64 {
			il.Metrics = append(il.Metrics, genInt64Gauge(startTime, i, valuesPerTimeseries))
		}
		if histogram {
			il.Metrics = append(il.Metrics, genHistogram(startTime, i, valuesPerTimeseries))
		}
		if summary {
			il.Metrics = append(il.Metrics, genSummary(startTime, i, valuesPerTimeseries))
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

type SpanTranslator struct {
}

func (st *SpanTranslator) TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) core.ExportRequest {
	return batch
}
