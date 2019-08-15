package encodings

import (
	"log"
	"testing"

	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"

	"github.com/tigrannajaryan/exp-otelproto/core"

	"github.com/golang/protobuf/proto"

	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
)

var tests = []struct {
	name string
	gen  func() core.Generator
}{
	{
		name: "opencensus",
		gen:  func() core.Generator { return octraceprotobuf.NewGenerator() },
	},
	{
		name: "OTLP_A    ",
		gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
	},
	{
		name: "OTLP_B    ",
		gen:  func() core.Generator { return traceprotobuf.NewGeneratorB() },
	},
}

func BenchmarkEncode(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				generateAndEncode(test.gen())
			}
		})
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name+"+decode", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				batch, bytes := generateAndEncode(test.gen())
				decode(bytes, batch.(proto.Message))
			}
		})
	}
}

func generateAndEncode(gen core.Generator) (core.ExportRequest, []byte) {
	batch := gen.GenerateBatch(100, 2)
	bytes, err := proto.Marshal(batch.(proto.Message))
	if err != nil {
		log.Fatal(err)
	}
	return batch, bytes
}

func decode(bytes []byte, pb proto.Message) {
	err := proto.Unmarshal(bytes, pb)
	if err != nil {
		log.Fatal(err)
	}
}
