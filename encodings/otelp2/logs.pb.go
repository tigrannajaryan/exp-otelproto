// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logs.proto

package otelp2

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

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST not be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "SEVERITY_NUMBER_UNSPECIFIED",
	1:  "SEVERITY_NUMBER_TRACE",
	2:  "SEVERITY_NUMBER_TRACE2",
	3:  "SEVERITY_NUMBER_TRACE3",
	4:  "SEVERITY_NUMBER_TRACE4",
	5:  "SEVERITY_NUMBER_DEBUG",
	6:  "SEVERITY_NUMBER_DEBUG2",
	7:  "SEVERITY_NUMBER_DEBUG3",
	8:  "SEVERITY_NUMBER_DEBUG4",
	9:  "SEVERITY_NUMBER_INFO",
	10: "SEVERITY_NUMBER_INFO2",
	11: "SEVERITY_NUMBER_INFO3",
	12: "SEVERITY_NUMBER_INFO4",
	13: "SEVERITY_NUMBER_WARN",
	14: "SEVERITY_NUMBER_WARN2",
	15: "SEVERITY_NUMBER_WARN3",
	16: "SEVERITY_NUMBER_WARN4",
	17: "SEVERITY_NUMBER_ERROR",
	18: "SEVERITY_NUMBER_ERROR2",
	19: "SEVERITY_NUMBER_ERROR3",
	20: "SEVERITY_NUMBER_ERROR4",
	21: "SEVERITY_NUMBER_FATAL",
	22: "SEVERITY_NUMBER_FATAL2",
	23: "SEVERITY_NUMBER_FATAL3",
	24: "SEVERITY_NUMBER_FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"SEVERITY_NUMBER_UNSPECIFIED": 0,
	"SEVERITY_NUMBER_TRACE":       1,
	"SEVERITY_NUMBER_TRACE2":      2,
	"SEVERITY_NUMBER_TRACE3":      3,
	"SEVERITY_NUMBER_TRACE4":      4,
	"SEVERITY_NUMBER_DEBUG":       5,
	"SEVERITY_NUMBER_DEBUG2":      6,
	"SEVERITY_NUMBER_DEBUG3":      7,
	"SEVERITY_NUMBER_DEBUG4":      8,
	"SEVERITY_NUMBER_INFO":        9,
	"SEVERITY_NUMBER_INFO2":       10,
	"SEVERITY_NUMBER_INFO3":       11,
	"SEVERITY_NUMBER_INFO4":       12,
	"SEVERITY_NUMBER_WARN":        13,
	"SEVERITY_NUMBER_WARN2":       14,
	"SEVERITY_NUMBER_WARN3":       15,
	"SEVERITY_NUMBER_WARN4":       16,
	"SEVERITY_NUMBER_ERROR":       17,
	"SEVERITY_NUMBER_ERROR2":      18,
	"SEVERITY_NUMBER_ERROR3":      19,
	"SEVERITY_NUMBER_ERROR4":      20,
	"SEVERITY_NUMBER_FATAL":       21,
	"SEVERITY_NUMBER_FATAL2":      22,
	"SEVERITY_NUMBER_FATAL3":      23,
	"SEVERITY_NUMBER_FATAL4":      24,
}

func (x SeverityNumber) String() string {
	return proto.EnumName(SeverityNumber_name, int32(x))
}

func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_LOG_RECORD_FLAG_UNSPECIFIED      LogRecordFlags = 0
	LogRecordFlags_LOG_RECORD_FLAG_TRACE_FLAGS_MASK LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "LOG_RECORD_FLAG_UNSPECIFIED",
	255: "LOG_RECORD_FLAG_TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"LOG_RECORD_FLAG_UNSPECIFIED":      0,
	"LOG_RECORD_FLAG_TRACE_FLAGS_MASK": 255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{1}
}

