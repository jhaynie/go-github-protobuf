// Code generated by protoc-gen-gogo.
// source: fork_apply_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ForkApplyEvent struct {
	Head         string        `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Header       string        `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	After        string        `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	Installation *Installation `protobuf:"bytes,4,opt,name=installation" json:"installation,omitempty"`
}

func (m *ForkApplyEvent) Reset()                    { *m = ForkApplyEvent{} }
func (m *ForkApplyEvent) String() string            { return proto.CompactTextString(m) }
func (*ForkApplyEvent) ProtoMessage()               {}
func (*ForkApplyEvent) Descriptor() ([]byte, []int) { return fileDescriptorForkApplyEvent, []int{0} }

func init() {
	proto.RegisterType((*ForkApplyEvent)(nil), "github.ForkApplyEvent")
}
func (m *ForkApplyEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ForkApplyEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Head) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintForkApplyEvent(data, i, uint64(len(m.Head)))
		i += copy(data[i:], m.Head)
	}
	if len(m.Header) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintForkApplyEvent(data, i, uint64(len(m.Header)))
		i += copy(data[i:], m.Header)
	}
	if len(m.After) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintForkApplyEvent(data, i, uint64(len(m.After)))
		i += copy(data[i:], m.After)
	}
	if m.Installation != nil {
		data[i] = 0x22
		i++
		i = encodeVarintForkApplyEvent(data, i, uint64(m.Installation.Size()))
		n1, err := m.Installation.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64ForkApplyEvent(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ForkApplyEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintForkApplyEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ForkApplyEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Head)
	if l > 0 {
		n += 1 + l + sovForkApplyEvent(uint64(l))
	}
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sovForkApplyEvent(uint64(l))
	}
	l = len(m.After)
	if l > 0 {
		n += 1 + l + sovForkApplyEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovForkApplyEvent(uint64(l))
	}
	return n
}

func sovForkApplyEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozForkApplyEvent(x uint64) (n int) {
	return sovForkApplyEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForkApplyEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowForkApplyEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ForkApplyEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForkApplyEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthForkApplyEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Head = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthForkApplyEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Header = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthForkApplyEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.After = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthForkApplyEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Installation == nil {
				m.Installation = &Installation{}
			}
			if err := m.Installation.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipForkApplyEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthForkApplyEvent
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
func skipForkApplyEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowForkApplyEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowForkApplyEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthForkApplyEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowForkApplyEvent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipForkApplyEvent(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthForkApplyEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowForkApplyEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fork_apply_event.proto", fileDescriptorForkApplyEvent) }

var fileDescriptorForkApplyEvent = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcb, 0x2f, 0xca,
	0x8e, 0x4f, 0x2c, 0x28, 0xc8, 0xa9, 0x8c, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xd2, 0x85, 0xd0, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe9, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc,
	0x01, 0xb3, 0x20, 0xda, 0xa4, 0x84, 0x32, 0xf3, 0x8a, 0x4b, 0x12, 0x73, 0x72, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x20, 0x62, 0x4a, 0x3d, 0x8c, 0x5c, 0x7c, 0x6e, 0x40, 0x5b, 0x1c, 0x41, 0x96, 0xb8,
	0x82, 0xec, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x4d, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x40, 0x74, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x14, 0xca,
	0x13, 0x12, 0xe1, 0x62, 0x4d, 0x4c, 0x2b, 0x01, 0x0a, 0x33, 0x83, 0x85, 0x21, 0x1c, 0x21, 0x0b,
	0x2e, 0x1e, 0x64, 0xab, 0x24, 0x58, 0x80, 0x92, 0xdc, 0x46, 0x22, 0x7a, 0x50, 0xe7, 0x7a, 0x22,
	0xc9, 0x05, 0xa1, 0xa8, 0x74, 0x12, 0x39, 0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x40, 0xfc, 0x00, 0x88, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xbb, 0xd5, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x67, 0xe5, 0x7e, 0x4a, 0x10, 0x01, 0x00, 0x00,
}
