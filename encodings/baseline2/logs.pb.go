// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logs.proto

package baseline2

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
	SeverityNumber_UNDEFINED_SEVERITY_NUMBER SeverityNumber = 0
	SeverityNumber_TRACE                     SeverityNumber = 1
	SeverityNumber_TRACE2                    SeverityNumber = 2
	SeverityNumber_TRACE3                    SeverityNumber = 3
	SeverityNumber_TRACE4                    SeverityNumber = 4
	SeverityNumber_DEBUG                     SeverityNumber = 5
	SeverityNumber_DEBUG2                    SeverityNumber = 6
	SeverityNumber_DEBUG3                    SeverityNumber = 7
	SeverityNumber_DEBUG4                    SeverityNumber = 8
	SeverityNumber_INFO                      SeverityNumber = 9
	SeverityNumber_INFO2                     SeverityNumber = 10
	SeverityNumber_INFO3                     SeverityNumber = 11
	SeverityNumber_INFO4                     SeverityNumber = 12
	SeverityNumber_WARN                      SeverityNumber = 13
	SeverityNumber_WARN2                     SeverityNumber = 14
	SeverityNumber_WARN3                     SeverityNumber = 15
	SeverityNumber_WARN4                     SeverityNumber = 16
	SeverityNumber_ERROR                     SeverityNumber = 17
	SeverityNumber_ERROR2                    SeverityNumber = 18
	SeverityNumber_ERROR3                    SeverityNumber = 19
	SeverityNumber_ERROR4                    SeverityNumber = 20
	SeverityNumber_FATAL                     SeverityNumber = 21
	SeverityNumber_FATAL2                    SeverityNumber = 22
	SeverityNumber_FATAL3                    SeverityNumber = 23
	SeverityNumber_FATAL4                    SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "UNDEFINED_SEVERITY_NUMBER",
	1:  "TRACE",
	2:  "TRACE2",
	3:  "TRACE3",
	4:  "TRACE4",
	5:  "DEBUG",
	6:  "DEBUG2",
	7:  "DEBUG3",
	8:  "DEBUG4",
	9:  "INFO",
	10: "INFO2",
	11: "INFO3",
	12: "INFO4",
	13: "WARN",
	14: "WARN2",
	15: "WARN3",
	16: "WARN4",
	17: "ERROR",
	18: "ERROR2",
	19: "ERROR3",
	20: "ERROR4",
	21: "FATAL",
	22: "FATAL2",
	23: "FATAL3",
	24: "FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"UNDEFINED_SEVERITY_NUMBER": 0,
	"TRACE":                     1,
	"TRACE2":                    2,
	"TRACE3":                    3,
	"TRACE4":                    4,
	"DEBUG":                     5,
	"DEBUG2":                    6,
	"DEBUG3":                    7,
	"DEBUG4":                    8,
	"INFO":                      9,
	"INFO2":                     10,
	"INFO3":                     11,
	"INFO4":                     12,
	"WARN":                      13,
	"WARN2":                     14,
	"WARN3":                     15,
	"WARN4":                     16,
	"ERROR":                     17,
	"ERROR2":                    18,
	"ERROR3":                    19,
	"ERROR4":                    20,
	"FATAL":                     21,
	"FATAL2":                    22,
	"FATAL3":                    23,
	"FATAL4":                    24,
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
	LogRecordFlags_UNDEFINED_LOG_RECORD_FLAG LogRecordFlags = 0
	LogRecordFlags_TRACE_FLAGS_MASK          LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "UNDEFINED_LOG_RECORD_FLAG",
	255: "TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"UNDEFINED_LOG_RECORD_FLAG": 0,
	"TRACE_FLAGS_MASK":          255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{1}
}

