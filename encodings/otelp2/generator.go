package otelp2

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

func GenResource(deltaDict map[string]uint32) *Resource {
	return &Resource{
		Attributes: []*KeyValue{
			{
				KeyRef: getStringRef(deltaDict, "StartTimeUnixnano"),
				Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: 12345678}},
			},
			{KeyRef: getStringRef(deltaDict, "Pid"), Value: &AnyValue{Value: &AnyValue_IntValue{IntValue: 1234}}},
			{KeyRef: getStringRef(deltaDict, "HostName"), ValueRef: getStringRef(deltaDict, "fakehost")},
			{
				KeyRef:   getStringRef(deltaDict, "ServiceName"),
				ValueRef: getStringRef(deltaDict, "generator"),
			},
		},
	}
}

var builtInDict = createBuiltInDict()

var FirstStringRef = uint32(len(builtInDict) + 1)

func createBuiltInDict() map[string]uint32 {
	m := map[string]uint32{}
	//	for _, str := range core.ExampleAttributeNames {
	//m[str] = uint32(len(m) + 1)
	//}
	return m
}

func getStringRef(dict map[string]uint32, str string) uint32 {
	if ref, found := builtInDict[str]; found {
		return ref
	}

	if ref, found := dict[str]; found {
		return ref
	}

	ref := uint32(FirstStringRef + uint32(len(dict)))
	dict[str] = ref
	return ref
}

func createDict(dict map[string]uint32) *StringDict {
	r := &StringDict{
		StartIndex: FirstStringRef,
		Values:     make([]string, len(dict)),
	}
	for k, v := range dict {
		r.Values[v-FirstStringRef] = k
	}
	//for _, v := range r.Values {
	//	if v == "" {
	//		panic("Empty string in the dictionary")
	//	}
	//}

	return r
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batchStartTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

	deltaDict := map[string]uint32{}
	il := &InstrumentationLibrarySpans{
		InstrumentationLibrary: &InstrumentationLibrary{NameRef: getStringRef(deltaDict, "io.opentelemetry")},
	}
	batch := &TraceExportRequest{
		ResourceSpans: []*ResourceSpans{
			{
				Resource:                    GenResource(deltaDict),
				InstrumentationLibrarySpans: []*InstrumentationLibrarySpans{il},
			},
		},
		StartTimeUnixNano: core.TimeToTimestamp(batchStartTime),
	}

	for i := 0; i < spansPerBatch; i++ {

		spanID := atomic.AddUint64(&g.spansSent, 1)
		startTime := batchStartTime.Add(time.Duration(i) * time.Millisecond)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              Span_SPAN_KIND_CLIENT,
			StartTimeUnixNano: startTime.Sub(batchStartTime).Nanoseconds(),
			DurationNano:      uint64((time.Duration(i) * time.Millisecond).Nanoseconds()),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*KeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, "load_generator.span_seq_num"),
						Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}},
					})
				span.Attributes = append(span.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, "load_generator.trace_seq_num"),
						Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(traceID)}},
					})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := core.GenRandAttrName(g.random)
				attrVal := g.genRandByteString(g.random.Intn(20) + 1)
				span.Attributes = append(span.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, attrName),
						Value:  &AnyValue{Value: &AnyValue_StringValue{StringValue: attrVal}},
					})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.Events = append(span.Events, &Span_Event{
					TimeUnixNano: (time.Duration(i) * time.Millisecond).Nanoseconds(),
					// TimeStartDeltaNano: (time.Duration(i) * time.Millisecond).Nanoseconds(),
					Attributes: []*KeyValue{
						{
							KeyRef: getStringRef(deltaDict, "te"),
							Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}},
						},
					},
				})
			}
		}

		il.Spans = append(il.Spans, span)
	}

	batch.StringDict = createDict(deltaDict)

	return batch
}

func (g *Generator) GenerateLogBatch(logsPerBatch int, attrsPerLog int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batchStartTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

	deltaDict := map[string]uint32{}
	il := &InstrumentationLibraryLogs{
		InstrumentationLibrary: &InstrumentationLibrary{NameRef: getStringRef(deltaDict, "io.opentelemetry")},
	}
	batch := &ExportLogsServiceRequest{
		ResourceLogs: []*ResourceLogs{
			{
				Resource:                   GenResource(deltaDict),
				InstrumentationLibraryLogs: []*InstrumentationLibraryLogs{il},
			},
		},
		StartTimeUnixNano: core.TimeToTimestamp(batchStartTime),
	}

	logs := []*LogRecord{}
	for i := 0; i < logsPerBatch; i++ {

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a log.
		log := &LogRecord{
			TimeUnixNano:   (time.Duration(i) * time.Millisecond).Nanoseconds(),
			TraceId:        core.GenerateTraceID(traceID),
			SpanId:         core.GenerateSpanID(spanID),
			SeverityNumber: SeverityNumber_SEVERITY_NUMBER_INFO,
			SeverityText:   "info",
			Name:           "ProcessStarted",
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
				log.Attributes = append(
					log.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, "load_generator.span_seq_num"),
						Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(spanID)}},
					},
				)
				log.Attributes = append(
					log.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, "load_generator.trace_seq_num"),
						Value:  &AnyValue{Value: &AnyValue_IntValue{IntValue: int64(traceID)}},
					},
				)
			}

			for j := len(log.Attributes); j < attrsPerLog; j++ {
				attrName := core.GenRandAttrName(g.random)
				log.Attributes = append(
					log.Attributes,
					&KeyValue{
						KeyRef: getStringRef(deltaDict, attrName),
						Value:  &AnyValue{Value: &AnyValue_StringValue{StringValue: g.genRandByteString(g.random.Intn(20) + 1)}},
					},
				)
			}

			log.Attributes = append(
				log.Attributes,
				&KeyValue{
					KeyRef:   getStringRef(deltaDict, "event_type"),
					ValueRef: getStringRef(deltaDict, "auto_generated_event"),
				},
			)

		}

		logs = append(logs, log)
	}

	il.Logs = logs
	batch.StringDict = createDict(deltaDict)

	return batch
}

