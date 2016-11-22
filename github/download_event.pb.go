// Code generated by protoc-gen-gogo.
// source: download_event.proto
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

type Download struct {
	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	HtmlUrl       string `protobuf:"bytes,2,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	Id            int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	FileSize      int32  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	DownloadCount int32  `protobuf:"varint,7,opt,name=download_count,json=downloadCount,proto3" json:"download_count,omitempty"`
	ContentType   string `protobuf:"bytes,8,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (m *Download) Reset()                    { *m = Download{} }
func (m *Download) String() string            { return proto.CompactTextString(m) }
func (*Download) ProtoMessage()               {}
func (*Download) Descriptor() ([]byte, []int) { return fileDescriptorDownloadEvent, []int{0} }

type DownloadEvent struct {
	Download     *Download     `protobuf:"bytes,1,opt,name=download" json:"download,omitempty"`
	Installation *Installation `protobuf:"bytes,2,opt,name=installation" json:"installation,omitempty"`
}

func (m *DownloadEvent) Reset()                    { *m = DownloadEvent{} }
func (m *DownloadEvent) String() string            { return proto.CompactTextString(m) }
func (*DownloadEvent) ProtoMessage()               {}
func (*DownloadEvent) Descriptor() ([]byte, []int) { return fileDescriptorDownloadEvent, []int{1} }

func init() {
	proto.RegisterType((*Download)(nil), "github.Download")
	proto.RegisterType((*DownloadEvent)(nil), "github.DownloadEvent")
}
func (m *Download) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Download) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(len(m.Url)))
		i += copy(data[i:], m.Url)
	}
	if len(m.HtmlUrl) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(len(m.HtmlUrl)))
		i += copy(data[i:], m.HtmlUrl)
	}
	if m.Id != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.Description) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	if m.FileSize != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(m.FileSize))
	}
	if m.DownloadCount != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(m.DownloadCount))
	}
	if len(m.ContentType) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(len(m.ContentType)))
		i += copy(data[i:], m.ContentType)
	}
	return i, nil
}

