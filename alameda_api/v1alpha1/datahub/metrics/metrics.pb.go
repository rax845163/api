// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/metrics/metrics.proto

package metrics

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
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

// Represents metric data of a container
type ContainerMetric struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           []*common.MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{0}
}

func (m *ContainerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerMetric.Unmarshal(m, b)
}
func (m *ContainerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerMetric.Marshal(b, m, deterministic)
}
func (m *ContainerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerMetric.Merge(m, src)
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

func (m *ContainerMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// Represents metric data of a pod
type PodMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	ContainerMetrics     []*ContainerMetric    `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PodMetric) Reset()         { *m = PodMetric{} }
func (m *PodMetric) String() string { return proto.CompactTextString(m) }
func (*PodMetric) ProtoMessage()    {}
func (*PodMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{1}
}

func (m *PodMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodMetric.Unmarshal(m, b)
}
func (m *PodMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodMetric.Marshal(b, m, deterministic)
}
func (m *PodMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodMetric.Merge(m, src)
}
func (m *PodMetric) XXX_Size() int {
	return xxx_messageInfo_PodMetric.Size(m)
}
func (m *PodMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_PodMetric.DiscardUnknown(m)
}

var xxx_messageInfo_PodMetric proto.InternalMessageInfo

func (m *PodMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *PodMetric) GetContainerMetrics() []*ContainerMetric {
	if m != nil {
		return m.ContainerMetrics
	}
	return nil
}

type ControllerMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                 resources.Kind        `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	MetricData           []*common.MetricData  `protobuf:"bytes,3,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ControllerMetric) Reset()         { *m = ControllerMetric{} }
func (m *ControllerMetric) String() string { return proto.CompactTextString(m) }
func (*ControllerMetric) ProtoMessage()    {}
func (*ControllerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{2}
}

func (m *ControllerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerMetric.Unmarshal(m, b)
}
func (m *ControllerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerMetric.Marshal(b, m, deterministic)
}
func (m *ControllerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerMetric.Merge(m, src)
}
func (m *ControllerMetric) XXX_Size() int {
	return xxx_messageInfo_ControllerMetric.Size(m)
}
func (m *ControllerMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerMetric proto.InternalMessageInfo

func (m *ControllerMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ControllerMetric) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (m *ControllerMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

type ApplicationMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData           []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApplicationMetric) Reset()         { *m = ApplicationMetric{} }
func (m *ApplicationMetric) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetric) ProtoMessage()    {}
func (*ApplicationMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{3}
}

func (m *ApplicationMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMetric.Unmarshal(m, b)
}
func (m *ApplicationMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMetric.Marshal(b, m, deterministic)
}
func (m *ApplicationMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMetric.Merge(m, src)
}
func (m *ApplicationMetric) XXX_Size() int {
	return xxx_messageInfo_ApplicationMetric.Size(m)
}
func (m *ApplicationMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMetric proto.InternalMessageInfo

func (m *ApplicationMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ApplicationMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

type NamespaceMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData           []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NamespaceMetric) Reset()         { *m = NamespaceMetric{} }
func (m *NamespaceMetric) String() string { return proto.CompactTextString(m) }
func (*NamespaceMetric) ProtoMessage()    {}
func (*NamespaceMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{4}
}

func (m *NamespaceMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceMetric.Unmarshal(m, b)
}
func (m *NamespaceMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceMetric.Marshal(b, m, deterministic)
}
func (m *NamespaceMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceMetric.Merge(m, src)
}
func (m *NamespaceMetric) XXX_Size() int {
	return xxx_messageInfo_NamespaceMetric.Size(m)
}
func (m *NamespaceMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceMetric.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceMetric proto.InternalMessageInfo

func (m *NamespaceMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NamespaceMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// Represents metric data of a node
type NodeMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData           []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NodeMetric) Reset()         { *m = NodeMetric{} }
func (m *NodeMetric) String() string { return proto.CompactTextString(m) }
func (*NodeMetric) ProtoMessage()    {}
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{5}
}

func (m *NodeMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetric.Unmarshal(m, b)
}
func (m *NodeMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetric.Marshal(b, m, deterministic)
}
func (m *NodeMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetric.Merge(m, src)
}
func (m *NodeMetric) XXX_Size() int {
	return xxx_messageInfo_NodeMetric.Size(m)
}
func (m *NodeMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetric.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetric proto.InternalMessageInfo

func (m *NodeMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NodeMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

type ClusterMetric struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData           []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterMetric) Reset()         { *m = ClusterMetric{} }
func (m *ClusterMetric) String() string { return proto.CompactTextString(m) }
func (*ClusterMetric) ProtoMessage()    {}
func (*ClusterMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb3a291d3897af, []int{6}
}

func (m *ClusterMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterMetric.Unmarshal(m, b)
}
func (m *ClusterMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterMetric.Marshal(b, m, deterministic)
}
func (m *ClusterMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMetric.Merge(m, src)
}
func (m *ClusterMetric) XXX_Size() int {
	return xxx_messageInfo_ClusterMetric.Size(m)
}
func (m *ClusterMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMetric proto.InternalMessageInfo

func (m *ClusterMetric) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ClusterMetric) GetMetricData() []*common.MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric")
	proto.RegisterType((*PodMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.PodMetric")
	proto.RegisterType((*ControllerMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric")
	proto.RegisterType((*ApplicationMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric")
	proto.RegisterType((*NamespaceMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric")
	proto.RegisterType((*NodeMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.NodeMetric")
	proto.RegisterType((*ClusterMetric)(nil), "containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/metrics/metrics.proto", fileDescriptor_2dbb3a291d3897af)
}

var fileDescriptor_2dbb3a291d3897af = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcf, 0x6a, 0xf2, 0x40,
	0x10, 0x67, 0x55, 0x3e, 0x70, 0xc2, 0x57, 0x35, 0x27, 0xf1, 0x24, 0x9e, 0xbc, 0xb8, 0xc1, 0x14,
	0xa1, 0x50, 0x28, 0x58, 0x7b, 0x29, 0x45, 0x5b, 0x72, 0x6b, 0x29, 0x84, 0x71, 0xb3, 0xd4, 0xad,
	0x49, 0x36, 0x24, 0x6b, 0xcf, 0x7d, 0x8a, 0xbe, 0x51, 0xa1, 0x94, 0x9e, 0xfb, 0x3c, 0x25, 0xc9,
	0x9a, 0xa2, 0x07, 0x51, 0xa8, 0x82, 0xa7, 0x0c, 0xcb, 0xfe, 0xfe, 0xcd, 0x24, 0x13, 0xb0, 0xd1,
	0xc7, 0x80, 0x7b, 0xe8, 0x62, 0x24, 0xac, 0x97, 0x3e, 0xfa, 0xd1, 0x0c, 0xfb, 0x96, 0x87, 0x0a,
	0x67, 0x8b, 0xa9, 0x15, 0x70, 0x15, 0x0b, 0x96, 0x2c, 0x9f, 0x34, 0x8a, 0xa5, 0x92, 0x66, 0x8f,
	0xc9, 0x50, 0xa1, 0x08, 0x79, 0x9c, 0xa0, 0xa0, 0x9a, 0x80, 0x2e, 0xc1, 0x54, 0x83, 0xa9, 0x06,
	0xb5, 0xfa, 0x1b, 0x25, 0x98, 0x0c, 0x02, 0x19, 0xae, 0x2a, 0xb4, 0x06, 0x1b, 0x21, 0x31, 0x4f,
	0xe4, 0x22, 0x66, 0x3c, 0xf3, 0x85, 0xe9, 0x69, 0x0e, 0xeb, 0xbc, 0x12, 0xa8, 0x8d, 0x96, 0xde,
	0xc6, 0x19, 0xa3, 0x69, 0x42, 0x25, 0xc4, 0x80, 0x37, 0x49, 0x9b, 0x74, 0xab, 0x4e, 0x56, 0x9b,
	0xf7, 0x60, 0xe4, 0x7a, 0x6e, 0x0a, 0x6e, 0x96, 0xda, 0xe5, 0xae, 0x61, 0x9f, 0xd1, 0xed, 0x62,
	0xe5, 0x86, 0x69, 0x4e, 0x7f, 0x85, 0x0a, 0x1d, 0x08, 0x8a, 0xba, 0xf3, 0x4d, 0xa0, 0x7a, 0x27,
	0x3d, 0x2d, 0xfe, 0x08, 0x86, 0x9c, 0x3e, 0x73, 0xa6, 0xdc, 0xd4, 0x69, 0xe6, 0xc1, 0xb0, 0xcf,
	0xb7, 0x14, 0x2a, 0x62, 0xd2, 0xdb, 0x8c, 0x63, 0xcc, 0x53, 0x2d, 0x59, 0xd4, 0xe6, 0x1c, 0x1a,
	0x05, 0x93, 0xab, 0x1b, 0xa8, 0xc3, 0x5c, 0xd0, 0x9d, 0x66, 0x44, 0xd7, 0xba, 0xe6, 0xd4, 0xd9,
	0xea, 0x41, 0xd2, 0x79, 0x2b, 0x41, 0x3d, 0xbd, 0x15, 0x4b, 0xdf, 0x2f, 0x9a, 0xbb, 0xdf, 0x7c,
	0xd7, 0x50, 0x99, 0x8b, 0xd0, 0x6b, 0x96, 0xda, 0xa4, 0x7b, 0x62, 0x0f, 0x76, 0xa6, 0xbd, 0x11,
	0xa1, 0xe7, 0x64, 0x14, 0xeb, 0x13, 0x2f, 0xff, 0xe1, 0xc4, 0xbf, 0x08, 0x34, 0x86, 0x51, 0xe4,
	0x0b, 0x86, 0x4a, 0xc8, 0xf0, 0x20, 0x9d, 0xd9, 0xe3, 0x0b, 0xfc, 0x49, 0xa0, 0x36, 0xc1, 0x80,
	0x27, 0x11, 0x32, 0x7e, 0xec, 0x61, 0xde, 0x09, 0xc0, 0x44, 0x7a, 0x47, 0x9f, 0xe3, 0x83, 0xc0,
	0xff, 0x91, 0xbf, 0x48, 0xd4, 0x81, 0xbe, 0xbc, 0xfd, 0x45, 0xb9, 0x1c, 0x3d, 0x0c, 0x9f, 0x84,
	0xd2, 0xf7, 0xac, 0x5f, 0xc6, 0x1e, 0x0a, 0x2b, 0xdd, 0xf6, 0xdb, 0xfc, 0x8f, 0xa6, 0xff, 0xb2,
	0x7d, 0x7f, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x12, 0xfc, 0x41, 0x7c, 0xbe, 0x06, 0x00, 0x00,
}