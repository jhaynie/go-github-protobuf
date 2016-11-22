// Code generated by protoc-gen-gogo.
// source: page_build_event.proto
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

type BuildError struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *BuildError) Reset()                    { *m = BuildError{} }
func (m *BuildError) String() string            { return proto.CompactTextString(m) }
func (*BuildError) ProtoMessage()               {}
func (*BuildError) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{0} }

type Build struct {
	Url       string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Status    string      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error     *BuildError `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Pusher    *User       `protobuf:"bytes,4,opt,name=pusher" json:"pusher,omitempty"`
	Commit    string      `protobuf:"bytes,5,opt,name=commit,proto3" json:"commit,omitempty"`
	Duration  int32       `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt string      `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{1} }

type PageBuildEvent struct {
	Id           int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Build        *Build        `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	Repository   *Repository   `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Sender       *User         `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,5,opt,name=installation" json:"installation,omitempty"`
}

func (m *PageBuildEvent) Reset()                    { *m = PageBuildEvent{} }
func (m *PageBuildEvent) String() string            { return proto.CompactTextString(m) }
func (*PageBuildEvent) ProtoMessage()               {}
func (*PageBuildEvent) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{2} }

func init() {
	proto.RegisterType((*BuildError)(nil), "github.BuildError")
	proto.RegisterType((*Build)(nil), "github.Build")
	proto.RegisterType((*PageBuildEvent)(nil), "github.PageBuildEvent")
}
func (m *BuildError) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BuildError) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Message)))
		i += copy(data[i:], m.Message)
	}
	return i, nil
}

func (m *Build) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Build) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Url)))
		i += copy(data[i:], m.Url)
	}
	if len(m.Status) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Status)))
		i += copy(data[i:], m.Status)
	}
	if m.Error != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Pusher != nil {
		data[i] = 0x22
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Pusher.Size()))
		n2, err := m.Pusher.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Commit) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Commit)))
		i += copy(data[i:], m.Commit)
	}
	if m.Duration != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Duration))
	}
	if len(m.CreatedAt) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.CreatedAt)))
		i += copy(data[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.UpdatedAt)))
		i += copy(data[i:], m.UpdatedAt)
	}
	return i, nil
}

