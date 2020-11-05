package core

import (
	otlpmetriccol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
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
	TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) ExportRequest
}

type MetricTranslator interface {
	TranslateMetrics(batch *otlpmetriccol.ExportMetricsServiceRequest) ExportRequest
}
