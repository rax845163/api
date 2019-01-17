// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/metric/v1alpha2/metric.proto

package v1alpha2 // import "github.com/containers-ai/api/datahub/metric/v1alpha2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha2 "github.com/containers-ai/api/datahub/resource/metadata/v1alpha2"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Metric type. A metric may be either CPU or memory.
type MetricType int32

const (
	MetricType_UNDEFINED                    MetricType = 0
	MetricType_CPU_USAGE_SECONDS_PERCENTAGE MetricType = 1
	MetricType_MEMORY_USAGE_BYTES           MetricType = 2
)

var MetricType_name = map[int32]string{
	0: "UNDEFINED",
	1: "CPU_USAGE_SECONDS_PERCENTAGE",
	2: "MEMORY_USAGE_BYTES",
}
var MetricType_value = map[string]int32{
	"UNDEFINED":                    0,
	"CPU_USAGE_SECONDS_PERCENTAGE": 1,
	"MEMORY_USAGE_BYTES":           2,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{0}
}

// *
// Represents metric data of a container
type ContainerMetric struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           map[string]*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{0}
}
func (m *ContainerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerMetric.Unmarshal(m, b)
}
func (m *ContainerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerMetric.Marshal(b, m, deterministic)
}
func (dst *ContainerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerMetric.Merge(dst, src)
}
func (m *ContainerMetric) XXX_Size() int {
	return xxx_messageInfo_ContainerMetric.Size(m)
}
func (m *ContainerMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerMetric proto.InternalMessageInfo

func (m *ContainerMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerMetric) GetMetricData() map[string]*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// *
// Represents metric data of a pod
type PodMetric struct {
	NamespacedName       *v1alpha2.NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ContainerMetrics     []*ContainerMetric       `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PodMetric) Reset()         { *m = PodMetric{} }
func (m *PodMetric) String() string { return proto.CompactTextString(m) }
func (*PodMetric) ProtoMessage()    {}
func (*PodMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{1}
}
func (m *PodMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodMetric.Unmarshal(m, b)
}
func (m *PodMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodMetric.Marshal(b, m, deterministic)
}
func (dst *PodMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodMetric.Merge(dst, src)
}
func (m *PodMetric) XXX_Size() int {
	return xxx_messageInfo_PodMetric.Size(m)
}
func (m *PodMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_PodMetric.DiscardUnknown(m)
}

var xxx_messageInfo_PodMetric proto.InternalMessageInfo

func (m *PodMetric) GetNamespacedName() *v1alpha2.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodMetric) GetContainerMetrics() []*ContainerMetric {
	if m != nil {
		return m.ContainerMetrics
	}
	return nil
}

// *
// Represents metric data of a node
type NodeMetric struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           map[string]*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NodeMetric) Reset()         { *m = NodeMetric{} }
func (m *NodeMetric) String() string { return proto.CompactTextString(m) }
func (*NodeMetric) ProtoMessage()    {}
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{2}
}
func (m *NodeMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetric.Unmarshal(m, b)
}
func (m *NodeMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetric.Marshal(b, m, deterministic)
}
func (dst *NodeMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetric.Merge(dst, src)
}
func (m *NodeMetric) XXX_Size() int {
	return xxx_messageInfo_NodeMetric.Size(m)
}
func (m *NodeMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetric.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetric proto.InternalMessageInfo

func (m *NodeMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetric) GetMetricData() map[string]*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// *
// Represents a data point of time-series metric data
type Sample struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	NumValue             string               `protobuf:"bytes,2,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{3}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (dst *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(dst, src)
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

func (m *Sample) GetNumValue() string {
	if m != nil {
		return m.NumValue
	}
	return ""
}

// *
// Represents a time range definition
//
type TimeRange struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration   `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{4}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (dst *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(dst, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

// *
// Represents a piece of metreic data
type MetricData struct {
	// data can be time series or non-time series
	Data                 []*Sample `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c133c177247e156c, []int{5}
}
func (m *MetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricData.Unmarshal(m, b)
}
func (m *MetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricData.Marshal(b, m, deterministic)
}
func (dst *MetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricData.Merge(dst, src)
}
func (m *MetricData) XXX_Size() int {
	return xxx_messageInfo_MetricData.Size(m)
}
func (m *MetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricData proto.InternalMessageInfo

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetric)(nil), "containersai.datahub.metric.v1alpha2.ContainerMetric")
	proto.RegisterMapType((map[string]*MetricData)(nil), "containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry")
	proto.RegisterType((*PodMetric)(nil), "containersai.datahub.metric.v1alpha2.PodMetric")
	proto.RegisterType((*NodeMetric)(nil), "containersai.datahub.metric.v1alpha2.NodeMetric")
	proto.RegisterMapType((map[string]*MetricData)(nil), "containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry")
	proto.RegisterType((*Sample)(nil), "containersai.datahub.metric.v1alpha2.Sample")
	proto.RegisterType((*TimeRange)(nil), "containersai.datahub.metric.v1alpha2.TimeRange")
	proto.RegisterType((*MetricData)(nil), "containersai.datahub.metric.v1alpha2.MetricData")
	proto.RegisterEnum("containersai.datahub.metric.v1alpha2.MetricType", MetricType_name, MetricType_value)
}

