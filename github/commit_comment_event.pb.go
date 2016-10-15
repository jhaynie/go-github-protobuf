// Code generated by protoc-gen-go.
// source: commit_comment_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommitCommentEvent struct {
	Action     string      `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Comment    *Comment    `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Repository *Repository `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
}

func (m *CommitCommentEvent) Reset()                    { *m = CommitCommentEvent{} }
func (m *CommitCommentEvent) String() string            { return proto.CompactTextString(m) }
func (*CommitCommentEvent) ProtoMessage()               {}
func (*CommitCommentEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommitCommentEvent) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *CommitCommentEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CommitCommentEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitCommentEvent)(nil), "github.CommitCommentEvent")
}

func init() { proto.RegisterFile("commit_comment_event.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0x2c, 0x89, 0x07, 0x51, 0xa9, 0x79, 0x25, 0xf1, 0xa9, 0x65, 0x40, 0x52, 0xaf, 0xa0, 0x28,
	0xbf, 0x24, 0x5f, 0x88, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x8a, 0xab, 0xb4, 0x38, 0xb5,
	0x08, 0x22, 0x26, 0xc5, 0x0b, 0x55, 0x08, 0xe5, 0x0a, 0x14, 0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96,
	0xe4, 0x17, 0x55, 0x42, 0x44, 0x94, 0x36, 0x32, 0x72, 0x09, 0x39, 0x83, 0xcd, 0x74, 0x86, 0xa8,
	0x74, 0x05, 0x99, 0x28, 0x24, 0xc6, 0xc5, 0x96, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x69, 0x72, 0xb1, 0x43, 0x4d, 0x94, 0x60, 0x02, 0x4a,
	0x70, 0x1b, 0xf1, 0xeb, 0x41, 0x6c, 0xd5, 0x83, 0x6a, 0x0f, 0x82, 0xc9, 0x0b, 0x19, 0x71, 0x71,
	0x21, 0x6c, 0x93, 0x60, 0x06, 0xab, 0x16, 0x82, 0xa9, 0x0e, 0x82, 0xcb, 0x04, 0x21, 0xa9, 0x12,
	0x52, 0xe1, 0x62, 0x2b, 0x4e, 0xcd, 0x4b, 0x49, 0x2d, 0x92, 0x60, 0x01, 0xab, 0xe7, 0x81, 0xa9,
	0x0f, 0x05, 0x7a, 0x29, 0x08, 0x2a, 0x97, 0xc4, 0x06, 0x76, 0xba, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0xb3, 0xc8, 0xca, 0x0d, 0x01, 0x00, 0x00,
}
