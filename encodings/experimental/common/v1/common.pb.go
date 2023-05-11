// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: common/v1/common.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AnyValue is used to represent any type of attribute value. AnyValue may contain a
// primitive value such as a string or integer or it may contain an arbitrary nested
// object containing arrays, key-value lists and primitives.
type AnyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value is one of the listed fields. It is valid for all values to be unspecified
	// in which case this AnyValue is considered to be "empty".
	//
	// Types that are assignable to Value:
	//	*AnyValue_StringValue
	//	*AnyValue_BoolValue
	//	*AnyValue_IntValue
	//	*AnyValue_DoubleValue
	//	*AnyValue_ArrayValue
	//	*AnyValue_KvlistValue
	//	*AnyValue_BytesValue
	Value isAnyValue_Value `protobuf_oneof:"value"`
}

func (x *AnyValue) Reset() {
	*x = AnyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyValue) ProtoMessage() {}

func (x *AnyValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyValue.ProtoReflect.Descriptor instead.
func (*AnyValue) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (m *AnyValue) GetValue() isAnyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *AnyValue) GetStringValue() string {
	if x, ok := x.GetValue().(*AnyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *AnyValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*AnyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *AnyValue) GetIntValue() int64 {
	if x, ok := x.GetValue().(*AnyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *AnyValue) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*AnyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *AnyValue) GetArrayValue() *ArrayValue {
	if x, ok := x.GetValue().(*AnyValue_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (x *AnyValue) GetKvlistValue() *KeyValueList {
	if x, ok := x.GetValue().(*AnyValue_KvlistValue); ok {
		return x.KvlistValue
	}
	return nil
}

func (x *AnyValue) GetBytesValue() []byte {
	if x, ok := x.GetValue().(*AnyValue_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

type isAnyValue_Value interface {
	isAnyValue_Value()
}

type AnyValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AnyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type AnyValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AnyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type AnyValue_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,5,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type AnyValue_KvlistValue struct {
	KvlistValue *KeyValueList `protobuf:"bytes,6,opt,name=kvlist_value,json=kvlistValue,proto3,oneof"`
}

type AnyValue_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*AnyValue_StringValue) isAnyValue_Value() {}

func (*AnyValue_BoolValue) isAnyValue_Value() {}

func (*AnyValue_IntValue) isAnyValue_Value() {}

func (*AnyValue_DoubleValue) isAnyValue_Value() {}

func (*AnyValue_ArrayValue) isAnyValue_Value() {}

func (*AnyValue_KvlistValue) isAnyValue_Value() {}

func (*AnyValue_BytesValue) isAnyValue_Value() {}

// ArrayValue is a list of AnyValue messages. We need ArrayValue as a message
// since oneof in AnyValue does not allow repeated fields.
type ArrayValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of values. The array may be empty (contain 0 elements).
	Values []*AnyValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ArrayValue) Reset() {
	*x = ArrayValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayValue) ProtoMessage() {}

func (x *ArrayValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayValue.ProtoReflect.Descriptor instead.
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ArrayValue) GetValues() []*AnyValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// KeyValueList is a list of KeyValue messages. We need KeyValueList as a message
// since `oneof` in AnyValue does not allow repeated fields. Everywhere else where we need
// a list of KeyValue messages (e.g. in Span) we use `repeated KeyValue` directly to
// avoid unnecessary extra wrapping (which slows down the protocol). The 2 approaches
// are semantically equivalent.
type KeyValueList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of key/value pairs of key-value pairs. The list may be empty (may
	// contain 0 elements).
	// The keys MUST be unique (it is not allowed to have more than one
	// value with the same key).
	Values []*KeyValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *KeyValueList) Reset() {
	*x = KeyValueList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueList) ProtoMessage() {}

func (x *KeyValueList) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueList.ProtoReflect.Descriptor instead.
func (*KeyValueList) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueList) GetValues() []*KeyValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// KeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *AnyValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() *AnyValue {
	if x != nil {
		return x.Value
	}
	return nil
}

// InstrumentationScope is a message representing the instrumentation scope information
// such as the fully qualified name and version.
type InstrumentationScope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An empty instrumentation scope name means the name is unknown.
	Name                   string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version                string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Attributes             []*KeyValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32      `protobuf:"varint,4,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
}

func (x *InstrumentationScope) Reset() {
	*x = InstrumentationScope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstrumentationScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentationScope) ProtoMessage() {}

func (x *InstrumentationScope) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentationScope.ProtoReflect.Descriptor instead.
func (*InstrumentationScope) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *InstrumentationScope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstrumentationScope) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InstrumentationScope) GetAttributes() []*KeyValue {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *InstrumentationScope) GetDroppedAttributesCount() uint32 {
	if x != nil {
		return x.DroppedAttributesCount
	}
	return 0
}

var File_common_v1_common_proto protoreflect.FileDescriptor

var file_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x22, 0xd2, 0x02, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49,
	0x0a, 0x0c, 0x6b, 0x76, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x76,
	0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x48, 0x0a,
	0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc0, 0x01,
	0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65,
	0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x8d, 0x01, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x6e,
	0x6e, 0x61, 0x6a, 0x61, 0x72, 0x79, 0x61, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x2d, 0x6f, 0x74, 0x65,
	0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x16, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_common_proto_rawDescOnce sync.Once
	file_common_v1_common_proto_rawDescData = file_common_v1_common_proto_rawDesc
)

func file_common_v1_common_proto_rawDescGZIP() []byte {
	file_common_v1_common_proto_rawDescOnce.Do(func() {
		file_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_common_proto_rawDescData)
	})
	return file_common_v1_common_proto_rawDescData
}

var file_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_v1_common_proto_goTypes = []interface{}{
	(*AnyValue)(nil),             // 0: experimental.common.v1.AnyValue
	(*ArrayValue)(nil),           // 1: experimental.common.v1.ArrayValue
	(*KeyValueList)(nil),         // 2: experimental.common.v1.KeyValueList
	(*KeyValue)(nil),             // 3: experimental.common.v1.KeyValue
	(*InstrumentationScope)(nil), // 4: experimental.common.v1.InstrumentationScope
}
var file_common_v1_common_proto_depIdxs = []int32{
	1, // 0: experimental.common.v1.AnyValue.array_value:type_name -> experimental.common.v1.ArrayValue
	2, // 1: experimental.common.v1.AnyValue.kvlist_value:type_name -> experimental.common.v1.KeyValueList
	0, // 2: experimental.common.v1.ArrayValue.values:type_name -> experimental.common.v1.AnyValue
	3, // 3: experimental.common.v1.KeyValueList.values:type_name -> experimental.common.v1.KeyValue
	0, // 4: experimental.common.v1.KeyValue.value:type_name -> experimental.common.v1.AnyValue
	3, // 5: experimental.common.v1.InstrumentationScope.attributes:type_name -> experimental.common.v1.KeyValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_v1_common_proto_init() }
func file_common_v1_common_proto_init() {
	if File_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstrumentationScope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_v1_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AnyValue_StringValue)(nil),
		(*AnyValue_BoolValue)(nil),
		(*AnyValue_IntValue)(nil),
		(*AnyValue_DoubleValue)(nil),
		(*AnyValue_ArrayValue)(nil),
		(*AnyValue_KvlistValue)(nil),
		(*AnyValue_BytesValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_common_proto_goTypes,
		DependencyIndexes: file_common_v1_common_proto_depIdxs,
		MessageInfos:      file_common_v1_common_proto_msgTypes,
	}.Build()
	File_common_v1_common_proto = out.File
	file_common_v1_common_proto_rawDesc = nil
	file_common_v1_common_proto_goTypes = nil
	file_common_v1_common_proto_depIdxs = nil
}
