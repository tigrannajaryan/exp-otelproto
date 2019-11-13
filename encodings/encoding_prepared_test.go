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

func encodeUnprepared(metricCount int) proto.Message {
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

func encodePrepared(metricCount int) proto.Message {
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

func BenchmarkEncode100Single(b *testing.B) {
	tests := []struct {
		name    string
		encoder func(metricCount int) proto.Message
	}{
		{
			name:    "Unprepared",
			encoder: encodeUnprepared,
		},
		{
			name:    "Prepared",
			encoder: encodePrepared,
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
