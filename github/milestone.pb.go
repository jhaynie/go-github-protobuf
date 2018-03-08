// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: milestone.proto

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

type Milestone struct {
	Url          string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	HtmlUrl      string `protobuf:"bytes,2,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	LabelsUrl    string `protobuf:"bytes,3,opt,name=labels_url,json=labelsUrl,proto3" json:"labels_url,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Number       int32  `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	State        string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Title        string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Creator      *User  `protobuf:"bytes,9,opt,name=creator" json:"creator,omitempty"`
	OpenIssues   int32  `protobuf:"varint,10,opt,name=open_issues,json=openIssues,proto3" json:"open_issues,omitempty"`
	ClosedIssues int32  `protobuf:"varint,11,opt,name=closed_issues,json=closedIssues,proto3" json:"closed_issues,omitempty"`
	CreatedAt    string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ClosedAt     string `protobuf:"bytes,14,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
	DueOn        string `protobuf:"bytes,15,opt,name=due_on,json=dueOn,proto3" json:"due_on,omitempty"`
}

func (m *Milestone) Reset()                    { *m = Milestone{} }
func (m *Milestone) String() string            { return proto.CompactTextString(m) }
func (*Milestone) ProtoMessage()               {}
func (*Milestone) Descriptor() ([]byte, []int) { return fileDescriptorMilestone, []int{0} }

func init() {
	proto.RegisterType((*Milestone)(nil), "github.Milestone")
}
func (m *Milestone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Milestone) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.HtmlUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.HtmlUrl)))
		i += copy(dAtA[i:], m.HtmlUrl)
	}
	if len(m.LabelsUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.LabelsUrl)))
		i += copy(dAtA[i:], m.LabelsUrl)
	}
	if m.Id != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(m.Id))
	}
	if m.Number != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(m.Number))
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.Creator != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(m.Creator.Size()))
		n1, err := m.Creator.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.OpenIssues != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(m.OpenIssues))
	}
	if m.ClosedIssues != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(m.ClosedIssues))
	}
	if len(m.CreatedAt) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.CreatedAt)))
		i += copy(dAtA[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.UpdatedAt)))
		i += copy(dAtA[i:], m.UpdatedAt)
	}
	if len(m.ClosedAt) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.ClosedAt)))
		i += copy(dAtA[i:], m.ClosedAt)
	}
	if len(m.DueOn) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintMilestone(dAtA, i, uint64(len(m.DueOn)))
		i += copy(dAtA[i:], m.DueOn)
	}
	return i, nil
}

func encodeFixed64Milestone(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Milestone(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMilestone(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Milestone) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.HtmlUrl)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.LabelsUrl)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovMilestone(uint64(m.Id))
	}
	if m.Number != 0 {
		n += 1 + sovMilestone(uint64(m.Number))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	if m.Creator != nil {
		l = m.Creator.Size()
		n += 1 + l + sovMilestone(uint64(l))
	}
	if m.OpenIssues != 0 {
		n += 1 + sovMilestone(uint64(m.OpenIssues))
	}
	if m.ClosedIssues != 0 {
		n += 1 + sovMilestone(uint64(m.ClosedIssues))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.ClosedAt)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	l = len(m.DueOn)
	if l > 0 {
		n += 1 + l + sovMilestone(uint64(l))
	}
	return n
}

func sovMilestone(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMilestone(x uint64) (n int) {
	return sovMilestone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Milestone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMilestone
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
			return fmt.Errorf("proto: Milestone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Milestone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HtmlUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HtmlUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LabelsUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LabelsUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Creator == nil {
				m.Creator = &User{}
			}
			if err := m.Creator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenIssues", wireType)
			}
			m.OpenIssues = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpenIssues |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedIssues", wireType)
			}
			m.ClosedIssues = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClosedIssues |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClosedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DueOn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMilestone
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
				return ErrInvalidLengthMilestone
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DueOn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMilestone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMilestone
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
func skipMilestone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMilestone
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
					return 0, ErrIntOverflowMilestone
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
					return 0, ErrIntOverflowMilestone
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
				return 0, ErrInvalidLengthMilestone
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMilestone
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
				next, err := skipMilestone(dAtA[start:])
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
	ErrInvalidLengthMilestone = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMilestone   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("milestone.proto", fileDescriptorMilestone) }

var fileDescriptorMilestone = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0xaa, 0xd3, 0x40,
	0x10, 0x87, 0x4f, 0x5a, 0x9b, 0x36, 0x93, 0x9e, 0x53, 0x59, 0xaa, 0xac, 0x95, 0xc6, 0xa0, 0x20,
	0xbd, 0xb1, 0x05, 0x7d, 0x82, 0x7a, 0xe7, 0x85, 0x08, 0x85, 0x5e, 0x87, 0x24, 0x3b, 0xb6, 0x0b,
	0x9b, 0x6c, 0xd8, 0x3f, 0xef, 0xe2, 0x4b, 0xf8, 0x1e, 0xbd, 0xf4, 0x11, 0xb4, 0xbe, 0x88, 0x64,
	0x36, 0x01, 0xaf, 0x32, 0xbf, 0xef, 0x9b, 0xcc, 0x2e, 0xb3, 0xb0, 0x6a, 0xa4, 0x42, 0xeb, 0x74,
	0x8b, 0xfb, 0xce, 0x68, 0xa7, 0x59, 0x7c, 0x91, 0xee, 0xea, 0xab, 0xcd, 0x87, 0xf0, 0xdd, 0xd7,
	0xba, 0x39, 0x5c, 0xf4, 0x45, 0x1f, 0x48, 0x57, 0xfe, 0x3b, 0x25, 0x0a, 0x54, 0x85, 0xdf, 0x36,
	0xe0, 0x2d, 0x9a, 0x50, 0xbf, 0xfd, 0x39, 0x85, 0xe4, 0xeb, 0x38, 0x96, 0x3d, 0x87, 0xa9, 0x37,
	0x8a, 0x47, 0x79, 0xb4, 0x4b, 0x4e, 0x7d, 0xc9, 0x5e, 0xc1, 0xe2, 0xea, 0x1a, 0x55, 0xf4, 0x78,
	0x42, 0x78, 0xde, 0xe7, 0xb3, 0x51, 0x6c, 0x0b, 0xa0, 0xca, 0x0a, 0x95, 0x25, 0x39, 0x25, 0x99,
	0x04, 0xd2, 0xeb, 0x27, 0x98, 0x48, 0xc1, 0x9f, 0xe5, 0xd1, 0x6e, 0x76, 0x9a, 0x48, 0xc1, 0x5e,
	0x42, 0xdc, 0xfa, 0xa6, 0x42, 0xc3, 0x67, 0xc4, 0x86, 0xc4, 0xd6, 0x30, 0xb3, 0xae, 0x74, 0xc8,
	0x63, 0x9a, 0x10, 0x42, 0x4f, 0x9d, 0x74, 0x0a, 0xf9, 0x3c, 0x50, 0x0a, 0x2c, 0x87, 0x54, 0xa0,
	0xad, 0x8d, 0xec, 0x9c, 0xd4, 0x2d, 0x5f, 0x90, 0xfb, 0x1f, 0xb1, 0xf7, 0x30, 0xaf, 0x0d, 0x96,
	0x4e, 0x1b, 0x9e, 0xe4, 0xd1, 0x2e, 0xfd, 0xb8, 0xdc, 0x0f, 0xcb, 0x39, 0x5b, 0x34, 0xa7, 0x51,
	0xb2, 0x37, 0x90, 0xea, 0x0e, 0xdb, 0x42, 0x5a, 0xeb, 0xd1, 0x72, 0xa0, 0x2b, 0x41, 0x8f, 0xbe,
	0x10, 0x61, 0xef, 0xe0, 0xb1, 0x56, 0xda, 0xa2, 0x18, 0x5b, 0x52, 0x6a, 0x59, 0x06, 0x38, 0x34,
	0x6d, 0x01, 0x68, 0x20, 0x8a, 0xa2, 0x74, 0x7c, 0x19, 0x56, 0x30, 0x90, 0xa3, 0xeb, 0xb5, 0xef,
	0xc4, 0xa8, 0x1f, 0x83, 0x1e, 0xc8, 0xd1, 0xb1, 0xd7, 0x90, 0x0c, 0x47, 0x94, 0x8e, 0x3f, 0x91,
	0x5d, 0x04, 0x70, 0x74, 0xec, 0x05, 0xc4, 0xc2, 0x63, 0xa1, 0x5b, 0xbe, 0x0a, 0x1b, 0x10, 0x1e,
	0xbf, 0xb5, 0x9f, 0xd7, 0xb7, 0x3f, 0xd9, 0xc3, 0xed, 0x9e, 0x45, 0xbf, 0xee, 0x59, 0xf4, 0xfb,
	0x9e, 0x45, 0x3f, 0xfe, 0x66, 0x0f, 0x55, 0x4c, 0x8f, 0xf9, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xc1, 0xd9, 0xc5, 0x22, 0x02, 0x00, 0x00,
}
