// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nullabletypes.proto

package xbospb

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

type Int32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_nullabletypes_ffbf8acdca2a5424, []int{0}
}
func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (dst *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(dst, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_nullabletypes_ffbf8acdca2a5424, []int{1}
}
func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (dst *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(dst, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint64 struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint64) Reset()         { *m = Uint64{} }
func (m *Uint64) String() string { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()    {}
func (*Uint64) Descriptor() ([]byte, []int) {
	return fileDescriptor_nullabletypes_ffbf8acdca2a5424, []int{2}
}
func (m *Uint64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint64.Unmarshal(m, b)
}
func (m *Uint64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint64.Marshal(b, m, deterministic)
}
func (dst *Uint64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint64.Merge(dst, src)
}
func (m *Uint64) XXX_Size() int {
	return xxx_messageInfo_Uint64.Size(m)
}
func (m *Uint64) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint64.DiscardUnknown(m)
}

var xxx_messageInfo_Uint64 proto.InternalMessageInfo

func (m *Uint64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Double struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Double) Reset()         { *m = Double{} }
func (m *Double) String() string { return proto.CompactTextString(m) }
func (*Double) ProtoMessage()    {}
func (*Double) Descriptor() ([]byte, []int) {
	return fileDescriptor_nullabletypes_ffbf8acdca2a5424, []int{3}
}
func (m *Double) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Double.Unmarshal(m, b)
}
func (m *Double) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Double.Marshal(b, m, deterministic)
}
func (dst *Double) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Double.Merge(dst, src)
}
func (m *Double) XXX_Size() int {
	return xxx_messageInfo_Double.Size(m)
}
func (m *Double) XXX_DiscardUnknown() {
	xxx_messageInfo_Double.DiscardUnknown(m)
}

var xxx_messageInfo_Double proto.InternalMessageInfo

func (m *Double) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_nullabletypes_ffbf8acdca2a5424, []int{4}
}
func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (dst *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(dst, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*Int32)(nil), "xbospb.Int32")
	proto.RegisterType((*Int64)(nil), "xbospb.Int64")
	proto.RegisterType((*Uint64)(nil), "xbospb.Uint64")
	proto.RegisterType((*Double)(nil), "xbospb.Double")
	proto.RegisterType((*Bool)(nil), "xbospb.Bool")
}

func init() { proto.RegisterFile("nullabletypes.proto", fileDescriptor_nullabletypes_ffbf8acdca2a5424) }

var fileDescriptor_nullabletypes_ffbf8acdca2a5424 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2b, 0xcd, 0xc9,
	0x49, 0x4c, 0xca, 0x49, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xab, 0x48, 0xca, 0x2f, 0x2e, 0x48, 0x52, 0x92, 0xe5, 0x62, 0xf5, 0xcc, 0x2b, 0x31, 0x36,
	0x12, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d,
	0x82, 0x70, 0xa0, 0xd2, 0x66, 0x26, 0xa8, 0xd2, 0xcc, 0x30, 0x69, 0x39, 0x2e, 0xb6, 0xd0, 0x4c,
	0x4c, 0x79, 0x16, 0x24, 0x79, 0x97, 0xfc, 0x52, 0xa0, 0xd5, 0xa8, 0xf2, 0x8c, 0x30, 0x79, 0x19,
	0x2e, 0x16, 0xa7, 0xfc, 0xfc, 0x1c, 0x54, 0x59, 0x0e, 0xa8, 0x6c, 0x12, 0x1b, 0xd8, 0xa9, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x2e, 0x5c, 0xff, 0xc1, 0x00, 0x00, 0x00,
}
