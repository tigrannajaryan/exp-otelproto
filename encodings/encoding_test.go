package encodings

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"

	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
)

var tests = []struct {
	name string
	gen  func() core.Generator
}{
	{
		name: "OpenCensus",
		gen:  func() core.Generator { return octraceprotobuf.NewGenerator() },
	},
	{
		name: "OTLP",
		gen:  func() core.Generator { return otlp.NewGenerator() },
	},
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
	{name: "Attributes", batchGen: generateAttrBatches},
	{name: "TimedEvent", batchGen: generateTimedEventBatches},
	{name: "Metrics", batchGen: generateMetricBatches},
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

func generateMetricBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(100)
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

	variation := []struct {
		name                 string
		trace                bool
		firstUncompessedSize int
		firstCompressedSize  int
	}{
		{
			name:  "Trace",
			trace: true,
		},
		{
			name:  "Metric",
			trace: false,
		},
	}

	fmt.Println("===== Encoded sizes")

	for _, v := range variation {
		fmt.Println("Encoding                       Uncompressed  Improved        Compressed  Improved")
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				gen := test.gen()

				var batch core.ExportRequest
				if v.trace {
					batch = gen.GenerateSpanBatch(100, 3, 0)
				} else {
					batch = gen.GenerateMetricBatch(100)
				}
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