func GenInt64Timeseries(
	startTime time.Time, offset int, valuesPerTimeseries int, deltaDict map[string]uint32,
) *Metric_IntGauge {
	var timeseries []*IntDataPoint
	for j := 0; j < 5; j++ {
		var points []*IntDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := (time.Duration(j*k) * time.Millisecond).Nanoseconds()

			point := IntDataPoint{
				TimeUnixNano: pointTs,
				Value:        int64(offset * j * k),
				Labels: []*StringKeyValue{
					{
						KeyRef:   getStringRef(deltaDict, "label1"),
						ValueRef: getStringRef(deltaDict, "val1"),
					},
					{
						KeyRef:   getStringRef(deltaDict, "label2"),
						ValueRef: getStringRef(deltaDict, "val2"),
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

	return &Metric_IntGauge{IntGauge: &IntGauge{DataPoints: timeseries}}
}

func genInt64Gauge(
	startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int, deltaDict map[string]uint32,
) *Metric {
	descr := GenMetricDescriptor(i, deltaDict)
	descr.Data = GenInt64Timeseries(startTime, i, valuesPerTimeseries, deltaDict)
	return descr
}

func GenMetricDescriptor(i int, deltaDict map[string]uint32) *Metric {
	descr := &Metric{
		NameRef:        getStringRef(deltaDict, "metric"+strconv.Itoa(i)),
		DescriptionRef: getStringRef(deltaDict, "some description: "+strconv.Itoa(i)),
	}
	return descr
}

func genHistogram(
	startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int, deltaDict map[string]uint32,
) *Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i, deltaDict)

	var timeseries2 []*DoubleHistogramDataPoint
	for j := 0; j < 1; j++ {
		var points []*DoubleHistogramDataPoint

		//prevPointTs := int64(0)
		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := (time.Duration(j*k) * time.Millisecond).Nanoseconds()
			//diffTs := pointTs - prevPointTs
			//prevPointTs = pointTs
			val := float64(i * j * k)
			point := DoubleHistogramDataPoint{
				TimeUnixNano: pointTs,
				Count:        1,
				Sum:          val,
				BucketCounts: []uint64{12, 345},
				Exemplars: []*DoubleExemplar{
					{
						Value:        val,
						TimeUnixNano: pointTs,
					},
				},
				ExplicitBounds: []float64{0, 1000000},
				Labels: []*StringKeyValue{
					{
						KeyRef:   getStringRef(deltaDict, "label1"),
						ValueRef: getStringRef(deltaDict, "val1"),
					},
					{
						KeyRef:   getStringRef(deltaDict, "label2"),
						ValueRef: getStringRef(deltaDict, "val2"),
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

	descr.Data = &Metric_DoubleHistogram{DoubleHistogram: &DoubleHistogram{DataPoints: timeseries2}}

	return descr
}

func genSummary(
	startTime time.Time, i int, labelKeys []string, valuesPerTimeseries int, deltaDict map[string]uint32,
) *Metric {
	// Add Histogram
	descr := GenMetricDescriptor(i, deltaDict)

	var timeseries2 []*DoubleDataPoint
	for j := 0; j < 1; j++ {
		var points []*DoubleDataPoint

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := (time.Duration(j*k) * time.Millisecond).Nanoseconds()
			val := float64(i * j * k)
			point := DoubleDataPoint{
				TimeUnixNano: pointTs,
				Value:        val,
				Labels: []*StringKeyValue{
					{
						KeyRef:   getStringRef(deltaDict, "label1"),
						ValueRef: getStringRef(deltaDict, "val1"),
					},
					{
						KeyRef:   getStringRef(deltaDict, "label2"),
						ValueRef: getStringRef(deltaDict, "val2"),
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

	descr.Data = &Metric_DoubleSum{DoubleSum: &DoubleSum{DataPoints: timeseries2}}

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
	deltaDict := map[string]uint32{}
	batch := &MetricExportRequest{
		ResourceMetrics: []*ResourceMetrics{
			{
				Resource:                      GenResource(deltaDict),
				InstrumentationLibraryMetrics: []*InstrumentationLibraryMetrics{il},
			},
		},
		StartTimeUnixNano: core.TimeToTimestamp(batchStartTime),
	}

	for i := 0; i < metricsPerBatch; i++ {
		startTime := batchStartTime.Add(time.Duration(i) * time.Millisecond)

		labelKeys := []string{
			"label1",
			"label2",
		}

		if int64 {
			il.Metrics = append(
				il.Metrics, genInt64Gauge(startTime, i, labelKeys, valuesPerTimeseries, deltaDict),
			)
		}
		if histogram {
			il.Metrics = append(
				il.Metrics, genHistogram(startTime, i, labelKeys, valuesPerTimeseries, deltaDict),
			)
		}
		if summary {
			il.Metrics = append(il.Metrics, genSummary(startTime, i, labelKeys, valuesPerTimeseries, deltaDict))
		}
	}

	batch.StringDict = createDict(deltaDict)

	return batch
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}