// A collection of InstrumentationLibraryLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then no resource info is known.
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibraryLogs that originate from a resource.
	InstrumentationLibraryLogs []*InstrumentationLibraryLogs `protobuf:"bytes,2,rep,name=instrumentation_library_logs,json=instrumentationLibraryLogs,proto3" json:"instrumentation_library_logs,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *ResourceLogs) Reset()         { *m = ResourceLogs{} }
func (m *ResourceLogs) String() string { return proto.CompactTextString(m) }
func (*ResourceLogs) ProtoMessage()    {}
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{0}
}

func (m *ResourceLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLogs.Unmarshal(m, b)
}
func (m *ResourceLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLogs.Marshal(b, m, deterministic)
}
func (m *ResourceLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLogs.Merge(m, src)
}
func (m *ResourceLogs) XXX_Size() int {
	return xxx_messageInfo_ResourceLogs.Size(m)
}
func (m *ResourceLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLogs proto.InternalMessageInfo

func (m *ResourceLogs) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceLogs) GetInstrumentationLibraryLogs() []*InstrumentationLibraryLogs {
	if m != nil {
		return m.InstrumentationLibraryLogs
	}
	return nil
}

// A collection of Logs produced by an InstrumentationLibrary.
type InstrumentationLibraryLogs struct {
	// The instrumentation library information for the logs in this message.
	// If this field is not set then no library info is known.
	InstrumentationLibrary *InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of log records.
	Logs                 []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InstrumentationLibraryLogs) Reset()         { *m = InstrumentationLibraryLogs{} }
func (m *InstrumentationLibraryLogs) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibraryLogs) ProtoMessage()    {}
func (*InstrumentationLibraryLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{1}
}

func (m *InstrumentationLibraryLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibraryLogs.Unmarshal(m, b)
}
func (m *InstrumentationLibraryLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibraryLogs.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibraryLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibraryLogs.Merge(m, src)
}
func (m *InstrumentationLibraryLogs) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibraryLogs.Size(m)
}
func (m *InstrumentationLibraryLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibraryLogs.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibraryLogs proto.InternalMessageInfo

func (m *InstrumentationLibraryLogs) GetInstrumentationLibrary() *InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/master/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano int64 `protobuf:"varint,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=otelp2.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. Name describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. [Optional].
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	Attributes             []*KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32      `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. [Optional].
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. [Optional].
	SpanId               []byte   `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{2}
}

func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetTimeUnixNano() int64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogRecord) GetBody() *AnyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *LogRecord) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func init() {
	proto.RegisterEnum("otelp2.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("otelp2.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "otelp2.ResourceLogs")
	proto.RegisterType((*InstrumentationLibraryLogs)(nil), "otelp2.InstrumentationLibraryLogs")
	proto.RegisterType((*LogRecord)(nil), "otelp2.LogRecord")
}

func init() { proto.RegisterFile("logs.proto", fileDescriptor_782e6d65c19305b4) }

