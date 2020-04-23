package intotlp

import (
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

type TraceExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceSpans.
	ResourceSpans []*ResourceSpans `protobuf:"bytes,2,rep,name=resourceSpans,proto3" json:"resourceSpans,omitempty"`
}

// A collection of spans from a Resource.
type ResourceSpans struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Spans    []*Span   `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

type AttributesMap map[string]AttributeValue

// Resource information. This describes the source of telemetry data.
type Resource struct {
	orig *experimental.Resource
	// labels is a collection of attributes that describe the resource. See OpenTelemetry
	// specification semantic conventions for standardized label names:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-semantic-conventions.md
	Labels AttributesMap
}

type AttributeValue struct {
	orig *experimental.AttributeKeyValue
}

func (a AttributeValue) Type() experimental.AttributeKeyValue_ValueType {
	return a.orig.Type
}

func (a AttributeValue) String() string {
	return a.orig.StringValue
}

func (a AttributeValue) Int() int64 {
	return a.orig.IntValue
}

func (a AttributeValue) Double() float64 {
	return a.orig.DoubleValue
}

func (a AttributeValue) Bool() bool {
	return a.orig.BoolValue
}

type Span struct {
	orig *experimental.Span

	// attributes is a collection of key/value pairs. The value can be a string,
	// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "abc.com/myattribute": true
	//     "abc.com/score": 10.239
	attributes AttributesMap

	// events is a collection of Event items.
	events []*Span_Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`

	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	links []*Span_Link `protobuf:"bytes,13,rep,name=links,proto3" json:"links,omitempty"`
}

type Span_Event struct {
	orig *experimental.Span_Event

	// attributes is a collection of attribute key/value pairs on the event.
	attributes AttributesMap
}

type Span_Link struct {
	orig *experimental.Span_Link

	// attributes is a collection of attribute key/value pairs on the link.
	attributes AttributesMap
}

//func FromOtlp(tes *experimental.TraceExportRequest) *TraceExportRequest {
//	r := &TraceExportRequest{}
//	r.ResourceSpans = make([]*ResourceSpans, len(tes.ResourceSpans))
//	for i, s := range tes.ResourceSpans {
//		r.ResourceSpans[i] = ResourceSpansFromOtlp(s)
//	}
//	return r
//}

//func ResourceSpansFromOtlp(spans *experimental.ResourceSpans) *ResourceSpans {
//	return &ResourceSpans{
//		Resource: ResourceFromOtlp(spans.Resource),
//		Spans:    SpansFromOtlp(spans.Spans),
//	}
//}

func ResourceFromOtlp(resource *experimental.Resource) *Resource {
	return &Resource{
		orig:   resource,
		Labels: AttrsFromOtlp(resource.Attributes),
	}
}

func AttrsFromOtlp(attrs []*experimental.AttributeKeyValue) AttributesMap {
	m := make(AttributesMap, len(attrs))
	for _, attr := range attrs {
		m[attr.Key] = AttributeValue{attr}
	}
	return m
}

func SpansFromOtlp(spans []*experimental.Span) []*Span {
	ptrs := make([]*Span, len(spans))
	content := make([]Span, len(spans))
	for i, s := range spans {
		SpanFromOtlp(s, &content[i])
		ptrs[i] = &content[i]
	}
	return ptrs
}

func SpanFromOtlp(src *experimental.Span, dest *Span) {
	dest.orig = src
	dest.attributes = AttrsFromOtlp(src.Attributes)
	dest.events = EventsFromOtlp(src.Events)
	dest.links = LinksFromOtlp(src.Links)
}

func EventsFromOtlp(events []*experimental.Span_Event) []*Span_Event {
	r := make([]*Span_Event, len(events))
	for i, e := range events {
		r[i] = EventFromOtlp(e)
	}
	return r
}

func EventFromOtlp(e *experimental.Span_Event) *Span_Event {
	return &Span_Event{
		orig:       e,
		attributes: AttrsFromOtlp(e.Attributes),
	}
}

func LinksFromOtlp(links []*experimental.Span_Link) []*Span_Link {
	r := make([]*Span_Link, len(links))
	for i, e := range links {
		r[i] = LinkFromOtlp(e)
	}
	return r
}

func LinkFromOtlp(l *experimental.Span_Link) *Span_Link {
	return &Span_Link{
		orig:       l,
		attributes: AttrsFromOtlp(l.Attributes),
	}
}

//func ToOtlp(tes *TraceExportRequest) *experimental.TraceExportRequest {
//	r := make([]*experimental.ResourceSpans, len(tes.ResourceSpans))
//	for i, s := range tes.ResourceSpans {
//		r[i] = ResourceSpansToOtlp(s)
//	}
//	return &experimental.TraceExportRequest{
//		ResourceSpans: r,
//	}
//}

//func ResourceSpansToOtlp(spans *ResourceSpans) *experimental.ResourceSpans {
//	return &experimental.ResourceSpans{
//		Resource: ResourceToOtlp(spans.Resource),
//		Spans:    SpansToOtlp(spans.Spans),
//	}
//}

func ResourceToOtlp(resource *Resource) *experimental.Resource {
	AttrsToOtlp(&resource.orig.Attributes, resource.Labels)
	return resource.orig
}

func AttrsToOtlp(dest *[]*experimental.AttributeKeyValue, src AttributesMap) {
	if len(*dest) < len(src) {
		*dest = append(*dest, make([]*experimental.AttributeKeyValue, len(src)-len(*dest))...)
	}
	i := 0
	for key, attr := range src {
		// Key in the map is the source of truth.
		attr.orig.Key = key
		(*dest)[i] = attr.orig
		i++
	}
}

func SpansToOtlp(spans []*Span) []*experimental.Span {
	ptrs := make([]*experimental.Span, len(spans))
	for i, s := range spans {
		ptrs[i] = SpanToOtlp(s)
	}
	return ptrs
}

func SpanToOtlp(src *Span) *experimental.Span {
	AttrsToOtlp(&src.orig.Attributes, src.attributes)
	EventsToOtlp(&src.orig.Events, src.events)
	LinksToOtlp(&src.orig.Links, src.links)
	return src.orig
}

func EventsToOtlp(dest *[]*experimental.Span_Event, src []*Span_Event) {
	if len(*dest) < len(src) {
		*dest = append(*dest, make([]*experimental.Span_Event, len(src)-len(*dest))...)
	}
	for i, event := range src {
		AttrsToOtlp(&event.orig.Attributes, event.attributes)
		(*dest)[i] = event.orig
	}
}

func LinksToOtlp(dest *[]*experimental.Span_Link, src []*Span_Link) {
	if len(*dest) < len(src) {
		*dest = append(*dest, make([]*experimental.Span_Link, len(src)-len(*dest))...)
	}
	for i, link := range src {
		AttrsToOtlp(&link.orig.Attributes, link.attributes)
		(*dest)[i] = link.orig
	}
}