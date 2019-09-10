package encodings

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/tigrannajaryan/exp-otelproto/encodings/otlptimewrapped"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"

	"github.com/golang/protobuf/proto"

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
		name: "OTLP/AttrAsMap",
		gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
	},
	{
		name: "OTLP/AttrAsList",
		gen:  func() core.Generator { return otlp.NewGenerator() },
	},
	{
		name: "OTLP/AttrAsList/TimeWrapped",
		gen:  func() core.Generator { return otlptimewrapped.NewGenerator() },
	},
}

var batchTypes = []struct {
	name     string
	batchGen func(gen core.Generator) []core.ExportRequest
}{
	{name: "Attributes", batchGen: generateAttrBatches},
	{name: "TimedEvent", batchGen: generateTimedEventBatches},
}

const BatchCount = 1000

func BenchmarkEncode(b *testing.B) {

	for _, test := range tests {
		for _, batchType := range batchTypes {
			b.Run(test.name+"/"+batchType.name, func(b *testing.B) {
				b.StopTimer()
				batches := batchType.batchGen(test.gen())

				runtime.GC()
				b.StartTimer()
				for i := 0; i < b.N; i++ {
					for _, batch := range batches {
						encode(batch)
					}
				}
			})
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for _, test := range tests {
		for _, batchType := range batchTypes {
			b.Run(test.name+"/"+batchType.name, func(b *testing.B) {
				b.StopTimer()
				batches := batchType.batchGen(test.gen())
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
	}
}

func generateAttrBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateBatch(100, 3, 0))
	}
	return batches
}

func generateTimedEventBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateBatch(100, 0, 3))
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gen := test.gen()
			batch := gen.GenerateBatch(100, 3, 0)
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

			fmt.Printf("%-15v size %5d bytes, gzip size %5d\n", test.name, len(bodyBytes), len(compressedBytes))
		})
	}
}
