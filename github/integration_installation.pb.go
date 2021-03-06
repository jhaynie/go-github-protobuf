// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: integration_installation.proto

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

type IntegrationInstallation struct {
	Action       string        `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Installation *Installation `protobuf:"bytes,2,opt,name=installation" json:"installation,omitempty"`
	Sender       *User         `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
}

func (m *IntegrationInstallation) Reset()         { *m = IntegrationInstallation{} }
func (m *IntegrationInstallation) String() string { return proto.CompactTextString(m) }
func (*IntegrationInstallation) ProtoMessage()    {}
func (*IntegrationInstallation) Descriptor() ([]byte, []int) {
	return fileDescriptorIntegrationInstallation, []int{0}
}

func init() {
	proto.RegisterType((*IntegrationInstallation)(nil), "github.IntegrationInstallation")
}
func (m *IntegrationInstallation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntegrationInstallation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIntegrationInstallation(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if m.Installation != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIntegrationInstallation(dAtA, i, uint64(m.Installation.Size()))
		n1, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Sender != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintIntegrationInstallation(dAtA, i, uint64(m.Sender.Size()))
		n2, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64IntegrationInstallation(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32IntegrationInstallation(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIntegrationInstallation(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IntegrationInstallation) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovIntegrationInstallation(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovIntegrationInstallation(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovIntegrationInstallation(uint64(l))
	}
	return n
}

func sovIntegrationInstallation(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIntegrationInstallation(x uint64) (n int) {
	return sovIntegrationInstallation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntegrationInstallation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntegrationInstallation
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
			return fmt.Errorf("proto: IntegrationInstallation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntegrationInstallation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationInstallation
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
				return ErrInvalidLengthIntegrationInstallation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationInstallation
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
				return ErrInvalidLengthIntegrationInstallation
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntegrationInstallation
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
				return ErrInvalidLengthIntegrationInstallation
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
		default:
			iNdEx = preIndex
			skippy, err := skipIntegrationInstallation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIntegrationInstallation
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
func skipIntegrationInstallation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIntegrationInstallation
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
					return 0, ErrIntOverflowIntegrationInstallation
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
					return 0, ErrIntOverflowIntegrationInstallation
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
				return 0, ErrInvalidLengthIntegrationInstallation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIntegrationInstallation
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
				next, err := skipIntegrationInstallation(dAtA[start:])
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
	ErrInvalidLengthIntegrationInstallation = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIntegrationInstallation   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("integration_installation.proto", fileDescriptorIntegrationInstallation)
}

var fileDescriptorIntegrationInstallation = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcc, 0x2b, 0x49,
	0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x01,
	0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xa4,
	0x74, 0x21, 0xb4, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e, 0x58, 0x3a, 0xa9,
	0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0x88, 0x36, 0x29, 0xae, 0xd2, 0xe2, 0xd4, 0x22, 0x28,
	0x5b, 0x08, 0xd3, 0x58, 0xa5, 0x89, 0x8c, 0x5c, 0xe2, 0x9e, 0x08, 0x9b, 0x3d, 0x91, 0x54, 0x08,
	0x89, 0x71, 0xb1, 0x25, 0x26, 0x83, 0x58, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e,
	0x90, 0x05, 0x17, 0x0f, 0xb2, 0x49, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x50,
	0x97, 0x21, 0x9b, 0x11, 0x84, 0xa2, 0x52, 0x48, 0x85, 0x8b, 0xad, 0x38, 0x35, 0x2f, 0x25, 0xb5,
	0x48, 0x82, 0x19, 0xac, 0x87, 0x07, 0xa6, 0x27, 0xb4, 0x38, 0xb5, 0x28, 0x08, 0x2a, 0xe7, 0x24,
	0x72, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0xc1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x62, 0xce, 0x4e, 0x40, 0x29, 0x01, 0x00, 0x00,
}
