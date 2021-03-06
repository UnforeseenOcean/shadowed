// Code generated by protoc-gen-go. DO NOT EDIT.
// source: music.proto

package class

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

type MusicGroup struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tracks               []string `protobuf:"bytes,2,rep,name=tracks,proto3" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MusicGroup) Reset()         { *m = MusicGroup{} }
func (m *MusicGroup) String() string { return proto.CompactTextString(m) }
func (*MusicGroup) ProtoMessage()    {}
func (*MusicGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88ba70d01ec9147, []int{0}
}

func (m *MusicGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MusicGroup.Unmarshal(m, b)
}
func (m *MusicGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MusicGroup.Marshal(b, m, deterministic)
}
func (m *MusicGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MusicGroup.Merge(m, src)
}
func (m *MusicGroup) XXX_Size() int {
	return xxx_messageInfo_MusicGroup.Size(m)
}
func (m *MusicGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MusicGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MusicGroup proto.InternalMessageInfo

func (m *MusicGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MusicGroup) GetTracks() []string {
	if m != nil {
		return m.Tracks
	}
	return nil
}

type MusicLib struct {
	Groups               []*MusicGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MusicLib) Reset()         { *m = MusicLib{} }
func (m *MusicLib) String() string { return proto.CompactTextString(m) }
func (*MusicLib) ProtoMessage()    {}
func (*MusicLib) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88ba70d01ec9147, []int{1}
}

func (m *MusicLib) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MusicLib.Unmarshal(m, b)
}
func (m *MusicLib) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MusicLib.Marshal(b, m, deterministic)
}
func (m *MusicLib) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MusicLib.Merge(m, src)
}
func (m *MusicLib) XXX_Size() int {
	return xxx_messageInfo_MusicLib.Size(m)
}
func (m *MusicLib) XXX_DiscardUnknown() {
	xxx_messageInfo_MusicLib.DiscardUnknown(m)
}

var xxx_messageInfo_MusicLib proto.InternalMessageInfo

func (m *MusicLib) GetGroups() []*MusicGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*MusicGroup)(nil), "class.MusicGroup")
	proto.RegisterType((*MusicLib)(nil), "class.MusicLib")
}

func init() { proto.RegisterFile("music.proto", fileDescriptor_c88ba70d01ec9147) }

var fileDescriptor_c88ba70d01ec9147 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2d, 0x2d, 0xce,
	0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xce, 0x49, 0x2c, 0x2e, 0x56, 0xb2,
	0xe0, 0xe2, 0xf2, 0x05, 0x89, 0xba, 0x17, 0xe5, 0x97, 0x16, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x25,
	0x45, 0x89, 0xc9, 0xd9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x29,
	0x17, 0x07, 0x58, 0xa7, 0x4f, 0x66, 0x92, 0x90, 0x26, 0x17, 0x5b, 0x3a, 0xc8, 0x80, 0x62, 0x09,
	0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x41, 0x3d, 0xb0, 0xe9, 0x7a, 0x08, 0xa3, 0x83, 0xa0, 0x0a,
	0x92, 0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xf0, 0xd3, 0xd7, 0x8d,
	0x00, 0x00, 0x00,
}
