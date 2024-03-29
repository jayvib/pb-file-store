// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Article struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func init() {
	proto.RegisterType((*Article)(nil), "pb.Article")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0x2a, 0xc9,
	0x4c, 0xce, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x8a, 0xe5,
	0x62, 0x77, 0x84, 0x08, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30,
	0x07, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79,
	0x25, 0x12, 0xcc, 0x60, 0x71, 0x18, 0x57, 0x48, 0x88, 0x8b, 0xa5, 0xb4, 0x38, 0xb5, 0x48, 0x82,
	0x05, 0x2c, 0x0c, 0x66, 0x27, 0xb1, 0x81, 0x6d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0x8a, 0x15, 0x78, 0x7a, 0x00, 0x00, 0x00,
}
