package experimental2

import (
	"encoding/binary"
	"sort"

	v12 "go.opentelemetry.io/proto/otlp/trace/v1"

	otlptracecol "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
	otlptracecolexp "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/collector/trace/v1"
	otlpcommon "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/common/v1"
	otlpresource "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/resource/v1"
	otlptrace "github.com/tigrannajaryan/exp-otelproto/encodings/experimental2/trace/v1"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

type dictElem struct {
	ref   uint32
	count int
}
type dictType map[string]dictElem

type translatorPass int

const countStringsPass translatorPass = 1
const replacePass translatorPass = 2

type spanTranslator struct {
	pass          translatorPass
	keyDict       dictType
	valDict       dictType
	spanNameDict  dictType
	eventNameDict dictType
}

func NewSpanTranslator() *spanTranslator {
	return &spanTranslator{}
}

func (st *spanTranslator) TranslateSpans(batch *otlptracecol.ExportTraceServiceRequest) core.ExportRequest {

	res := &otlptracecolexp.ExportTraceServiceRequest{}
	st.keyDict = dictType{}
	st.valDict = dictType{}
	st.spanNameDict = dictType{}
	st.eventNameDict = dictType{}

	st.pass = countStringsPass
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

	res = &otlptracecolexp.ExportTraceServiceRequest{}
	st.pass = replacePass
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

	return res
}

const firstStringRef = uint32(1)

func (st *spanTranslator) dictionizeStr(dict dictType, str *string, ref *uint32) bool {
	var idx dictElem
	var ok bool

	if st.pass == countStringsPass {
		if idx, ok = dict[*str]; !ok {
			idx.ref = uint32(len(dict) + 1)
			idx.count = 1
			dict[*str] = idx
		} else {
			idx.count++
			dict[*str] = idx
		}
		return false
	} else if st.pass == replacePass {
		if idx, ok = dict[*str]; !ok {
			return false
		}
		*str = ""
		*ref = idx.ref
		return true
	} else {
		return false
	}
}

func createDict(dict dictType) *otlpcommon.StringDict {

	var freqs []struct {
		str   string
		count int
	}
	for k, v := range dict {
		if v.count < 2 {
			// Not worth dictionary encoding since only one occurrence of this string.
			continue
		}
		freqs = append(
			freqs, struct {
				str   string
				count int
			}{str: k, count: v.count},
		)
	}

	sort.Slice(
		freqs, func(i, j int) bool {
			return freqs[i].count > freqs[j].count
		},
	)
	dict = dictType{}
	ref := firstStringRef
	for _, e := range freqs {

		buf := make([]byte, 10)
		n := binary.PutUvarint(buf, uint64(ref))
		if n > len(e.str) {
			// Not worth using ref since numeric encoding of the ref is larger than the string itself
			continue
		}
		dict[e.str] = dictElem{
			ref:   ref,
			count: e.count,
		}
		ref++
	}

	r := &otlpcommon.StringDict{
		StartIndex: firstStringRef,
		Values:     make([]string, len(dict)),
	}
	for k, v := range dict {
		r.Values[v.ref-firstStringRef] = k
	}

	return r
}

func (st *spanTranslator) translateAttrs(attrs []*v1.KeyValue) (r []*otlpcommon.KeyValue) {
	for _, attr := range attrs {
		var ref uint32
		var kv *otlpcommon.KeyValue
		if st.dictionizeStr(st.keyDict, &attr.Key, &ref) {
			kv = &otlpcommon.KeyValue{KeyRef: ref}
		} else {
			kv = &otlpcommon.KeyValue{Key: attr.Key}
		}

		var v *otlpcommon.AnyValue
		switch iv := attr.Value.Value.(type) {
		case *v1.AnyValue_StringValue:
			var ref uint32
			if st.dictionizeStr(st.valDict, &iv.StringValue, &ref) {
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
	st.dictionizeStr(st.spanNameDict, &s.Name, &s.NameRef)
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
	st.dictionizeStr(st.valDict, &is.Name, &is.NameRef)
	st.dictionizeStr(st.valDict, &is.Version, &is.VersionRef)
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
	e1 := &otlptrace.Span_Event{
		TimeUnixNano:           e.TimeUnixNano,
		Name:                   e.Name,
		Attributes:             st.translateAttrs(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
	st.dictionizeStr(st.eventNameDict, &e1.Name, &e1.NameRef)
	return e1
}

func (st *spanTranslator) translateStatus(status *v12.Status) *otlptrace.Status {
	return &otlptrace.Status{
		Message: status.Message,
		Code:    otlptrace.Status_StatusCode(status.Code),
	}
}