func (m *DownloadEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DownloadEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Download != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(m.Download.Size()))
		n1, err := m.Download.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Installation != nil {
		data[i] = 0x12
		i++
		i = encodeVarintDownloadEvent(data, i, uint64(m.Installation.Size()))
		n2, err := m.Installation.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64DownloadEvent(data []byte, offset int, v uint64) int {
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
func encodeFixed32DownloadEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDownloadEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Download) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	l = len(m.HtmlUrl)
	if l > 0 {
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovDownloadEvent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	if m.FileSize != 0 {
		n += 1 + sovDownloadEvent(uint64(m.FileSize))
	}
	if m.DownloadCount != 0 {
		n += 1 + sovDownloadEvent(uint64(m.DownloadCount))
	}
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	return n
}

func (m *DownloadEvent) Size() (n int) {
	var l int
	_ = l
	if m.Download != nil {
		l = m.Download.Size()
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovDownloadEvent(uint64(l))
	}
	return n
}

func sovDownloadEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDownloadEvent(x uint64) (n int) {
	return sovDownloadEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Download) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDownloadEvent
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
			return fmt.Errorf("proto: Download: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Download: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HtmlUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HtmlUrl = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			m.FileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FileSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownloadCount", wireType)
			}
			m.DownloadCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DownloadCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDownloadEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDownloadEvent
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
func (m *DownloadEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDownloadEvent
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
			return fmt.Errorf("proto: DownloadEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DownloadEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Download", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Download == nil {
				m.Download = &Download{}
			}
			if err := m.Download.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownloadEvent
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
				return ErrInvalidLengthDownloadEvent
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
			skippy, err := skipDownloadEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDownloadEvent
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
func skipDownloadEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDownloadEvent
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
					return 0, ErrIntOverflowDownloadEvent
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
					return 0, ErrIntOverflowDownloadEvent
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
				return 0, ErrInvalidLengthDownloadEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDownloadEvent
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
				next, err := skipDownloadEvent(data[start:])
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
	ErrInvalidLengthDownloadEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDownloadEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("download_event.proto", fileDescriptorDownloadEvent) }

var fileDescriptorDownloadEvent = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x9b, 0x5e, 0xf3, 0x4f, 0x2f, 0xaa, 0xac, 0x2e, 0xfc, 0x77, 0x51, 0x4a, 0x25, 0x24,
	0x16, 0xd0, 0x4a, 0x65, 0xc3, 0xba, 0x5c, 0x24, 0xb6, 0x01, 0xd6, 0x55, 0x9a, 0x98, 0xd6, 0x92,
	0x6b, 0x57, 0x89, 0x43, 0x05, 0x4f, 0xc2, 0x23, 0x75, 0xc9, 0x13, 0x20, 0x28, 0x2b, 0xde, 0x02,
	0x7b, 0x92, 0x54, 0x61, 0x61, 0x65, 0xe6, 0x3b, 0xc7, 0x1e, 0xcd, 0x09, 0xf4, 0x42, 0xb5, 0x95,
	0x42, 0xf9, 0xe1, 0x9c, 0x3d, 0x33, 0xa9, 0xc7, 0x9b, 0x48, 0x69, 0x45, 0xea, 0x4b, 0xae, 0x57,
	0xc9, 0xa2, 0x7f, 0x9e, 0x7e, 0xc7, 0x81, 0x5a, 0x4f, 0x96, 0x6a, 0xa9, 0x26, 0x28, 0x2f, 0x92,
	0x27, 0xec, 0xb0, 0xc1, 0x2a, 0xbd, 0xd6, 0x27, 0x5c, 0xc6, 0xda, 0x17, 0xc2, 0xd7, 0x5c, 0xc9,
	0x94, 0x8d, 0x7e, 0x1c, 0x70, 0xaf, 0xb3, 0x19, 0xa4, 0x0b, 0x95, 0x24, 0x12, 0xd4, 0x19, 0x3a,
	0xa7, 0xff, 0x3c, 0x5b, 0x92, 0xff, 0xe0, 0xae, 0xf4, 0x5a, 0xcc, 0x2d, 0x2e, 0x23, 0x6e, 0xd8,
	0xfe, 0xd1, 0x48, 0x1d, 0x28, 0xf3, 0x90, 0x56, 0x0c, 0xac, 0x79, 0xa6, 0x22, 0x04, 0xaa, 0xd2,
	0x5f, 0x33, 0x5a, 0x45, 0x1b, 0xd6, 0x64, 0x08, 0xcd, 0x90, 0xc5, 0x41, 0xc4, 0x37, 0x76, 0x24,
	0xad, 0xa1, 0x54, 0x44, 0xc6, 0x51, 0x8d, 0xf9, 0x2b, 0xa3, 0x75, 0xfb, 0xce, 0xac, 0xb5, 0xff,
	0x38, 0x72, 0x6f, 0xb9, 0x60, 0xf7, 0x86, 0x79, 0xa8, 0x90, 0x13, 0xe8, 0x1c, 0x42, 0x08, 0x54,
	0x22, 0x35, 0x6d, 0xe0, 0xcc, 0x76, 0x4e, 0xaf, 0x2c, 0x24, 0xc7, 0xd0, 0x0a, 0x94, 0xd4, 0x26,
	0xa4, 0xb9, 0x7e, 0xd9, 0x30, 0xea, 0xa6, 0xb3, 0x32, 0xf6, 0x60, 0xd0, 0x68, 0x0b, 0xed, 0x7c,
	0xd5, 0x1b, 0x9b, 0x26, 0x39, 0x03, 0x37, 0x7f, 0x04, 0x97, 0x6e, 0x4e, 0xbb, 0xe3, 0x2c, 0xd2,
	0xdc, 0xe8, 0x1d, 0x1c, 0xe4, 0x12, 0x5a, 0xc5, 0x00, 0x31, 0x8f, 0xe6, 0xb4, 0x97, 0xdf, 0xb8,
	0x2b, 0x68, 0xde, 0x1f, 0xe7, 0xac, 0xb7, 0xfb, 0x1a, 0x94, 0x76, 0xfb, 0x81, 0xf3, 0x6e, 0xce,
	0xa7, 0x39, 0x6f, 0xdf, 0x83, 0xd2, 0xa2, 0x8e, 0x7f, 0xe0, 0xe2, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x3b, 0x2e, 0x05, 0xe4, 0x01, 0x00, 0x00,
}
