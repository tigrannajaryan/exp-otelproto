// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_data.proto

package otlp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ValueType is the enumeration of possible types that value can have.
type AttributeKeyValue_ValueType int32

const (
	AttributeKeyValue_STRING AttributeKeyValue_ValueType = 0
	AttributeKeyValue_BOOL   AttributeKeyValue_ValueType = 1
	AttributeKeyValue_INT64  AttributeKeyValue_ValueType = 2
	AttributeKeyValue_DOUBLE AttributeKeyValue_ValueType = 3
)

var AttributeKeyValue_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
	3: "DOUBLE",
}

var AttributeKeyValue_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
	"DOUBLE": 3,
}

func (x AttributeKeyValue_ValueType) String() string {
	return proto.EnumName(AttributeKeyValue_ValueType_name, int32(x))
}

func (AttributeKeyValue_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{2, 0}
}

// SpanKind is the type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type Span_SpanKind int32

const (
	// Unspecified. Do NOT use as default.
	// Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
	Span_SPAN_KIND_UNSPECIFIED Span_SpanKind = 0
	// Indicates that the span represents an internal operation within an application,
	// as opposed to an operations happening at the boundaries. Default value.
	Span_INTERNAL Span_SpanKind = 1
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	Span_SERVER Span_SpanKind = 2
	// Indicates that the span describes a request to some remote service.
	Span_CLIENT Span_SpanKind = 3
	// Indicates that the span describes a producer sending a message to a broker.
	// Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
	// between producer and consumer spans. A PRODUCER span ends when the message was accepted
	// by the broker while the logical processing of the message might span a much longer time.
	Span_PRODUCER Span_SpanKind = 4
	// Indicates that the span describes consumer receiving a message from a broker.
	// Like the PRODUCER kind, there is often no direct critical path latency relationship
	// between producer and consumer spans.
	Span_CONSUMER Span_SpanKind = 5
)

var Span_SpanKind_name = map[int32]string{
	0: "SPAN_KIND_UNSPECIFIED",
	1: "INTERNAL",
	2: "SERVER",
	3: "CLIENT",
	4: "PRODUCER",
	5: "CONSUMER",
}

var Span_SpanKind_value = map[string]int32{
	"SPAN_KIND_UNSPECIFIED": 0,
	"INTERNAL":              1,
	"SERVER":                2,
	"CLIENT":                3,
	"PRODUCER":              4,
	"CONSUMER":              5,
}

func (x Span_SpanKind) String() string {
	return proto.EnumName(Span_SpanKind_name, int32(x))
}

func (Span_SpanKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{3, 0}
}

// StatusCode mirrors the codes defined at
// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/api-tracing.md#statuscanonicalcode
type Status_StatusCode int32

const (
	Status_Ok                 Status_StatusCode = 0
	Status_Cancelled          Status_StatusCode = 1
	Status_UnknownError       Status_StatusCode = 2
	Status_InvalidArgument    Status_StatusCode = 3
	Status_DeadlineExceeded   Status_StatusCode = 4
	Status_NotFound           Status_StatusCode = 5
	Status_AlreadyExists      Status_StatusCode = 6
	Status_PermissionDenied   Status_StatusCode = 7
	Status_ResourceExhausted  Status_StatusCode = 8
	Status_FailedPrecondition Status_StatusCode = 9
	Status_Aborted            Status_StatusCode = 10
	Status_OutOfRange         Status_StatusCode = 11
	Status_Unimplemented      Status_StatusCode = 12
	Status_InternalError      Status_StatusCode = 13
	Status_Unavailable        Status_StatusCode = 14
	Status_DataLoss           Status_StatusCode = 15
	Status_Unauthenticated    Status_StatusCode = 16
)

var Status_StatusCode_name = map[int32]string{
	0:  "Ok",
	1:  "Cancelled",
	2:  "UnknownError",
	3:  "InvalidArgument",
	4:  "DeadlineExceeded",
	5:  "NotFound",
	6:  "AlreadyExists",
	7:  "PermissionDenied",
	8:  "ResourceExhausted",
	9:  "FailedPrecondition",
	10: "Aborted",
	11: "OutOfRange",
	12: "Unimplemented",
	13: "InternalError",
	14: "Unavailable",
	15: "DataLoss",
	16: "Unauthenticated",
}

