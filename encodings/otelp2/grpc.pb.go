// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package otelp2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x2f, 0x49, 0xcd, 0x29, 0x30, 0x92, 0xe2,
	0x4b, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x85, 0x88, 0x1b, 0x85, 0x72, 0xf1, 0x86, 0xe6,
	0x25, 0x16, 0x55, 0xba, 0x56, 0x14, 0xe4, 0x17, 0x95, 0xa4, 0x16, 0x09, 0xb9, 0x70, 0xf1, 0x40,
	0xd8, 0x21, 0x45, 0x89, 0xc9, 0xa9, 0xc5, 0x42, 0x52, 0x7a, 0x10, 0x9d, 0x7a, 0x60, 0x3e, 0x44,
	0x2a, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x4a, 0x0c, 0x26, 0x07, 0x13, 0x2e, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x55, 0x62, 0x30, 0x8a, 0xe2, 0xe2, 0x0b, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x85,
	0x9b, 0xeb, 0x41, 0x0d, 0x73, 0x35, 0x18, 0x0d, 0x18, 0x9d, 0x3c, 0xb8, 0x64, 0x32, 0xf3, 0xf5,
	0xf2, 0x0b, 0x52, 0xf3, 0x92, 0x53, 0xf3, 0x8a, 0x4b, 0x8b, 0x21, 0x3e, 0xd1, 0x2b, 0x01, 0x99,
	0xa3, 0x57, 0x66, 0xe8, 0xc4, 0x05, 0x36, 0x31, 0x00, 0x24, 0x18, 0xc0, 0xf8, 0x8a, 0x49, 0xd2,
	0xbf, 0x20, 0x35, 0xcf, 0x19, 0xa2, 0x12, 0x2c, 0x08, 0xb1, 0x51, 0x2f, 0xcc, 0x30, 0x89, 0x0d,
	0xac, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x95, 0x39, 0xbd, 0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UnaryExporterClient is the client API for UnaryExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UnaryExporterClient interface {
	ExportTraces(ctx context.Context, in *TraceExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
}

type unaryExporterClient struct {
	cc *grpc.ClientConn
}

func NewUnaryExporterClient(cc *grpc.ClientConn) UnaryExporterClient {
	return &unaryExporterClient{cc}
}

func (c *unaryExporterClient) ExportTraces(ctx context.Context, in *TraceExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, "/otelp2.UnaryExporter/ExportTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnaryExporterServer is the server API for UnaryExporter service.
type UnaryExporterServer interface {
	ExportTraces(context.Context, *TraceExportRequest) (*ExportResponse, error)
}

// UnimplementedUnaryExporterServer can be embedded to have forward compatible implementations.
type UnimplementedUnaryExporterServer struct {
}

func (*UnimplementedUnaryExporterServer) ExportTraces(ctx context.Context, req *TraceExportRequest) (*ExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportTraces not implemented")
}

func RegisterUnaryExporterServer(s *grpc.Server, srv UnaryExporterServer) {
	s.RegisterService(&_UnaryExporter_serviceDesc, srv)
}

func _UnaryExporter_ExportTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnaryExporterServer).ExportTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otelp2.UnaryExporter/ExportTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryExporterServer).ExportTraces(ctx, req.(*TraceExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnaryExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "otelp2.UnaryExporter",
	HandlerType: (*UnaryExporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportTraces",
			Handler:    _UnaryExporter_ExportTraces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// StreamExporterClient is the client API for StreamExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamExporterClient interface {
	ExportTraces(ctx context.Context, opts ...grpc.CallOption) (StreamExporter_ExportTracesClient, error)
}

type streamExporterClient struct {
	cc *grpc.ClientConn
}

func NewStreamExporterClient(cc *grpc.ClientConn) StreamExporterClient {
	return &streamExporterClient{cc}
}

func (c *streamExporterClient) ExportTraces(ctx context.Context, opts ...grpc.CallOption) (StreamExporter_ExportTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamExporter_serviceDesc.Streams[0], "/otelp2.StreamExporter/ExportTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamExporterExportTracesClient{stream}
	return x, nil
}

type StreamExporter_ExportTracesClient interface {
	Send(*TraceExportRequest) error
	Recv() (*ExportResponse, error)
	grpc.ClientStream
}

type streamExporterExportTracesClient struct {
	grpc.ClientStream
}

func (x *streamExporterExportTracesClient) Send(m *TraceExportRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamExporterExportTracesClient) Recv() (*ExportResponse, error) {
	m := new(ExportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamExporterServer is the server API for StreamExporter service.
type StreamExporterServer interface {
	ExportTraces(StreamExporter_ExportTracesServer) error
}

// UnimplementedStreamExporterServer can be embedded to have forward compatible implementations.
type UnimplementedStreamExporterServer struct {
}

func (*UnimplementedStreamExporterServer) ExportTraces(srv StreamExporter_ExportTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method ExportTraces not implemented")
}

func RegisterStreamExporterServer(s *grpc.Server, srv StreamExporterServer) {
	s.RegisterService(&_StreamExporter_serviceDesc, srv)
}

func _StreamExporter_ExportTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamExporterServer).ExportTraces(&streamExporterExportTracesServer{stream})
}

type StreamExporter_ExportTracesServer interface {
	Send(*ExportResponse) error
	Recv() (*TraceExportRequest, error)
	grpc.ServerStream
}

type streamExporterExportTracesServer struct {
	grpc.ServerStream
}

func (x *streamExporterExportTracesServer) Send(m *ExportResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamExporterExportTracesServer) Recv() (*TraceExportRequest, error) {
	m := new(TraceExportRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "otelp2.StreamExporter",
	HandlerType: (*StreamExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExportTraces",
			Handler:       _StreamExporter_ExportTraces_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}