// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource.proto

package traceprotobuf

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

// Resource information.
type Resource struct {
	// Type identifier for the resource.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Set of labels that describe the resource.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b72f771c35e3b8, []int{0}
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

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "traceprotobuf.Resource")
	proto.RegisterMapType((map[string]string)(nil), "traceprotobuf.Resource.LabelsEntry")
}

func init() { proto.RegisterFile("resource.proto", fileDescriptor_d1b72f771c35e3b8) }

var fileDescriptor_d1b72f771c35e3b8 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x29, 0x4a, 0x4c,
	0x4e, 0x05, 0xb3, 0x93, 0x4a, 0xd3, 0x94, 0xa6, 0x31, 0x72, 0x71, 0x04, 0x41, 0x55, 0x08, 0x09,
	0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42,
	0xd6, 0x5c, 0x6c, 0x39, 0x89, 0x49, 0xa9, 0x39, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0xca, 0x7a, 0x28, 0x06, 0xe8, 0xc1, 0x34, 0xeb, 0xf9, 0x80, 0x55, 0xb9, 0xe6, 0x95, 0x14, 0x55,
	0x06, 0x41, 0xb5, 0x48, 0x59, 0x72, 0x71, 0x23, 0x09, 0x0b, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56,
	0x42, 0x8d, 0x07, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0,
	0x62, 0x10, 0x8e, 0x15, 0x93, 0x05, 0xa3, 0x53, 0x20, 0x97, 0x7c, 0x66, 0xbe, 0x5e, 0x7e, 0x41,
	0x6a, 0x5e, 0x72, 0x6a, 0x5e, 0x71, 0x69, 0x31, 0xc4, 0xf9, 0x7a, 0x70, 0xdf, 0x94, 0x19, 0x3a,
	0xf1, 0xc2, 0xec, 0x0e, 0x00, 0x49, 0x05, 0x30, 0xbe, 0x62, 0x92, 0xf1, 0x2f, 0x48, 0xcd, 0x73,
	0x86, 0xa8, 0x07, 0x0b, 0x22, 0x9c, 0x17, 0x66, 0x98, 0xc4, 0x06, 0x36, 0xc2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x19, 0xf5, 0x7d, 0x99, 0x13, 0x01, 0x00, 0x00,
}
