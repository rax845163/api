// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/score/v1alpha2/score.proto

package v1alpha2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//*
// Represents a system score before and after pod scheduled on node
//
type SimulatedSchedulingScore struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	ScoreBefore          float64              `protobuf:"fixed64,2,opt,name=score_before,json=scoreBefore,proto3" json:"score_before,omitempty"`
	ScoreAfter           float64              `protobuf:"fixed64,3,opt,name=score_after,json=scoreAfter,proto3" json:"score_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SimulatedSchedulingScore) Reset()         { *m = SimulatedSchedulingScore{} }
func (m *SimulatedSchedulingScore) String() string { return proto.CompactTextString(m) }
func (*SimulatedSchedulingScore) ProtoMessage()    {}
func (*SimulatedSchedulingScore) Descriptor() ([]byte, []int) {
	return fileDescriptor_95e3aed5b59bfd54, []int{0}
}

func (m *SimulatedSchedulingScore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimulatedSchedulingScore.Unmarshal(m, b)
}
func (m *SimulatedSchedulingScore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimulatedSchedulingScore.Marshal(b, m, deterministic)
}
func (m *SimulatedSchedulingScore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulatedSchedulingScore.Merge(m, src)
}
func (m *SimulatedSchedulingScore) XXX_Size() int {
	return xxx_messageInfo_SimulatedSchedulingScore.Size(m)
}
func (m *SimulatedSchedulingScore) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulatedSchedulingScore.DiscardUnknown(m)
}

var xxx_messageInfo_SimulatedSchedulingScore proto.InternalMessageInfo

func (m *SimulatedSchedulingScore) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *SimulatedSchedulingScore) GetScoreBefore() float64 {
	if m != nil {
		return m.ScoreBefore
	}
	return 0
}

func (m *SimulatedSchedulingScore) GetScoreAfter() float64 {
	if m != nil {
		return m.ScoreAfter
	}
	return 0
}

func init() {
	proto.RegisterType((*SimulatedSchedulingScore)(nil), "containersai.datahub.score.v1alpha2.SimulatedSchedulingScore")
}

func init() { proto.RegisterFile("datahub/score/v1alpha2/score.proto", fileDescriptor_95e3aed5b59bfd54) }

var fileDescriptor_95e3aed5b59bfd54 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0x75, 0x80, 0x18, 0xae, 0x4c, 0x99, 0xa2, 0x2e, 0x2d, 0x65, 0xe9, 0xc2, 0x9d, 0x68,
	0xc5, 0x03, 0xd0, 0x47, 0x68, 0x98, 0x58, 0x22, 0x27, 0x71, 0x12, 0x4b, 0x49, 0x1c, 0x5d, 0x1c,
	0x1e, 0x83, 0x67, 0x46, 0xf1, 0x11, 0xb1, 0x74, 0xf4, 0xaf, 0xef, 0xff, 0xee, 0x6c, 0x7b, 0xa8,
	0x40, 0xa0, 0x9d, 0x0b, 0x3f, 0x95, 0x1c, 0xd0, 0x7f, 0xbf, 0x41, 0x37, 0xb6, 0x70, 0x8a, 0xa3,
	0x1b, 0x03, 0x0b, 0x27, 0x2f, 0x25, 0x0f, 0x02, 0x34, 0x60, 0x98, 0x80, 0xdc, 0x5f, 0xc1, 0x45,
	0x62, 0x2d, 0x6c, 0x77, 0x0d, 0x73, 0xd3, 0xa1, 0xd7, 0x4a, 0x31, 0xd7, 0x5e, 0xa8, 0xc7, 0x49,
	0xa0, 0x1f, 0xa3, 0xe5, 0xf0, 0x63, 0x6c, 0x9a, 0x51, 0x3f, 0x77, 0x20, 0x58, 0x65, 0x65, 0x8b,
	0xd5, 0xdc, 0xd1, 0xd0, 0x64, 0x8b, 0x26, 0x71, 0xf6, 0x61, 0xe1, 0x53, 0xb3, 0x37, 0xc7, 0xcd,
	0x69, 0xeb, 0xa2, 0xcc, 0xad, 0x32, 0xf7, 0xb9, 0xca, 0xae, 0xca, 0x25, 0xcf, 0xf6, 0x49, 0xdf,
	0xcf, 0x0b, 0xac, 0x39, 0x60, 0x7a, 0xb7, 0x37, 0x47, 0x73, 0xdd, 0x68, 0x76, 0xd1, 0x28, 0xd9,
	0xd9, 0x38, 0xe6, 0x50, 0x0b, 0x86, 0xf4, 0x5e, 0x09, 0xab, 0xd1, 0xc7, 0x92, 0x5c, 0xde, 0xbf,
	0xce, 0x0d, 0xc9, 0xb2, 0x4a, 0xc9, 0xbd, 0xff, 0xdf, 0xf1, 0x15, 0xc8, 0xc3, 0x48, 0xfe, 0xf6,
	0x65, 0x8a, 0x47, 0xfd, 0xd4, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x08, 0x5d, 0xe2, 0x3a,
	0x01, 0x00, 0x00,
}