func init() {
	proto.RegisterFile("datahub/metric/v1alpha2/metric.proto", fileDescriptor_metric_c133c177247e156c)
}

var fileDescriptor_metric_c133c177247e156c = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0xcd, 0xaf, 0xd4, 0x13, 0x41, 0xc3, 0x1e, 0x50, 0x08, 0x08, 0xa2, 0xa8, 0x87,
	0x08, 0x51, 0x9b, 0x06, 0x8a, 0x80, 0x0b, 0x6d, 0x93, 0x6d, 0xc5, 0x21, 0x6e, 0xb4, 0x49, 0x90,
	0xca, 0xc5, 0xda, 0xd8, 0xdb, 0xc4, 0x22, 0xde, 0xb5, 0xec, 0x75, 0xa5, 0x7c, 0x27, 0xbe, 0x10,
	0xdf, 0x83, 0x33, 0x42, 0xbb, 0xfe, 0x13, 0x25, 0x14, 0x88, 0x38, 0x71, 0xf2, 0xee, 0x78, 0xde,
	0x9b, 0x37, 0x6f, 0x34, 0x0b, 0x07, 0x3e, 0x95, 0x74, 0x9e, 0x4e, 0xed, 0x90, 0xc9, 0x38, 0xf0,
	0xec, 0x9b, 0x23, 0xba, 0x88, 0xe6, 0xb4, 0x9b, 0xdf, 0xad, 0x28, 0x16, 0x52, 0xa0, 0x03, 0x4f,
	0x70, 0x49, 0x03, 0xce, 0xe2, 0x84, 0x06, 0x56, 0x0e, 0xb1, 0xf2, 0x94, 0x02, 0xd2, 0x7c, 0x3a,
	0x13, 0x62, 0xb6, 0x60, 0xb6, 0xc6, 0x4c, 0xd3, 0x6b, 0x5b, 0x06, 0x21, 0x4b, 0x24, 0x0d, 0xa3,
	0x8c, 0xa6, 0xf9, 0x64, 0x33, 0xc1, 0x4f, 0x63, 0x2a, 0x03, 0xc1, 0xf3, 0xff, 0x47, 0x85, 0x98,
	0x98, 0x25, 0x22, 0x8d, 0x3d, 0xa6, 0x54, 0x50, 0x15, 0x5c, 0xd3, 0xa5, 0x23, 0x19, 0xa4, 0xfd,
	0xdd, 0x80, 0xfd, 0x5e, 0x21, 0x6e, 0xa0, 0x05, 0x21, 0x04, 0x55, 0x4e, 0x43, 0xd6, 0x30, 0x5a,
	0x46, 0xc7, 0x24, 0xfa, 0x8c, 0xae, 0xa1, 0x96, 0xc9, 0x75, 0x15, 0xb8, 0x51, 0x69, 0xed, 0x74,
	0x6a, 0x5d, 0x6c, 0x6d, 0xd3, 0x97, 0xb5, 0xc1, 0x6f, 0x65, 0x9f, 0x3e, 0x95, 0x14, 0x73, 0x19,
	0x2f, 0x09, 0x84, 0x65, 0xa0, 0x29, 0x60, 0x7f, 0xe3, 0x37, 0xaa, 0xc3, 0xce, 0x67, 0xb6, 0xcc,
	0xd5, 0xa8, 0x23, 0x3a, 0x87, 0xff, 0x6f, 0xe8, 0x22, 0x65, 0x8d, 0x4a, 0xcb, 0xe8, 0xd4, 0xba,
	0x2f, 0xb6, 0x93, 0xb1, 0xe2, 0x25, 0x19, 0xfc, 0x5d, 0xe5, 0x8d, 0xd1, 0xfe, 0x6a, 0x80, 0x39,
	0x14, 0x7e, 0xde, 0xfa, 0x1c, 0xf6, 0x55, 0xbb, 0x49, 0x44, 0x3d, 0xe6, 0xbb, 0xa5, 0x0b, 0xb5,
	0xee, 0xfb, 0xdb, 0x6b, 0x14, 0x46, 0x5b, 0xa5, 0xad, 0x65, 0x39, 0xa7, 0xe4, 0x51, 0x27, 0x72,
	0x8f, 0xaf, 0xdd, 0xd1, 0x14, 0xee, 0x97, 0x8c, 0x6e, 0x26, 0x35, 0xc9, 0x6d, 0x3d, 0xfe, 0x2b,
	0x5b, 0x49, 0xdd, 0x5b, 0x0f, 0x24, 0xed, 0x6f, 0x06, 0x80, 0x23, 0x7c, 0xf6, 0x9b, 0xb9, 0xd2,
	0xdb, 0xe6, 0x7a, 0xb2, 0x9d, 0x80, 0x15, 0xf5, 0xbf, 0x35, 0xd2, 0x09, 0xec, 0x8e, 0x68, 0x18,
	0x2d, 0x18, 0xb2, 0xa0, 0xaa, 0x76, 0x28, 0x9f, 0x61, 0xd3, 0xca, 0xf6, 0xc7, 0x2a, 0xf6, 0xc7,
	0x1a, 0x17, 0x0b, 0x46, 0x74, 0x1e, 0x7a, 0x04, 0x26, 0x4f, 0x43, 0x77, 0xa5, 0xc4, 0x24, 0x7b,
	0x3c, 0x0d, 0x3f, 0xaa, 0x7b, 0xfb, 0x8b, 0x01, 0xa6, 0x02, 0x10, 0xca, 0x67, 0x0c, 0xbd, 0x05,
	0x48, 0x24, 0x8d, 0xa5, 0xbb, 0x65, 0x01, 0x53, 0x67, 0xab, 0x3b, 0x3a, 0x86, 0x3d, 0xc6, 0xfd,
	0x0c, 0x58, 0xf9, 0x23, 0xf0, 0x0e, 0xe3, 0xbe, 0x86, 0x1d, 0x42, 0x35, 0x91, 0x2c, 0x6a, 0xec,
	0x68, 0xc8, 0xc3, 0x9f, 0x20, 0xfd, 0xfc, 0x31, 0x20, 0x3a, 0xad, 0xed, 0x00, 0xac, 0xec, 0x41,
	0x27, 0x50, 0xd5, 0x03, 0x36, 0xf4, 0x80, 0x9f, 0x6f, 0x67, 0x6f, 0xe6, 0x22, 0xd1, 0xc8, 0x67,
	0x93, 0x82, 0x6f, 0xbc, 0x8c, 0x18, 0xba, 0x0b, 0xe6, 0xc4, 0xe9, 0xe3, 0xf3, 0x0f, 0x0e, 0xee,
	0xd7, 0xff, 0x43, 0x2d, 0x78, 0xdc, 0x1b, 0x4e, 0xdc, 0xc9, 0xe8, 0xf4, 0x02, 0xbb, 0x23, 0xdc,
	0xbb, 0x74, 0xfa, 0x23, 0x77, 0x88, 0x49, 0x0f, 0x3b, 0xe3, 0xd3, 0x0b, 0x5c, 0x37, 0xd0, 0x03,
	0x40, 0x03, 0x3c, 0xb8, 0x24, 0x57, 0x79, 0xd2, 0xd9, 0xd5, 0x18, 0x8f, 0xea, 0x95, 0xb3, 0xd7,
	0x9f, 0x5e, 0xcd, 0x02, 0xa9, 0xaa, 0x7b, 0x22, 0xb4, 0x57, 0xb2, 0x0e, 0x69, 0x60, 0xd3, 0x28,
	0xb0, 0x7f, 0xf1, 0xbe, 0x4e, 0x77, 0x75, 0xdf, 0x2f, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0xaf, 0x02, 0xab, 0x81, 0x05, 0x00, 0x00,
}
