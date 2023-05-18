package core

import (
	otlpmetriccol "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

// SpanGenerator allows to generate a ExportRequest containing a batch of spans.
type SpanGenerator interface {
	GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) ExportRequest
}

// MetricGenerator allows to generate a ExportRequest containing a batch of metrics.
type MetricGenerator interface {
	GenerateMetricBatch(
		metricsPerBatch int,
		valuesPerTimeseries int,
		int64 bool,
		histogram bool,
		summary bool,
	) ExportRequest
}

// LogGenerator allows to generate a ExportRequest containing a batch of log.
type LogGenerator interface {
	GenerateLogBatch(logsPerBatch int, attrsPerLog int) ExportRequest
}

type Generator interface {
	SpanGenerator
	MetricGenerator
	LogGenerator
}

type SpanTranslator interface {
	FromOTLPSpans(batch *otlptracecol.ExportTraceServiceRequest) ExportRequest
}

type MetricTranslator interface {
	FromOTLPMetrics(batch *otlpmetriccol.ExportMetricsServiceRequest) ExportRequest
}
