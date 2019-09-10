package core

// Generator allows to generate a ExportRequest.
type Generator interface {
	GenerateBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) ExportRequest
}