// A collection of logs from a Resource.
type ResourceLogs struct {
	// The resource for the spans in this message.
	// If this field is not set then no resource info is known.
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of log records.
	Logs                 []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func (m *ResourceLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

// A log record according to OpenTelemetry Log Data Model: https://github.com/open-telemetry/oteps/pull/97
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixnano uint64 `protobuf:"fixed64,1,opt,name=time_unixnano,json=timeUnixnano,proto3" json:"time_unixnano,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. Optional.
	TraceId []byte `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. Optional.
	SpanId []byte `protobuf:"bytes,3,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Optional flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers that must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). Optional.
	Flags uint32 `protobuf:"fixed32,4,opt,name=flags,proto3" json:"flags,omitempty"`
	// Numerical value of the severity, normalized to values described in
	// https://github.com/open-telemetry/oteps/pull/97. Optional.
	SeverityNumber SeverityNumber `protobuf:"varint,5,opt,name=severity_number,json=severityNumber,proto3,enum=baseline2.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. Optional.
	SeverityText string `protobuf:"bytes,6,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. ShortName describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. Optional.
	ShortName string `protobuf:"bytes,7,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. Optional.
	Body *AttributeKeyValue `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. Optional.
	Attributes             []*AttributeKeyValue `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32               `protobuf:"varint,10,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_782e6d65c19305b4, []int{1}
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

func (m *LogRecord) GetTimeUnixnano() uint64 {
	if m != nil {
		return m.TimeUnixnano
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

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_UNDEFINED_SEVERITY_NUMBER
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *LogRecord) GetBody() *AttributeKeyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*AttributeKeyValue {
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

func init() {
	proto.RegisterEnum("baseline2.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("baseline2.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "baseline2.ResourceLogs")
	proto.RegisterType((*LogRecord)(nil), "baseline2.LogRecord")
}

func init() { proto.RegisterFile("logs.proto", fileDescriptor_782e6d65c19305b4) }

var fileDescriptor_782e6d65c19305b4 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xda, 0x4a,
	0x14, 0xc6, 0x63, 0xfe, 0x73, 0x42, 0xc8, 0xb9, 0x93, 0x7f, 0xce, 0xd5, 0xcd, 0x15, 0xca, 0xdd,
	0x58, 0x91, 0x2e, 0x6d, 0x0d, 0x8b, 0x2e, 0xba, 0x31, 0xc1, 0x44, 0x28, 0xc4, 0x44, 0x03, 0xa4,
	0xea, 0xca, 0x32, 0x78, 0x4a, 0x2d, 0x61, 0x0f, 0xb2, 0x87, 0x28, 0xbc, 0x53, 0x5f, 0xa2, 0x4f,
	0xd6, 0x6a, 0x26, 0x78, 0x4a, 0x36, 0xdd, 0xfd, 0xce, 0xf7, 0x7d, 0x33, 0x67, 0x74, 0x34, 0x07,
	0x60, 0xc5, 0x97, 0x59, 0x7b, 0x9d, 0x72, 0xc1, 0x49, 0x7d, 0x1e, 0x64, 0x6c, 0x15, 0x25, 0xcc,
	0xfe, 0xbb, 0xb1, 0xe0, 0x71, 0xcc, 0x93, 0x57, 0xe3, 0x3a, 0x82, 0x06, 0x65, 0x19, 0xdf, 0xa4,
	0x0b, 0x36, 0xe2, 0xcb, 0x8c, 0xbc, 0x83, 0x5a, 0xba, 0xab, 0x4d, 0xa3, 0x65, 0x58, 0x87, 0xf6,
	0x49, 0x5b, 0x9f, 0x6d, 0xe7, 0x51, 0xaa, 0x43, 0xc4, 0x82, 0x92, 0xec, 0x63, 0x16, 0x5a, 0x45,
	0xeb, 0xd0, 0x3e, 0xdd, 0x0b, 0x8f, 0xf8, 0x92, 0xb2, 0x05, 0x4f, 0x43, 0xaa, 0x12, 0xd7, 0xdf,
	0x8b, 0x50, 0xd7, 0x1a, 0xf9, 0x0f, 0x8e, 0x44, 0x14, 0x33, 0x7f, 0x93, 0x44, 0x2f, 0x49, 0x90,
	0x70, 0xd5, 0xad, 0x42, 0x1b, 0x52, 0x9c, 0xed, 0x34, 0x72, 0x09, 0x35, 0x91, 0x06, 0x0b, 0xe6,
	0x47, 0xa1, 0x59, 0x68, 0x19, 0x56, 0x83, 0x56, 0x55, 0x3d, 0x0c, 0xc9, 0x05, 0x54, 0xb3, 0x75,
	0x90, 0x48, 0xa7, 0xa8, 0x9c, 0x8a, 0x2c, 0x87, 0x21, 0x39, 0x85, 0xf2, 0xd7, 0x55, 0xb0, 0xcc,
	0xcc, 0x52, 0xcb, 0xb0, 0xaa, 0xf4, 0xb5, 0x20, 0x3d, 0x38, 0xce, 0xd8, 0x33, 0x4b, 0x23, 0xb1,
	0xf5, 0x93, 0x4d, 0x3c, 0x67, 0xa9, 0x59, 0x6e, 0x19, 0x56, 0xd3, 0xbe, 0xdc, 0x7b, 0xf1, 0x64,
	0x97, 0xf0, 0x54, 0x80, 0x36, 0xb3, 0x37, 0xb5, 0x7c, 0xb2, 0xbe, 0x43, 0xb0, 0x17, 0x61, 0x56,
	0x5a, 0x86, 0x55, 0xa7, 0x8d, 0x5c, 0x9c, 0xb2, 0x17, 0x41, 0xae, 0x00, 0xb2, 0x6f, 0x3c, 0x15,
	0x7e, 0x12, 0xc4, 0xcc, 0xac, 0xaa, 0x44, 0x5d, 0x29, 0x5e, 0x10, 0x33, 0xf2, 0x1e, 0x4a, 0x73,
	0x1e, 0x6e, 0xcd, 0x9a, 0x9a, 0xed, 0x3f, 0x7b, 0xcd, 0x1d, 0x21, 0xd2, 0x68, 0xbe, 0x11, 0xec,
	0x9e, 0x6d, 0x9f, 0x82, 0xd5, 0x86, 0x51, 0x95, 0x24, 0x9f, 0x00, 0x82, 0xdc, 0xca, 0xcc, 0xba,
	0x1a, 0xf3, 0x9f, 0xcf, 0xed, 0xe5, 0xc9, 0x47, 0x30, 0xc3, 0x94, 0xaf, 0xd7, 0x2c, 0xf4, 0x7f,
	0xab, 0xfe, 0x82, 0x6f, 0x12, 0x61, 0x42, 0xcb, 0xb0, 0x8e, 0xe8, 0xf9, 0xce, 0xd7, 0xf7, 0x64,
	0xb7, 0xd2, 0xbd, 0xf9, 0x51, 0x80, 0xe6, 0xdb, 0x81, 0x90, 0x2b, 0xb8, 0x9c, 0x79, 0x7d, 0x77,
	0x30, 0xf4, 0xdc, 0xbe, 0x3f, 0x71, 0x9f, 0x5c, 0x3a, 0x9c, 0x7e, 0xf1, 0xbd, 0xd9, 0x43, 0xcf,
	0xa5, 0x78, 0x40, 0xea, 0x50, 0x9e, 0x52, 0xe7, 0xd6, 0x45, 0x83, 0x00, 0x54, 0x14, 0xda, 0x58,
	0xd0, 0xdc, 0xc1, 0xa2, 0xe6, 0x2e, 0x96, 0x64, 0xbc, 0xef, 0xf6, 0x66, 0x77, 0x58, 0x96, 0xb2,
	0x42, 0x1b, 0x2b, 0x9a, 0x3b, 0x58, 0xd5, 0xdc, 0xc5, 0x1a, 0xa9, 0x41, 0x69, 0xe8, 0x0d, 0xc6,
	0x58, 0x97, 0x07, 0x25, 0xd9, 0x08, 0x39, 0x76, 0xf0, 0x30, 0xc7, 0x2e, 0x36, 0x64, 0xf4, 0xb3,
	0x43, 0x3d, 0x3c, 0x92, 0xa2, 0x24, 0x1b, 0x9b, 0x39, 0x76, 0xf0, 0x38, 0xc7, 0x2e, 0xa2, 0x44,
	0x97, 0xd2, 0x31, 0xc5, 0xbf, 0x64, 0x33, 0x85, 0x36, 0x12, 0xcd, 0x1d, 0x3c, 0xd1, 0xdc, 0xc5,
	0x53, 0x19, 0x1f, 0x38, 0x53, 0x67, 0x84, 0x67, 0x52, 0x56, 0x68, 0xe3, 0xb9, 0xe6, 0x0e, 0x5e,
	0x68, 0xee, 0xa2, 0x79, 0x33, 0x80, 0xa6, 0xfe, 0xf1, 0x03, 0xf5, 0x0f, 0xdf, 0x8c, 0x70, 0x34,
	0xbe, 0xf3, 0xa9, 0x7b, 0x3b, 0xa6, 0x7d, 0x7f, 0x30, 0x72, 0xee, 0xf0, 0x80, 0x9c, 0x01, 0xaa,
	0xf9, 0xa8, 0x7a, 0xe2, 0x3f, 0x38, 0x93, 0x7b, 0xfc, 0x69, 0xf4, 0xfe, 0x87, 0x7f, 0x23, 0xde,
	0xe6, 0x6b, 0x96, 0x08, 0xb6, 0x62, 0x31, 0x13, 0xe9, 0xf6, 0x75, 0x7f, 0xdb, 0x6a, 0xc7, 0x9f,
	0x3f, 0xf4, 0xe4, 0x66, 0x65, 0x8f, 0x52, 0x7a, 0x34, 0xe6, 0x15, 0xe5, 0x75, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x2a, 0x4c, 0x5a, 0x02, 0x04, 0x00, 0x00,
}