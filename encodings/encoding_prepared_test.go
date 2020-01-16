package encodings

import (
	"bytes"
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

func createLabels() []string {
	return []string{"val1", "val2"}
}

func TestPreparedMetric(t *testing.T) {

	descr := &otlp.MetricDescriptor{
		Name:        "metric1",
		Description: "some description: 1",
		Type:        otlp.MetricDescriptor_GAUGE_INT64,
		LabelKeys: []string{
			"label1",
			"label2",
		},
	}

	resource := otlp.GenResource()

	data := []*otlp.Int64Value{
		{
			TimestampUnixnano: uint64(time.Now().UnixNano()),
			Value:             123,
		},
	}

	labelValues := createLabels()

	metric1 := &otlp.Metric{
		MetricDescriptor: descr,
		Resource:         resource,
		Int64Timeseries: []*otlp.Int64TimeSeries{
			{
				LabelValues: labelValues,
				Points:      data,
			},
		},
	}

	descrBytes, err := proto.Marshal(descr)
	if err != nil {
		t.Fatal()
	}

	resourceBytes, err := proto.Marshal(resource)
	if err != nil {
		t.Fatal()
	}

	metric2 := &otlp.MetricPrepared{
		MetricDescriptor: descrBytes,
		Resource:         resourceBytes,
		Int64Timeseries: []*otlp.Int64TimeSeries{
			{
				LabelValues: labelValues,
				Points:      data,
			},
		},
	}

	b1, err := proto.Marshal(metric1)
	if err != nil {
		t.Fatal()
	}

	b2, err := proto.Marshal(metric2)
	if err != nil {
		t.Fatal()
	}

	if c := bytes.Compare(b1, b2); c != 0 {
		t.Fatal()
	}

	var metric3 otlp.Metric
	if err := proto.Unmarshal(b2, &metric3); err != nil {
		t.Fatal()
	}
}

func TestPreparedTrace(t *testing.T) {

	g := otlp.NewGenerator()
	request := g.GenerateSpanBatch(100, 5, 0).(*otlp.TraceExportRequest)

	requestBytes, err := proto.Marshal(request)
	if err != nil {
		t.Fatal()
	}

	var preparedRequest otlp.TraceExportRequestPrepared
	if err := proto.Unmarshal(requestBytes, &preparedRequest); err != nil {
		t.Fatal()
	}

	for i, spans := range preparedRequest.ResourceSpans {

		var resource otlp.Resource
		if err := proto.Unmarshal(spans.Resource, &resource); err != nil {
			t.Fatal()
		}

		if !proto.Equal(&resource, request.ResourceSpans[i].Resource) {
			t.Fatal()
		}

		for j, span := range spans.Spans {
			for k, attrBytes := range span.Attributes {
				var attr otlp.AttributeKeyValue
				if err := proto.Unmarshal(attrBytes, &attr); err != nil {
					t.Fatal()
				}

				if !proto.Equal(&attr, request.ResourceSpans[i].Spans[j].Attributes[k]) {
					t.Fatal()
				}
			}
		}
	}

	preparedRequestBytes, err := proto.Marshal(&preparedRequest)
	if err != nil {
		t.Fatal()
	}

	if c := bytes.Compare(requestBytes, preparedRequestBytes); c != 0 {
		t.Fatal()
	}
}

func genInt64DataPoints(offset int) []*otlp.Int64Value {
	var points []*otlp.Int64Value

	for k := 0; k < 1; k++ {
		pointTs := core.TimeToTimestamp(time.Now().Add(time.Duration(k) * time.Millisecond))

		point := otlp.Int64Value{
			TimestampUnixnano: pointTs,
			Value:             int64(offset * k),
		}

		//sz := unsafe.Sizeof(SummaryValue{})
		//log.Printf("size=%v", sz)
		if k == 0 {
			point.StartTimeUnixnano = pointTs
		}

		points = append(points, &point)
	}

	return points
}

func encodeUnpreparedMetrics(metricCount int) proto.Message {
	batch := &otlp.ResourceMetrics{}
	for i := 0; i < metricCount; i++ {
		metric := &otlp.Metric{
			MetricDescriptor: otlp.GenMetricDescriptor(1),
			Int64Timeseries: []*otlp.Int64TimeSeries{
				{
					LabelValues: createLabels(),
					Points:      genInt64DataPoints(i),
				},
			},
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

func encodeLabelValues(labelValues []*otlp.LabelValue) [][]byte {
	var arr [][]byte
	for i := 0; i < len(labelValues); i++ {
		b, err := proto.Marshal(labelValues[i])
		if err != nil {
			log.Fatal("Encoding failed")
		}
		arr = append(arr, b)
	}
	return arr
}

func encodePreparedMetrics(metricCount int) proto.Message {
	batch := &otlp.ResourceMetricsPrepared{}
	descr := otlp.GenMetricDescriptor(1)
	descrBytes, err := proto.Marshal(descr)
	if err != nil {
		log.Fatal("Cannot marshal")
	}

	labelValues := createLabels()

	for i := 0; i < metricCount; i++ {
		metric := &otlp.MetricPrepared{
			MetricDescriptor: descrBytes,
			Int64Timeseries: []*otlp.Int64TimeSeries{
				{
					LabelValues: labelValues,
					Points:      genInt64DataPoints(i),
				},
			},
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

func BenchmarkEncode100SingleMetrics(b *testing.B) {
	tests := []struct {
		name    string
		encoder func(metricCount int) proto.Message
	}{
		{
			name:    "Unprepared",
			encoder: encodeUnpreparedMetrics,
		},
		{
			name:    "Prepared",
			encoder: encodePreparedMetrics,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batch := test.encoder(100)
			runtime.GC()
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				var err error
				bytes, err := proto.Marshal(batch)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}

func encodeUnpreparedTraces(spanCount int) proto.Message {
	g := otlp.NewGenerator()
	request := g.GenerateSpanBatch(spanCount, 5, 0).(*otlp.TraceExportRequest)
	return request
}

func encodePreparedTraces(spanCount int) proto.Message {
	request := encodeUnpreparedTraces(spanCount)

	requestBytes, err := proto.Marshal(request)
	if err != nil {
		log.Fatal()
	}

	var preparedRequest otlp.TraceExportRequestPrepared
	if err := proto.Unmarshal(requestBytes, &preparedRequest); err != nil {
		log.Fatal()
	}

	return &preparedRequest
}

func BenchmarkEncodeTraces(b *testing.B) {
	tests := []struct {
		name    string
		encoder func(spanCount int) proto.Message
	}{
		{
			name:    "Unprepared",
			encoder: encodeUnpreparedTraces,
		},
		{
			name:    "Prepared",
			encoder: encodePreparedTraces,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batch := test.encoder(100)
			runtime.GC()
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				var err error
				bytes, err := proto.Marshal(batch)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}

func BenchmarkDecodeEncodeTraces(b *testing.B) {
	tests := []struct {
		name            string
		emptyMsgCreator func() proto.Message
	}{
		{
			name:            "Full",
			emptyMsgCreator: func() proto.Message { return &otlp.TraceExportRequest{} },
		},
		{
			name:            "Partial",
			emptyMsgCreator: func() proto.Message { return &otlp.TraceExportRequestPrepared{} },
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batch := encodeUnpreparedTraces(100)
			runtime.GC()

			var err error
			bytes, err := proto.Marshal(batch)
			if err != nil || len(bytes) == 0 {
				log.Fatal("Cannot encode batch")
			}

			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				msg := test.emptyMsgCreator()
				err = proto.Unmarshal(bytes, msg)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot decode batch")
				}

				bytes, err := proto.Marshal(msg)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}
