package core

// SpanGenerator allows to generate a ExportRequest containing a batch of spans.
type SpanGenerator interface {
	GenerateBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) ExportRequest
}

// MetricGenerator allows to generate a ExportRequest containing a batch of metrics.
type MetricGenerator interface {
	GenerateMetricBatch(metricsPerBatch int) ExportRequest
}

type Generator interface {
	SpanGenerator
	MetricGenerator
}
