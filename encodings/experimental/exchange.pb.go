// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

package experimental

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

type Capabilities int32

const (
	Capabilities_SKIP             Capabilities = 0
	Capabilities_ZLIB_COMPRESSION Capabilities = 1
	Capabilities_LZ4_COMPRESSION  Capabilities = 2
)

var Capabilities_name = map[int32]string{
	0: "SKIP",
	1: "ZLIB_COMPRESSION",
	2: "LZ4_COMPRESSION",
}

var Capabilities_value = map[string]int32{
	"SKIP":             0,
	"ZLIB_COMPRESSION": 1,
	"LZ4_COMPRESSION":  2,
}

func (x Capabilities) String() string {
	return proto.EnumName(Capabilities_name, int32(x))
}

func (Capabilities) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

type CompressionMethod int32

const (
	CompressionMethod_NONE CompressionMethod = 0
	CompressionMethod_LZ4  CompressionMethod = 1
	CompressionMethod_ZLIB CompressionMethod = 2
	CompressionMethod_ZSTD CompressionMethod = 3
)

var CompressionMethod_name = map[int32]string{
	0: "NONE",
	1: "LZ4",
	2: "ZLIB",
}

var CompressionMethod_value = map[string]int32{
	"NONE": 0,
	"LZ4":  1,
	"ZLIB": 2,
}

func (x CompressionMethod) String() string {
	return proto.EnumName(CompressionMethod_name, int32(x))
}

func (CompressionMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

type RequestType int32

const (
	RequestType__           RequestType = 0
	RequestType_TraceExport RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "_",
	1: "TraceExport",
}

var RequestType_value = map[string]int32{
	"_":           0,
	"TraceExport": 1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}

func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2}
}

type ExportResponse_ResultCode int32

const (
	// Telemetry data is successfully processed by the server.
	ExportResponse_Success ExportResponse_ResultCode = 0
	// processing of telemetry data failed. The client MUST NOT retry
	// sending the same telemetry data. The telemetry data MUST be dropped.
	// This for example can happen when the request contains bad data and
	// cannot be deserialized or otherwise processed by the server.
	ExportResponse_FailedNoneRetryable ExportResponse_ResultCode = 1
	// Processing of telemetry data failed. The client SHOULD record the
	// error and MAY retry exporting the same data after some time. This
	// for example can happen when the server is overloaded.
	ExportResponse_FailedRetryable ExportResponse_ResultCode = 2
)

var ExportResponse_ResultCode_name = map[int32]string{
	0: "Success",
	1: "FailedNoneRetryable",
	2: "FailedRetryable",
}

var ExportResponse_ResultCode_value = map[string]int32{
	"Success":             0,
	"FailedNoneRetryable": 1,
	"FailedRetryable":     2,
}

func (x ExportResponse_ResultCode) String() string {
	return proto.EnumName(ExportResponse_ResultCode_name, int32(x))
}

func (ExportResponse_ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2, 0}
}

// A request from client to server containing trace data to export.
type TraceExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceSpans.
	ResourceSpans        []*ResourceSpans `protobuf:"bytes,2,rep,name=resourceSpans,proto3" json:"resourceSpans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TraceExportRequest) Reset()         { *m = TraceExportRequest{} }
func (m *TraceExportRequest) String() string { return proto.CompactTextString(m) }
func (*TraceExportRequest) ProtoMessage()    {}
func (*TraceExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

func (m *TraceExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceExportRequest.Unmarshal(m, b)
}
func (m *TraceExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceExportRequest.Marshal(b, m, deterministic)
}
func (m *TraceExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceExportRequest.Merge(m, src)
}
func (m *TraceExportRequest) XXX_Size() int {
	return xxx_messageInfo_TraceExportRequest.Size(m)
}
func (m *TraceExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceExportRequest proto.InternalMessageInfo

func (m *TraceExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TraceExportRequest) GetResourceSpans() []*ResourceSpans {
	if m != nil {
		return m.ResourceSpans
	}
	return nil
}

// A request from client to server containing metric data to export.
type MetricExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceMetrics.
	ResourceMetrics      []*ResourceMetrics `protobuf:"bytes,2,rep,name=resourceMetrics,proto3" json:"resourceMetrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricExportRequest) Reset()         { *m = MetricExportRequest{} }
