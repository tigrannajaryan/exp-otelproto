package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"

	v1 "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo/common/v1"
)

type AnyValueCustom struct {
	// type of the value.
	Type         v1.ValueType           `protobuf:"varint,1,opt,name=type,proto3,enum=opentelemetrygogo2.proto.common.v1.v1.ValueType" json:"type,omitempty"`
	BoolValue    bool                `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue  string              `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue     int64               `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue  float64             `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	ListValues   []AnyValueCustom          `protobuf:"bytes,6,rep,name=list_values,json=listValues,proto3" json:"list_values"`
	KvlistValues []v1.AttributeKeyValue `protobuf:"bytes,7,rep,name=kvlist_values,json=kvlistValues,proto3" json:"kvlist_values"`
	BytesValue   []byte              `protobuf:"bytes,8,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
}

//func (m *AnyValueCustom) XXX_Unmarshal(b []byte) error {
//	return m.Unmarshal(b)
//}
//func (m *AnyValueCustom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	if deterministic {
//		return xxx_messageInfo_AnyValueCustom.Marshal(b, m, deterministic)
//	} else {
//		b = b[:cap(b)]
//		n, err := m.MarshalToSizedBuffer(b)
//		if err != nil {
//			return nil, err
//		}
//		return b[:n], nil
//	}
//}
//func (m *AnyValueCustom) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_AnyValueCustom.Merge(m, src)
//}
//func (m *AnyValueCustom) XXX_Size() int {
//	return m.Size()
//}
//func (m *AnyValueCustom) XXX_DiscardUnknown() {
//	xxx_messageInfo_AnyValueCustom.DiscardUnknown(m)
//}

var xxx_messageInfo_AnyValueCustom proto.InternalMessageInfo

func (m *AnyValueCustom) GetType() v1.ValueType {
	if m != nil {
		return m.Type
	}
	return v1.ValueType_STRING
}

func (m *AnyValueCustom) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AnyValueCustom) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AnyValueCustom) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AnyValueCustom) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AnyValueCustom) GetListValues() []AnyValueCustom {
	if m != nil {
		return m.ListValues
	}
	return nil
}

func (m *AnyValueCustom) GetKvlistValues() []v1.AttributeKeyValue {
	if m != nil {
		return m.KvlistValues
	}
	return nil
}

func (m *AnyValueCustom) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func (m *AnyValueCustom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyValueCustom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyValueCustom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AnyValueCustom) Size() (n int) {
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

func (m *AnyValueCustom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AnyValueCustom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyValueCustom: illegal tag %d (wire type %d)", fieldNum, wire)
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
				m.Type |= v1.ValueType(b&0x7F) << shift
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
			m.ListValues = append(m.ListValues, AnyValueCustom{})
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
			m.KvlistValues = append(m.KvlistValues, v1.AttributeKeyValue{})
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
