// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package experimental2

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
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0xad, 0x28, 0x48, 0x2d, 0xca, 0xcc, 0x4d,
	0xcd, 0x2b, 0x49, 0xcc, 0x31, 0x92, 0xe2, 0x4b, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x85,
	0x48, 0x1b, 0x25, 0x73, 0xf1, 0x86, 0xe6, 0x25, 0x16, 0x55, 0xba, 0x56, 0x14, 0xe4, 0x17, 0x95,
	0xa4, 0x16, 0x09, 0x05, 0x71, 0xf1, 0x40, 0xd8, 0x21, 0x45, 0x89, 0xc9, 0xa9, 0xc5, 0x42, 0x8a,
	0x7a, 0x28, 0x06, 0xe8, 0x81, 0x85, 0x21, 0x2a, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4,
	0x64, 0xd1, 0x94, 0xc0, 0x64, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x8c, 0x32, 0xb8,
	0xf8, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0xe1, 0xb6, 0x84, 0x51, 0xdf, 0x16, 0x0d, 0x46, 0x03,
	0x46, 0x27, 0x0f, 0x2e, 0x99, 0xcc, 0x7c, 0xbd, 0xfc, 0x82, 0xd4, 0xbc, 0xe4, 0xd4, 0xbc, 0xe2,
	0xd2, 0x62, 0x88, 0x2f, 0xf5, 0x4a, 0x40, 0xc6, 0xe9, 0x95, 0x19, 0x3a, 0x71, 0x81, 0x0d, 0x0e,
	0x00, 0x09, 0x06, 0x30, 0xbe, 0x62, 0x92, 0xf4, 0x2f, 0x48, 0xcd, 0x73, 0x86, 0xa8, 0x04, 0x0b,
	0x42, 0x2c, 0xd6, 0x0b, 0x33, 0x4c, 0x62, 0x03, 0xeb, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x02, 0x16, 0x74, 0xb6, 0x4c, 0x01, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/experimental2.UnaryExporter/ExportTraces", in, out, opts...)
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
		FullMethod: "/experimental2.UnaryExporter/ExportTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryExporterServer).ExportTraces(ctx, req.(*TraceExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnaryExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "experimental2.UnaryExporter",
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
	stream, err := c.cc.NewStream(ctx, &_StreamExporter_serviceDesc.Streams[0], "/experimental2.StreamExporter/ExportTraces", opts...)
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
	ServiceName: "experimental2.StreamExporter",
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