func (m *MetricExportRequest) String() string { return proto.CompactTextString(m) }
func (*MetricExportRequest) ProtoMessage()    {}
func (*MetricExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

func (m *MetricExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricExportRequest.Unmarshal(m, b)
}
func (m *MetricExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricExportRequest.Marshal(b, m, deterministic)
}
func (m *MetricExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricExportRequest.Merge(m, src)
}
func (m *MetricExportRequest) XXX_Size() int {
	return xxx_messageInfo_MetricExportRequest.Size(m)
}
func (m *MetricExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricExportRequest proto.InternalMessageInfo

func (m *MetricExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MetricExportRequest) GetResourceMetrics() []*ResourceMetrics {
	if m != nil {
		return m.ResourceMetrics
	}
	return nil
}

// A response to ExportRequest.
type ExportResponse struct {
	// ID of a response that the server acknowledges.
	Id                   uint64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResultCode           ExportResponse_ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=experimental.ExportResponse_ResultCode" json:"result_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExportResponse) Reset()         { *m = ExportResponse{} }
func (m *ExportResponse) String() string { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()    {}
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2}
}

func (m *ExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportResponse.Unmarshal(m, b)
}
func (m *ExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportResponse.Marshal(b, m, deterministic)
}
func (m *ExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportResponse.Merge(m, src)
}
func (m *ExportResponse) XXX_Size() int {
	return xxx_messageInfo_ExportResponse.Size(m)
}
func (m *ExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportResponse proto.InternalMessageInfo

func (m *ExportResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportResponse) GetResultCode() ExportResponse_ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ExportResponse_Success
}

// RequestHeader is used by transports that unlike gRPC don't have built-in request
// compression such as WebSocket. Request body typically follows the header.
type RequestHeader struct {
	// Compression method used for body.
	Compression CompressionMethod `protobuf:"varint,1,opt,name=compression,proto3,enum=experimental.CompressionMethod" json:"compression,omitempty"`
	// Compression level as defined by the compression method.
	CompressionLevel     int32    `protobuf:"varint,2,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{3}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetCompression() CompressionMethod {
	if m != nil {
		return m.Compression
	}
	return CompressionMethod_NONE
}

func (m *RequestHeader) GetCompressionLevel() int32 {
	if m != nil {
		return m.CompressionLevel
	}
	return 0
}

// RequestBody is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type RequestBody struct {
	RequestType          RequestType         `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=experimental.RequestType" json:"request_type,omitempty"`
	Export               *TraceExportRequest `protobuf:"bytes,2,opt,name=export,proto3" json:"export,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RequestBody) Reset()         { *m = RequestBody{} }
func (m *RequestBody) String() string { return proto.CompactTextString(m) }
func (*RequestBody) ProtoMessage()    {}
func (*RequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{4}
}

func (m *RequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestBody.Unmarshal(m, b)
}
func (m *RequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestBody.Marshal(b, m, deterministic)
}
func (m *RequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBody.Merge(m, src)
}
func (m *RequestBody) XXX_Size() int {
	return xxx_messageInfo_RequestBody.Size(m)
}
func (m *RequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBody proto.InternalMessageInfo

func (m *RequestBody) GetRequestType() RequestType {
	if m != nil {
		return m.RequestType
	}
	return RequestType__
}

func (m *RequestBody) GetExport() *TraceExportRequest {
	if m != nil {
		return m.Export
	}
	return nil
}

// Response is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type Response struct {
	ResponseType         RequestType     `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=experimental.RequestType" json:"response_type,omitempty"`
	Export               *ExportResponse `protobuf:"bytes,2,opt,name=export,proto3" json:"export,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseType() RequestType {
	if m != nil {
		return m.ResponseType
	}
	return RequestType__
}

func (m *Response) GetExport() *ExportResponse {
	if m != nil {
		return m.Export
	}
	return nil
}

func init() {
	proto.RegisterEnum("experimental.Capabilities", Capabilities_name, Capabilities_value)
	proto.RegisterEnum("experimental.CompressionMethod", CompressionMethod_name, CompressionMethod_value)
	proto.RegisterEnum("experimental.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("experimental.ExportResponse_ResultCode", ExportResponse_ResultCode_name, ExportResponse_ResultCode_value)
	proto.RegisterType((*TraceExportRequest)(nil), "experimental.TraceExportRequest")
	proto.RegisterType((*MetricExportRequest)(nil), "experimental.MetricExportRequest")
	proto.RegisterType((*ExportResponse)(nil), "experimental.ExportResponse")
	proto.RegisterType((*RequestHeader)(nil), "experimental.RequestHeader")
	proto.RegisterType((*RequestBody)(nil), "experimental.RequestBody")
	proto.RegisterType((*Response)(nil), "experimental.Response")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_e0328a4f16f87ea1) }

var fileDescriptor_e0328a4f16f87ea1 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x18, 0xb4, 0x94, 0x5f, 0x3e, 0x39, 0x8e, 0xb2, 0x09, 0x34, 0x49, 0x53, 0x6a, 0x04, 0xa5, 0x21,
	0x05, 0xd1, 0xb8, 0x39, 0xf4, 0x50, 0x0a, 0xb1, 0x48, 0x13, 0x53, 0xff, 0xb1, 0x0e, 0x3d, 0xf8,
	0x22, 0x64, 0xe9, 0x23, 0x11, 0xc8, 0x5a, 0x75, 0x77, 0x15, 0xec, 0x53, 0x7b, 0xe9, 0x13, 0xf5,
	0xad, 0xfa, 0x14, 0x45, 0x2b, 0xff, 0x48, 0x76, 0x09, 0xbd, 0xed, 0xce, 0xce, 0xec, 0x8c, 0xbe,
	0x59, 0x41, 0x0d, 0x27, 0xfe, 0xa3, 0x17, 0x3f, 0xa0, 0x9d, 0x70, 0x26, 0x19, 0xa9, 0xe2, 0x24,
	0x41, 0x1e, 0x8e, 0x31, 0x96, 0x5e, 0x74, 0x7a, 0x24, 0x31, 0xc2, 0x31, 0x4a, 0x3e, 0x75, 0x03,
	0x4f, 0x7a, 0x39, 0xe7, 0xf4, 0x20, 0x43, 0x42, 0xbf, 0x00, 0x59, 0x0f, 0x40, 0xee, 0xb9, 0xe7,
	0xe3, 0xcd, 0x24, 0x61, 0x5c, 0x52, 0xfc, 0x9e, 0xa2, 0x90, 0xa4, 0x06, 0x7a, 0x18, 0x1c, 0x6b,
	0x75, 0xed, 0x7c, 0x93, 0xea, 0x61, 0x40, 0xae, 0x61, 0x8f, 0xa3, 0x60, 0x29, 0xf7, 0x71, 0x90,
	0x78, 0xb1, 0x38, 0xd6, 0xeb, 0x1b, 0xe7, 0x46, 0xe3, 0xa5, 0x5d, 0x34, 0xb5, 0x69, 0x91, 0x42,
	0xcb, 0x0a, 0x2b, 0x86, 0xc3, 0x8e, 0x72, 0x7f, 0xde, 0xe9, 0x16, 0xf6, 0xe7, 0xba, 0x9c, 0x3e,
	0xf7, 0x7a, 0xf5, 0x6f, 0xaf, 0x19, 0x89, 0xae, 0xaa, 0xac, 0xdf, 0x1a, 0xd4, 0xe6, 0x56, 0x22,
	0x61, 0xb1, 0xc0, 0x35, 0xaf, 0x3b, 0x30, 0x38, 0x8a, 0x34, 0x92, 0xae, 0xcf, 0x02, 0x3c, 0xd6,
	0xeb, 0xda, 0x79, 0xad, 0xf1, 0xb6, 0xec, 0x53, 0xbe, 0x22, 0xb3, 0x4d, 0x23, 0xe9, 0xb0, 0x00,
	0x29, 0xf0, 0xc5, 0xda, 0xba, 0x05, 0x58, 0x9e, 0x10, 0x03, 0x76, 0x06, 0xa9, 0xef, 0xa3, 0x10,
	0x66, 0x85, 0xbc, 0x80, 0xc3, 0x2f, 0x5e, 0x18, 0x61, 0xd0, 0x65, 0x31, 0xd2, 0xac, 0x11, 0x6f,
	0x14, 0xa1, 0xa9, 0x91, 0x43, 0xd8, 0xcf, 0x0f, 0x96, 0xa0, 0x6e, 0xfd, 0x80, 0xbd, 0xd9, 0x64,
	0xee, 0xd0, 0x0b, 0x90, 0x93, 0x6b, 0x30, 0x7c, 0x36, 0x4e, 0x38, 0x0a, 0x11, 0xb2, 0x58, 0x85,
	0xaf, 0x35, 0x5e, 0x97, 0x33, 0x3a, 0x4b, 0x42, 0x07, 0xe5, 0x23, 0x0b, 0x68, 0x51, 0x43, 0xde,
	0xc1, 0x41, 0x61, 0xeb, 0x46, 0xf8, 0x84, 0x91, 0xfa, 0xd8, 0x2d, 0x6a, 0x16, 0x0e, 0xda, 0x19,
	0x6e, 0xfd, 0xd2, 0xc0, 0x98, 0x25, 0x68, 0xb2, 0x60, 0x4a, 0x3e, 0x41, 0x95, 0xe7, 0x5b, 0x57,
	0x4e, 0x13, 0x9c, 0x05, 0x38, 0x59, 0x2d, 0x43, 0x31, 0xee, 0xa7, 0x09, 0x52, 0x83, 0x2f, 0x37,
	0xe4, 0x23, 0x6c, 0xa3, 0x1a, 0xa0, 0xf2, 0x33, 0x1a, 0xf5, 0xb2, 0x6e, 0xfd, 0xe5, 0xd1, 0x19,
	0xdf, 0xfa, 0xa9, 0xc1, 0xee, 0xa2, 0xb8, 0xcf, 0xea, 0xf9, 0xa9, 0xf5, 0x7f, 0xa6, 0xa8, 0xce,
	0xf9, 0x2a, 0xc6, 0xd5, 0x4a, 0x8c, 0xb3, 0xe7, 0x3a, 0x9e, 0x47, 0xb8, 0x70, 0xa0, 0xea, 0x78,
	0x89, 0x37, 0x0a, 0xa3, 0x50, 0x86, 0x28, 0xc8, 0x2e, 0x6c, 0x0e, 0xbe, 0xb6, 0xfa, 0x66, 0x85,
	0x1c, 0x81, 0x39, 0x6c, 0xb7, 0x9a, 0xae, 0xd3, 0xeb, 0xf4, 0xe9, 0xcd, 0x60, 0xd0, 0xea, 0x75,
	0xf3, 0x42, 0xdb, 0xc3, 0xab, 0x12, 0xa8, 0x5f, 0xbc, 0x87, 0x83, 0xb5, 0x7a, 0xb2, 0x9b, 0xba,
	0xbd, 0xee, 0x8d, 0x59, 0x21, 0x3b, 0xb0, 0xd1, 0x1e, 0x5e, 0x99, 0x5a, 0x06, 0x65, 0x57, 0x9a,
	0xfa, 0xc5, 0x9b, 0x45, 0x01, 0x2a, 0xfb, 0x16, 0x68, 0xae, 0x59, 0x21, 0xfb, 0x60, 0x14, 0xa6,
	0x65, 0x6a, 0xcd, 0x3b, 0x38, 0x0b, 0x99, 0xcd, 0x12, 0x8c, 0x7d, 0x8c, 0x45, 0x2a, 0xf2, 0xff,
	0xd9, 0x96, 0x19, 0xc9, 0x7e, 0xba, 0x6c, 0x82, 0xa2, 0xf7, 0x33, 0xb0, 0xaf, 0xfd, 0xd1, 0x4f,
	0x7a, 0x09, 0xc6, 0x4e, 0xce, 0x54, 0x60, 0x3e, 0x7c, 0xfb, 0xdb, 0xe5, 0x68, 0x5b, 0x29, 0x3f,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x27, 0x20, 0x50, 0xce, 0x52, 0x04, 0x00, 0x00,
}
