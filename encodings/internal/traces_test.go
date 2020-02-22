package internal

import (
	"testing"

	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

const BatchCount = 1000

func BenchmarkFromOtlpToInternal(b *testing.B) {
	b.StopTimer()
	g := otlp.NewGenerator()

	var batch []*otlp.TraceExportRequest
	for i := 0; i < BatchCount; i++ {
		batch = append(batch,
			g.GenerateSpanBatch(100, 5, 0).(*otlp.TraceExportRequest))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < BatchCount; i++ {
			FromOtlp(batch[i].ResourceSpans)
		}
	}
}
