// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dentmeter.proto

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

type DentMeterState struct {
	// unit:ns
	Time                 uint64        `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Phases               []*PhaseState `protobuf:"bytes,2,rep,name=phases,proto3" json:"phases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DentMeterState) Reset()         { *m = DentMeterState{} }
func (m *DentMeterState) String() string { return proto.CompactTextString(m) }
func (*DentMeterState) ProtoMessage()    {}
func (*DentMeterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dentmeter_0503a193ebc33e9f, []int{0}
}
func (m *DentMeterState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DentMeterState.Unmarshal(m, b)
}
func (m *DentMeterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DentMeterState.Marshal(b, m, deterministic)
}
func (dst *DentMeterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DentMeterState.Merge(dst, src)
}
func (m *DentMeterState) XXX_Size() int {
	return xxx_messageInfo_DentMeterState.Size(m)
}
func (m *DentMeterState) XXX_DiscardUnknown() {
	xxx_messageInfo_DentMeterState.DiscardUnknown(m)
}

var xxx_messageInfo_DentMeterState proto.InternalMessageInfo

func (m *DentMeterState) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DentMeterState) GetPhases() []*PhaseState {
	if m != nil {
		return m.Phases
	}
	return nil
}

type PhaseState struct {
	Phase      string `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	Annotation string `protobuf:"bytes,14,opt,name=annotation,proto3" json:"annotation,omitempty"`
	// unit: kWh
	TrueEnergy float64 `protobuf:"fixed64,2,opt,name=true_energy,json=trueEnergy,proto3" json:"true_energy,omitempty"`
	// unit: kVARh
	ReactiveEnergy float64 `protobuf:"fixed64,3,opt,name=reactive_energy,json=reactiveEnergy,proto3" json:"reactive_energy,omitempty"`
	// unit: kVAh
	ApparentEnergy float64 `protobuf:"fixed64,4,opt,name=apparent_energy,json=apparentEnergy,proto3" json:"apparent_energy,omitempty"`
	// unit: kW
	TruePower float64 `protobuf:"fixed64,5,opt,name=true_power,json=truePower,proto3" json:"true_power,omitempty"`
	// unit: kVAR
	ReactivePower float64 `protobuf:"fixed64,6,opt,name=reactive_power,json=reactivePower,proto3" json:"reactive_power,omitempty"`
	// unit: kVA
	ApparentPower float64 `protobuf:"fixed64,7,opt,name=apparent_power,json=apparentPower,proto3" json:"apparent_power,omitempty"`
	// unit: PF
	DisplacementPf float64 `protobuf:"fixed64,8,opt,name=displacement_pf,json=displacementPf,proto3" json:"displacement_pf,omitempty"`
	// unit: PF
	ApparentPf float64 `protobuf:"fixed64,9,opt,name=apparent_pf,json=apparentPf,proto3" json:"apparent_pf,omitempty"`
	// unit: A
	Current float64 `protobuf:"fixed64,10,opt,name=current,proto3" json:"current,omitempty"`
	// unit: HZ
	LineFrequency float64 `protobuf:"fixed64,11,opt,name=line_frequency,json=lineFrequency,proto3" json:"line_frequency,omitempty"`
	// unit: V
	Volts float64 `protobuf:"fixed64,12,opt,name=volts,proto3" json:"volts,omitempty"`
	// unit: V
	PhaseNeutralVoltage  float64  `protobuf:"fixed64,13,opt,name=phase_neutral_voltage,json=phaseNeutralVoltage,proto3" json:"phase_neutral_voltage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhaseState) Reset()         { *m = PhaseState{} }
func (m *PhaseState) String() string { return proto.CompactTextString(m) }
func (*PhaseState) ProtoMessage()    {}
func (*PhaseState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dentmeter_0503a193ebc33e9f, []int{1}
}
func (m *PhaseState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhaseState.Unmarshal(m, b)
}
func (m *PhaseState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhaseState.Marshal(b, m, deterministic)
}
func (dst *PhaseState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhaseState.Merge(dst, src)
}
func (m *PhaseState) XXX_Size() int {
	return xxx_messageInfo_PhaseState.Size(m)
}
func (m *PhaseState) XXX_DiscardUnknown() {
	xxx_messageInfo_PhaseState.DiscardUnknown(m)
}

var xxx_messageInfo_PhaseState proto.InternalMessageInfo

func (m *PhaseState) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func (m *PhaseState) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

func (m *PhaseState) GetTrueEnergy() float64 {
	if m != nil {
		return m.TrueEnergy
	}
	return 0
}

func (m *PhaseState) GetReactiveEnergy() float64 {
	if m != nil {
		return m.ReactiveEnergy
	}
	return 0
}

func (m *PhaseState) GetApparentEnergy() float64 {
	if m != nil {
		return m.ApparentEnergy
	}
	return 0
}

func (m *PhaseState) GetTruePower() float64 {
	if m != nil {
		return m.TruePower
	}
	return 0
}

func (m *PhaseState) GetReactivePower() float64 {
	if m != nil {
		return m.ReactivePower
	}
	return 0
}

func (m *PhaseState) GetApparentPower() float64 {
	if m != nil {
		return m.ApparentPower
	}
	return 0
}

func (m *PhaseState) GetDisplacementPf() float64 {
	if m != nil {
		return m.DisplacementPf
	}
	return 0
}

func (m *PhaseState) GetApparentPf() float64 {
	if m != nil {
		return m.ApparentPf
	}
	return 0
}

func (m *PhaseState) GetCurrent() float64 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *PhaseState) GetLineFrequency() float64 {
	if m != nil {
		return m.LineFrequency
	}
	return 0
}

