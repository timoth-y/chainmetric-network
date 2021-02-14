// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: requirement.proto

package model

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Requirement struct {
	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Property     string  `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
	MinThreshold float64 `protobuf:"fixed64,3,opt,name=minThreshold,proto3" json:"minThreshold,omitempty"`
	MaxThreshold float64 `protobuf:"fixed64,4,opt,name=maxThreshold,proto3" json:"maxThreshold,omitempty"`
}

func (m *Requirement) Reset()         { *m = Requirement{} }
func (m *Requirement) String() string { return proto.CompactTextString(m) }
func (*Requirement) ProtoMessage()    {}
func (*Requirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b809697927d7988, []int{0}
}
func (m *Requirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Requirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Requirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Requirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Requirement.Merge(m, src)
}
func (m *Requirement) XXX_Size() int {
	return m.Size()
}
func (m *Requirement) XXX_DiscardUnknown() {
	xxx_messageInfo_Requirement.DiscardUnknown(m)
}

var xxx_messageInfo_Requirement proto.InternalMessageInfo

func (m *Requirement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Requirement) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *Requirement) GetMinThreshold() float64 {
	if m != nil {
		return m.MinThreshold
	}
	return 0
}

func (m *Requirement) GetMaxThreshold() float64 {
	if m != nil {
		return m.MaxThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*Requirement)(nil), "model.Requirement")
}

func init() { proto.RegisterFile("requirement.proto", fileDescriptor_6b809697927d7988) }

var fileDescriptor_6b809697927d7988 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x2d, 0x2c,
	0xcd, 0x2c, 0x4a, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd,
	0xcd, 0x4f, 0x49, 0xcd, 0x51, 0x6a, 0x64, 0xe4, 0xe2, 0x0e, 0x42, 0x48, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x49, 0x71, 0x71,
	0x14, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x54, 0x4a, 0x30, 0x81, 0x45, 0xe1, 0x7c, 0x21, 0x25,
	0x2e, 0x9e, 0xdc, 0xcc, 0xbc, 0x90, 0x8c, 0xa2, 0xd4, 0xe2, 0x8c, 0xfc, 0x9c, 0x14, 0x09, 0x66,
	0x05, 0x46, 0x0d, 0xc6, 0x20, 0x14, 0x31, 0xb0, 0x9a, 0xc4, 0x0a, 0x84, 0x1a, 0x16, 0xa8, 0x1a,
	0x24, 0x31, 0x27, 0x9d, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x42,
	0x72, 0x77, 0xb1, 0x3e, 0xd8, 0xc5, 0x49, 0x6c, 0x60, 0xf7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xea, 0x3e, 0x09, 0x39, 0xd4, 0x00, 0x00, 0x00,
}

func (m *Requirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Requirement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRequirement(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Property) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRequirement(dAtA, i, uint64(len(m.Property)))
		i += copy(dAtA[i:], m.Property)
	}
	if m.MinThreshold != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MinThreshold))))
		i += 8
	}
	if m.MaxThreshold != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MaxThreshold))))
		i += 8
	}
	return i, nil
}

func encodeVarintRequirement(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Requirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRequirement(uint64(l))
	}
	l = len(m.Property)
	if l > 0 {
		n += 1 + l + sovRequirement(uint64(l))
	}
	if m.MinThreshold != 0 {
		n += 9
	}
	if m.MaxThreshold != 0 {
		n += 9
	}
	return n
}

func sovRequirement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequirement(x uint64) (n int) {
	return sovRequirement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Requirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequirement
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
			return fmt.Errorf("proto: Requirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Requirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequirement
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
				return ErrInvalidLengthRequirement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequirement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Property", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequirement
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
				return ErrInvalidLengthRequirement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequirement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Property = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinThreshold", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MinThreshold = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxThreshold", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MaxThreshold = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRequirement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRequirement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRequirement
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
func skipRequirement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequirement
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
					return 0, ErrIntOverflowRequirement
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequirement
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
				return 0, ErrInvalidLengthRequirement
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRequirement
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRequirement
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRequirement(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRequirement
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRequirement = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequirement   = fmt.Errorf("proto: integer overflow")
)
