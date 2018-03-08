// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: organization_event.proto

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

type Invitation struct {
	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role  string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *Invitation) Reset()                    { *m = Invitation{} }
func (m *Invitation) String() string            { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()               {}
func (*Invitation) Descriptor() ([]byte, []int) { return fileDescriptorOrganizationEvent, []int{0} }

type Membership struct {
	Url             string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	State           string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Role            string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	OrganizationUrl string `protobuf:"bytes,4,opt,name=organization_url,json=organizationUrl,proto3" json:"organization_url,omitempty"`
	User            *User  `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
}

func (m *Membership) Reset()                    { *m = Membership{} }
func (m *Membership) String() string            { return proto.CompactTextString(m) }
func (*Membership) ProtoMessage()               {}
func (*Membership) Descriptor() ([]byte, []int) { return fileDescriptorOrganizationEvent, []int{1} }

type OrganizationEvent struct {
	Action       string        `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Invitation   *Invitation   `protobuf:"bytes,2,opt,name=invitation" json:"invitation,omitempty"`
	Membership   *Membership   `protobuf:"bytes,3,opt,name=membership" json:"membership,omitempty"`
	Organization *User         `protobuf:"bytes,4,opt,name=organization" json:"organization,omitempty"`
	Sender       *User         `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,6,opt,name=installation" json:"installation,omitempty"`
}

func (m *OrganizationEvent) Reset()         { *m = OrganizationEvent{} }
func (m *OrganizationEvent) String() string { return proto.CompactTextString(m) }
func (*OrganizationEvent) ProtoMessage()    {}
func (*OrganizationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorOrganizationEvent, []int{2}
}

func init() {
	proto.RegisterType((*Invitation)(nil), "github.Invitation")
	proto.RegisterType((*Membership)(nil), "github.Membership")
	proto.RegisterType((*OrganizationEvent)(nil), "github.OrganizationEvent")
}
func (m *Invitation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Invitation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Id))
	}
	if len(m.Login) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Login)))
		i += copy(dAtA[i:], m.Login)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}

func (m *Membership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Membership) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	if len(m.OrganizationUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.OrganizationUrl)))
		i += copy(dAtA[i:], m.OrganizationUrl)
	}
	if m.User != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *OrganizationEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrganizationEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if m.Invitation != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Invitation.Size()))
		n2, err := m.Invitation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Membership != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Membership.Size()))
		n3, err := m.Membership.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Organization != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Organization.Size()))
		n4, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Sender != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Sender.Size()))
		n5, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Installation != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOrganizationEvent(dAtA, i, uint64(m.Installation.Size()))
		n6, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeFixed64OrganizationEvent(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32OrganizationEvent(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintOrganizationEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Invitation) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovOrganizationEvent(uint64(m.Id))
	}
	l = len(m.Login)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	return n
}

func (m *Membership) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	l = len(m.OrganizationUrl)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	return n
}

