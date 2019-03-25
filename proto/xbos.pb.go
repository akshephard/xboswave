// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xbos.proto

package xbospb

/*
This defines the protocol buffers used by XBOS with WaveMQ. The schema for
a wavemq xbos proto message is:
  xbosproto/RootMessageType

Version 1.0
*/

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Root union type
// We reserve 50 fields per imported file
type XBOS struct {
	// Hamilton Fields
	HamiltonData        *HamiltonData        `protobuf:"bytes,50,opt,name=HamiltonData,proto3" json:"HamiltonData,omitempty"`
	HamiltonBRLinkStats *HamiltonBRLinkStats `protobuf:"bytes,51,opt,name=HamiltonBRLinkStats,proto3" json:"HamiltonBRLinkStats,omitempty"`
	HamiltonBRMessage   *HamiltonBRMessage   `protobuf:"bytes,52,opt,name=HamiltonBRMessage,proto3" json:"HamiltonBRMessage,omitempty"`
	// IoT Fields
	XBOSIoTDeviceState     *XBOSIoTDeviceState     `protobuf:"bytes,100,opt,name=XBOSIoTDeviceState,proto3" json:"XBOSIoTDeviceState,omitempty"`
	XBOSIoTDeviceActuation *XBOSIoTDeviceActuation `protobuf:"bytes,101,opt,name=XBOSIoTDeviceActuation,proto3" json:"XBOSIoTDeviceActuation,omitempty"`
	XBOSIoTContext         *XBOSIoTContext         `protobuf:"bytes,102,opt,name=XBOSIoTContext,proto3" json:"XBOSIoTContext,omitempty"`
	// Dent Fields
	DentMeterState *DentMeterState `protobuf:"bytes,150,opt,name=DentMeterState,proto3" json:"DentMeterState,omitempty"`
	// System Status Fields
	BasicServerStatus    *BasicServerStatus `protobuf:"bytes,200,opt,name=BasicServerStatus,proto3" json:"BasicServerStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *XBOS) Reset()         { *m = XBOS{} }
func (m *XBOS) String() string { return proto.CompactTextString(m) }
func (*XBOS) ProtoMessage()    {}
func (*XBOS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dcd15659c7161cd, []int{0}
}
func (m *XBOS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOS.Unmarshal(m, b)
}
func (m *XBOS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOS.Marshal(b, m, deterministic)
}
func (m *XBOS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOS.Merge(m, src)
}
func (m *XBOS) XXX_Size() int {
	return xxx_messageInfo_XBOS.Size(m)
}
func (m *XBOS) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOS.DiscardUnknown(m)
}

var xxx_messageInfo_XBOS proto.InternalMessageInfo

func (m *XBOS) GetHamiltonData() *HamiltonData {
	if m != nil {
		return m.HamiltonData
	}
	return nil
}

func (m *XBOS) GetHamiltonBRLinkStats() *HamiltonBRLinkStats {
	if m != nil {
		return m.HamiltonBRLinkStats
	}
	return nil
}

func (m *XBOS) GetHamiltonBRMessage() *HamiltonBRMessage {
	if m != nil {
		return m.HamiltonBRMessage
	}
	return nil
}

func (m *XBOS) GetXBOSIoTDeviceState() *XBOSIoTDeviceState {
	if m != nil {
		return m.XBOSIoTDeviceState
	}
	return nil
}

func (m *XBOS) GetXBOSIoTDeviceActuation() *XBOSIoTDeviceActuation {
	if m != nil {
		return m.XBOSIoTDeviceActuation
	}
	return nil
}

func (m *XBOS) GetXBOSIoTContext() *XBOSIoTContext {
	if m != nil {
		return m.XBOSIoTContext
	}
	return nil
}

func (m *XBOS) GetDentMeterState() *DentMeterState {
	if m != nil {
		return m.DentMeterState
	}
	return nil
}

func (m *XBOS) GetBasicServerStatus() *BasicServerStatus {
	if m != nil {
		return m.BasicServerStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*XBOS)(nil), "xbospb.XBOS")
}

func init() { proto.RegisterFile("xbos.proto", fileDescriptor_3dcd15659c7161cd) }

var fileDescriptor_3dcd15659c7161cd = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0x19, 0x8c, 0xc1, 0x9b, 0x57, 0x26, 0x8b, 0x63, 0xc4, 0x0a, 0x22, 0x5e, 0x79, 0xd5,
	0x8b, 0xcd, 0x0b, 0xaf, 0x14, 0x6b, 0xc1, 0x29, 0x16, 0x21, 0x15, 0xf1, 0x4e, 0xd2, 0xee, 0xa8,
	0x41, 0x9b, 0x8c, 0xe6, 0x6c, 0xcc, 0x2f, 0xe2, 0x57, 0xd2, 0x8f, 0x25, 0x4d, 0x33, 0x5d, 0xff,
	0x78, 0x97, 0x3e, 0xe7, 0xf7, 0xfc, 0xca, 0x69, 0x43, 0xc8, 0x2a, 0xd1, 0xc6, 0x9f, 0xe7, 0x1a,
	0x35, 0xed, 0x15, 0xe7, 0x79, 0xe2, 0xf5, 0x5f, 0x44, 0x26, 0xdf, 0x50, 0xab, 0x32, 0xf7, 0xfe,
	0x49, 0x8d, 0xee, 0xb8, 0x3d, 0x03, 0x85, 0x19, 0x20, 0xe4, 0x2e, 0x18, 0x9a, 0x77, 0x83, 0x90,
	0x3d, 0x66, 0x5a, 0x49, 0xd4, 0x2e, 0x3d, 0xfc, 0xec, 0x92, 0xee, 0x43, 0x70, 0x1b, 0xd3, 0x13,
	0xb2, 0x35, 0x75, 0xb2, 0x50, 0xa0, 0x60, 0xe3, 0x83, 0xce, 0xd1, 0xff, 0xf1, 0xd0, 0x2f, 0xdf,
	0xe4, 0x6f, 0xce, 0x78, 0x85, 0xa4, 0x11, 0xd9, 0x59, 0x3f, 0x07, 0xfc, 0x46, 0xaa, 0xd7, 0x18,
	0x05, 0x1a, 0x36, 0xb1, 0x82, 0xbd, 0xba, 0x60, 0x03, 0xe1, 0x6d, 0x3d, 0x7a, 0x49, 0x06, 0xbf,
	0x71, 0x04, 0xc6, 0x88, 0x67, 0x60, 0xc7, 0x56, 0xb6, 0xdb, 0x94, 0x39, 0x80, 0x37, 0x3b, 0xf4,
	0x9a, 0xd0, 0x62, 0xb3, 0x2b, 0x7d, 0x17, 0xc2, 0x52, 0xa6, 0x50, 0xe8, 0x81, 0xcd, 0xac, 0xc9,
	0x5b, 0x9b, 0x9a, 0x04, 0x6f, 0x69, 0xd1, 0x7b, 0x32, 0xaa, 0xa4, 0xe7, 0x29, 0x2e, 0x04, 0x4a,
	0xad, 0x18, 0x58, 0xdf, 0x7e, 0xab, 0xef, 0x87, 0xe2, 0x7f, 0xb4, 0xe9, 0x29, 0xe9, 0xbb, 0xc9,
	0x85, 0x56, 0x08, 0x2b, 0x64, 0x4f, 0xd6, 0x37, 0xaa, 0xf9, 0xdc, 0x94, 0xd7, 0x68, 0x7a, 0x46,
	0xfa, 0x21, 0x28, 0x8c, 0x8a, 0xff, 0x5c, 0xee, 0xf7, 0xd1, 0xa9, 0x0a, 0xaa, 0x63, 0x5e, 0xc3,
	0xe9, 0x94, 0x0c, 0x02, 0x61, 0x64, 0x1a, 0x43, 0xbe, 0x2c, 0xb3, 0x85, 0x61, 0x5f, 0x9d, 0xea,
	0xe7, 0x6e, 0x10, 0xbc, 0x59, 0x4a, 0x7a, 0xf6, 0x42, 0x4d, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x2a, 0x27, 0x17, 0xa8, 0x02, 0x00, 0x00,
}