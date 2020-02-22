package internal

import "github.com/tigrannajaryan/exp-otelproto/encodings/otlp"

// A collection of spans from a Resource.
type ResourceSpans struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Spans    []*Span   `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

// Resource information. This describes the source of telemetry data.
type Resource struct {
	// labels is a collection of attributes that describe the resource. See OpenTelemetry
	// specification semantic conventions for standardized label names:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-semantic-conventions.md
	Labels map[string]*otlp.AttributeKeyValue `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_labels_count is the number of dropped labels. If the value is 0, then
	// no labels were dropped.
	DroppedLabelsCount int32 `protobuf:"varint,2,opt,name=dropped_labels_count,json=droppedLabelsCount,proto3" json:"dropped_labels_count,omitempty"`
}

/*type AttributeKeyValue struct {
	// type of the value.
	Type        otlp.AttributeKeyValue_ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=experimental.AttributeKeyValue_ValueType" json:"type,omitempty"`
	StringValue string                           `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue    int64                            `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue float64                          `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue   bool                             `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
}*/

type Span struct {
	// A unique identifier for a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random trace_id if empty or invalid trace_id was received.
	//
	// This field is required.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random span_id if empty or invalid span_id was received.
	//
	// This field is required.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// tracestate conveys information about request position in multiple distributed tracing graphs.
	// It is a tracestate in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
	// See also https://github.com/w3c/distributed-tracing for more details about this field.
	Tracestate string `protobuf:"bytes,3,opt,name=tracestate,proto3" json:"tracestate,omitempty"`
	// The `span_id` of this span's parent span. If this is a root span, then this
	// field must be empty. The ID is an 8-byte array.
	ParentSpanId []byte `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// A description of the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	// When null or empty string received - receiver may use string "name"
	// as a replacement. There might be smarted algorithms implemented by
	// receiver to fix the empty span name.
	//
	// This field is required.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `CLIENT` (caller)
	// and `SERVER` (callee) to identify queueing latency associated with the span.
	Kind otlp.Span_SpanKind `protobuf:"varint,6,opt,name=kind,proto3,enum=experimental.Span_SpanKind" json:"kind,omitempty"`
	// start_time_unixnano is the start time of the span. On the client side, this is the time
	// kept by the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	StartTimeUnixnano uint64 `protobuf:"fixed64,7,opt,name=start_time_unixnano,json=startTimeUnixnano,proto3" json:"start_time_unixnano,omitempty"`
	// end_time_unixnano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	EndTimeUnixnano uint64 `protobuf:"fixed64,8,opt,name=end_time_unixnano,json=endTimeUnixnano,proto3" json:"end_time_unixnano,omitempty"`
	// attributes is a collection of key/value pairs. The value can be a string,
	// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "abc.com/myattribute": true
	//     "abc.com/score": 10.239
	Attributes map[string]*otlp.AttributeKeyValue `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_attributes_count is the number of attributes that were discarded. Attributes
	// can be discarded because their keys are too long or because there are too many
	// attributes. If this value is 0, then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,10,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// events is a collection of Event items.
	Events []*Span_Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`
	// dropped_events_count is the number of dropped events. If the value is 0, then no
	// events were dropped.
	DroppedEventsCount uint32 `protobuf:"varint,12,opt,name=dropped_events_count,json=droppedEventsCount,proto3" json:"dropped_events_count,omitempty"`
	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	Links []*Span_Link `protobuf:"bytes,13,rep,name=links,proto3" json:"links,omitempty"`
	// dropped_links_count is the number of dropped links after the maximum size was
	// enforced. If this value is 0, then no links were dropped.
	DroppedLinksCount uint32 `protobuf:"varint,14,opt,name=dropped_links_count,json=droppedLinksCount,proto3" json:"dropped_links_count,omitempty"`
	// An optional final status for this span. Semantically when Status
	// wasn't set it is means span ended without errors and assume
	// Status.Ok (code = 0).
	Status *otlp.Status `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	// An optional number of local child spans that were generated while this span
	// was active. Value of -1 indicates that the number of local child spans is unknown.
	// If local_child_span_count>=0, allows an implementation to detect missing child spans.
	LocalChildSpanCount int32 `protobuf:"fixed32,16,opt,name=local_child_span_count,json=localChildSpanCount,proto3" json:"local_child_span_count,omitempty"`
}