func (m *PhaseState) GetVolts() float64 {
	if m != nil {
		return m.Volts
	}
	return 0
}

func (m *PhaseState) GetPhaseNeutralVoltage() float64 {
	if m != nil {
		return m.PhaseNeutralVoltage
	}
	return 0
}

func init() {
	proto.RegisterType((*DentMeterState)(nil), "xbospb.DentMeterState")
	proto.RegisterType((*PhaseState)(nil), "xbospb.PhaseState")
}

func init() { proto.RegisterFile("dentmeter.proto", fileDescriptor_dentmeter_0503a193ebc33e9f) }

var fileDescriptor_dentmeter_0503a193ebc33e9f = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x92, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x33, 0x81, 0x21, 0x0f, 0x81, 0xa4, 0x6a, 0xd2, 0x8b, 0x4a, 0x48, 0x8c, 0xc4, 0x03,
	0x07, 0xfc, 0x0a, 0xea, 0x4d, 0xb3, 0xcc, 0xc4, 0xeb, 0x52, 0xc6, 0x1b, 0x2e, 0x19, 0x5d, 0xed,
	0x3a, 0x94, 0x4f, 0xee, 0xd5, 0xf5, 0x95, 0x0e, 0x6e, 0x7b, 0xbf, 0xff, 0xaf, 0xaf, 0x6f, 0x6d,
	0x61, 0xb2, 0x46, 0x69, 0xb6, 0x68, 0x50, 0x2f, 0x94, 0x2e, 0x4d, 0xc9, 0xc2, 0xdf, 0x55, 0x59,
	0xa9, 0xd5, 0x2c, 0x82, 0xf1, 0x73, 0x13, 0xbd, 0xd9, 0xe8, 0xc3, 0x08, 0x83, 0x8c, 0x41, 0xd7,
	0xe4, 0x5b, 0xe4, 0xc1, 0x34, 0x98, 0x77, 0x63, 0xfa, 0x66, 0x8f, 0x10, 0xaa, 0x2f, 0x51, 0x61,
	0xc5, 0xcf, 0xa6, 0x9d, 0xf9, 0x70, 0xc9, 0x16, 0x6e, 0xf9, 0x22, 0xb2, 0x94, 0xd6, 0xc5, 0x07,
	0x63, 0xf6, 0xd7, 0x01, 0x38, 0x62, 0x76, 0x05, 0x3d, 0x0a, 0xa8, 0xdf, 0x20, 0x76, 0x05, 0xbb,
	0x05, 0x10, 0x52, 0x96, 0x8d, 0x91, 0x97, 0x92, 0x8f, 0x29, 0x3a, 0x21, 0xec, 0x0e, 0x86, 0x46,
	0xd7, 0x98, 0xa0, 0x44, 0xbd, 0xd9, 0x37, 0xbb, 0x06, 0xf3, 0x20, 0x06, 0x8b, 0x5e, 0x88, 0xb0,
	0x07, 0x98, 0x68, 0x14, 0xa9, 0xc9, 0x77, 0xad, 0xd4, 0x21, 0x69, 0xec, 0xf1, 0x51, 0x14, 0x4a,
	0x09, 0xdd, 0xfc, 0xa4, 0x17, 0xbb, 0x4e, 0xf4, 0xf8, 0x20, 0xde, 0x00, 0xf5, 0x4f, 0x54, 0xf9,
	0x83, 0x9a, 0xf7, 0xc8, 0x19, 0x58, 0x12, 0x59, 0xc0, 0xee, 0xa1, 0xed, 0x7c, 0x50, 0x42, 0x52,
	0x46, 0x9e, 0xb6, 0x5a, 0xbb, 0x9d, 0xd3, 0xfa, 0x4e, 0xf3, 0xd4, 0x69, 0xcd, 0x54, 0xeb, 0xbc,
	0x52, 0x85, 0x48, 0x71, 0x4b, 0x6a, 0xc6, 0xcf, 0xdd, 0x54, 0xa7, 0x38, 0xca, 0xec, 0x41, 0x1c,
	0xfb, 0x65, 0x7c, 0xe0, 0x0e, 0xa2, 0x6d, 0x96, 0x31, 0x0e, 0xfd, 0xb4, 0xd6, 0xb6, 0xe0, 0x40,
	0xa1, 0x2f, 0xed, 0x28, 0x45, 0x2e, 0x31, 0xc9, 0x34, 0x7e, 0xd7, 0x28, 0xd3, 0x3d, 0x1f, 0xba,
	0x51, 0x2c, 0x7d, 0xf5, 0xd0, 0x5e, 0xd0, 0xae, 0x2c, 0x4c, 0xc5, 0x2f, 0x28, 0x75, 0x05, 0x5b,
	0xc2, 0x35, 0xdd, 0x54, 0x22, 0xb1, 0x36, 0x5a, 0x14, 0x89, 0xc5, 0x62, 0x83, 0x7c, 0x44, 0xd6,
	0x25, 0x85, 0xef, 0x2e, 0xfb, 0x74, 0xd1, 0x2a, 0xa4, 0xa7, 0xf5, 0xf4, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x61, 0xc4, 0xbf, 0x6d, 0x02, 0x00, 0x00,
}
