// Code generated by protoc-gen-go. DO NOT EDIT.
// source: complex/complex.proto

package complexpb

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

type ComplexMessage struct {
	OneDummy             *DummyMessage   `protobuf:"bytes,2,opt,name=one_dummy,json=oneDummy,proto3" json:"one_dummy,omitempty"`
	MultipleDummy        []*DummyMessage `protobuf:"bytes,3,rep,name=multiple_dummy,json=multipleDummy,proto3" json:"multiple_dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ComplexMessage) Reset()         { *m = ComplexMessage{} }
func (m *ComplexMessage) String() string { return proto.CompactTextString(m) }
func (*ComplexMessage) ProtoMessage()    {}
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_complex_c2b6e3162031a7a9, []int{0}
}
func (m *ComplexMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexMessage.Unmarshal(m, b)
}
func (m *ComplexMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexMessage.Marshal(b, m, deterministic)
}
func (dst *ComplexMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexMessage.Merge(dst, src)
}
func (m *ComplexMessage) XXX_Size() int {
	return xxx_messageInfo_ComplexMessage.Size(m)
}
func (m *ComplexMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexMessage proto.InternalMessageInfo

func (m *ComplexMessage) GetOneDummy() *DummyMessage {
	if m != nil {
		return m.OneDummy
	}
	return nil
}

func (m *ComplexMessage) GetMultipleDummy() []*DummyMessage {
	if m != nil {
		return m.MultipleDummy
	}
	return nil
}

type DummyMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyMessage) Reset()         { *m = DummyMessage{} }
func (m *DummyMessage) String() string { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()    {}
func (*DummyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_complex_c2b6e3162031a7a9, []int{1}
}
func (m *DummyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyMessage.Unmarshal(m, b)
}
func (m *DummyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyMessage.Marshal(b, m, deterministic)
}
func (dst *DummyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyMessage.Merge(dst, src)
}
func (m *DummyMessage) XXX_Size() int {
	return xxx_messageInfo_DummyMessage.Size(m)
}
func (m *DummyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DummyMessage proto.InternalMessageInfo

func (m *DummyMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DummyMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ComplexMessage)(nil), "example.complex.ComplexMessage")
	proto.RegisterType((*DummyMessage)(nil), "example.complex.DummyMessage")
}

func init() { proto.RegisterFile("complex/complex.proto", fileDescriptor_complex_c2b6e3162031a7a9) }

var fileDescriptor_complex_c2b6e3162031a7a9 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2d,
	0xc8, 0x49, 0xad, 0xd0, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xfc, 0xa9, 0x15,
	0x89, 0x20, 0xbe, 0x1e, 0x54, 0x58, 0x69, 0x12, 0x23, 0x17, 0x9f, 0x33, 0x84, 0xed, 0x9b, 0x5a,
	0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x64, 0xc5, 0xc5, 0x99, 0x9f, 0x97, 0x1a, 0x9f, 0x52, 0x9a, 0x9b,
	0x5b, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xab, 0x87, 0xa6, 0x4f, 0xcf, 0x05, 0x24,
	0x0b, 0xd5, 0x11, 0xc4, 0x91, 0x9f, 0x97, 0x0a, 0x16, 0x10, 0x72, 0xe1, 0xe2, 0xcb, 0x2d, 0xcd,
	0x29, 0xc9, 0x2c, 0xc8, 0x81, 0x19, 0xc0, 0xac, 0xc0, 0x4c, 0xd8, 0x00, 0x5e, 0x98, 0x26, 0xb0,
	0xa8, 0x92, 0x11, 0x17, 0x0f, 0xb2, 0xb4, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0x2a, 0xd8,
	0x71, 0x9c, 0x41, 0x60, 0xb6, 0x13, 0x77, 0x14, 0x27, 0xd4, 0xe8, 0x82, 0xa4, 0x24, 0x36, 0xb0,
	0x6f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x59, 0x14, 0x8b, 0x06, 0x01, 0x00, 0x00,
}