type Span_Event struct {
	// time_unixnano is the time the event occurred.
	TimeUnixnano uint64 `protobuf:"fixed64,1,opt,name=time_unixnano,json=timeUnixnano,proto3" json:"time_unixnano,omitempty"`
	// description is a user-supplied text.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// attributes is a collection of attribute key/value pairs on the event.
	Attributes map[string]*otlp.AttributeKeyValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,4,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
}

type Span_Link struct {
	// A unique identifier of a trace that this linked span is part of. The ID is a
	// 16-byte array.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for the linked span. The ID is an 8-byte array.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// The tracestate associated with the link.
	Tracestate string `protobuf:"bytes,3,opt,name=tracestate,proto3" json:"tracestate,omitempty"`
	// attributes is a collection of attribute key/value pairs on the link.
	Attributes map[string]*otlp.AttributeKeyValue `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,5,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
}

func FromOtlp(rs []*otlp.ResourceSpans) (r []*ResourceSpans) {
	r = make([]*ResourceSpans, len(rs))
	for i, s := range rs {
		r[i] = ResourceSpansFromOtlp(s)
	}
	return
}

func ResourceSpansFromOtlp(spans *otlp.ResourceSpans) *ResourceSpans {
	return &ResourceSpans{
		Resource: ResourceFromOtlp(spans.Resource),
		Spans:    SpansFromOtlp(spans.Spans),
	}
}

func ResourceFromOtlp(resource *otlp.Resource) *Resource {
	return &Resource{
		Labels:             AttrsFromOtlp(resource.Labels),
		DroppedLabelsCount: resource.DroppedLabelsCount,
	}
}

func AttrsFromOtlp(attrs []*otlp.AttributeKeyValue) (m map[string]*otlp.AttributeKeyValue) {
	m = make(map[string]*otlp.AttributeKeyValue, len(attrs))
	for _, a := range attrs {
		m[a.Key] = a
	}
	return
}

/*func AttrFromOtlp(a *otlp.AttributeKeyValue) *AttributeKeyValue {
	return &AttributeKeyValue{
		Type:        a.Type,
		StringValue: a.StringValue,
		IntValue:    a.IntValue,
		DoubleValue: a.DoubleValue,
		BoolValue:   a.BoolValue,
	}
}*/

func SpansFromOtlp(spans []*otlp.Span) (r []*Span) {
	r = make([]*Span, len(spans))
	for i, s := range spans {
		r[i] = SpanFromOtlp(s)
	}
	return
}

func SpanFromOtlp(s *otlp.Span) *Span {
	return &Span{
		TraceId:                s.TraceId,
		SpanId:                 s.SpanId,
		Tracestate:             s.Tracestate,
		ParentSpanId:           s.ParentSpanId,
		Name:                   s.Name,
		Kind:                   s.Kind,
		StartTimeUnixnano:      s.StartTimeUnixnano,
		EndTimeUnixnano:        s.EndTimeUnixnano,
		Attributes:             AttrsFromOtlp(s.Attributes),
		DroppedAttributesCount: s.DroppedAttributesCount,
		Events:                 EventsFromOtlp(s.Events),
		DroppedEventsCount:     s.DroppedEventsCount,
		Links:                  LinksFromOtlp(s.Links),
		DroppedLinksCount:      s.DroppedLinksCount,
		Status:                 s.Status,
		LocalChildSpanCount:    s.LocalChildSpanCount,
	}
}

func EventsFromOtlp(events []*otlp.Span_Event) (r []*Span_Event) {
	r = make([]*Span_Event, len(events))
	for i, e := range events {
		r[i] = EventFromOtlp(e)
	}
	return
}

func EventFromOtlp(e *otlp.Span_Event) *Span_Event {
	return &Span_Event{
		TimeUnixnano:           e.TimeUnixnano,
		Description:            e.Description,
		Attributes:             AttrsFromOtlp(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
}

func LinksFromOtlp(links []*otlp.Span_Link) (r []*Span_Link) {
	r = make([]*Span_Link, len(links))
	for i, e := range links {
		r[i] = LinkFromOtlp(e)
	}
	return
}

func LinkFromOtlp(l *otlp.Span_Link) *Span_Link {
	return &Span_Link{
		TraceId:                l.TraceId,
		SpanId:                 l.SpanId,
		Attributes:             AttrsFromOtlp(l.Attributes),
		DroppedAttributesCount: l.DroppedAttributesCount,
	}
}

func ToOtlp(rs []*ResourceSpans) (r []*otlp.ResourceSpans) {
	r = make([]*otlp.ResourceSpans, len(rs))
	for i, s := range rs {
		r[i] = ResourceSpansToOtlp(s)
	}
	return
}

func ResourceSpansToOtlp(spans *ResourceSpans) *otlp.ResourceSpans {
	return &otlp.ResourceSpans{
		Resource: ResourceToOtlp(spans.Resource),
		Spans:    SpansToOtlp(spans.Spans),
	}
}

func ResourceToOtlp(resource *Resource) *otlp.Resource {
	return &otlp.Resource{
		Labels:             AttrsToOtlp(resource.Labels),
		DroppedLabelsCount: resource.DroppedLabelsCount,
	}
}

func AttrsToOtlp(attrs map[string]*otlp.AttributeKeyValue) (m []*otlp.AttributeKeyValue) {
	m = make([]*otlp.AttributeKeyValue, 0, len(attrs))
	for _, a := range attrs {
		m = append(m, a)
		// a.Key = k
	}
	return
}

/*func AttrToOtlp(a *otlp.AttributeKeyValue) *AttributeKeyValue {
	return &AttributeKeyValue{
		Type:        a.Type,
		StringValue: a.StringValue,
		IntValue:    a.IntValue,
		DoubleValue: a.DoubleValue,
		BoolValue:   a.BoolValue,
	}
}*/

func SpansToOtlp(spans []*Span) (r []*otlp.Span) {
	r = make([]*otlp.Span, len(spans))
	for i, s := range spans {
		r[i] = SpanToOtlp(s)
	}
	return
}

func SpanToOtlp(s *Span) *otlp.Span {
	return &otlp.Span{
		TraceId:                s.TraceId,
		SpanId:                 s.SpanId,
		Tracestate:             s.Tracestate,
		ParentSpanId:           s.ParentSpanId,
		Name:                   s.Name,
		Kind:                   s.Kind,
		StartTimeUnixnano:      s.StartTimeUnixnano,
		EndTimeUnixnano:        s.EndTimeUnixnano,
		Attributes:             AttrsToOtlp(s.Attributes),
		DroppedAttributesCount: s.DroppedAttributesCount,
		Events:                 EventsToOtlp(s.Events),
		DroppedEventsCount:     s.DroppedEventsCount,
		Links:                  LinksToOtlp(s.Links),
		DroppedLinksCount:      s.DroppedLinksCount,
		Status:                 s.Status,
		LocalChildSpanCount:    s.LocalChildSpanCount,
	}
}

func EventsToOtlp(events []*Span_Event) (r []*otlp.Span_Event) {
	r = make([]*otlp.Span_Event, len(events))
	for i, e := range events {
		r[i] = EventToOtlp(e)
	}
	return
}

func EventToOtlp(e *Span_Event) *otlp.Span_Event {
	return &otlp.Span_Event{
		TimeUnixnano:           e.TimeUnixnano,
		Description:            e.Description,
		Attributes:             AttrsToOtlp(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
}

func LinksToOtlp(links []*Span_Link) (r []*otlp.Span_Link) {
	r = make([]*otlp.Span_Link, len(links))
	for i, e := range links {
		r[i] = LinkToOtlp(e)
	}
	return
}

func LinkToOtlp(l *Span_Link) *otlp.Span_Link {
	return &otlp.Span_Link{
		TraceId:                l.TraceId,
		SpanId:                 l.SpanId,
		Attributes:             AttrsToOtlp(l.Attributes),
		DroppedAttributesCount: l.DroppedAttributesCount,
	}
}
