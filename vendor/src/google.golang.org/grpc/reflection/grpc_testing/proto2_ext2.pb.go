// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2_ext2.proto

package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AnotherExtension struct {
	Whatchamacallit      *int32   `protobuf:"varint,1,opt,name=whatchamacallit" json:"whatchamacallit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnotherExtension) Reset()         { *m = AnotherExtension{} }
func (m *AnotherExtension) String() string { return proto.CompactTextString(m) }
func (*AnotherExtension) ProtoMessage()    {}
func (*AnotherExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto2_ext2_039d342873655470, []int{0}
}
func (m *AnotherExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnotherExtension.Unmarshal(m, b)
}
func (m *AnotherExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnotherExtension.Marshal(b, m, deterministic)
}
func (dst *AnotherExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnotherExtension.Merge(dst, src)
}
func (m *AnotherExtension) XXX_Size() int {
	return xxx_messageInfo_AnotherExtension.Size(m)
}
func (m *AnotherExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_AnotherExtension.DiscardUnknown(m)
}

var xxx_messageInfo_AnotherExtension proto.InternalMessageInfo

func (m *AnotherExtension) GetWhatchamacallit() int32 {
	if m != nil && m.Whatchamacallit != nil {
		return *m.Whatchamacallit
	}
	return 0
}

var E_Frob = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtended)(nil),
	ExtensionType: (*string)(nil),
	Field:         23,
	Name:          "grpc.testing.frob",
	Tag:           "bytes,23,opt,name=frob",
	Filename:      "proto2_ext2.proto",
}

var E_Nitz = &proto.ExtensionDesc{
	ExtendedType:  (*ToBeExtended)(nil),
	ExtensionType: (*AnotherExtension)(nil),
	Field:         29,
	Name:          "grpc.testing.nitz",
	Tag:           "bytes,29,opt,name=nitz",
	Filename:      "proto2_ext2.proto",
}

func init() {
	proto.RegisterType((*AnotherExtension)(nil), "grpc.testing.AnotherExtension")
	proto.RegisterExtension(E_Frob)
	proto.RegisterExtension(E_Nitz)
}

func init() { proto.RegisterFile("proto2_ext2.proto", fileDescriptor_proto2_ext2_039d342873655470) }

var fileDescriptor_proto2_ext2_039d342873655470 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x8a, 0x4f, 0xad, 0x28, 0x31, 0xd2, 0x03, 0xb3, 0x85, 0x78, 0xd2, 0x8b, 0x0a, 0x92,
	0xf5, 0x4a, 0x52, 0x8b, 0x4b, 0x32, 0xf3, 0xd2, 0xa5, 0x78, 0x20, 0x0a, 0x20, 0x72, 0x4a, 0x36,
	0x5c, 0x02, 0x8e, 0x79, 0xf9, 0x25, 0x19, 0xa9, 0x45, 0xae, 0x15, 0x25, 0xa9, 0x79, 0xc5, 0x99,
	0xf9, 0x79, 0x42, 0x1a, 0x5c, 0xfc, 0xe5, 0x19, 0x89, 0x25, 0xc9, 0x19, 0x89, 0xb9, 0x89, 0xc9,
	0x89, 0x39, 0x39, 0x99, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xe8, 0xc2, 0x56, 0x7a,
	0x5c, 0x2c, 0x69, 0x45, 0xf9, 0x49, 0x42, 0x52, 0x7a, 0xc8, 0x56, 0xe8, 0x85, 0xe4, 0x3b, 0xa5,
	0x82, 0x8d, 0x4b, 0x49, 0x4d, 0x91, 0x10, 0x57, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xab, 0xb3, 0xf2,
	0xe3, 0x62, 0xc9, 0xcb, 0x2c, 0xa9, 0xc2, 0xab, 0x5e, 0x56, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x0e,
	0x55, 0x05, 0xba, 0x1b, 0x83, 0xc0, 0xe6, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x7e, 0x0d,
	0x26, 0xed, 0x00, 0x00, 0x00,
}
