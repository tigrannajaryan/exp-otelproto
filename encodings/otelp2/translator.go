package otelp2

import (
	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	v12 "go.opentelemetry.io/proto/otlp/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type spanTranslator struct {
	dict map[string]uint32
}

func NewSpanTranslator() *spanTranslator {
	return &spanTranslator{dict: map[string]uint32{}}
}

func (st *spanTranslator) TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) core.ExportRequest {

	res := &TraceExportRequest{}
	deltaDict := map[string]uint32{}

	for _, rssi := range batch.ResourceSpans {
		rsso := &ResourceSpans{
			Resource: &Resource{
				Attributes: translateAttrs(st.dict, deltaDict, rssi.Resource.Attributes),
			},
			InstrumentationLibrarySpans: translateInstrumentationLibrarySpans(st.dict, deltaDict, rssi.ScopeSpans),
		}
		res.ResourceSpans = append(res.ResourceSpans, rsso)

	}

	res.StringDict = createDict(deltaDict)
	for k, v := range deltaDict {
		st.dict[k] = v
	}

	return res
}

func translateAttrs(dict map[string]uint32, deltaDict map[string]uint32, attrs []*v1.KeyValue) (r []*KeyValue) {
	for _, attr := range attrs {
		kv := &KeyValue{
			//Key:               attr.Key,
			KeyRef: getStringRef(dict, deltaDict, attr.Key),
		}

		var v *AnyValue
		switch iv := attr.Value.Value.(type) {
		case *v1.AnyValue_StringValue:
			kv.ValueRef = getStringRef(dict, deltaDict, iv.StringValue)
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
	dict, deltaDict map[string]uint32,
	in []*v12.ScopeSpans,
) (r []*InstrumentationLibrarySpans) {

	for _, ils := range in {
		out := &InstrumentationLibrarySpans{
			InstrumentationLibrary: translateInstrumentationLibrary(dict, deltaDict, ils.Scope),
		}

		for _, span := range ils.Spans {
			outSpan := translateSpan(dict, deltaDict, span)
			out.Spans = append(out.Spans, outSpan)
		}

		r = append(r, out)
	}

	return r
}

func translateSpan(dict, deltaDict map[string]uint32, span *v12.Span) *Span {
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
		Attributes:             translateAttrs(dict, deltaDict, span.Attributes),
		DroppedAttributesCount: span.DroppedAttributesCount,
		Events:                 nil,
		DroppedEventsCount:     0,
		Links:                  nil,
		DroppedLinksCount:      0,
		Status:                 nil,
	}
}

func translateInstrumentationLibrary(
	dict, deltaDict map[string]uint32, in *v1.InstrumentationScope,
) *InstrumentationLibrary {
	if in == nil {
		return nil
	}
	return &InstrumentationLibrary{
		NameRef:    getStringRef(dict, deltaDict, in.Name),
		VersionRef: getStringRef(dict, deltaDict, in.Version),
	}
}
