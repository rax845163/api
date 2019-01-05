// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/metric.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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
	return fileDescriptor_metric_c7e7a966514b8b14, []int{0}
}

type StrOp int32

const (
	StrOp_EQUAL     StrOp = 0
	StrOp_NOT_EQUAL StrOp = 1
)

var StrOp_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
}
var StrOp_value = map[string]int32{
	"EQUAL":     0,
	"NOT_EQUAL": 1,
}

func (x StrOp) String() string {
	return proto.EnumName(StrOp_name, int32(x))
}
func (StrOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{1}
}

// *
// Represents metric data of a container
type ContainerMetric struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           []*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{0}
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

func (m *ContainerMetric) GetMetricData() []*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// *
// Represents metric data of a pod
type PodMetric struct {
	NamespacedName       *NamespacedName    `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ContainerMetrics     []*ContainerMetric `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PodMetric) Reset()         { *m = PodMetric{} }
func (m *PodMetric) String() string { return proto.CompactTextString(m) }
func (*PodMetric) ProtoMessage()    {}
func (*PodMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{1}
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

func (m *PodMetric) GetNamespacedName() *NamespacedName {
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
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           []*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeMetric) Reset()         { *m = NodeMetric{} }
func (m *NodeMetric) String() string { return proto.CompactTextString(m) }
func (*NodeMetric) ProtoMessage()    {}
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{2}
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

func (m *NodeMetric) GetMetricData() []*MetricData {
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
	return fileDescriptor_metric_c7e7a966514b8b14, []int{3}
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

type MetricResult struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Samples              []*Sample         `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricResult) Reset()         { *m = MetricResult{} }
func (m *MetricResult) String() string { return proto.CompactTextString(m) }
func (*MetricResult) ProtoMessage()    {}
func (*MetricResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{4}
}
func (m *MetricResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricResult.Unmarshal(m, b)
}
func (m *MetricResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricResult.Marshal(b, m, deterministic)
}
func (dst *MetricResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricResult.Merge(dst, src)
}
func (m *MetricResult) XXX_Size() int {
	return xxx_messageInfo_MetricResult.Size(m)
}
func (m *MetricResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricResult.DiscardUnknown(m)
}

var xxx_messageInfo_MetricResult proto.InternalMessageInfo

func (m *MetricResult) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricResult) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

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
	return fileDescriptor_metric_c7e7a966514b8b14, []int{5}
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

type LabelSelector struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op                   StrOp    `protobuf:"varint,2,opt,name=op,proto3,enum=containers_ai.alameda.v1alpha1.datahub.StrOp" json:"op,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{6}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelector.Unmarshal(m, b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
}
func (dst *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(dst, src)
}
func (m *LabelSelector) XXX_Size() int {
	return xxx_messageInfo_LabelSelector.Size(m)
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelector) GetOp() StrOp {
	if m != nil {
		return m.Op
	}
	return StrOp_EQUAL
}

func (m *LabelSelector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// *
// Represents a piece of metreic data
type MetricData struct {
	MetricType MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containers_ai.alameda.v1alpha1.datahub.MetricType" json:"metric_type,omitempty"`
	// data can be time series or non-time series
	Data                 []*Sample `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_c7e7a966514b8b14, []int{7}
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

func (m *MetricData) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_UNDEFINED
}

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerMetric")
	proto.RegisterType((*PodMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.PodMetric")
	proto.RegisterType((*NodeMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.NodeMetric")
	proto.RegisterType((*Sample)(nil), "containers_ai.alameda.v1alpha1.datahub.Sample")
	proto.RegisterType((*MetricResult)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricResult")
	proto.RegisterMapType((map[string]string)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry")
	proto.RegisterType((*TimeRange)(nil), "containers_ai.alameda.v1alpha1.datahub.TimeRange")
	proto.RegisterType((*LabelSelector)(nil), "containers_ai.alameda.v1alpha1.datahub.LabelSelector")
	proto.RegisterType((*MetricData)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricData")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.MetricType", MetricType_name, MetricType_value)
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.StrOp", StrOp_name, StrOp_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/metric.proto", fileDescriptor_metric_c7e7a966514b8b14)
}

var fileDescriptor_metric_c7e7a966514b8b14 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xf4, 0xe2, 0x09, 0x6d, 0xc3, 0x0a, 0xa1, 0x10, 0x10, 0x14, 0x3f, 0xa0, 0x52,
	0xd4, 0x8d, 0x1a, 0xc4, 0xa5, 0x48, 0x48, 0xb4, 0x89, 0x29, 0x48, 0xad, 0x53, 0xd6, 0x09, 0xa2,
	0x4f, 0xab, 0x4d, 0xbc, 0xb4, 0x11, 0xbe, 0xac, 0xec, 0x75, 0xa5, 0xf0, 0x3f, 0xbc, 0xf1, 0x35,
	0xfc, 0x07, 0xff, 0x80, 0xbc, 0x6b, 0xa7, 0x69, 0x41, 0x6a, 0xfa, 0xc2, 0xdb, 0x4c, 0x32, 0x67,
	0xe6, 0x9c, 0xb3, 0xc7, 0xf0, 0x94, 0x05, 0x2c, 0xe4, 0x3e, 0xa3, 0x4c, 0x8c, 0x5b, 0xe7, 0x3b,
	0x2c, 0x10, 0x67, 0x6c, 0xa7, 0xe5, 0x33, 0xc9, 0xce, 0xb2, 0x61, 0x2b, 0xe4, 0x32, 0x19, 0x8f,
	0xb0, 0x48, 0x62, 0x19, 0xa3, 0x27, 0xa3, 0x38, 0x92, 0x6c, 0x1c, 0xf1, 0x24, 0xa5, 0x6c, 0x8c,
	0x0b, 0x20, 0x2e, 0x41, 0xb8, 0x00, 0x35, 0x1f, 0x9d, 0xc6, 0xf1, 0x69, 0xc0, 0x5b, 0x0a, 0x35,
	0xcc, 0xbe, 0xb6, 0xe4, 0x38, 0xe4, 0xa9, 0x64, 0xa1, 0xd0, 0x8b, 0x9a, 0x0f, 0xaf, 0x0e, 0xf8,
	0x59, 0xc2, 0xe4, 0x38, 0x8e, 0x8a, 0xff, 0x9f, 0x5d, 0xc7, 0x89, 0xe5, 0xb5, 0x1e, 0xb6, 0xbf,
	0xc3, 0x7a, 0xa7, 0xe4, 0x75, 0xa4, 0xe8, 0x22, 0x04, 0xd5, 0x88, 0x85, 0xbc, 0x61, 0x6c, 0x18,
	0x9b, 0x16, 0x51, 0x35, 0xf2, 0xa0, 0xa6, 0xc5, 0xd0, 0x1c, 0xdb, 0x30, 0x37, 0x2a, 0x9b, 0xb5,
	0x76, 0x1b, 0xcf, 0x27, 0x09, 0xeb, 0xc5, 0x5d, 0x26, 0x19, 0x81, 0x70, 0x5a, 0xdb, 0xbf, 0x0c,
	0xb0, 0x8e, 0x63, 0xbf, 0x38, 0x4b, 0x61, 0x3d, 0x3f, 0x95, 0x0a, 0x36, 0xe2, 0x3e, 0x9d, 0x32,
	0xa8, 0xb5, 0x5f, 0xce, 0x7b, 0xc6, 0x9d, 0xc2, 0xf3, 0x8a, 0xac, 0x45, 0x97, 0x7a, 0xe4, 0xc3,
	0xed, 0xe9, 0x22, 0xaa, 0x69, 0xa4, 0x85, 0x92, 0x57, 0xf3, 0x9e, 0xb8, 0xe2, 0x15, 0xa9, 0x8f,
	0x2e, 0xff, 0x90, 0xda, 0x19, 0x80, 0x1b, 0xfb, 0xfc, 0x7f, 0x7b, 0x39, 0x80, 0x25, 0x8f, 0x85,
	0x22, 0xe0, 0x08, 0x43, 0x35, 0x4f, 0x4c, 0x61, 0x5e, 0x13, 0xeb, 0xb4, 0xe0, 0x32, 0x2d, 0xb8,
	0x5f, 0xc6, 0x89, 0xa8, 0x39, 0x74, 0x1f, 0xac, 0x28, 0x0b, 0xe9, 0x39, 0x0b, 0x32, 0xde, 0x30,
	0x15, 0xcf, 0x95, 0x28, 0x0b, 0x3f, 0xe7, 0xbd, 0xfd, 0xdb, 0x80, 0x5b, 0x85, 0x54, 0x9e, 0x66,
	0x81, 0x44, 0x5f, 0x60, 0x29, 0x60, 0x43, 0x1e, 0xa4, 0x0d, 0x43, 0xf1, 0x7e, 0x77, 0x33, 0xde,
	0x7a, 0x0b, 0x3e, 0x54, 0x2b, 0x9c, 0x48, 0x26, 0x13, 0x52, 0xec, 0x43, 0x1f, 0x60, 0x39, 0x55,
	0x0a, 0xca, 0x47, 0xc1, 0xf3, 0xae, 0xd6, 0xc2, 0x49, 0x09, 0x6f, 0xee, 0x42, 0x6d, 0xe6, 0x00,
	0xaa, 0x43, 0xe5, 0x1b, 0x9f, 0x14, 0x4f, 0x90, 0x97, 0xe8, 0x0e, 0x2c, 0xce, 0xca, 0xd5, 0xcd,
	0x1b, 0xf3, 0xb5, 0x61, 0xff, 0x34, 0xc0, 0xca, 0x0d, 0x22, 0x2c, 0x3a, 0xe5, 0x68, 0x17, 0x20,
	0x95, 0x2c, 0x91, 0x74, 0x4e, 0x43, 0x2d, 0x35, 0x9d, 0xf7, 0xe8, 0x05, 0xac, 0xf0, 0xc8, 0xd7,
	0x40, 0xf3, 0x5a, 0xe0, 0x32, 0x8f, 0x7c, 0x05, 0xdb, 0x86, 0x6a, 0x2a, 0xb9, 0x68, 0x54, 0x14,
	0xe4, 0xde, 0x5f, 0x90, 0x6e, 0xf1, 0xa9, 0x13, 0x35, 0x66, 0x9f, 0xc3, 0xaa, 0x52, 0xea, 0xf1,
	0x80, 0x8f, 0x64, 0x9c, 0xfc, 0x43, 0xeb, 0x5b, 0x30, 0x63, 0xa1, 0x28, 0xac, 0xb5, 0xb7, 0xe7,
	0x76, 0x54, 0x26, 0x3d, 0x41, 0xcc, 0x58, 0x5c, 0x58, 0x55, 0x99, 0xb1, 0xca, 0xfe, 0x61, 0x00,
	0x5c, 0x04, 0x71, 0x26, 0xd1, 0x72, 0x22, 0xb4, 0x51, 0x6b, 0x37, 0x4d, 0x74, 0x7f, 0x22, 0x78,
	0x99, 0xe8, 0xbc, 0x46, 0xfb, 0x50, 0x9d, 0xf9, 0x3e, 0x6e, 0x1a, 0x06, 0x85, 0xdd, 0x1a, 0x94,
	0x34, 0xd5, 0xc6, 0x55, 0xb0, 0x06, 0x6e, 0xd7, 0x79, 0xff, 0xd1, 0x75, 0xba, 0xf5, 0x05, 0xb4,
	0x01, 0x0f, 0x3a, 0xc7, 0x03, 0x3a, 0xf0, 0xf6, 0x0e, 0x1c, 0xea, 0x39, 0x9d, 0x9e, 0xdb, 0xf5,
	0xe8, 0xb1, 0x43, 0x3a, 0x8e, 0xdb, 0xdf, 0x3b, 0x70, 0xea, 0x06, 0xba, 0x0b, 0xe8, 0xc8, 0x39,
	0xea, 0x91, 0x93, 0x62, 0x68, 0xff, 0xa4, 0xef, 0x78, 0x75, 0x73, 0xeb, 0x31, 0x2c, 0x2a, 0x87,
	0x90, 0x05, 0x8b, 0xce, 0xa7, 0xc1, 0xde, 0x61, 0x7d, 0x21, 0x5f, 0xee, 0xf6, 0xfa, 0x54, 0xb7,
	0xc6, 0x70, 0x49, 0x3d, 0xd9, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xd4, 0xa5, 0x0a,
	0x21, 0x06, 0x00, 0x00,
}
