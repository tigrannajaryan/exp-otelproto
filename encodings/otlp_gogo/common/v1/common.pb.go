// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/common/v1/common.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ValueType is the enumeration of possible types that value can have.
type ValueType int32

const (
	ValueType_STRING ValueType = 0
	ValueType_INT    ValueType = 1
	ValueType_DOUBLE ValueType = 2
	ValueType_BOOL   ValueType = 3
)

var ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
}

var ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62ba46dcb97aa817, []int{0}
}

type AnyValue struct {
	// type of the value.
	Type         ValueType            `protobuf:"varint,1,opt,name=type,proto3,enum=opentelemetrygogo.proto.common.v1.ValueType" json:"type,omitempty"`
	BoolValue    bool                 `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue  string               `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue     int64                `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue  float64              `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	ListValues   []*AnyValue          `protobuf:"bytes,6,rep,name=list_values,json=listValues,proto3" json:"list_values,omitempty"`
	KvlistValues []*AttributeKeyValue `protobuf:"bytes,7,rep,name=kvlist_values,json=kvlistValues,proto3" json:"kvlist_values,omitempty"`
	BytesValue   []byte               `protobuf:"bytes,8,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
}

func (m *AnyValue) Reset()         { *m = AnyValue{} }
func (m *AnyValue) String() string { return proto.CompactTextString(m) }
func (*AnyValue) ProtoMessage()    {}
func (*AnyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ba46dcb97aa817, []int{0}
}
func (m *AnyValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyValue.Merge(m, src)
}
func (m *AnyValue) XXX_Size() int {
	return m.Size()
}
func (m *AnyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AnyValue proto.InternalMessageInfo

func (m *AnyValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_STRING
}

func (m *AnyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AnyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AnyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AnyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AnyValue) GetListValues() []*AnyValue {
	if m != nil {
		return m.ListValues
	}
	return nil
}

func (m *AnyValue) GetKvlistValues() []*AttributeKeyValue {
	if m != nil {
		return m.KvlistValues
	}
	return nil
}

func (m *AnyValue) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	//  string string_value = 3;
	//  int64 int_value = 4;
	//  double double_value = 5;
	//  bool bool_value = 6;
	Value AnyValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ba46dcb97aa817, []int{1}
}
func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return m.Size()
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

func (m *AttributeKeyValue) GetValue() AnyValue {
	if m != nil {
		return m.Value
	}
	return AnyValue{}
}

// StringKeyValue is a pair of key/value strings. This is the simpler (and faster) version
// of AttributeKeyValue that only supports string values.
type StringKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringKeyValue) Reset()         { *m = StringKeyValue{} }
func (m *StringKeyValue) String() string { return proto.CompactTextString(m) }
func (*StringKeyValue) ProtoMessage()    {}
func (*StringKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ba46dcb97aa817, []int{2}
}
func (m *StringKeyValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringKeyValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringKeyValue.Merge(m, src)
}
func (m *StringKeyValue) XXX_Size() int {
	return m.Size()
}
func (m *StringKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringKeyValue proto.InternalMessageInfo

func (m *StringKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// InstrumentationLibrary is a message representing the instrumentation library information
// such as the fully qualified name and version.
type InstrumentationLibrary struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *InstrumentationLibrary) Reset()         { *m = InstrumentationLibrary{} }
func (m *InstrumentationLibrary) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibrary) ProtoMessage()    {}
func (*InstrumentationLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ba46dcb97aa817, []int{3}
}
func (m *InstrumentationLibrary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstrumentationLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstrumentationLibrary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstrumentationLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibrary.Merge(m, src)
}
func (m *InstrumentationLibrary) XXX_Size() int {
	return m.Size()
}
func (m *InstrumentationLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibrary proto.InternalMessageInfo

func (m *InstrumentationLibrary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstrumentationLibrary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("opentelemetrygogo.proto.common.v1.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*AnyValue)(nil), "opentelemetrygogo.proto.common.v1.AnyValue")
	proto.RegisterType((*AttributeKeyValue)(nil), "opentelemetrygogo.proto.common.v1.AttributeKeyValue")
	proto.RegisterType((*StringKeyValue)(nil), "opentelemetrygogo.proto.common.v1.StringKeyValue")
	proto.RegisterType((*InstrumentationLibrary)(nil), "opentelemetrygogo.proto.common.v1.InstrumentationLibrary")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/common/v1/common.proto", fileDescriptor_62ba46dcb97aa817)
}

var fileDescriptor_62ba46dcb97aa817 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x33, 0x75, 0x9a, 0xc4, 0xd7, 0xf9, 0xab, 0xfc, 0x23, 0x84, 0x22, 0x10, 0xae, 0x9b,
	0x95, 0x55, 0x68, 0xac, 0x16, 0x84, 0x58, 0x42, 0xa0, 0x54, 0x11, 0x51, 0x53, 0xb9, 0x01, 0x09,
	0x36, 0x95, 0xdd, 0x0e, 0x66, 0xa8, 0x33, 0x63, 0xd9, 0x63, 0x0b, 0xbf, 0x45, 0x97, 0x3c, 0x52,
	0x97, 0x5d, 0xb2, 0x42, 0x28, 0x79, 0x11, 0xe4, 0x99, 0x49, 0x9a, 0x82, 0x04, 0x65, 0x77, 0x7d,
	0xe6, 0x9c, 0xef, 0x6a, 0xee, 0xf8, 0xc2, 0x36, 0x4f, 0x08, 0x13, 0x24, 0x26, 0x53, 0x22, 0xd2,
	0xd2, 0x4b, 0x52, 0x2e, 0xb8, 0x77, 0xca, 0xa7, 0x53, 0xce, 0xbc, 0x62, 0x57, 0x57, 0x7d, 0x29,
	0xe3, 0xad, 0x1b, 0xde, 0x88, 0x47, 0x5c, 0x1d, 0xf4, 0xb5, 0xab, 0xd8, 0xbd, 0xb7, 0x13, 0x51,
	0xf1, 0x29, 0x0f, 0x2b, 0xc5, 0xab, 0x0c, 0x0a, 0x18, 0xe6, 0x1f, 0xe5, 0x97, 0xa2, 0x5f, 0x07,
	0x7b, 0x5f, 0x0d, 0x68, 0xbd, 0x60, 0xe5, 0xbb, 0x20, 0xce, 0x09, 0x7e, 0x0e, 0x75, 0x51, 0x26,
	0xa4, 0x8b, 0x1c, 0xe4, 0x6e, 0xec, 0x3d, 0xea, 0xff, 0xb5, 0x5b, 0x5f, 0xe6, 0x26, 0x65, 0x42,
	0x7c, 0x99, 0xc4, 0x0f, 0x00, 0x42, 0xce, 0xe3, 0x93, 0xa2, 0xd2, 0xbb, 0x6b, 0x0e, 0x72, 0x5b,
	0xbe, 0x59, 0x29, 0xaa, 0xc1, 0x16, 0xb4, 0x33, 0x91, 0x52, 0x16, 0x69, 0x83, 0xe1, 0x20, 0xd7,
	0xf4, 0x2d, 0xa5, 0x29, 0xcb, 0x7d, 0x30, 0x29, 0x13, 0xfa, 0xbc, 0xee, 0x20, 0xd7, 0xf0, 0x5b,
	0x94, 0x89, 0x65, 0xfe, 0x8c, 0xe7, 0x61, 0x4c, 0xf4, 0xf9, 0xba, 0x83, 0x5c, 0xe4, 0x5b, 0x4a,
	0x53, 0x96, 0x11, 0x58, 0x31, 0xcd, 0x34, 0x20, 0xeb, 0x36, 0x1c, 0xc3, 0xb5, 0xf6, 0x1e, 0xde,
	0xe2, 0x2a, 0x8b, 0x29, 0xf8, 0x50, 0xe5, 0x65, 0x99, 0xe1, 0xf7, 0xf0, 0xdf, 0x79, 0xb1, 0xca,
	0x6b, 0x4a, 0xde, 0x93, 0xdb, 0xf0, 0x84, 0x48, 0x69, 0x98, 0x0b, 0xf2, 0x86, 0x68, 0x70, 0x5b,
	0xa1, 0x34, 0x7a, 0x13, 0xac, 0xb0, 0x14, 0x24, 0xd3, 0x57, 0x69, 0x39, 0xc8, 0x6d, 0xfb, 0x20,
	0x25, 0xe9, 0xe8, 0x31, 0xf8, 0xff, 0x37, 0x06, 0xee, 0x80, 0x71, 0x4e, 0x4a, 0xf9, 0x42, 0xa6,
	0x5f, 0x95, 0xf8, 0x00, 0xd6, 0xaf, 0xa7, 0xfd, 0x6f, 0x57, 0x1d, 0xd4, 0x2f, 0xbf, 0x6f, 0xd6,
	0x7c, 0x95, 0xef, 0x3d, 0x83, 0x8d, 0x63, 0xf9, 0x10, 0x7f, 0x68, 0x76, 0x67, 0xb5, 0x99, 0xb9,
	0x48, 0xbe, 0x86, 0xbb, 0x43, 0x96, 0x89, 0x34, 0x9f, 0x12, 0x26, 0x02, 0x41, 0x39, 0x1b, 0xd1,
	0x30, 0x0d, 0xd2, 0x12, 0x63, 0xa8, 0xb3, 0x60, 0x4a, 0x34, 0x42, 0xd6, 0xb8, 0x0b, 0xcd, 0x82,
	0xa4, 0x19, 0xe5, 0x4c, 0x53, 0x16, 0x9f, 0xdb, 0x4f, 0xc1, 0x5c, 0xfe, 0x50, 0x18, 0xa0, 0x71,
	0x3c, 0xf1, 0x87, 0x87, 0x07, 0x9d, 0x1a, 0x6e, 0x82, 0x31, 0x3c, 0x9c, 0x74, 0x50, 0x25, 0xbe,
	0x1a, 0xbf, 0x1d, 0x8c, 0xf6, 0x3b, 0x6b, 0xb8, 0x05, 0xf5, 0xc1, 0x78, 0x3c, 0xea, 0x18, 0x83,
	0x0b, 0x74, 0x39, 0xb3, 0xd1, 0xd5, 0xcc, 0x46, 0x3f, 0x66, 0x36, 0xba, 0x98, 0xdb, 0xb5, 0xab,
	0xb9, 0x5d, 0xfb, 0x36, 0xb7, 0x6b, 0xe0, 0x50, 0x7e, 0x73, 0x20, 0xbf, 0x0e, 0x63, 0x60, 0xbd,
	0x94, 0xe5, 0x51, 0x25, 0x1f, 0xa1, 0x0f, 0xfb, 0x2b, 0xfb, 0x23, 0x68, 0x94, 0x06, 0x8c, 0x05,
	0x9f, 0x83, 0xb4, 0x0c, 0x98, 0x47, 0xbe, 0x24, 0x3b, 0x5c, 0x90, 0x58, 0x2d, 0x11, 0x61, 0xa7,
	0xfc, 0x8c, 0xb2, 0x28, 0xf3, 0xb8, 0x88, 0x93, 0x13, 0xb9, 0x6b, 0xcb, 0xb5, 0x0d, 0x1b, 0xd2,
	0xf4, 0xf8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xa6, 0x04, 0xb7, 0xde, 0x03, 0x00, 0x00,
}

func (m *AnyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BytesValue) > 0 {
		i -= len(m.BytesValue)
		copy(dAtA[i:], m.BytesValue)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.BytesValue)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.KvlistValues) > 0 {
		for iNdEx := len(m.KvlistValues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KvlistValues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ListValues) > 0 {
		for iNdEx := len(m.ListValues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListValues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.DoubleValue != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DoubleValue))))
		i--
		dAtA[i] = 0x29
	}
	if m.IntValue != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.IntValue))
		i--
		dAtA[i] = 0x20
	}
	if len(m.StringValue) > 0 {
		i -= len(m.StringValue)
		copy(dAtA[i:], m.StringValue)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.StringValue)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BoolValue {
		i--
		if m.BoolValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AttributeKeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttributeKeyValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttributeKeyValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StringKeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringKeyValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringKeyValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InstrumentationLibrary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstrumentationLibrary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstrumentationLibrary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnyValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	if m.BoolValue {
		n += 2
	}
	l = len(m.StringValue)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.IntValue != 0 {
		n += 1 + sovCommon(uint64(m.IntValue))
	}
	if m.DoubleValue != 0 {
		n += 9
	}
	if len(m.ListValues) > 0 {
		for _, e := range m.ListValues {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.KvlistValues) > 0 {
		for _, e := range m.KvlistValues {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	l = len(m.BytesValue)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *AttributeKeyValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = m.Value.Size()
	n += 1 + l + sovCommon(uint64(l))
	return n
}

func (m *StringKeyValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *InstrumentationLibrary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AnyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ValueType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BoolValue = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntValue", wireType)
			}
			m.IntValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntValue |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DoubleValue = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListValues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListValues = append(m.ListValues, &AnyValue{})
			if err := m.ListValues[len(m.ListValues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KvlistValues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KvlistValues = append(m.KvlistValues, &AttributeKeyValue{})
			if err := m.KvlistValues[len(m.KvlistValues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesValue = append(m.BytesValue[:0], dAtA[iNdEx:postIndex]...)
			if m.BytesValue == nil {
				m.BytesValue = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AttributeKeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AttributeKeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttributeKeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StringKeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringKeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringKeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstrumentationLibrary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstrumentationLibrary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstrumentationLibrary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
