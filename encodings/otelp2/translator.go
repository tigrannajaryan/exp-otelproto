package otelp2

import (
	otlptracecol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	v12 "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type SpanTranslator struct {
}

func (st *SpanTranslator) TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) core.ExportRequest {

	res := &TraceExportRequest{}
	dict := map[string]uint32{}

	for _, rssi := range batch.ResourceSpans {
		rsso := &ResourceSpans{
			Resource: &Resource{
				Attributes: translateAttrs(dict, rssi.Resource.Attributes),
			},
			InstrumentationLibrarySpans: translateInstrumentationLibrarySpans(dict, rssi.InstrumentationLibrarySpans),
		}
		res.ResourceSpans = append(res.ResourceSpans, rsso)

	}

	res.StringDict = createDict(dict)

	return res
}

func translateAttrs(dict map[string]uint32, attrs []*v1.KeyValue) (r []*KeyValue) {
	for _, attr := range attrs {
		kv := &KeyValue{
			//Key:               attr.Key,
			KeyRef: getStringRef(dict, attr.Key),
		}

		var v *AnyValue
		switch iv := attr.Value.Value.(type) {
		case *v1.AnyValue_StringValue:
			kv.ValueRef = getStringRef(dict, iv.StringValue)
			//v = &AnyValue{Value:&AnyValue_StringValue{StringValue:iv.StringValue}}
		case *v1.AnyValue_BoolValue:
			v = &AnyValue{Value: &AnyValue_BoolValue{BoolValue: iv.BoolValue}}
		case *v1.AnyValue_IntValue:
			v = &AnyValue{Value: &AnyValue_IntValue{IntValue: iv.IntValue}}
		case *v1.AnyValue_DoubleValue:
			v = &AnyValue{Value: &AnyValue_DoubleValue{DoubleValue: iv.DoubleValue}}
		default:
			panic("not implemented")
		}

		kv.Value = v

		r = append(r, kv)
	}
	return r
}

func translateInstrumentationLibrarySpans(
	dict map[string]uint32,
	in []*v12.InstrumentationLibrarySpans,
) (r []*InstrumentationLibrarySpans) {

	for _, ils := range in {
		out := &InstrumentationLibrarySpans{
			InstrumentationLibrary: translateInstrumentationLibrary(dict, ils.InstrumentationLibrary),
		}

		for _, span := range ils.Spans {
			outSpan := translateSpan(dict, span)
			out.Spans = append(out.Spans, outSpan)
		}

		r = append(r, out)
	}

	return r
}

func translateSpan(dict map[string]uint32, span *v12.Span) *Span {
	if span == nil {
		return nil
	}

	return &Span{
		TraceId:                span.TraceId,
		SpanId:                 span.SpanId,
		TraceState:             span.TraceState,
		ParentSpanId:           span.ParentSpanId,
		Name:                   span.Name,
		Kind:                   Span_SpanKind(span.Kind),
		StartTimeUnixNano:      int64(span.StartTimeUnixNano),
		DurationNano:           span.EndTimeUnixNano - span.StartTimeUnixNano,
		Attributes:             translateAttrs(dict, span.Attributes),
		DroppedAttributesCount: span.DroppedAttributesCount,
		Events:                 nil,
		DroppedEventsCount:     0,
		Links:                  nil,
		DroppedLinksCount:      0,
		Status:                 nil,
	}
}

func translateInstrumentationLibrary(dict map[string]uint32, in *v1.InstrumentationLibrary) *InstrumentationLibrary {
	if in == nil {
		return nil
	}
	return &InstrumentationLibrary{
		NameRef:    getStringRef(dict, in.Name),
		VersionRef: getStringRef(dict, in.Version),
	}
}
