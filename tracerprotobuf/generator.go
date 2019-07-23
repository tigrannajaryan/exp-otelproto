package tracerprotobuf

import "github.com/tigrannajaryan/exp-otelproto/core"

// Generator allows to generate a SpanBatch.
type Generator struct {
}

func (g *Generator) GenerateBatch() core.SpanBatch {
	batch := &SpanBatch{
		Name: "generated span",
	}
	return batch
}