var Status_StatusCode_value = map[string]int32{
	"Ok":                 0,
	"Cancelled":          1,
	"UnknownError":       2,
	"InvalidArgument":    3,
	"DeadlineExceeded":   4,
	"NotFound":           5,
	"AlreadyExists":      6,
	"PermissionDenied":   7,
	"ResourceExhausted":  8,
	"FailedPrecondition": 9,
	"Aborted":            10,
	"OutOfRange":         11,
	"Unimplemented":      12,
	"InternalError":      13,
	"Unavailable":        14,
	"DataLoss":           15,
	"Unauthenticated":    16,
}

func (x Status_StatusCode) String() string {
	return proto.EnumName(Status_StatusCode_name, int32(x))
}

func (Status_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{4, 0}
}

// A collection of spans from a Resource.
type ResourceSpans struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Spans                []*Span   `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceSpans) Reset()         { *m = ResourceSpans{} }
func (m *ResourceSpans) String() string { return proto.CompactTextString(m) }
func (*ResourceSpans) ProtoMessage()    {}
func (*ResourceSpans) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{0}
}

func (m *ResourceSpans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceSpans.Unmarshal(m, b)
}
func (m *ResourceSpans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceSpans.Marshal(b, m, deterministic)
}
func (m *ResourceSpans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceSpans.Merge(m, src)
}
func (m *ResourceSpans) XXX_Size() int {
	return xxx_messageInfo_ResourceSpans.Size(m)
}
func (m *ResourceSpans) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceSpans.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceSpans proto.InternalMessageInfo

func (m *ResourceSpans) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceSpans) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

// Resource information. This describes the source of telemetry data.
type Resource struct {
	// labels is a collection of attributes that describe the resource. See OpenTelemetry
	// specification semantic conventions for standardized label names:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-semantic-conventions.md
	Labels []*AttributeKeyValue `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// dropped_labels_count is the number of dropped labels. If the value is 0, then
	// no labels were dropped.
	DroppedLabelsCount   int32    `protobuf:"varint,2,opt,name=dropped_labels_count,json=droppedLabelsCount,proto3" json:"dropped_labels_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{1}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetLabels() []*AttributeKeyValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Resource) GetDroppedLabelsCount() int32 {
	if m != nil {
		return m.DroppedLabelsCount
	}
	return 0
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Resource
// labels, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type AttributeKeyValue_ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=otlp.AttributeKeyValue_ValueType" json:"type,omitempty"`
	// A string up to 256 bytes long.
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	// A 64-bit signed integer.
	Int64Value int64 `protobuf:"varint,4,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	// A Boolean value represented by `true` or `false`.
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	// A double value.
	DoubleValue          float64  `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{2}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() AttributeKeyValue_ValueType {
	if m != nil {
		return m.Type
	}
	return AttributeKeyValue_STRING
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

// Span represents a single operation within a trace. Spans can be
// nested to form a trace tree. Spans may also be linked to other spans
// from the same or different trace and form graphs. Often, a trace
// contains a root span that describes the end-to-end latency, and one
// or more subspans for its sub-operations. A trace can also contain
// multiple root spans, or none at all. Spans do not need to be
// contiguous - there may be gaps or overlaps between spans in a trace.
//
// The next field id is 18.
type Span struct {
	// trace_id is the unique identifier of a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid.
	//
	// This field is semantically required. If empty or invalid trace_id is received:
	// - The receiver MAY reject the invalid data and respond with the appropriate error
	//   code to the sender.
	// - The receiver MAY accept the invalid data and attempt to correct it.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// span_id is a unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid.
	//
	// This field is semantically required. If empty or invalid span_id is received:
	// - The receiver MAY reject the invalid data and respond with the appropriate error
	//   code to the sender.
	// - The receiver MAY accept the invalid data and attempt to correct it.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// tracestate conveys information about request position in multiple distributed tracing graphs.
	// It is a collection of TracestateEntry with a maximum of 32 members in the collection.
	//
	// See the https://github.com/w3c/distributed-tracing for more details about this field.
	Tracestate []*Span_TraceStateEntry `protobuf:"bytes,3,rep,name=tracestate,proto3" json:"tracestate,omitempty"`
	// parent_span_id is the `span_id` of this span's parent span. If this is a root span, then this
	// field must be omitted. The ID is an 8-byte array.
	ParentSpanId []byte `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// resource that is associated with this span. Optional. If not set, this span
	// should be part of a ResourceSpans message that does include the resource information,
	// unless resource information is unknown.
	Resource *Resource `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
	// name describes the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	//
	// This field is required.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// kind field distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `CLIENT` (caller)
	// and `SERVER` (callee) to identify network latency associated with the span.
	Kind Span_SpanKind `protobuf:"varint,7,opt,name=kind,proto3,enum=otlp.Span_SpanKind" json:"kind,omitempty"`
	// start_time_unixnano is the start time of the span. On the client side, this is the time
	// kept by the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	//
	// This field is required.
	StartTimeUnixnano int64 `protobuf:"varint,8,opt,name=start_time_unixnano,json=startTimeUnixnano,proto3" json:"start_time_unixnano,omitempty"`
	// end_time_unixnano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	//
	// This field is required.
	EndTimeUnixnano int64 `protobuf:"varint,9,opt,name=end_time_unixnano,json=endTimeUnixnano,proto3" json:"end_time_unixnano,omitempty"`
	// attributes is a collection of key/value pairs. The value can be a string,
	// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "abc.com/myattribute": true
	//     "abc.com/score": 10.239
	Attributes []*AttributeKeyValue `protobuf:"bytes,10,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of attributes that were discarded. Attributes
	// can be discarded because their keys are too long or because there are too many
	// attributes. If this value is 0, then no attributes were dropped.
	DroppedAttributesCount int32 `protobuf:"varint,11,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// timed_events is a collection of TimedEvent items.
	TimedEvents []*Span_TimedEvent `protobuf:"bytes,12,rep,name=timed_events,json=timedEvents,proto3" json:"timed_events,omitempty"`
	// dropped_timed_events_count is the number of dropped timed events. If the value is 0,
	// then no events were dropped.
	DroppedTimedEventsCount int32 `protobuf:"varint,13,opt,name=dropped_timed_events_count,json=droppedTimedEventsCount,proto3" json:"dropped_timed_events_count,omitempty"`
	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	Links []*Span_Link `protobuf:"bytes,14,rep,name=links,proto3" json:"links,omitempty"`
	// dropped_links_count is the number of dropped links after the maximum size was
	// enforced. If this value is 0, then no links were dropped.
	DroppedLinksCount int32 `protobuf:"varint,15,opt,name=dropped_links_count,json=droppedLinksCount,proto3" json:"dropped_links_count,omitempty"`
	// status is an optional final status for this span. Semantically when status
	// wasn't set it is means span ended without errors and assume Status.Ok (code = 0).
	Status *Status `protobuf:"bytes,16,opt,name=status,proto3" json:"status,omitempty"`
	// child_span_count is an optional number of local child spans that were generated while this
	// span was active. If set, allows an implementation to detect missing child spans.
	ChildSpanCount       int32    `protobuf:"varint,17,opt,name=child_span_count,json=childSpanCount,proto3" json:"child_span_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{3}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Span) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Span) GetTracestate() []*Span_TraceStateEntry {
	if m != nil {
		return m.Tracestate
	}
	return nil
}

func (m *Span) GetParentSpanId() []byte {
	if m != nil {
		return m.ParentSpanId
	}
	return nil
}

func (m *Span) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetKind() Span_SpanKind {
	if m != nil {
		return m.Kind
	}
	return Span_SPAN_KIND_UNSPECIFIED
}

func (m *Span) GetStartTimeUnixnano() int64 {
	if m != nil {
		return m.StartTimeUnixnano
	}
	return 0
}

func (m *Span) GetEndTimeUnixnano() int64 {
	if m != nil {
		return m.EndTimeUnixnano
	}
	return 0
}

func (m *Span) GetAttributes() []*AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span) GetDroppedAttributesCount() int32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *Span) GetTimedEvents() []*Span_TimedEvent {
	if m != nil {
		return m.TimedEvents
	}
	return nil
}

func (m *Span) GetDroppedTimedEventsCount() int32 {
	if m != nil {
		return m.DroppedTimedEventsCount
	}
	return 0
}

func (m *Span) GetLinks() []*Span_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Span) GetDroppedLinksCount() int32 {
	if m != nil {
		return m.DroppedLinksCount
	}
	return 0
}

func (m *Span) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Span) GetChildSpanCount() int32 {
	if m != nil {
		return m.ChildSpanCount
	}
	return 0
}

// TraceStateEntry is the entry that is repeated in tracestate field (see below).
type Span_TraceStateEntry struct {
	// key must begin with a lowercase letter, and can only contain
	// lowercase letters 'a'-'z', digits '0'-'9', underscores '_', dashes
	// '-', asterisks '*', and forward slashes '/'.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is opaque string up to 256 characters printable ASCII
	// RFC0020 characters (i.e., the range 0x20 to 0x7E) except ',' and '='.
	// Note that this also excludes tabs, newlines, carriage returns, etc.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span_TraceStateEntry) Reset()         { *m = Span_TraceStateEntry{} }
func (m *Span_TraceStateEntry) String() string { return proto.CompactTextString(m) }
func (*Span_TraceStateEntry) ProtoMessage()    {}
func (*Span_TraceStateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{3, 0}
}

func (m *Span_TraceStateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_TraceStateEntry.Unmarshal(m, b)
}
func (m *Span_TraceStateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_TraceStateEntry.Marshal(b, m, deterministic)
}
func (m *Span_TraceStateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_TraceStateEntry.Merge(m, src)
}
func (m *Span_TraceStateEntry) XXX_Size() int {
	return xxx_messageInfo_Span_TraceStateEntry.Size(m)
}
func (m *Span_TraceStateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_TraceStateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Span_TraceStateEntry proto.InternalMessageInfo

func (m *Span_TraceStateEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Span_TraceStateEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// TimedEvent is a time-stamped annotation of the span, consisting of either
// user-supplied key-value pairs, or details of a message sent/received between Spans.
type Span_TimedEvent struct {
	// time_unixnano is the time the event occurred.
	TimeUnixnano int64 `protobuf:"varint,1,opt,name=time_unixnano,json=timeUnixnano,proto3" json:"time_unixnano,omitempty"`
	// name is a user-supplied description of the event.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// attributes is a collection of attribute key/value pairs on the event.
	Attributes []*AttributeKeyValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount int32    `protobuf:"varint,4,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Span_TimedEvent) Reset()         { *m = Span_TimedEvent{} }
func (m *Span_TimedEvent) String() string { return proto.CompactTextString(m) }
func (*Span_TimedEvent) ProtoMessage()    {}
func (*Span_TimedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{3, 1}
}

func (m *Span_TimedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_TimedEvent.Unmarshal(m, b)
}
func (m *Span_TimedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_TimedEvent.Marshal(b, m, deterministic)
}
func (m *Span_TimedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_TimedEvent.Merge(m, src)
}
func (m *Span_TimedEvent) XXX_Size() int {
	return xxx_messageInfo_Span_TimedEvent.Size(m)
}
func (m *Span_TimedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_TimedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_Span_TimedEvent proto.InternalMessageInfo

func (m *Span_TimedEvent) GetTimeUnixnano() int64 {
	if m != nil {
		return m.TimeUnixnano
	}
	return 0
}

func (m *Span_TimedEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span_TimedEvent) GetAttributes() []*AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span_TimedEvent) GetDroppedAttributesCount() int32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// Link is a pointer from the current span to another span in the same trace or in a
// different trace. For example, this can be used in batching operations,
// where a single batch handler processes multiple requests from different
// traces or when the handler receives a request from a different project.
// See also Links specification:
// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/overview.md#links-between-spans
type Span_Link struct {
	// trace_id is a unique identifier of a trace that this linked span is part of.
	// The ID is a 16-byte array.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// span_id is a unique identifier for the linked span. The ID is an 8-byte array.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// tracestate is the trace state associated with the link.
	Tracestate []*Span_TraceStateEntry `protobuf:"bytes,3,rep,name=tracestate,proto3" json:"tracestate,omitempty"`
	// attributes is a collection of attribute key/value pairs on the link.
	Attributes []*AttributeKeyValue `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount int32    `protobuf:"varint,5,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Span_Link) Reset()         { *m = Span_Link{} }
func (m *Span_Link) String() string { return proto.CompactTextString(m) }
func (*Span_Link) ProtoMessage()    {}
func (*Span_Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{3, 2}
}

func (m *Span_Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Link.Unmarshal(m, b)
}
func (m *Span_Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Link.Marshal(b, m, deterministic)
}
func (m *Span_Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Link.Merge(m, src)
}
func (m *Span_Link) XXX_Size() int {
	return xxx_messageInfo_Span_Link.Size(m)
}
func (m *Span_Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Link proto.InternalMessageInfo

func (m *Span_Link) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Span_Link) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Span_Link) GetTracestate() []*Span_TraceStateEntry {
	if m != nil {
		return m.Tracestate
	}
	return nil
}

func (m *Span_Link) GetAttributes() []*AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span_Link) GetDroppedAttributesCount() int32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// The Status type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs.
type Status struct {
	// The status code. This is optional field. It is safe to assume 0 (OK)
	// when not set.
	Code Status_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=otlp.Status_StatusCode" json:"code,omitempty"`
	// A developer-facing error message, which should be in English.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_9161fdd1e0292445, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Status_StatusCode {
	if m != nil {
		return m.Code
	}
	return Status_Ok
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("otlp.AttributeKeyValue_ValueType", AttributeKeyValue_ValueType_name, AttributeKeyValue_ValueType_value)
	proto.RegisterEnum("otlp.Span_SpanKind", Span_SpanKind_name, Span_SpanKind_value)
	proto.RegisterEnum("otlp.Status_StatusCode", Status_StatusCode_name, Status_StatusCode_value)
	proto.RegisterType((*ResourceSpans)(nil), "otlp.ResourceSpans")
	proto.RegisterType((*Resource)(nil), "otlp.Resource")
	proto.RegisterType((*AttributeKeyValue)(nil), "otlp.AttributeKeyValue")
	proto.RegisterType((*Span)(nil), "otlp.Span")
	proto.RegisterType((*Span_TraceStateEntry)(nil), "otlp.Span.TraceStateEntry")
	proto.RegisterType((*Span_TimedEvent)(nil), "otlp.Span.TimedEvent")
	proto.RegisterType((*Span_Link)(nil), "otlp.Span.Link")
	proto.RegisterType((*Status)(nil), "otlp.Status")
}

func init() { proto.RegisterFile("telemetry_data.proto", fileDescriptor_9161fdd1e0292445) }

var fileDescriptor_9161fdd1e0292445 = []byte{
	// 1122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x92, 0x1a, 0x45,
	0x14, 0xce, 0xc0, 0xc0, 0xc2, 0xe1, 0xaf, 0xe9, 0x6c, 0x92, 0xc9, 0x96, 0x96, 0x04, 0x63, 0x49,
	0xc5, 0x2a, 0x34, 0x31, 0xc6, 0xa8, 0x57, 0x2c, 0x4c, 0x94, 0x0a, 0x02, 0xd5, 0x40, 0xee, 0x2c,
	0xaa, 0xa1, 0xdb, 0x4d, 0x17, 0x43, 0x0f, 0x35, 0xd3, 0xb3, 0x2e, 0x4f, 0xe1, 0x0b, 0xf8, 0x20,
	0xde, 0xe4, 0x81, 0xf4, 0x29, 0xac, 0xee, 0x9e, 0x61, 0xd9, 0xa8, 0xf1, 0x26, 0xe5, 0xcd, 0x30,
	0x7d, 0xbe, 0xaf, 0xbf, 0xd3, 0xa7, 0xcf, 0xd7, 0xcd, 0xc0, 0xa9, 0xe2, 0x01, 0xdf, 0x72, 0x15,
	0xed, 0x97, 0x8c, 0x2a, 0xda, 0xdd, 0x45, 0xa1, 0x0a, 0xb1, 0x1b, 0xaa, 0x60, 0xd7, 0xfe, 0x09,
	0x6a, 0x84, 0xc7, 0x61, 0x12, 0xad, 0xf9, 0x6c, 0x47, 0x65, 0x8c, 0x1f, 0x41, 0x29, 0x4a, 0x03,
	0x9e, 0xd3, 0x72, 0x3a, 0x95, 0x27, 0xf5, 0xae, 0x66, 0x76, 0x33, 0x1a, 0x39, 0xe0, 0xb8, 0x05,
	0x85, 0x58, 0x4f, 0xf2, 0x72, 0xad, 0x7c, 0xa7, 0xf2, 0x04, 0x2c, 0x51, 0xeb, 0x10, 0x0b, 0xb4,
	0xb7, 0x50, 0xca, 0xe6, 0xe1, 0xcf, 0xa1, 0x18, 0xd0, 0x15, 0x0f, 0x62, 0xcf, 0x31, 0xf4, 0x7b,
	0x96, 0xde, 0x53, 0x2a, 0x12, 0xab, 0x44, 0xf1, 0x97, 0x7c, 0xff, 0x8a, 0x06, 0x09, 0x27, 0x29,
	0x0d, 0x7f, 0x01, 0xa7, 0x2c, 0x0a, 0x77, 0x3b, 0xce, 0x96, 0x36, 0xb2, 0x5c, 0x87, 0x89, 0x54,
	0x5e, 0xae, 0xe5, 0x74, 0x0a, 0x04, 0xa7, 0xd8, 0xc8, 0x40, 0x7d, 0x8d, 0xb4, 0x7f, 0xcb, 0x41,
	0xf3, 0x6f, 0x7a, 0x18, 0x41, 0x7e, 0xc3, 0xf7, 0xa6, 0x9a, 0x32, 0xd1, 0xaf, 0xf8, 0x2b, 0x70,
	0xd5, 0x7e, 0xc7, 0x8d, 0x52, 0xfd, 0xc9, 0x83, 0x7f, 0x59, 0x48, 0xd7, 0x3c, 0xe7, 0xfb, 0x1d,
	0x27, 0x86, 0x8e, 0x1f, 0x40, 0x35, 0x56, 0x91, 0x90, 0x17, 0xcb, 0x4b, 0x8d, 0x78, 0x79, 0xa3,
	0x58, 0xb1, 0x31, 0x9b, 0xeb, 0x23, 0xa8, 0x08, 0xa9, 0x9e, 0x3d, 0x4d, 0x19, 0x6e, 0xcb, 0xe9,
	0xe4, 0x09, 0x98, 0x90, 0x25, 0x7c, 0x08, 0xb0, 0x0a, 0xc3, 0x20, 0xc5, 0x0b, 0x2d, 0xa7, 0x53,
	0x22, 0x65, 0x1d, 0xb1, 0xf0, 0x03, 0xa8, 0xb2, 0x30, 0x59, 0x05, 0x3c, 0x25, 0x14, 0x5b, 0x4e,
	0xc7, 0x21, 0x15, 0x1b, 0x33, 0x94, 0xf6, 0x73, 0x28, 0x1f, 0x16, 0x86, 0x01, 0x8a, 0xb3, 0x39,
	0x19, 0x8e, 0xbf, 0x47, 0xb7, 0x70, 0x09, 0xdc, 0xf3, 0xc9, 0x64, 0x84, 0x1c, 0x5c, 0x86, 0xc2,
	0x70, 0x3c, 0x7f, 0xf6, 0x14, 0xe5, 0x34, 0x61, 0x30, 0x59, 0x9c, 0x8f, 0x7c, 0x94, 0x6f, 0xbf,
	0x01, 0x70, 0x75, 0x77, 0xf0, 0x7d, 0x28, 0xa9, 0x88, 0xae, 0xf9, 0x52, 0x30, 0xb3, 0x2d, 0x55,
	0x72, 0x62, 0xc6, 0x43, 0x86, 0xef, 0xc1, 0x89, 0x6e, 0x9d, 0x46, 0x72, 0x06, 0x29, 0xea, 0xe1,
	0x90, 0xe1, 0x6f, 0x01, 0x0c, 0x27, 0x56, 0x54, 0xe9, 0xd2, 0x75, 0x0b, 0xcf, 0xae, 0x3b, 0xde,
	0x9d, 0x6b, 0x70, 0xa6, 0x41, 0x5f, 0xaa, 0x68, 0x4f, 0x8e, 0xd8, 0xf8, 0x21, 0xd4, 0x77, 0x34,
	0xe2, 0x52, 0x2d, 0x33, 0x6d, 0xd7, 0x68, 0x57, 0x6d, 0x74, 0x66, 0x33, 0x1c, 0x5b, 0xaf, 0xf0,
	0x1f, 0xd6, 0xc3, 0xe0, 0x4a, 0xba, 0xb5, 0xfb, 0x53, 0x26, 0xe6, 0x1d, 0x7f, 0x0a, 0xee, 0x46,
	0x48, 0xe6, 0x9d, 0x98, 0xae, 0xde, 0x3e, 0x5a, 0x9b, 0x7e, 0xbc, 0x14, 0x92, 0x11, 0x43, 0xc0,
	0x5d, 0xb8, 0x1d, 0x2b, 0x1a, 0xa9, 0xa5, 0x12, 0x5b, 0xbe, 0x4c, 0xa4, 0xb8, 0x92, 0x54, 0x86,
	0x5e, 0xc9, 0x34, 0xab, 0x69, 0xa0, 0xb9, 0xd8, 0xf2, 0x45, 0x0a, 0xe0, 0x47, 0xd0, 0xe4, 0x92,
	0xbd, 0xc5, 0x2e, 0x1b, 0x76, 0x83, 0x4b, 0x76, 0x83, 0xfb, 0x35, 0x00, 0xcd, 0x8c, 0x14, 0x7b,
	0xf0, 0x6e, 0xa7, 0x1f, 0x51, 0xf1, 0x73, 0xf0, 0x32, 0xb7, 0x5f, 0x47, 0x53, 0xc7, 0x57, 0x8c,
	0xe3, 0xef, 0xa6, 0xf8, 0x41, 0xc7, 0xba, 0x1e, 0x3f, 0x87, 0xaa, 0x5e, 0x1a, 0x5b, 0xf2, 0x4b,
	0x2e, 0x55, 0xec, 0x55, 0x4d, 0xd2, 0x3b, 0xc7, 0xbd, 0xd1, 0xb0, 0xaf, 0x51, 0x52, 0x51, 0x87,
	0xf7, 0x18, 0x7f, 0x07, 0x67, 0x59, 0xce, 0x63, 0x85, 0x34, 0x6b, 0xcd, 0x64, 0xbd, 0x97, 0x32,
	0xae, 0x35, 0xd2, 0xb4, 0x9f, 0x40, 0x21, 0x10, 0x72, 0x13, 0x7b, 0x75, 0x93, 0xaf, 0x71, 0x94,
	0x6f, 0x24, 0xe4, 0x86, 0x58, 0x54, 0x6f, 0xf6, 0xe1, 0x14, 0xeb, 0x40, 0x2a, 0xde, 0x30, 0xe2,
	0xcd, 0xec, 0x10, 0x6b, 0xc4, 0xca, 0x3e, 0x84, 0xa2, 0x36, 0x4d, 0x12, 0x7b, 0xc8, 0x78, 0xa0,
	0x9a, 0xea, 0x9a, 0x18, 0x49, 0x31, 0xdc, 0x01, 0xb4, 0x7e, 0x2d, 0x02, 0x66, 0x0d, 0x65, 0x25,
	0x9b, 0x46, 0xb2, 0x6e, 0xe2, 0x7a, 0x19, 0x46, 0xef, 0xec, 0x1b, 0x68, 0xbc, 0x65, 0xcd, 0x7f,
	0xb8, 0x10, 0x4e, 0xa1, 0x60, 0xcf, 0x5b, 0xce, 0xc4, 0xec, 0xe0, 0xec, 0x77, 0x07, 0xe0, 0xba,
	0x6c, 0xfc, 0x31, 0xd4, 0x6e, 0x5a, 0xc0, 0x31, 0x16, 0x30, 0x9b, 0x7f, 0xe8, 0x7f, 0x66, 0xcc,
	0xdc, 0x91, 0x31, 0x6f, 0x7a, 0x22, 0xff, 0x7e, 0x3c, 0xe1, 0xbe, 0xcb, 0x13, 0x67, 0x7f, 0x38,
	0xe0, 0xea, 0x4d, 0xfd, 0xdf, 0x8f, 0xfa, 0xcd, 0x5a, 0xdd, 0xf7, 0x53, 0x6b, 0xe1, 0x5d, 0xb5,
	0xb6, 0x2f, 0xa0, 0x94, 0x1d, 0x70, 0x7c, 0x1f, 0xee, 0xcc, 0xa6, 0xbd, 0xf1, 0xf2, 0xe5, 0x70,
	0x3c, 0x58, 0x2e, 0xc6, 0xb3, 0xa9, 0xdf, 0x1f, 0xbe, 0x18, 0xfa, 0x03, 0x74, 0x0b, 0x57, 0xa1,
	0x34, 0x1c, 0xcf, 0x7d, 0x32, 0xee, 0xe9, 0x2b, 0x52, 0x5f, 0x9c, 0x3e, 0x79, 0xe5, 0x13, 0x7b,
	0x47, 0xf6, 0x47, 0x43, 0x7f, 0x3c, 0x47, 0x79, 0xcd, 0x9a, 0x92, 0xc9, 0x60, 0xd1, 0xf7, 0x09,
	0x72, 0xf5, 0xa8, 0x3f, 0x19, 0xcf, 0x16, 0x3f, 0xfa, 0x04, 0x15, 0xda, 0xbf, 0xe6, 0xa1, 0x68,
	0x7d, 0x88, 0x3f, 0x03, 0x77, 0x1d, 0x32, 0xfb, 0x17, 0x59, 0xcf, 0x0a, 0xb4, 0x58, 0xfa, 0xd3,
	0x0f, 0x19, 0x27, 0x86, 0x84, 0x3d, 0x38, 0xd9, 0xf2, 0x38, 0xa6, 0x17, 0x99, 0x2d, 0xb2, 0x61,
	0xfb, 0x4d, 0x0e, 0xe0, 0x9a, 0x8e, 0x8b, 0x90, 0x9b, 0x6c, 0xd0, 0x2d, 0x5c, 0x83, 0x72, 0x9f,
	0xca, 0x35, 0x0f, 0x02, 0xce, 0x90, 0x83, 0x11, 0x54, 0x17, 0x72, 0x23, 0xc3, 0x5f, 0xa4, 0x1f,
	0x45, 0x61, 0x84, 0x72, 0xf8, 0x36, 0x34, 0x86, 0xf2, 0x92, 0x06, 0x82, 0xf5, 0xa2, 0x8b, 0x64,
	0xcb, 0xa5, 0x42, 0x79, 0x7c, 0x0a, 0x68, 0xc0, 0x29, 0x0b, 0x84, 0xe4, 0xfe, 0xd5, 0x9a, 0x73,
	0xc6, 0x99, 0x2d, 0x61, 0x1c, 0xaa, 0x17, 0x61, 0x22, 0x19, 0x2a, 0xe0, 0x26, 0xd4, 0x7a, 0x41,
	0xc4, 0x29, 0xdb, 0xfb, 0x57, 0x22, 0x56, 0x31, 0x2a, 0xea, 0x69, 0x53, 0x1e, 0x6d, 0x45, 0x1c,
	0x8b, 0x50, 0x0e, 0xb8, 0x14, 0x9c, 0xa1, 0x13, 0x7c, 0x07, 0x9a, 0xd9, 0xb5, 0xeb, 0x5f, 0xbd,
	0xa6, 0x49, 0xac, 0x38, 0x43, 0x25, 0x7c, 0x17, 0xf0, 0x0b, 0x2a, 0x02, 0xce, 0xa6, 0x11, 0x5f,
	0x87, 0x92, 0x09, 0x25, 0x42, 0x89, 0xca, 0xb8, 0x02, 0x27, 0xbd, 0x55, 0x18, 0x69, 0x12, 0xe0,
	0x3a, 0xc0, 0x24, 0x51, 0x93, 0x9f, 0x09, 0x95, 0x17, 0x1c, 0x55, 0x74, 0xd2, 0x85, 0x14, 0xdb,
	0x9d, 0xfe, 0x0a, 0x91, 0x9a, 0x52, 0xd5, 0xa1, 0xa1, 0x54, 0x3c, 0x92, 0x34, 0xb0, 0x35, 0xd5,
	0x70, 0x03, 0x2a, 0x0b, 0x49, 0x2f, 0xa9, 0x08, 0xe8, 0x2a, 0xe0, 0xa8, 0xae, 0x57, 0x3e, 0xa0,
	0x8a, 0x8e, 0xc2, 0x38, 0x46, 0x0d, 0x5d, 0xf2, 0x42, 0xd2, 0x44, 0xbd, 0xe6, 0x52, 0x89, 0x35,
	0xd5, 0x32, 0xe8, 0xfc, 0x07, 0xf8, 0x40, 0x84, 0xdd, 0x70, 0xc7, 0xe5, 0x9a, 0xcb, 0x38, 0x89,
	0xed, 0xb7, 0x4d, 0xd7, 0xf8, 0xb1, 0x7b, 0xf9, 0xf8, 0x1c, 0x8c, 0x55, 0xa7, 0x3a, 0x38, 0x75,
	0xfe, 0xcc, 0xdd, 0x9f, 0xec, 0xb8, 0xec, 0x5b, 0xa6, 0x09, 0x5a, 0x2b, 0x77, 0x5f, 0x3d, 0x5e,
	0x15, 0xcd, 0xcc, 0x2f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x12, 0x68, 0x74, 0x44, 0x2d, 0x09,
	0x00, 0x00,
}
