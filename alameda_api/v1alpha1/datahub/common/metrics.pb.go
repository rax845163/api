// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/metrics.proto

package common

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
// Metric type. A metric may be either CPU or memory.
type MetricType int32

const (
	MetricType_METRICS_TYPE_UNDEFINED       MetricType = 0
	MetricType_CPU_USAGE_SECONDS_PERCENTAGE MetricType = 1
	MetricType_MEMORY_USAGE_BYTES           MetricType = 2
	MetricType_POWER_USAGE_WATTS            MetricType = 3
	MetricType_TEMPERATURE_CELSIUS          MetricType = 4
	MetricType_DUTY_CYCLE                   MetricType = 5
)

var MetricType_name = map[int32]string{
	0: "METRICS_TYPE_UNDEFINED",
	1: "CPU_USAGE_SECONDS_PERCENTAGE",
	2: "MEMORY_USAGE_BYTES",
	3: "POWER_USAGE_WATTS",
	4: "TEMPERATURE_CELSIUS",
	5: "DUTY_CYCLE",
}

var MetricType_value = map[string]int32{
	"METRICS_TYPE_UNDEFINED":       0,
	"CPU_USAGE_SECONDS_PERCENTAGE": 1,
	"MEMORY_USAGE_BYTES":           2,
	"POWER_USAGE_WATTS":            3,
	"TEMPERATURE_CELSIUS":          4,
	"DUTY_CYCLE":                   5,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{0}
}

type ResourceName int32

const (
	ResourceName_RESOURCE_NAME_UNDEFINED ResourceName = 0
	ResourceName_CPU                     ResourceName = 1
	ResourceName_MEMORY                  ResourceName = 2
)

var ResourceName_name = map[int32]string{
	0: "RESOURCE_NAME_UNDEFINED",
	1: "CPU",
	2: "MEMORY",
}

var ResourceName_value = map[string]int32{
	"RESOURCE_NAME_UNDEFINED": 0,
	"CPU":                     1,
	"MEMORY":                  2,
}

func (x ResourceName) String() string {
	return proto.EnumName(ResourceName_name, int32(x))
}

func (ResourceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{1}
}

//*
// Represents a data point of time-series metric data
type Sample struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	NumValue             string               `protobuf:"bytes,3,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{0}
}

func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Sample) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Sample) GetNumValue() string {
	if m != nil {
		return m.NumValue
	}
	return ""
}

//*
// Represents a piece of metreic data
type MetricData struct {
	MetricType           MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	Data                 []*Sample  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Granularity          int64      `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{1}
}

func (m *MetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricData.Unmarshal(m, b)
}
func (m *MetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricData.Marshal(b, m, deterministic)
}
func (m *MetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricData.Merge(m, src)
}
func (m *MetricData) XXX_Size() int {
	return xxx_messageInfo_MetricData.Size(m)
}
func (m *MetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricData proto.InternalMessageInfo

func (m *MetricData) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_METRICS_TYPE_UNDEFINED
}

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MetricData) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.MetricType", MetricType_name, MetricType_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.ResourceName", ResourceName_name, ResourceName_value)
	proto.RegisterType((*Sample)(nil), "containersai.alameda.v1alpha1.datahub.common.Sample")
	proto.RegisterType((*MetricData)(nil), "containersai.alameda.v1alpha1.datahub.common.MetricData")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/metrics.proto", fileDescriptor_0dcb68d43798345b)
}

