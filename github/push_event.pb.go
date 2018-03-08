// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: push_event.proto

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

type PushCommit struct {
	Sha       string      `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	Id        string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TreeId    string      `protobuf:"bytes,3,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	Distinct  bool        `protobuf:"varint,4,opt,name=distinct,proto3" json:"distinct,omitempty"`
	Message   string      `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp string      `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Url       string      `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Author    *CommitUser `protobuf:"bytes,8,opt,name=author" json:"author,omitempty"`
	Committer *CommitUser `protobuf:"bytes,9,opt,name=committer" json:"committer,omitempty"`
	Added     []string    `protobuf:"bytes,10,rep,name=added" json:"added,omitempty"`
	Removed   []string    `protobuf:"bytes,11,rep,name=removed" json:"removed,omitempty"`
	Modified  []string    `protobuf:"bytes,12,rep,name=modified" json:"modified,omitempty"`
}

func (m *PushCommit) Reset()                    { *m = PushCommit{} }
func (m *PushCommit) String() string            { return proto.CompactTextString(m) }
func (*PushCommit) ProtoMessage()               {}
func (*PushCommit) Descriptor() ([]byte, []int) { return fileDescriptorPushEvent, []int{0} }

type PushEvent struct {
	Ref          string        `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Before       string        `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After        string        `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	Created      bool          `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Deleted      bool          `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Forced       bool          `protobuf:"varint,6,opt,name=forced,proto3" json:"forced,omitempty"`
	BaseRef      string        `protobuf:"bytes,7,opt,name=base_ref,json=baseRef,proto3" json:"base_ref,omitempty"`
	Compare      string        `protobuf:"bytes,8,opt,name=compare,proto3" json:"compare,omitempty"`
	Commits      []*PushCommit `protobuf:"bytes,9,rep,name=commits" json:"commits,omitempty"`
	HeadCommit   *PushCommit   `protobuf:"bytes,10,opt,name=head_commit,json=headCommit" json:"head_commit,omitempty"`
	Repository   *Repository   `protobuf:"bytes,11,opt,name=repository" json:"repository,omitempty"`
	Pusher       *CommitUser   `protobuf:"bytes,12,opt,name=pusher" json:"pusher,omitempty"`
	Sender       *User         `protobuf:"bytes,13,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation `protobuf:"bytes,14,opt,name=installation" json:"installation,omitempty"`
	Organization *User         `protobuf:"bytes,15,opt,name=organization" json:"organization,omitempty"`
}

func (m *PushEvent) Reset()                    { *m = PushEvent{} }
func (m *PushEvent) String() string            { return proto.CompactTextString(m) }
func (*PushEvent) ProtoMessage()               {}
func (*PushEvent) Descriptor() ([]byte, []int) { return fileDescriptorPushEvent, []int{1} }

func init() {
	proto.RegisterType((*PushCommit)(nil), "github.PushCommit")
	proto.RegisterType((*PushEvent)(nil), "github.PushEvent")
}
func (m *PushCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushCommit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sha) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Sha)))
		i += copy(dAtA[i:], m.Sha)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.TreeId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.TreeId)))
		i += copy(dAtA[i:], m.TreeId)
	}
	if m.Distinct {
		dAtA[i] = 0x20
		i++
		if m.Distinct {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Timestamp) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Timestamp)))
		i += copy(dAtA[i:], m.Timestamp)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.Author != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Author.Size()))
		n1, err := m.Author.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Committer != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Committer.Size()))
		n2, err := m.Committer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Added) > 0 {
		for _, s := range m.Added {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Removed) > 0 {
		for _, s := range m.Removed {
			dAtA[i] = 0x5a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Modified) > 0 {
		for _, s := range m.Modified {
			dAtA[i] = 0x62
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *PushEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ref) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	if len(m.Before) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Before)))
		i += copy(dAtA[i:], m.Before)
	}
	if len(m.After) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.After)))
		i += copy(dAtA[i:], m.After)
	}
	if m.Created {
		dAtA[i] = 0x20
		i++
		if m.Created {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Deleted {
		dAtA[i] = 0x28
		i++
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Forced {
		dAtA[i] = 0x30
		i++
		if m.Forced {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.BaseRef) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.BaseRef)))
		i += copy(dAtA[i:], m.BaseRef)
	}
	if len(m.Compare) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(len(m.Compare)))
		i += copy(dAtA[i:], m.Compare)
	}
	if len(m.Commits) > 0 {
		for _, msg := range m.Commits {
			dAtA[i] = 0x4a
			i++
			i = encodeVarintPushEvent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.HeadCommit != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.HeadCommit.Size()))
		n3, err := m.HeadCommit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Repository != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Repository.Size()))
		n4, err := m.Repository.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Pusher != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Pusher.Size()))
		n5, err := m.Pusher.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Sender != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Sender.Size()))
		n6, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Installation != nil {
		dAtA[i] = 0x72
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Installation.Size()))
		n7, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Organization != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintPushEvent(dAtA, i, uint64(m.Organization.Size()))
		n8, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeFixed64PushEvent(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32PushEvent(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPushEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PushCommit) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sha)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.TreeId)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Distinct {
		n += 2
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Author != nil {
		l = m.Author.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Committer != nil {
		l = m.Committer.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if len(m.Added) > 0 {
		for _, s := range m.Added {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		for _, s := range m.Removed {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if len(m.Modified) > 0 {
		for _, s := range m.Modified {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	return n
}

func (m *PushEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Before)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.After)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Created {
		n += 2
	}
	if m.Deleted {
		n += 2
	}
	if m.Forced {
		n += 2
	}
	l = len(m.BaseRef)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Compare)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if len(m.Commits) > 0 {
		for _, e := range m.Commits {
			l = e.Size()
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if m.HeadCommit != nil {
		l = m.HeadCommit.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Pusher != nil {
		l = m.Pusher.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	return n
}

func sovPushEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPushEvent(x uint64) (n int) {
	return sovPushEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PushCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushEvent
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
			return fmt.Errorf("proto: PushCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distinct", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Distinct = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Author == nil {
				m.Author = &CommitUser{}
			}
			if err := m.Author.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Committer == nil {
				m.Committer = &CommitUser{}
			}
			if err := m.Committer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Added", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Added = append(m.Added, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Removed = append(m.Removed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modified", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Modified = append(m.Modified, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPushEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushEvent
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
func (m *PushEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushEvent
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
			return fmt.Errorf("proto: PushEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Before", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Before = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.After = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forced", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Forced = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseRef", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseRef = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Compare = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commits = append(m.Commits, &PushCommit{})
			if err := m.Commits[len(m.Commits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HeadCommit == nil {
				m.HeadCommit = &PushCommit{}
			}
			if err := m.HeadCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pusher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pusher == nil {
				m.Pusher = &CommitUser{}
			}
			if err := m.Pusher.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
			skippy, err := skipPushEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushEvent
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
func skipPushEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPushEvent
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
					return 0, ErrIntOverflowPushEvent
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
					return 0, ErrIntOverflowPushEvent
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
				return 0, ErrInvalidLengthPushEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPushEvent
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
				next, err := skipPushEvent(dAtA[start:])
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
	ErrInvalidLengthPushEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPushEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("push_event.proto", fileDescriptorPushEvent) }

var fileDescriptorPushEvent = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x37, 0xed, 0x36, 0x6d, 0xa6, 0x65, 0x29, 0x56, 0x05, 0xa6, 0x42, 0x55, 0xb5, 0xe2,
	0x50, 0x21, 0xe8, 0xae, 0x76, 0x2f, 0x9c, 0x41, 0x1c, 0xf6, 0x86, 0x2c, 0x71, 0xae, 0xdc, 0x7a,
	0xd2, 0x58, 0x6a, 0xe2, 0xc8, 0x76, 0x56, 0x82, 0x27, 0xe1, 0xca, 0xdb, 0xec, 0x91, 0x1b, 0x57,
	0x28, 0x2f, 0x82, 0x6c, 0x27, 0x6d, 0x2a, 0xb1, 0xa7, 0xe6, 0x9b, 0xff, 0x1f, 0x79, 0xfc, 0x77,
	0x0c, 0xe3, 0xb2, 0x32, 0xd9, 0x0a, 0xef, 0xb1, 0xb0, 0xcb, 0x52, 0x2b, 0xab, 0x48, 0xbc, 0x95,
	0x36, 0xab, 0xd6, 0xd3, 0x77, 0xe1, 0x77, 0xb9, 0x51, 0xf9, 0xd5, 0x56, 0x6d, 0xd5, 0x95, 0x97,
	0xd7, 0x55, 0xea, 0xc9, 0x83, 0xff, 0x0a, 0x6d, 0x53, 0xa8, 0x0c, 0xea, 0xfa, 0xfb, 0xd9, 0x46,
	0xe5, 0xb9, 0xb4, 0xab, 0x56, 0x69, 0xac, 0xb1, 0x54, 0x46, 0x5a, 0xa5, 0xbf, 0xd6, 0x15, 0x22,
	0x0b, 0x63, 0xf9, 0x6e, 0xc7, 0xad, 0x54, 0x45, 0xa8, 0x5d, 0xfe, 0xea, 0x00, 0x7c, 0xae, 0x4c,
	0xf6, 0xd1, 0xf7, 0x93, 0x31, 0x74, 0x4d, 0xc6, 0x69, 0x34, 0x8f, 0x16, 0x09, 0x73, 0x9f, 0xe4,
	0x02, 0x3a, 0x52, 0xd0, 0x8e, 0x2f, 0x74, 0xa4, 0x20, 0x2f, 0xa0, 0x6f, 0x35, 0xe2, 0x4a, 0x0a,
	0xda, 0xf5, 0xc5, 0xd8, 0xe1, 0x9d, 0x20, 0x53, 0x18, 0x08, 0x69, 0xac, 0x2c, 0x36, 0x96, 0x9e,
	0xcf, 0xa3, 0xc5, 0x80, 0x1d, 0x98, 0x50, 0xe8, 0xe7, 0x68, 0x0c, 0xdf, 0x22, 0xed, 0xf9, 0xa6,
	0x06, 0xc9, 0x2b, 0x48, 0xac, 0xcc, 0xd1, 0x58, 0x9e, 0x97, 0x34, 0xf6, 0xda, 0xb1, 0xe0, 0xc6,
	0xa9, 0xf4, 0x8e, 0xf6, 0xc3, 0x38, 0x95, 0xde, 0x91, 0x37, 0x10, 0xf3, 0xca, 0x66, 0x4a, 0xd3,
	0xc1, 0x3c, 0x5a, 0x0c, 0x6f, 0xc8, 0xb2, 0x0e, 0x2d, 0x5c, 0xe0, 0x8b, 0x41, 0xcd, 0x6a, 0x07,
	0xb9, 0x86, 0x24, 0xc4, 0x62, 0x51, 0xd3, 0xe4, 0x51, 0xfb, 0xd1, 0x44, 0x26, 0xd0, 0xe3, 0x42,
	0xa0, 0xa0, 0x30, 0xef, 0x2e, 0x12, 0x16, 0xc0, 0x4d, 0xaf, 0x31, 0x57, 0xf7, 0x28, 0xe8, 0xd0,
	0xd7, 0x1b, 0x74, 0x77, 0xce, 0x95, 0x90, 0xa9, 0x44, 0x41, 0x47, 0x5e, 0x3a, 0xf0, 0xe5, 0x8f,
	0x73, 0x48, 0x5c, 0xb2, 0x9f, 0xdc, 0x3f, 0xed, 0x6e, 0xa2, 0x31, 0x6d, 0x82, 0xd5, 0x98, 0x92,
	0xe7, 0x10, 0xaf, 0x31, 0x55, 0x1a, 0xeb, 0x70, 0x6b, 0xf2, 0x33, 0xa4, 0x6e, 0xe2, 0x10, 0x6f,
	0x00, 0x37, 0xc3, 0x46, 0x23, 0xb7, 0x28, 0xea, 0x70, 0x1b, 0x74, 0x8a, 0xc0, 0x1d, 0x3a, 0xa5,
	0x17, 0x94, 0x1a, 0xdd, 0x09, 0xa9, 0xd2, 0x1b, 0x14, 0x3e, 0xd8, 0x01, 0xab, 0x89, 0xbc, 0x84,
	0xc1, 0x9a, 0x1b, 0x5c, 0xb9, 0x81, 0x42, 0xb4, 0x7d, 0xc7, 0x0c, 0x53, 0x7f, 0x8c, 0xca, 0x4b,
	0xae, 0xd1, 0xe7, 0x9b, 0xb0, 0x06, 0xc9, 0x5b, 0xaf, 0xe4, 0xd2, 0x1a, 0x9a, 0xcc, 0xbb, 0xed,
	0x28, 0x8f, 0xeb, 0xc3, 0x1a, 0x0b, 0xb9, 0x85, 0x61, 0x86, 0x5c, 0xac, 0x02, 0x53, 0x38, 0x0d,
	0xbf, 0xd5, 0x01, 0xce, 0x56, 0x2f, 0xdf, 0x0d, 0xc0, 0x71, 0x67, 0xe9, 0xf0, 0xb4, 0x87, 0x1d,
	0x14, 0xd6, 0x72, 0xb9, 0x7d, 0x70, 0xef, 0x09, 0x35, 0x1d, 0x3d, 0xbe, 0x0f, 0xc1, 0x41, 0x5e,
	0x43, 0x6c, 0xb0, 0x10, 0xa8, 0xe9, 0x13, 0xef, 0x1d, 0x35, 0xde, 0xe0, 0x0a, 0x1a, 0x79, 0x0f,
	0xa3, 0xf6, 0x3b, 0xa1, 0x17, 0xde, 0x3b, 0x69, 0xbc, 0x77, 0x2d, 0x8d, 0x9d, 0x38, 0xc9, 0x35,
	0x8c, 0x94, 0xde, 0xf2, 0x42, 0x7e, 0x0b, 0x9d, 0x4f, 0xff, 0x73, 0xca, 0x89, 0xe3, 0xc3, 0xe4,
	0xe1, 0xcf, 0xec, 0xec, 0x61, 0x3f, 0x8b, 0x7e, 0xee, 0x67, 0xd1, 0xef, 0xfd, 0x2c, 0xfa, 0xfe,
	0x77, 0x76, 0xb6, 0x8e, 0xfd, 0xd3, 0xbc, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd1, 0x12,
	0x52, 0x2a, 0x04, 0x00, 0x00,
}
