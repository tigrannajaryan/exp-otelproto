package encodings

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"

	"github.com/tigrannajaryan/exp-otelproto/encodings/baseline"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/core"
)

var tests = []struct {
	name string
	gen  func() core.Generator
}{
	/*{
		name: "OpenCensus",
		gen:  func() core.Generator { return octraceprotobuf.NewGenerator() },
	},*/
	{
		name: "Baseline",
		gen:  func() core.Generator { return baseline.NewGenerator() },
	},
	{
		name: "Proposed",
		gen:  func() core.Generator { return experimental.NewGenerator() },
	},
	/*{
		name: "OTLP",
		gen:  func() core.Generator { return otlp.NewGenerator() },
	},*/
	/* These are historical experiments. Uncomment if interested to see results.
	{
		name: "OC+AttrAsMap",
		gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
	},
	{
		name: "OC+AttrAsList+TimeWrapped",
		gen:  func() core.Generator { return otlptimewrapped.NewGenerator() },
	},
	*/
}

var batchTypes = []struct {
	name     string
	batchGen func(gen core.Generator) []core.ExportRequest
}{
	//{name: "Attributes", batchGen: generateAttrBatches},
	//{name: "TimedEvent", batchGen: generateTimedEventBatches},
	{name: "MetricOne", batchGen: generateMetricOneBatches},
	{name: "MetricSeries", batchGen: generateMetricSeriesBatches},
}

const BatchCount = 1000

func BenchmarkEncode(b *testing.B) {

	for _, batchType := range batchTypes {
		for _, test := range tests {
			b.Run(test.name+"/"+batchType.name, func(b *testing.B) {
				b.StopTimer()
				gen := test.gen()
				batches := batchType.batchGen(gen)
				if batches == nil {
					// Unsupported test type and batch type combination.
					b.SkipNow()
					return
				}

				runtime.GC()
				b.StartTimer()
				for i := 0; i < b.N; i++ {
					for _, batch := range batches {
						encode(batch)
					}
				}
			})
		}
		fmt.Println("")
	}
}

func BenchmarkDecode(b *testing.B) {
	for _, batchType := range batchTypes {
		for _, test := range tests {
			b.Run(test.name+"/"+batchType.name, func(b *testing.B) {
				b.StopTimer()
				batches := batchType.batchGen(test.gen())
				if batches == nil {
					// Unsupported test type and batch type combination.
					b.SkipNow()
					return
				}

				var encodedBytes [][]byte
				for _, batch := range batches {
					encodedBytes = append(encodedBytes, encode(batch))
				}

				runtime.GC()
				b.StartTimer()
				for i := 0; i < b.N; i++ {
					for j, bytes := range encodedBytes {
						decode(bytes, batches[j].(proto.Message))
					}
				}
			})
		}
		fmt.Println("")
	}
}

func generateAttrBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateSpanBatch(100, 3, 0))
	}
	return batches
}

func generateMetricOneBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(100, 1)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricSeriesBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(100, 5)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateTimedEventBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateSpanBatch(100, 0, 3))
	}
	return batches
}

func encode(request core.ExportRequest) []byte {
	bytes, err := proto.Marshal(request.(proto.Message))
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func decode(bytes []byte, pb proto.Message) {
	err := proto.Unmarshal(bytes, pb)
	if err != nil {
		log.Fatal(err)
	}
}

func TestEncodeSize(t *testing.T) {

	const batchSize = 100

	variation := []struct {
		name                 string
		genFunc              func(gen core.Generator) core.ExportRequest
		firstUncompessedSize int
		firstCompressedSize  int
	}{
		/*{
			name: "Trace",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateSpanBatch(batchSize, 3, 0)
			},
		},
		{
			name: "Event",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateSpanBatch(batchSize, 0, 3)
			},
		},*/
		{
			name: "MetricOne",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateMetricBatch(batchSize, 1)
			},
		},
		{
			name: "MetricSeries",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateMetricBatch(batchSize, 5)
			},
		},
	}

	fmt.Println("===== Encoded sizes")

	for _, v := range variation {
		fmt.Println("Encoding                       Uncompressed  Improved        Compressed  Improved")
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				gen := test.gen()

				batch := v.genFunc(gen)
				if batch == nil {
					// Skip this case.
					return
				}

				bodyBytes, err := proto.Marshal(batch.(proto.Message))
				if err != nil {
					log.Fatal(err)
				}

				// Try to compress
				var b bytes.Buffer
				w := zlib.NewWriter(&b)
				w.Write(bodyBytes)
				w.Close()
				compressedBytes := b.Bytes()

				uncompressedSize := len(bodyBytes)
				compressedSize := len(compressedBytes)

				uncompressedRatioStr := "[1.000]"
				compressedRatioStr := "[1.000]"

				if v.firstUncompessedSize == 0 {
					v.firstUncompessedSize = uncompressedSize
				} else {
					uncompressedRatioStr = fmt.Sprintf(" [%1.3f]", float64(v.firstUncompessedSize)/float64(uncompressedSize))
				}

				if v.firstCompressedSize == 0 {
					v.firstCompressedSize = compressedSize
				} else {
					compressedRatioStr = fmt.Sprintf(" [%1.3f]", float64(v.firstCompressedSize)/float64(compressedSize))
				}

				fmt.Printf(
					"%-31v %5d bytes%9s, gziped %4d bytes%9s\n",
					test.name+"/"+v.name,
					uncompressedSize,
					uncompressedRatioStr,
					compressedSize,
					compressedRatioStr,
				)

			})
		}
		fmt.Println("")
	}
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

	data := []*otlp.Int64TimeSeries{
		{
			LabelValues: []*otlp.LabelValue{
				{Value: "val1"},
				{Value: "val2"},
			},
			Points: []*otlp.Int64Value{
				{
					TimestampUnixnano: time.Now().UnixNano(),
					Value:             123,
				},
			},
		},
	}

	metric1 := &otlp.Metric{
		MetricDescriptor: descr,
		Resource:         resource,
		Int64Timeseries:  data,
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
		Int64Timeseries:  data,
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

func encodeUnprepared(metricCount int) proto.Message {
	batch := &otlp.ResourceMetrics{}
	for i := 0; i < metricCount; i++ {
		metric := &otlp.Metric{
			MetricDescriptor: otlp.GenMetricDescriptor(1),
			Int64Timeseries:  otlp.GenInt64Timeseries(time.Now(), i, 1),
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

func encodePrepared(metricCount int) proto.Message {
	batch := &otlp.ResourceMetricsPrepared{}
	descr := otlp.GenMetricDescriptor(1)
	descrBytes, err := proto.Marshal(descr)
	if err != nil {
		log.Fatal("Cannot marshal")
	}
	for i := 0; i < metricCount; i++ {
		metric := &otlp.MetricPrepared{
			MetricDescriptor: descrBytes,
			Int64Timeseries:  otlp.GenInt64Timeseries(time.Now(), i, 1),
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

func BenchmarkEncodePreparedVsUnprepared(b *testing.B) {
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
			batch := test.encoder(1000)
			runtime.GC()
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