var fileDescriptor_0dcb68d43798345b = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6e, 0xd3, 0x4c,
	0x10, 0x85, 0x7f, 0x27, 0xf9, 0xd3, 0x76, 0x82, 0x2a, 0xb3, 0x88, 0x36, 0x0a, 0x48, 0x44, 0xbd,
	0x8a, 0x2a, 0x58, 0x2b, 0x01, 0x24, 0x2e, 0x9b, 0x38, 0x4b, 0x89, 0x54, 0x3b, 0xd6, 0xee, 0x9a,
	0xca, 0xdc, 0xac, 0x36, 0xc9, 0x92, 0x5a, 0xf2, 0xda, 0x96, 0xb3, 0xae, 0x94, 0x77, 0xe0, 0x29,
	0x78, 0x25, 0x5e, 0x08, 0xc5, 0x4e, 0x0a, 0x5c, 0x41, 0xef, 0x3c, 0xe3, 0xf9, 0x8e, 0xe7, 0x1c,
	0x0f, 0x0c, 0x65, 0x22, 0xb5, 0x5a, 0x49, 0x21, 0xf3, 0xd8, 0xb9, 0x1f, 0xca, 0x24, 0xbf, 0x93,
	0x43, 0x67, 0x25, 0x8d, 0xbc, 0x2b, 0x17, 0xce, 0x32, 0xd3, 0x3a, 0x4b, 0x1d, 0xad, 0x4c, 0x11,
	0x2f, 0x37, 0x38, 0x2f, 0x32, 0x93, 0xa1, 0xd7, 0xcb, 0x2c, 0x35, 0x32, 0x4e, 0x55, 0xb1, 0x91,
	0x31, 0xde, 0xf3, 0xf8, 0xc0, 0xe2, 0x3d, 0x8b, 0x6b, 0xb6, 0xf7, 0x6a, 0x9d, 0x65, 0xeb, 0x44,
	0x39, 0x15, 0xbb, 0x28, 0xbf, 0x3a, 0x26, 0xd6, 0x6a, 0x63, 0xa4, 0xce, 0x6b, 0xb9, 0x8b, 0x6f,
	0x16, 0xb4, 0x99, 0xd4, 0x79, 0xa2, 0x10, 0x86, 0xd6, 0xee, 0x6d, 0xd7, 0xea, 0x5b, 0x83, 0xce,
	0xa8, 0x87, 0x6b, 0x14, 0x1f, 0x50, 0xcc, 0x0f, 0x28, 0xad, 0xe6, 0xd0, 0x7b, 0x38, 0x56, 0xe9,
	0x4a, 0x54, 0x4c, 0xe3, 0xaf, 0xcc, 0x91, 0x4a, 0x57, 0xbb, 0x0a, 0xbd, 0x80, 0x93, 0xb4, 0xd4,
	0xe2, 0x5e, 0x26, 0xa5, 0xea, 0x36, 0xfb, 0xd6, 0xe0, 0x84, 0x1e, 0xa7, 0xa5, 0xfe, 0xbc, 0xab,
	0x2f, 0x7e, 0x58, 0x00, 0x5e, 0xe5, 0x77, 0x2a, 0x8d, 0x44, 0x11, 0x74, 0x6a, 0xf7, 0xc2, 0x6c,
	0xf3, 0x7a, 0xb3, 0xd3, 0xd1, 0x07, 0xfc, 0x98, 0x08, 0x70, 0x2d, 0xc7, 0xb7, 0xb9, 0xa2, 0xa0,
	0x1f, 0x9e, 0xd1, 0x27, 0x68, 0xed, 0x06, 0xbb, 0x8d, 0x7e, 0x73, 0xd0, 0x19, 0xbd, 0x7b, 0x9c,
	0x66, 0x9d, 0x18, 0xad, 0x14, 0x50, 0x1f, 0x3a, 0xeb, 0x42, 0xa6, 0x65, 0x22, 0x8b, 0xd8, 0x6c,
	0x2b, 0x4b, 0x4d, 0xfa, 0x7b, 0xeb, 0xf2, 0xfb, 0x83, 0xab, 0xea, 0xd3, 0x3d, 0x38, 0xf3, 0x08,
	0xa7, 0x33, 0x97, 0x09, 0x1e, 0x05, 0x44, 0x84, 0xfe, 0x94, 0x7c, 0x9c, 0xf9, 0x64, 0x6a, 0xff,
	0x87, 0xfa, 0xf0, 0xd2, 0x0d, 0x42, 0x11, 0xb2, 0xf1, 0x35, 0x11, 0x8c, 0xb8, 0x73, 0x7f, 0xca,
	0x44, 0x40, 0xa8, 0x4b, 0x7c, 0x3e, 0xbe, 0x26, 0xb6, 0x85, 0xce, 0x00, 0x79, 0xc4, 0x9b, 0xd3,
	0x68, 0x3f, 0x34, 0x89, 0x38, 0x61, 0x76, 0x03, 0x3d, 0x87, 0xa7, 0xc1, 0xfc, 0x96, 0xd0, 0x7d,
	0xfb, 0x76, 0xcc, 0x39, 0xb3, 0x9b, 0xe8, 0x1c, 0x9e, 0x71, 0xe2, 0x05, 0x84, 0x8e, 0x79, 0x48,
	0x89, 0x70, 0xc9, 0x0d, 0x9b, 0x85, 0xcc, 0x6e, 0xa1, 0x53, 0x80, 0x69, 0xc8, 0x23, 0xe1, 0x46,
	0xee, 0x0d, 0xb1, 0xff, 0xbf, 0xbc, 0x82, 0x27, 0x54, 0x6d, 0xb2, 0xb2, 0x58, 0x2a, 0x5f, 0x56,
	0xff, 0xe9, 0x9c, 0x12, 0x36, 0x0f, 0xa9, 0x4b, 0x84, 0x3f, 0xf6, 0xfe, 0x5c, 0xf3, 0x08, 0x9a,
	0x6e, 0x10, 0xda, 0x16, 0x02, 0x68, 0xd7, 0xdb, 0xd8, 0x8d, 0xc9, 0xe4, 0xcb, 0xd5, 0x3a, 0x36,
	0xfb, 0x98, 0x9c, 0x5f, 0x81, 0xbe, 0x91, 0xb1, 0xb3, 0x3b, 0xf0, 0x7f, 0x38, 0xf6, 0x45, 0xbb,
	0x3a, 0x9d, 0xb7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x77, 0x43, 0x04, 0x63, 0x1a, 0x03, 0x00,
	0x00,
}