var fileDescriptor_782e6d65c19305b4 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd5, 0xdd, 0x6e, 0xda, 0x3e,
	0x14, 0x00, 0xf0, 0x7f, 0xca, 0xf7, 0x29, 0xa5, 0xae, 0xff, 0x2d, 0x4d, 0xe9, 0xd4, 0x46, 0xac,
	0x95, 0x50, 0xb5, 0xa1, 0x2d, 0x70, 0xb1, 0xbb, 0x29, 0xd0, 0x50, 0xa1, 0x52, 0xa8, 0x0c, 0xb4,
	0xdb, 0x55, 0x14, 0x88, 0x87, 0x22, 0x41, 0x8c, 0x12, 0x53, 0xc1, 0x13, 0xec, 0x25, 0x76, 0xb5,
	0x57, 0xd9, 0x83, 0x6d, 0x8a, 0xf9, 0x58, 0x41, 0xf1, 0xee, 0x72, 0xce, 0xef, 0x1c, 0x1f, 0xc7,
	0x84, 0x04, 0x60, 0xcc, 0x46, 0x41, 0x79, 0xea, 0x33, 0xce, 0x70, 0x92, 0x71, 0x3a, 0x9e, 0xea,
	0x85, 0xec, 0x90, 0x4d, 0x26, 0xcc, 0x5b, 0x66, 0x8b, 0x3f, 0x15, 0xc8, 0x12, 0x1a, 0xb0, 0x99,
	0x3f, 0xa4, 0x2d, 0x36, 0x0a, 0xf0, 0x3b, 0x48, 0xfb, 0xab, 0x58, 0x55, 0x34, 0xa5, 0xb4, 0xaf,
	0xa3, 0xf2, 0xb2, 0xb3, 0xbc, 0xae, 0x23, 0x9b, 0x0a, 0xec, 0xc0, 0x1b, 0xd7, 0x0b, 0xb8, 0x3f,
	0x9b, 0x50, 0x8f, 0xdb, 0xdc, 0x65, 0x9e, 0x35, 0x76, 0x07, 0xbe, 0xed, 0x2f, 0xac, 0x70, 0xb4,
	0xba, 0xa7, 0xc5, 0x4a, 0xfb, 0x7a, 0x71, 0xbd, 0x42, 0x73, 0xbb, 0xb6, 0xb5, 0x2c, 0x0d, 0xe7,
	0x92, 0x82, 0x2b, 0xb5, 0xe2, 0x0f, 0x05, 0x0a, 0xf2, 0x56, 0xfc, 0x0c, 0xa7, 0x92, 0x4d, 0xac,
	0xee, 0xe0, 0xe2, 0xdf, 0xf3, 0x49, 0x3e, 0x7a, 0x36, 0xbe, 0x86, 0xf8, 0xab, 0xbb, 0x38, 0x5a,
	0xaf, 0xd2, 0x62, 0x23, 0x42, 0x87, 0xcc, 0x77, 0x88, 0xe0, 0xe2, 0xf7, 0x18, 0x64, 0x36, 0x39,
	0x7c, 0x05, 0x39, 0xee, 0x4e, 0xa8, 0x35, 0xf3, 0xdc, 0xb9, 0xe5, 0xd9, 0x1e, 0x13, 0x9b, 0x88,
	0x91, 0x6c, 0x98, 0xed, 0x7b, 0xee, 0xbc, 0x6d, 0x7b, 0x0c, 0x7f, 0x86, 0xc3, 0x80, 0xbe, 0x50,
	0xdf, 0xe5, 0x0b, 0xcb, 0x9b, 0x4d, 0x06, 0xd4, 0x57, 0xf7, 0x34, 0xa5, 0x94, 0xd3, 0xf3, 0xeb,
	0x29, 0xdd, 0x15, 0xb7, 0x85, 0x92, 0x5c, 0xb0, 0x15, 0xe3, 0xb7, 0x70, 0xb0, 0x59, 0x80, 0xd3,
	0x39, 0x57, 0x63, 0x9a, 0x52, 0xca, 0x90, 0xec, 0x3a, 0xd9, 0xa3, 0x73, 0x8e, 0x31, 0xc4, 0x3d,
	0x7b, 0x42, 0xd5, 0xb8, 0x30, 0x71, 0x8d, 0xaf, 0x20, 0x3e, 0x60, 0xce, 0x42, 0x4d, 0x6c, 0xff,
	0xb8, 0x86, 0xb7, 0x78, 0xb2, 0xc7, 0x33, 0x4a, 0x84, 0xe2, 0x0f, 0x00, 0x36, 0xe7, 0xbe, 0x3b,
	0x98, 0x71, 0x1a, 0xa8, 0x49, 0x71, 0x00, 0x9b, 0xda, 0x7b, 0xba, 0xaa, 0x7d, 0x55, 0x83, 0x3f,
	0x81, 0xea, 0xf8, 0x6c, 0x3a, 0xa5, 0x8e, 0xf5, 0x37, 0x6b, 0x0d, 0xd9, 0xcc, 0xe3, 0x6a, 0x4a,
	0x53, 0x4a, 0x07, 0x24, 0xbf, 0x72, 0x63, 0xc3, 0xf5, 0x50, 0xf1, 0x31, 0x24, 0xbe, 0x8d, 0xed,
	0x51, 0xa0, 0xa6, 0x35, 0xa5, 0x94, 0x22, 0xcb, 0x00, 0x9f, 0x41, 0x9a, 0xfb, 0xf6, 0x90, 0x5a,
	0xae, 0xa3, 0x66, 0x34, 0xa5, 0x94, 0x25, 0x29, 0x11, 0x37, 0x1d, 0x7c, 0x0a, 0xa9, 0x60, 0x6a,
	0x7b, 0xa1, 0x80, 0x90, 0x64, 0x18, 0x36, 0x9d, 0x9b, 0x5f, 0x09, 0xc8, 0x6d, 0x9f, 0x1b, 0xbe,
	0x84, 0xf3, 0xae, 0xf9, 0x64, 0x92, 0x66, 0xef, 0xab, 0xd5, 0xee, 0x3f, 0xd4, 0x4c, 0x62, 0xf5,
	0xdb, 0xdd, 0x47, 0xb3, 0xde, 0x6c, 0x34, 0xcd, 0x5b, 0xf4, 0x1f, 0x3e, 0x83, 0x93, 0xdd, 0x82,
	0x1e, 0x31, 0xea, 0x26, 0x52, 0x70, 0x01, 0xf2, 0x91, 0xa4, 0xa3, 0x3d, 0xa9, 0x55, 0x50, 0x4c,
	0x6a, 0x55, 0x14, 0x8f, 0x1a, 0x77, 0x6b, 0xd6, 0xfa, 0x77, 0x28, 0x11, 0xd5, 0x26, 0x48, 0x47,
	0x49, 0xa9, 0x55, 0x50, 0x4a, 0x6a, 0x55, 0x94, 0xc6, 0x2a, 0x1c, 0xef, 0x5a, 0xb3, 0xdd, 0xe8,
	0xa0, 0x4c, 0xd4, 0x46, 0x42, 0xd1, 0x11, 0xc8, 0xa8, 0x82, 0xf6, 0x65, 0x54, 0x45, 0xd9, 0xa8,
	0x51, 0xcf, 0x06, 0x69, 0xa3, 0x83, 0xa8, 0xa6, 0x50, 0x74, 0x94, 0x93, 0x51, 0x05, 0x1d, 0xca,
	0xa8, 0x8a, 0x50, 0x14, 0x99, 0x84, 0x74, 0x08, 0x3a, 0x8a, 0x3a, 0x0c, 0x41, 0x3a, 0xc2, 0x52,
	0xab, 0xa0, 0xff, 0xa5, 0x56, 0x45, 0xc7, 0x51, 0xe3, 0x1a, 0x46, 0xcf, 0x68, 0xa1, 0x93, 0xa8,
	0x36, 0x41, 0x3a, 0xca, 0x4b, 0xad, 0x82, 0x4e, 0xa5, 0x56, 0x45, 0xea, 0xcd, 0x17, 0xc8, 0x6d,
	0x5e, 0x27, 0x0d, 0xf1, 0x5f, 0xb8, 0x84, 0xf3, 0x56, 0xe7, 0xce, 0x22, 0x66, 0xbd, 0x43, 0x6e,
	0xad, 0x46, 0xcb, 0xb8, 0xdb, 0x79, 0x88, 0xaf, 0x41, 0xdb, 0x2d, 0x10, 0x4f, 0x9c, 0xb8, 0xec,
	0x5a, 0x0f, 0x46, 0xf7, 0x1e, 0xfd, 0x56, 0x6a, 0xef, 0xe1, 0xc2, 0x65, 0x65, 0x36, 0xa5, 0x1e,
	0xa7, 0x63, 0x3a, 0xa1, 0xdc, 0x5f, 0x2c, 0xbf, 0x03, 0x65, 0xf1, 0xa1, 0x78, 0xf9, 0x58, 0x0b,
	0x5f, 0x64, 0xc1, 0x63, 0x98, 0x7a, 0x54, 0x06, 0x49, 0x61, 0x95, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0x4c, 0xae, 0xa2, 0x47, 0x06, 0x00, 0x00,
}