func (m *OrganizationEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.Invitation != nil {
		l = m.Invitation.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.Membership != nil {
		l = m.Membership.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovOrganizationEvent(uint64(l))
	}
	return n
}

func sovOrganizationEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOrganizationEvent(x uint64) (n int) {
	return sovOrganizationEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Invitation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganizationEvent
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
			return fmt.Errorf("proto: Invitation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Invitation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Login", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Login = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrganizationEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrganizationEvent
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
func (m *Membership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganizationEvent
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
			return fmt.Errorf("proto: Membership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Membership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrganizationEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrganizationEvent
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
func (m *OrganizationEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrganizationEvent
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
			return fmt.Errorf("proto: OrganizationEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrganizationEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invitation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Invitation == nil {
				m.Invitation = &Invitation{}
			}
			if err := m.Invitation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Membership == nil {
				m.Membership = &Membership{}
			}
			if err := m.Membership.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
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
					return ErrIntOverflowOrganizationEvent
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
				return ErrInvalidLengthOrganizationEvent
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
		default:
			iNdEx = preIndex
			skippy, err := skipOrganizationEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrganizationEvent
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
func skipOrganizationEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrganizationEvent
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
					return 0, ErrIntOverflowOrganizationEvent
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
					return 0, ErrIntOverflowOrganizationEvent
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
				return 0, ErrInvalidLengthOrganizationEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOrganizationEvent
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
				next, err := skipOrganizationEvent(dAtA[start:])
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
	ErrInvalidLengthOrganizationEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrganizationEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("organization_event.proto", fileDescriptorOrganizationEvent) }

var fileDescriptorOrganizationEvent = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0x76, 0xa3, 0x06, 0x1c, 0xe5, 0xff, 0xed, 0x20, 0x25, 0x78, 0x08, 0x41, 0x7a, 0xb0, 0x87,
	0x6a, 0xb1, 0x97, 0x9e, 0x0b, 0x3d, 0xf4, 0x50, 0x0a, 0x01, 0x6f, 0x85, 0x92, 0xe8, 0x36, 0x2e,
	0x24, 0x59, 0xd9, 0x6c, 0x3c, 0xf4, 0x3d, 0x0a, 0x7d, 0x81, 0xbe, 0x8b, 0xc7, 0x3e, 0x42, 0x6b,
	0x5f, 0xa4, 0xec, 0x24, 0x9a, 0x58, 0xe8, 0xc9, 0xf9, 0xbe, 0x99, 0xf9, 0x76, 0xbe, 0xcf, 0x80,
	0x23, 0x55, 0x14, 0xa4, 0xe2, 0x25, 0xd0, 0x42, 0xa6, 0x4f, 0x7c, 0xc3, 0x53, 0x3d, 0x59, 0x2b,
	0xa9, 0x25, 0xda, 0x91, 0xd0, 0xab, 0x3c, 0x1c, 0x5e, 0x14, 0xbf, 0x93, 0x85, 0x4c, 0xa6, 0x91,
	0x8c, 0xe4, 0x94, 0xda, 0x61, 0xfe, 0x4c, 0x88, 0x00, 0x55, 0xc5, 0xda, 0x10, 0xf2, 0x8c, 0xab,
	0xb2, 0x46, 0x91, 0x66, 0x3a, 0x88, 0x63, 0x12, 0x2f, 0xb8, 0xd1, 0x23, 0xc0, 0x5d, 0xba, 0x11,
	0x9a, 0x38, 0xfc, 0x07, 0x96, 0x58, 0x3a, 0xcc, 0x63, 0xe3, 0xb6, 0x6f, 0x89, 0x25, 0x0e, 0xa0,
	0x1d, 0xcb, 0x48, 0xa4, 0x8e, 0xe5, 0xb1, 0x71, 0xc7, 0x2f, 0x80, 0x61, 0x79, 0x12, 0x88, 0xd8,
	0x69, 0x16, 0x2c, 0x01, 0x44, 0x68, 0x29, 0x19, 0x73, 0xa7, 0x45, 0x24, 0xd5, 0xa3, 0x57, 0x06,
	0x70, 0xcf, 0x93, 0x90, 0xab, 0x6c, 0x25, 0xd6, 0xd8, 0x87, 0x66, 0xae, 0x62, 0xd2, 0xef, 0xf8,
	0xa6, 0x34, 0x52, 0x99, 0x0e, 0x34, 0xdf, 0x3f, 0x40, 0xe0, 0x20, 0xd5, 0xac, 0xa4, 0xf0, 0x1c,
	0xfa, 0x47, 0xd9, 0x18, 0xa1, 0xe2, 0xa9, 0xff, 0x75, 0x7e, 0xae, 0x62, 0xf4, 0xa0, 0x65, 0x5c,
	0x3b, 0x6d, 0x8f, 0x8d, 0xbb, 0xb3, 0xde, 0xa4, 0x4c, 0x6c, 0x9e, 0x71, 0xe5, 0x53, 0x67, 0xf4,
	0x6e, 0xc1, 0xc9, 0x43, 0x6d, 0xeb, 0xd6, 0x04, 0x8d, 0xa7, 0x60, 0x07, 0x0b, 0x03, 0xcb, 0x0b,
	0x4b, 0x84, 0x33, 0x00, 0x71, 0xc8, 0x88, 0x2e, 0xed, 0xce, 0x70, 0xaf, 0x5a, 0xa5, 0xe7, 0xd7,
	0xa6, 0xcc, 0x4e, 0x72, 0x30, 0x4e, 0x46, 0x6a, 0x3b, 0x55, 0x24, 0x7e, 0x6d, 0x0a, 0x2f, 0xa1,
	0x57, 0xb7, 0x42, 0xf6, 0x7e, 0xdf, 0x7f, 0x34, 0x81, 0x67, 0x60, 0x67, 0x3c, 0x5d, 0xfe, 0xe1,
	0xb5, 0xec, 0xe1, 0x35, 0xf4, 0xea, 0xff, 0xbc, 0x63, 0xd3, 0xec, 0xa0, 0x72, 0x50, 0xf5, 0xfc,
	0xa3, 0xc9, 0x9b, 0xc1, 0xf6, 0xcb, 0x6d, 0x6c, 0x77, 0x2e, 0xfb, 0xd8, 0xb9, 0xec, 0x73, 0xe7,
	0xb2, 0xb7, 0x6f, 0xb7, 0x11, 0xda, 0xf4, 0xe9, 0x5c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x93,
	0x5f, 0x54, 0xe6, 0xad, 0x02, 0x00, 0x00,
}
