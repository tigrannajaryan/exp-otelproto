package octraceprotobuf

import (
	"encoding/binary"
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
		Labels: map[string]string{
			"StartTimeUnixnano": "12345678",
			"Pid":               "1234",
			"HostName":          "fakehost",
			"ServiceName":       "generator",
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)
	batch := &ExportRequest{
		Resource: genResource(),
	}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:    generateTraceID(traceID),
			SpanId:     generateSpanID(spanID),
			Name:       &TruncatableString{Value: "load-generator-span"},
			Kind:       Span_CLIENT,
			Attributes: &Span_Attributes{},
			StartTime:  timeToTimestamp(startTime),
			EndTime:    timeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes.AttributeMap = map[string]*AttributeValue{}

			if attrsPerSpan >= 2 {
				span.Attributes.AttributeMap["load_generator.span_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(spanID)}}
				span.Attributes.AttributeMap["load_generator.trace_seq_num"] = &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}}
			}

			for j := len(span.Attributes.AttributeMap); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes.AttributeMap[attrName] = &AttributeValue{
					Value: &AttributeValue_StringValue{
						StringValue: &TruncatableString{Value: g.genRandByteString(g.random.Intn(20) + 1)},
					},
				}
			}

		}

		if timedEventsPerSpan > 0 {
			span.TimeEvents = &Span_TimeEvents{}
			for i := 0; i < timedEventsPerSpan; i++ {
				ts := startTime.Add(time.Duration(i) * time.Millisecond)
				span.TimeEvents.TimeEvent = append(span.TimeEvents.TimeEvent, &Span_TimeEvent{
					Time: timeToTimestamp(ts),
					Value: &Span_TimeEvent_Annotation_{
						Annotation: &Span_TimeEvent_Annotation{
							Attributes: &Span_Attributes{
								AttributeMap: map[string]*AttributeValue{
									"te": {Value: &AttributeValue_IntValue{IntValue: int64(spanID)}},
								},
							},
						},
					},
				})
			}
		}

		batch.Spans = append(batch.Spans, span)
	}
	return batch
}

func generateTraceID(id uint64) []byte {
	var traceID [16]byte
	binary.PutUvarint(traceID[:], id)
	return traceID[:]
}

func generateSpanID(id uint64) []byte {
	var spanID [8]byte
	binary.PutUvarint(spanID[:], id)
	return spanID[:]
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}

func genInt64Gauge(startTime time.Time, i int, labelKeys []*LabelKey, valuesPerTimeseries int) *Metric {
	descr := &MetricDescriptor{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Type:        MetricDescriptor_GAUGE_INT64,
		LabelKeys:   labelKeys,
	}

	var timeseries []*TimeSeries
	for j := 0; j < 5; j++ {
		var points []*Point

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := startTime.Add(time.Duration(j*k) * time.Millisecond)
			point := Point{
				Timestamp: timeToTimestamp(pointTs),
				Value:     &Point_Int64Value{Int64Value: int64(i * j * k)},
			}
			points = append(points, &point)
		}

		ts := TimeSeries{
			StartTimestamp: timeToTimestamp(startTime),
			LabelValues: []*LabelValue{
				{Value: "val1", HasValue: true},
				{Value: "val2", HasValue: true},
			},
			Points: points,
		}
		timeseries = append(timeseries, &ts)
	}

	metric1 := &Metric{
		MetricDescriptor: descr,
		Timeseries:       timeseries,
	}

	return metric1
}

func genHistogram(startTime time.Time, i int, labelKeys []*LabelKey, valuesPerTimeseries int) *Metric {
	// Add Histogram
	descr := &MetricDescriptor{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Type:        MetricDescriptor_GAUGE_DISTRIBUTION,
		LabelKeys:   labelKeys,
	}

	timeseries := []*TimeSeries{}
	for j := 0; j < 1; j++ {
		var points []*Point

		for k := 0; k < valuesPerTimeseries; k++ {
			pointTs := timeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			val := float64(i * j * k)
			point := Point{
				Timestamp: pointTs,
				Value: &Point_DistributionValue{
					&DistributionValue{
						Count:                 1,
						Sum:                   val,
						SumOfSquaredDeviation: 123,
						BucketOptions: &DistributionValue_BucketOptions{
							Type: &DistributionValue_BucketOptions_Explicit_{
								Explicit: &DistributionValue_BucketOptions_Explicit{
									Bounds: []float64{0, val},
								},
							},
						},
						Buckets: []*DistributionValue_Bucket{
							{
								Count: 12,
								Exemplar: &DistributionValue_Exemplar{
									Value:     val,
									Timestamp: pointTs,
								},
							},
							{
								Count: 345,
							},
						},
					},
				},
			}
			points = append(points, &point)
		}

		ts := TimeSeries{
			StartTimestamp: timeToTimestamp(startTime),
			LabelValues: []*LabelValue{
				{Value: "val1", HasValue: true},
				{Value: "val2", HasValue: true},
			},
			Points: points,
		}
		timeseries = append(timeseries, &ts)
	}

	metric2 := &Metric{
		MetricDescriptor: descr,
		Timeseries:       timeseries,
	}

	return metric2
}

func (g *Generator) GenerateMetricBatch(metricsPerBatch int, valuesPerTimeseries int) core.ExportRequest {
	batch := &ExportMetricsServiceRequest{
		Resource: genResource(),
	}
	for i := 0; i < metricsPerBatch/2; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		labelKeys := []*LabelKey{
			{Key: "label1"},
			{Key: "label2"},
		}

		batch.Metrics = append(batch.Metrics, genInt64Gauge(startTime, i, labelKeys, valuesPerTimeseries))
		batch.Metrics = append(batch.Metrics, genHistogram(startTime, i, labelKeys, valuesPerTimeseries))
	}
	return batch
}