func (m *PageBuildEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PageBuildEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Id))
	}
	if m.Build != nil {
		data[i] = 0x12
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Build.Size()))
		n3, err := m.Build.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Repository != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Repository.Size()))
		n4, err := m.Repository.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Sender != nil {
		data[i] = 0x22
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Sender.Size()))
		n5, err := m.Sender.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Installation != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Installation.Size()))
		n6, err := m.Installation.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeFixed64PageBuildEvent(data []byte, offset int, v uint64) int {
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
func encodeFixed32PageBuildEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPageBuildEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *BuildError) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func (m *Build) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Pusher != nil {
		l = m.Pusher.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.Commit)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovPageBuildEvent(uint64(m.Duration))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func (m *PageBuildEvent) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPageBuildEvent(uint64(m.Id))
	}
	if m.Build != nil {
		l = m.Build.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func sovPageBuildEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPageBuildEvent(x uint64) (n int) {
	return sovPageBuildEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuildError) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: BuildError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuildError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func (m *Build) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: Build: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Build: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &BuildError{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pusher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pusher == nil {
				m.Pusher = &User{}
			}
			if err := m.Pusher.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commit = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func (m *PageBuildEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: PageBuildEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageBuildEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Build == nil {
				m.Build = &Build{}
			}
			if err := m.Build.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
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
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func skipPageBuildEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPageBuildEvent
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
					return 0, ErrIntOverflowPageBuildEvent
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
					return 0, ErrIntOverflowPageBuildEvent
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
				return 0, ErrInvalidLengthPageBuildEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPageBuildEvent
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
				next, err := skipPageBuildEvent(data[start:])
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
	ErrInvalidLengthPageBuildEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPageBuildEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("page_build_event.proto", fileDescriptorPageBuildEvent) }

var fileDescriptorPageBuildEvent = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xcb, 0x8a, 0xe2, 0x40,
	0x14, 0x35, 0x3a, 0x89, 0x7a, 0x75, 0x44, 0x0a, 0x91, 0x42, 0x18, 0x19, 0x9c, 0x61, 0x70, 0x33,
	0x0a, 0x99, 0xcd, 0x6c, 0x15, 0x66, 0x31, 0xbb, 0xa6, 0xa0, 0xd7, 0x92, 0x98, 0xea, 0x18, 0x88,
	0xa9, 0x50, 0x8f, 0x86, 0xfe, 0x93, 0xfe, 0x24, 0x97, 0xbd, 0xed, 0x5d, 0x3f, 0x3e, 0xa1, 0x7f,
	0xa0, 0x2b, 0x55, 0x15, 0x1f, 0x0d, 0xbd, 0x08, 0xde, 0xf3, 0xa8, 0xba, 0x75, 0xee, 0x15, 0xc6,
	0x65, 0x94, 0xd2, 0x4d, 0xac, 0xb2, 0x3c, 0xd9, 0xd0, 0x5b, 0x5a, 0xc8, 0x45, 0xc9, 0x99, 0x64,
	0x28, 0x48, 0x33, 0xb9, 0x53, 0xf1, 0xe4, 0xb7, 0xfd, 0x5d, 0x6c, 0xd9, 0x7e, 0x99, 0xb2, 0x94,
	0x2d, 0x8d, 0x1c, 0xab, 0x1b, 0x83, 0x0c, 0x30, 0x95, 0x3d, 0x36, 0x01, 0x25, 0x28, 0x77, 0xf5,
	0x90, 0xd3, 0x92, 0x89, 0x4c, 0x32, 0x7e, 0xe7, 0x18, 0x94, 0x15, 0x42, 0x46, 0x79, 0x1e, 0xc9,
	0x8c, 0x15, 0x96, 0x9b, 0xfd, 0x02, 0x58, 0x57, 0xdd, 0xff, 0x71, 0xce, 0x38, 0xc2, 0xd0, 0xde,
	0x53, 0x21, 0xf4, 0x9b, 0xb0, 0xf7, 0xdd, 0x9b, 0x77, 0x49, 0x0d, 0x67, 0x6f, 0x1e, 0xf8, 0xc6,
	0x88, 0x86, 0xd0, 0x52, 0x3c, 0x77, 0x7a, 0x55, 0xa2, 0x31, 0x04, 0xfa, 0x5e, 0xa9, 0x04, 0x6e,
	0x1a, 0xd2, 0x21, 0x34, 0x07, 0x9f, 0x56, 0xd7, 0xe2, 0x96, 0xa6, 0x7b, 0x21, 0x5a, 0xb8, 0x30,
	0xa7, 0x86, 0xc4, 0x1a, 0xd0, 0x4f, 0x08, 0x4a, 0x25, 0x76, 0x94, 0xe3, 0x2f, 0xc6, 0xda, 0xaf,
	0xad, 0xd7, 0x3a, 0x0f, 0x71, 0x5a, 0xd5, 0x47, 0xcf, 0x61, 0x9f, 0x49, 0xec, 0xdb, 0x3e, 0x16,
	0xa1, 0x09, 0x74, 0x12, 0xc5, 0x4d, 0x2a, 0x1c, 0x68, 0xc5, 0x27, 0x47, 0x8c, 0xbe, 0x01, 0x6c,
	0x39, 0x8d, 0x24, 0x4d, 0x36, 0x91, 0xc4, 0x6d, 0x73, 0xae, 0xeb, 0x98, 0x95, 0xac, 0x64, 0x55,
	0x26, 0xb5, 0xdc, 0xb1, 0xb2, 0x63, 0x56, 0x72, 0xf6, 0xe8, 0xc1, 0xe0, 0x4a, 0xc7, 0xb7, 0x2f,
	0xae, 0xf6, 0x83, 0x06, 0xd0, 0xcc, 0x12, 0x93, 0xde, 0x27, 0xba, 0x42, 0x3f, 0xc0, 0x37, 0xeb,
	0x33, 0xd9, 0x7b, 0xe1, 0xd7, 0x8b, 0x90, 0xc4, 0x6a, 0x28, 0x04, 0x38, 0x6d, 0xe3, 0xe3, 0x38,
	0xc8, 0x51, 0x21, 0x67, 0xae, 0x6a, 0x26, 0x82, 0x16, 0xc9, 0x67, 0x33, 0xb1, 0x1a, 0xfa, 0x0b,
	0xfd, 0xf3, 0xad, 0x9a, 0xc9, 0xf4, 0xc2, 0x51, 0xed, 0xfd, 0x7f, 0xa6, 0x91, 0x0b, 0xe7, 0x7a,
	0x74, 0x78, 0x9e, 0x36, 0x0e, 0x2f, 0x53, 0xef, 0x41, 0x7f, 0x4f, 0xfa, 0xbb, 0x7f, 0x9d, 0x36,
	0xe2, 0xc0, 0xfc, 0x2d, 0xfe, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x75, 0xf5, 0x3a, 0x99,
	0x02, 0x00, 0x00,
}
