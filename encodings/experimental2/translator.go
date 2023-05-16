package experimental2

import (
	"encoding/binary"

	v12 "go.opentelemetry.io/proto/otlp/trace/v1"

	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
	otlptracecolexp "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/collector/trace/v1"
	otlpcommon "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/common/v1"
	otlpresource "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/resource/v1"
	otlptrace "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/trace/v1"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

type spanTranslator struct {
	keyDict       map[string]uint32
	valDict       map[string]uint32
	spanNameDict  map[string]uint32
	eventNameDict map[string]uint32
}

func NewSpanTranslator() *spanTranslator {
	return &spanTranslator{}
}

func (st *spanTranslator) TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) core.ExportRequest {

	res := &otlptracecolexp.ExportTraceServiceRequest{}
	st.keyDict = map[string]uint32{}
	st.valDict = map[string]uint32{}
	st.spanNameDict = map[string]uint32{}
	st.eventNameDict = map[string]uint32{}

	for _, rssi := range batch.ResourceSpans {
		rsso := &otlptrace.ResourceSpans{
			Resource: &otlpresource.Resource{
				Attributes:             st.translateAttrs(rssi.Resource.Attributes),
				DroppedAttributesCount: rssi.Resource.DroppedAttributesCount,
			},
			ScopeSpans: st.translateInstrumentationLibrarySpans(rssi.ScopeSpans),
		}
		res.ResourceSpans = append(res.ResourceSpans, rsso)

	}

	res.KeyDict = createDict(st.keyDict)
	res.ValDict = createDict(st.valDict)
	res.SpanNameDict = createDict(st.spanNameDict)
	res.EventNameDict = createDict(st.eventNameDict)

	return res
}

var builtInDict = map[string]uint32{}

var FirstStringRef = uint32(len(builtInDict) + 1)

//func getStringRef(dict map[string]uint32, str string) uint32 {
//	if ref, found := dict[str]; found {
//		return ref
//	}
//	ref := FirstStringRef + uint32(len(dict))
//	dict[str] = ref
//	return ref
//}

func dictionizeStr(dict map[string]uint32, str *string, ref *uint32) bool {
	//return false
	var idx uint32
	var ok bool
	if idx, ok = dict[*str]; !ok {
		idx = uint32(len(dict) + 1)

		buf := make([]byte, 10)
		n := binary.PutUvarint(buf, uint64(idx))
		if n >= len(*str) {
			// Not worth using ref.
			return false
		}

		dict[*str] = idx
	}
	*str = ""
	*ref = idx
	return true
}

func createDict(dict map[string]uint32) *otlpcommon.StringDict {
	r := &otlpcommon.StringDict{
		StartIndex: FirstStringRef,
		Values:     make([]string, len(dict)),
	}
	for k, v := range dict {
		r.Values[v-FirstStringRef] = k
	}
	//for _, v := range r.Values {
	//	if v == "" {
	//		panic("Empty string in the dictionary")
	//	}
	//}

	return r
}

func (st *spanTranslator) translateAttrs(attrs []*v1.KeyValue) (r []*otlpcommon.KeyValue) {
	for _, attr := range attrs {
		var ref uint32
		var kv *otlpcommon.KeyValue
		if dictionizeStr(st.keyDict, &attr.Key, &ref) {
			kv = &otlpcommon.KeyValue{KeyRef: ref}
		} else {
			kv = &otlpcommon.KeyValue{Key: attr.Key}
		}

		var v *otlpcommon.AnyValue
		switch iv := attr.Value.Value.(type) {
		case *v1.AnyValue_StringValue:
			var ref uint32
			if dictionizeStr(st.valDict, &iv.StringValue, &ref) {
				v = &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringRefValue{StringRefValue: ref}}
			} else {
				v = &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: iv.StringValue}}
			}

		case *v1.AnyValue_BoolValue:
			v = &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_BoolValue{BoolValue: iv.BoolValue}}
		case *v1.AnyValue_IntValue:
			v = &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_IntValue{IntValue: iv.IntValue}}
		case *v1.AnyValue_DoubleValue:
			v = &otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_DoubleValue{DoubleValue: iv.DoubleValue}}
		default:
			panic("not implemented")
		}

		kv.Value = v

		r = append(r, kv)
	}
	return r
}

func (st *spanTranslator) translateInstrumentationLibrarySpans(
	in []*v12.ScopeSpans,
) (r []*otlptrace.ScopeSpans) {

	for _, ils := range in {
		out := &otlptrace.ScopeSpans{
			Scope: st.translateInstrumentationLibrary(ils.Scope),
		}

		for _, span := range ils.Spans {
			outSpan := st.translateSpan(span)
			out.Spans = append(out.Spans, outSpan)
		}

		r = append(r, out)
	}

	return r
}

func (st *spanTranslator) translateSpan(span *v12.Span) *otlptrace.Span {
	if span == nil {
		return nil
	}

	if span.Links != nil {
		panic("not implemented")
	}

	s := &otlptrace.Span{
		TraceId:                span.TraceId,
		SpanId:                 span.SpanId,
		TraceState:             span.TraceState,
		ParentSpanId:           span.ParentSpanId,
		Name:                   span.Name,
		Kind:                   otlptrace.Span_SpanKind(span.Kind),
		StartTimeUnixNano:      span.StartTimeUnixNano,
		EndTimeUnixNano:        span.EndTimeUnixNano,
		Attributes:             st.translateAttrs(span.Attributes),
		DroppedAttributesCount: span.DroppedAttributesCount,
		Events:                 st.translateEvents(span.Events),
		DroppedEventsCount:     span.DroppedEventsCount,
		Links:                  nil,
		DroppedLinksCount:      span.DroppedLinksCount,
		Status:                 st.translateStatus(span.Status),
	}
	dictionizeStr(st.spanNameDict, &s.Name, &s.NameRef)
	return s
}

func (st *spanTranslator) translateInstrumentationLibrary(in *v1.InstrumentationScope) *otlpcommon.InstrumentationScope {
	if in == nil {
		return nil
	}
	is := &otlpcommon.InstrumentationScope{
		Name:                   in.Name,
		Version:                in.Version,
		Attributes:             st.translateAttrs(in.Attributes),
		DroppedAttributesCount: in.DroppedAttributesCount,
	}
	dictionizeStr(st.valDict, &is.Name, &is.NameRef)
	dictionizeStr(st.valDict, &is.Version, &is.VersionRef)
	return is
}

func (st *spanTranslator) translateEvents(events []*v12.Span_Event) []*otlptrace.Span_Event {
	out := []*otlptrace.Span_Event{}
	for _, e := range events {
		out = append(out, st.translateEvent(e))
	}
	return out
}

func (st *spanTranslator) translateEvent(e *v12.Span_Event) *otlptrace.Span_Event {
	return &otlptrace.Span_Event{
		TimeUnixNano:           e.TimeUnixNano,
		Name:                   e.Name,
		Attributes:             st.translateAttrs(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
}

func (st *spanTranslator) translateStatus(status *v12.Status) *otlptrace.Status {
	return &otlptrace.Status{
		Message: status.Message,
		Code:    otlptrace.Status_StatusCode(status.Code),
	}
}
