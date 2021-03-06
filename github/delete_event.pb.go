// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_event.proto

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

type DeleteEvent struct {
	Ref          string        `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	RefType      string        `protobuf:"bytes,2,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
	PusherType   string        `protobuf:"bytes,3,opt,name=pusher_type,json=pusherType,proto3" json:"pusher_type,omitempty"`
	Repository   *Repository   `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Sender       *User         `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,6,opt,name=installation" json:"installation,omitempty"`
	Organization *User         `protobuf:"bytes,7,opt,name=organization" json:"organization,omitempty"`
}

func (m *DeleteEvent) Reset()                    { *m = DeleteEvent{} }
func (m *DeleteEvent) String() string            { return proto.CompactTextString(m) }
func (*DeleteEvent) ProtoMessage()               {}
func (*DeleteEvent) Descriptor() ([]byte, []int) { return fileDescriptorDeleteEvent, []int{0} }

func init() {
	proto.RegisterType((*DeleteEvent)(nil), "github.DeleteEvent")
}
func (m *DeleteEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ref) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	if len(m.RefType) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(len(m.RefType)))
		i += copy(dAtA[i:], m.RefType)
	}
	if len(m.PusherType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(len(m.PusherType)))
		i += copy(dAtA[i:], m.PusherType)
	}
	if m.Repository != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(m.Repository.Size()))
		n1, err := m.Repository.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Sender != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(m.Sender.Size()))
		n2, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Installation != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(m.Installation.Size()))
		n3, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Organization != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDeleteEvent(dAtA, i, uint64(m.Organization.Size()))
		n4, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64DeleteEvent(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32DeleteEvent(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeleteEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DeleteEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	l = len(m.RefType)
	if l > 0 {
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	l = len(m.PusherType)
	if l > 0 {
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovDeleteEvent(uint64(l))
	}
	return n
}

func sovDeleteEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeleteEvent(x uint64) (n int) {
	return sovDeleteEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeleteEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeleteEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PusherType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PusherType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Installation == nil {
				m.Installation = &Installation{}
			}
			if err := m.Installation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeleteEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeleteEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Organization == nil {
				m.Organization = &User{}
			}
			if err := m.Organization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeleteEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeleteEvent
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
func skipDeleteEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeleteEvent
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
					return 0, ErrIntOverflowDeleteEvent
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
					return 0, ErrIntOverflowDeleteEvent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDeleteEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeleteEvent
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
				next, err := skipDeleteEvent(dAtA[start:])
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
	ErrInvalidLengthDeleteEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeleteEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("delete_event.proto", fileDescriptorDeleteEvent) }

var fileDescriptorDeleteEvent = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0x7b, 0x2d, 0xa4, 0xe0, 0x76, 0xa8, 0x4e, 0x1d, 0x42, 0x87, 0x50, 0x21, 0x86, 0x2e,
	0xa4, 0xa8, 0x2c, 0xcc, 0x08, 0x06, 0xd6, 0x08, 0xe6, 0x2a, 0xa1, 0x4e, 0x1a, 0x29, 0xe4, 0x22,
	0xdf, 0x05, 0x29, 0x3c, 0x09, 0x2b, 0x6f, 0xd3, 0x91, 0x47, 0x80, 0xf0, 0x22, 0xa8, 0xbe, 0x14,
	0x52, 0x89, 0x29, 0xf6, 0xff, 0x7d, 0x96, 0xe3, 0x03, 0xb9, 0xc2, 0x0c, 0x0d, 0x2e, 0xf1, 0x05,
	0x73, 0xe3, 0x17, 0xa4, 0x8c, 0x92, 0x4e, 0x92, 0x9a, 0x75, 0x19, 0x4d, 0x2e, 0xec, 0xd7, 0x7f,
	0x52, 0xcf, 0xf3, 0x44, 0x25, 0x6a, 0xce, 0x38, 0x2a, 0x63, 0xee, 0xb8, 0xe1, 0xca, 0x8e, 0x4d,
	0xa0, 0xd4, 0x48, 0x4d, 0x3d, 0x22, 0x2c, 0x94, 0x4e, 0x8d, 0xa2, 0xaa, 0x49, 0x64, 0x9a, 0x6b,
	0x13, 0x66, 0x59, 0x68, 0x52, 0x95, 0xdb, 0xec, 0xec, 0xbd, 0x0b, 0x83, 0x5b, 0xde, 0x7f, 0xb7,
	0x5d, 0x2f, 0x47, 0xd0, 0x23, 0x8c, 0x5d, 0x31, 0x15, 0xb3, 0xe3, 0x60, 0x5b, 0xca, 0x13, 0x38,
	0x22, 0x8c, 0x97, 0xa6, 0x2a, 0xd0, 0xed, 0x72, 0xdc, 0x27, 0x8c, 0x1f, 0xaa, 0x02, 0xe5, 0x29,
	0x0c, 0x8a, 0x52, 0xaf, 0x91, 0x2c, 0xed, 0x31, 0x05, 0x1b, 0xb1, 0xb0, 0x00, 0xf8, 0xfb, 0x0b,
	0xf7, 0x60, 0x2a, 0x66, 0x83, 0x85, 0xf4, 0x9b, 0x9b, 0x82, 0x5f, 0x12, 0xb4, 0x2c, 0x79, 0x0e,
	0x8e, 0xc6, 0x7c, 0x85, 0xe4, 0x1e, 0xb2, 0x3f, 0xdc, 0xf9, 0x8f, 0x1a, 0x29, 0x68, 0x98, 0xbc,
	0x86, 0x61, 0xfb, 0x1a, 0xd7, 0x61, 0x77, 0xbc, 0x73, 0xef, 0x5b, 0x2c, 0xd8, 0x33, 0xe5, 0x25,
	0x0c, 0x15, 0x25, 0x61, 0x9e, 0xbe, 0xda, 0xc9, 0xfe, 0x3f, 0x5b, 0xf6, 0x8c, 0x9b, 0xf1, 0xe6,
	0xcb, 0xeb, 0x6c, 0x6a, 0x4f, 0x7c, 0xd4, 0x9e, 0xf8, 0xac, 0x3d, 0xf1, 0xf6, 0xed, 0x75, 0x22,
	0x87, 0x1f, 0xf0, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xee, 0x82, 0x7e, 0xbf, 0x01, 0x00,
	0x00,
}
