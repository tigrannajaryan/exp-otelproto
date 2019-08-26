package encodings

import (
	"fmt"
	"log"
	"runtime"
	"testing"

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
		name: "OTLP/AttrMap",
		gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
	},
	//{
	//	name: "OTLP_B",
	//	gen:  func() core.Generator { return traceprotobufb.NewGenerator() },
	//},
	{
		name: "OTLP/AttrList",
		gen:  func() core.Generator { return otlp.NewGenerator() },
	},
}

const BatchCount = 1000

func BenchmarkEncode(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batches := generateBatches(test.gen())

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

func BenchmarkDecode(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			batches := generateBatches(test.gen())
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

func generateBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateBatch(100, 3))
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
			batch := gen.GenerateBatch(100, 3)
			bytes, err := proto.Marshal(batch.(proto.Message))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v size %v bytes\n", test.name, len(bytes))
		})
	}
}
