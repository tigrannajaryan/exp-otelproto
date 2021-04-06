package baseline

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

func GenResource() *Resource {
	return &Resource{
		Attributes: []*KeyValue{
			{Key: "StartTimeUnixnano", Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: 12345678}}},
			{Key: "Pid", Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: 1234}}},
			{Key: "HostName", Value: &AnyValue{Value: &AnyValue_StringValue{StringValue: "fakehost"}}},
			{Key: "ServiceName", Value: &AnyValue{Value: &AnyValue_StringValue{StringValue: "generator"}}},
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batchStartTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

	il := &InstrumentationLibrarySpans{
		InstrumentationLibrary: &InstrumentationLibrary{Name: "io.opentelemetry"},
	}
	batch := &TraceExportRequest{
		ResourceSpans: []*ResourceSpans{
			{
				Resource:                    GenResource(),
				InstrumentationLibrarySpans: []*InstrumentationLibrarySpans{il},
			},
		},
	}

	for i := 0; i < spansPerBatch; i++ {
		startTime := batchStartTime.Add(time.Duration(i) * time.Millisecond)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              Span_SPAN_KIND_CLIENT,
			StartTimeUnixNano: core.TimeToTimestamp(startTime),
			EndTimeUnixNano:   core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*KeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&KeyValue{
					Key: "load_generator.span_seq_num",
					Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}},
					})
				span.Attributes = append(span.Attributes,
					&KeyValue{
					Key: "load_generator.trace_seq_num",
						Value: &AnyValue{Value: &AnyValue_IntValue{IntValue:  int64(traceID)}},
					})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := core.GenRandAttrName(g.random)
				span.Attributes = append(span.Attributes,
					&KeyValue{
						Key: attrName,
						Value: &AnyValue{Value: &AnyValue_StringValue{StringValue: g.genRandByteString(g.random.Intn(20) + 1)}},
					})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.Events = append(span.Events, &Span_Event{
					TimeUnixNano: core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
					// TimeStartDeltaNano: (time.Duration(i) * time.Millisecond).Nanoseconds(),
					Attributes: []*KeyValue{
						{Key: "te", Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}}},
					},
				})
			}
		}

		il.Spans = append(il.Spans, span)
	}
	return batch
}

func (g *Generator) GenerateLogBatch(logsPerBatch int, attrsPerLog int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	il := &InstrumentationLibraryLogs{
		InstrumentationLibrary: &InstrumentationLibrary{Name: "io.opentelemetry"},
	}
	batch := &ExportLogsServiceRequest{
		ResourceLogs: []*ResourceLogs{{
			Resource: GenResource(),
			InstrumentationLibraryLogs: []*InstrumentationLibraryLogs{il},
		}},
	}

	logs := []*LogRecord{}
	for i := 0; i < logsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a log.
		log := &LogRecord{
			TimeUnixNano:   core.TimeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
			TraceId:        core.GenerateTraceID(traceID),
			SpanId:         core.GenerateSpanID(spanID),
			SeverityNumber: SeverityNumber_SEVERITY_NUMBER_INFO,
			SeverityText:   "info",
			Name:      "ProcessStarted",
			Body:           &AnyValue{
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
			log.Attributes = []*KeyValue{}

			if attrsPerLog >= 2 {
				log.Attributes = append(log.Attributes,
					&KeyValue{Key: "load_generator.span_seq_num", Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}}})
				log.Attributes = append(log.Attributes,
					&KeyValue{Key: "load_generator.trace_seq_num", Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(traceID)}}})
			}

			for j := len(log.Attributes); j < attrsPerLog; j++ {
				attrName := core.GenRandAttrName(g.random)
				log.Attributes = append(log.Attributes,
					&KeyValue{Key: attrName, Value: &AnyValue{Value: &AnyValue_StringValue{StringValue: g.genRandByteString(g.random.Intn(20) + 1)}}})
			}

			log.Attributes = append(log.Attributes,
				&KeyValue{Key: "event_type", Value: &AnyValue{Value: &AnyValue_StringValue{StringValue: "auto_generated_event"}}})

		}

		logs = append(logs, log)
	}

	il.Logs = logs

	return batch
}

func GenInt64Timeseries(startTime time.Time, offset int, valuesPerTimeseries int) *Metric_IntGauge {
	var timeseries []*IntDataPoint
	for j := 0; j < 5; j++ {
		var points []*IntDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))

			point := IntDataPoint{
				TimeUnixNano: pointTs,
				Value:        int64(offset * j * k),
				Labels: []*StringKeyValue{
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

			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}

			points = append(points, &point)
		}

		timeseries = append(timeseries, points...)
	}

	return &Metric_IntGauge{IntGauge: &IntGauge{DataPoints:timeseries}}
}

func genInt64Gauge(startTime time.Time, i int, valuesPerTimeseries int) *Metric {
	descr := GenMetricDescriptor(i)
	descr.Data = GenInt64Timeseries(startTime, i, valuesPerTimeseries)
	return descr
}

func GenMetricDescriptor(i int) *Metric {
	descr := &Metric{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
	}
	return descr
}

func genHistogram(startTime time.Time, i int, valuesPerTimeseries int) *Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i)

	var timeseries2 []*HistogramDataPoint
	for j := 0; j < 1; j++ {
		var points []*HistogramDataPoint

		//prevPointTs := int64(0)
		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			//diffTs := pointTs - prevPointTs
			//prevPointTs = pointTs
			val := float64(i * j * k)
			point := HistogramDataPoint{
				TimeUnixNano: pointTs,
				Count:        1,
				Sum:          val,
				BucketCounts: []uint64{12,345},
				Exemplars: []*Exemplar{
						{
							Value:        &Exemplar_AsDouble{AsDouble: val},
							TimeUnixNano: pointTs,
						},
				},
				ExplicitBounds: []float64{0, 1000000},
				Labels: []*StringKeyValue{
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
			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}
			points = append(points, &point)
		}

		timeseries2 = append(timeseries2, points...)
	}

	descr.Data = &Metric_Histogram{Histogram:&Histogram{DataPoints:timeseries2}}

	return descr
}

func genSummary(startTime time.Time, i int, valuesPerTimeseries int) *Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i)

	var timeseries2 []*NumberDataPoint
	for j := 0; j < 1; j++ {
		var points []*NumberDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			val := float64(i * j * k)
			point := NumberDataPoint{
				TimeUnixNano: pointTs,
				Value:          &NumberDataPoint_AsDouble{AsDouble: val},
				Labels: []*StringKeyValue{
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
			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}
			points = append(points, &point)
		}

		timeseries2 = append(timeseries2, points...)
	}

	descr.Data = &Metric_Sum{Sum:&Sum{DataPoints:timeseries2}}

	return descr
}

func (g *Generator) GenerateMetricBatch(
	metricsPerBatch int,
	valuesPerTimeseries int,
	int64 bool,
	histogram bool,
	summary bool,
) core.ExportRequest {

	batchStartTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

	il := &InstrumentationLibraryMetrics{}
	batch := &MetricExportRequest{
		ResourceMetrics: []*ResourceMetrics{
			{
				Resource:                      GenResource(),
				InstrumentationLibraryMetrics: []*InstrumentationLibraryMetrics{il},
			},
		},
	}

	for i := 0; i < metricsPerBatch; i++ {
		startTime := batchStartTime.Add(time.Duration(i) * time.Millisecond)

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
