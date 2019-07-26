package core

// Generator allows to generate a SpanBatch.
type Generator interface {
	GenerateBatch(spansPerBatch int, attrsPerSpan int) SpanBatch
}
