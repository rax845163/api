// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/recommendation/v1alpha2/recommendation.proto

package v1alpha2 // import "github.com/containers-ai/api/datahub/recommendation/v1alpha2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha2 "github.com/containers-ai/api/datahub/metric/v1alpha2"
import v1alpha22 "github.com/containers-ai/api/datahub/resource/metadata/v1alpha2"
import v1alpha21 "github.com/containers-ai/api/datahub/resource/pod/assign/v1alpha2"
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
//  Recommendation policy. A policy may be either stable or compact.
type RecommendationPolicy int32

const (
	RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED RecommendationPolicy = 0
	RecommendationPolicy_STABLE                         RecommendationPolicy = 1
	RecommendationPolicy_COMPACT                        RecommendationPolicy = 2
)

var RecommendationPolicy_name = map[int32]string{
	0: "RECOMMENDATIONPOLICY_UNDEFINED",
	1: "STABLE",
	2: "COMPACT",
}
var RecommendationPolicy_value = map[string]int32{
	"RECOMMENDATIONPOLICY_UNDEFINED": 0,
	"STABLE":  1,
	"COMPACT": 2,
}

func (x RecommendationPolicy) String() string {
	return proto.EnumName(RecommendationPolicy_name, int32(x))
}
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_e67305e6a7503f64, []int{0}
}

// *
// Represents a resource configuration recommendation
//
// It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial stage
//
type ContainerRecommendation struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// use containersai.datahub.metric.v1alpha2.MetricType as key
	LimitRecommendations map[int32]*v1alpha2.MetricData `protobuf:"bytes,2,rep,name=limit_recommendations,json=limitRecommendations,proto3" json:"limit_recommendations,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// use containersai.datahub.metric.v1alpha2.MetricType as key
	RequestRecommendations map[int32]*v1alpha2.MetricData `protobuf:"bytes,3,rep,name=request_recommendations,json=requestRecommendations,proto3" json:"request_recommendations,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// use containersai.datahub.metric.v1alpha2.MetricType as key
	InitialLimitRecommendations map[int32]*v1alpha2.MetricData `protobuf:"bytes,4,rep,name=initial_limit_recommendations,json=initialLimitRecommendations,proto3" json:"initial_limit_recommendations,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// use containersai.datahub.metric.v1alpha2.MetricType as key
	InitialRequestRecommendations map[int32]*v1alpha2.MetricData `protobuf:"bytes,5,rep,name=initial_request_recommendations,json=initialRequestRecommendations,proto3" json:"initial_request_recommendations,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral          struct{}                       `json:"-"`
	XXX_unrecognized              []byte                         `json:"-"`
	XXX_sizecache                 int32                          `json:"-"`
}

func (m *ContainerRecommendation) Reset()         { *m = ContainerRecommendation{} }
func (m *ContainerRecommendation) String() string { return proto.CompactTextString(m) }
func (*ContainerRecommendation) ProtoMessage()    {}
func (*ContainerRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_e67305e6a7503f64, []int{0}
}
func (m *ContainerRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRecommendation.Unmarshal(m, b)
}
func (m *ContainerRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRecommendation.Marshal(b, m, deterministic)
}
func (dst *ContainerRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRecommendation.Merge(dst, src)
}
func (m *ContainerRecommendation) XXX_Size() int {
	return xxx_messageInfo_ContainerRecommendation.Size(m)
}
func (m *ContainerRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRecommendation proto.InternalMessageInfo

func (m *ContainerRecommendation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRecommendation) GetLimitRecommendations() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.LimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetRequestRecommendations() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.RequestRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialLimitRecommendations() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.InitialLimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialRequestRecommendations() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.InitialRequestRecommendations
	}
	return nil
}

// *
// Represents a recommended pod-to-node assignment (i.e. pod placement)
//
type AssignPodPolicy struct {
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are valid to be assigned to Policy:
	//	*AssignPodPolicy_NodePriority
	//	*AssignPodPolicy_NodeSelector
	//	*AssignPodPolicy_NodeName
	Policy               isAssignPodPolicy_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AssignPodPolicy) Reset()         { *m = AssignPodPolicy{} }
func (m *AssignPodPolicy) String() string { return proto.CompactTextString(m) }
func (*AssignPodPolicy) ProtoMessage()    {}
func (*AssignPodPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_e67305e6a7503f64, []int{1}
}
func (m *AssignPodPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPodPolicy.Unmarshal(m, b)
}
func (m *AssignPodPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPodPolicy.Marshal(b, m, deterministic)
}
func (dst *AssignPodPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPodPolicy.Merge(dst, src)
}
func (m *AssignPodPolicy) XXX_Size() int {
	return xxx_messageInfo_AssignPodPolicy.Size(m)
}
func (m *AssignPodPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPodPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPodPolicy proto.InternalMessageInfo

func (m *AssignPodPolicy) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type isAssignPodPolicy_Policy interface {
	isAssignPodPolicy_Policy()
}

type AssignPodPolicy_NodePriority struct {
	NodePriority *v1alpha21.NodePriority `protobuf:"bytes,2,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type AssignPodPolicy_NodeSelector struct {
	NodeSelector *v1alpha21.Selector `protobuf:"bytes,3,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
}

type AssignPodPolicy_NodeName struct {
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3,oneof"`
}

func (*AssignPodPolicy_NodePriority) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeSelector) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeName) isAssignPodPolicy_Policy() {}

func (m *AssignPodPolicy) GetPolicy() isAssignPodPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *AssignPodPolicy) GetNodePriority() *v1alpha21.NodePriority {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeSelector() *v1alpha21.Selector {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeName() string {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeName); ok {
		return x.NodeName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AssignPodPolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AssignPodPolicy_OneofMarshaler, _AssignPodPolicy_OneofUnmarshaler, _AssignPodPolicy_OneofSizer, []interface{}{
		(*AssignPodPolicy_NodePriority)(nil),
		(*AssignPodPolicy_NodeSelector)(nil),
		(*AssignPodPolicy_NodeName)(nil),
	}
}

func _AssignPodPolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodePriority); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeSelector:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeSelector); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NodeName)
	case nil:
	default:
		return fmt.Errorf("AssignPodPolicy.Policy has unexpected type %T", x)
	}
	return nil
}

func _AssignPodPolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AssignPodPolicy)
	switch tag {
	case 2: // policy.node_priority
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1alpha21.NodePriority)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodePriority{msg}
		return true, err
	case 3: // policy.node_selector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1alpha21.Selector)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodeSelector{msg}
		return true, err
	case 4: // policy.node_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Policy = &AssignPodPolicy_NodeName{x}
		return true, err
	default:
		return false, nil
	}
}

func _AssignPodPolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		s := proto.Size(x.NodePriority)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeSelector:
		s := proto.Size(x.NodeSelector)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NodeName)))
		n += len(x.NodeName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
// Represents a set of container resource configuration recommenations of a pod
//
type PodRecommendation struct {
	NamespacedName           *v1alpha22.NamespacedName  `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ApplyRecommendationNow   bool                       `protobuf:"varint,2,opt,name=apply_recommendation_now,json=applyRecommendationNow,proto3" json:"apply_recommendation_now,omitempty"`
	AssignPodPolicy          *AssignPodPolicy           `protobuf:"bytes,3,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	ContainerRecommendations []*ContainerRecommendation `protobuf:"bytes,4,rep,name=container_recommendations,json=containerRecommendations,proto3" json:"container_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *PodRecommendation) Reset()         { *m = PodRecommendation{} }
func (m *PodRecommendation) String() string { return proto.CompactTextString(m) }
func (*PodRecommendation) ProtoMessage()    {}
func (*PodRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_e67305e6a7503f64, []int{2}
}
func (m *PodRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodRecommendation.Unmarshal(m, b)
}
func (m *PodRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodRecommendation.Marshal(b, m, deterministic)
}
func (dst *PodRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodRecommendation.Merge(dst, src)
}
func (m *PodRecommendation) XXX_Size() int {
	return xxx_messageInfo_PodRecommendation.Size(m)
}
func (m *PodRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_PodRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_PodRecommendation proto.InternalMessageInfo

func (m *PodRecommendation) GetNamespacedName() *v1alpha22.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodRecommendation) GetApplyRecommendationNow() bool {
	if m != nil {
		return m.ApplyRecommendationNow
	}
	return false
}

func (m *PodRecommendation) GetAssignPodPolicy() *AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodRecommendation) GetContainerRecommendations() []*ContainerRecommendation {
	if m != nil {
		return m.ContainerRecommendations
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerRecommendation)(nil), "containersai.datahub.recommendation.v1alpha2.ContainerRecommendation")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialLimitRecommendationsEntry")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialRequestRecommendationsEntry")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.LimitRecommendationsEntry")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.RequestRecommendationsEntry")
	proto.RegisterType((*AssignPodPolicy)(nil), "containersai.datahub.recommendation.v1alpha2.AssignPodPolicy")
	proto.RegisterType((*PodRecommendation)(nil), "containersai.datahub.recommendation.v1alpha2.PodRecommendation")
	proto.RegisterEnum("containersai.datahub.recommendation.v1alpha2.RecommendationPolicy", RecommendationPolicy_name, RecommendationPolicy_value)
}

func init() {
	proto.RegisterFile("datahub/recommendation/v1alpha2/recommendation.proto", fileDescriptor_recommendation_e67305e6a7503f64)
}

var fileDescriptor_recommendation_e67305e6a7503f64 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0xb4, 0xb4, 0x13, 0xa0, 0xe9, 0xaa, 0xb4, 0x69, 0xaa, 0xd2, 0x28, 0xe2, 0x50,
	0x21, 0xb0, 0x69, 0xe0, 0x50, 0xf1, 0xa9, 0x34, 0x49, 0xd5, 0x48, 0xcd, 0x87, 0xdc, 0x70, 0x80,
	0x8b, 0xb5, 0xb1, 0xb7, 0xc9, 0x0a, 0xdb, 0x6b, 0xec, 0x4d, 0xab, 0x88, 0x0b, 0xea, 0x6f, 0xe0,
	0xc4, 0x1f, 0x40, 0xe2, 0xc4, 0x89, 0x3b, 0xff, 0x0c, 0x79, 0x6d, 0x27, 0x4d, 0x88, 0x1b, 0x15,
	0xa5, 0xb7, 0xdd, 0xd9, 0x99, 0x79, 0x6f, 0x9e, 0xd7, 0xcf, 0x86, 0x17, 0x06, 0xe6, 0xb8, 0xd7,
	0xef, 0x28, 0x2e, 0xd1, 0x99, 0x65, 0x11, 0xdb, 0xc0, 0x9c, 0x32, 0x5b, 0x39, 0xdf, 0xc7, 0xa6,
	0xd3, 0xc3, 0xc5, 0x89, 0xb8, 0xec, 0xb8, 0x8c, 0x33, 0xf4, 0x44, 0x67, 0x36, 0xc7, 0xd4, 0x26,
	0xae, 0x87, 0xa9, 0x1c, 0xb6, 0x90, 0x27, 0x52, 0xa3, 0x16, 0xb9, 0xdd, 0x2e, 0x63, 0x5d, 0x93,
	0x28, 0xa2, 0xb6, 0xd3, 0x3f, 0x53, 0x38, 0xb5, 0x88, 0xc7, 0xb1, 0xe5, 0x04, 0xed, 0x72, 0xfb,
	0x23, 0x12, 0x1e, 0xeb, 0xbb, 0x3a, 0x51, 0x2c, 0xc2, 0xb1, 0x1f, 0x1c, 0xf1, 0x88, 0x22, 0xb1,
	0x25, 0x0e, 0x33, 0x14, 0xec, 0x79, 0xb4, 0x7b, 0x85, 0x7c, 0xb0, 0x0f, 0x4b, 0x1e, 0x45, 0x25,
	0x16, 0xe1, 0x2e, 0xd5, 0xc7, 0x5a, 0xbb, 0x54, 0x0f, 0xb2, 0x0a, 0x7f, 0x56, 0x60, 0xb3, 0x1c,
	0x4d, 0xa7, 0x8e, 0x4d, 0x84, 0x10, 0xa4, 0x6c, 0x6c, 0x91, 0xac, 0x94, 0x97, 0xf6, 0x56, 0x54,
	0xb1, 0x46, 0xdf, 0x24, 0x78, 0x60, 0x52, 0x8b, 0x72, 0x6d, 0x7c, 0x7c, 0x2f, 0x9b, 0xc8, 0x27,
	0xf7, 0xd2, 0x45, 0x4d, 0xbe, 0x89, 0x56, 0x72, 0x0c, 0xb4, 0x7c, 0xe2, 0x43, 0x8c, 0xc7, 0xbc,
	0xaa, 0xcd, 0xdd, 0x81, 0xba, 0x6e, 0x4e, 0x39, 0x42, 0xdf, 0x25, 0xd8, 0x74, 0xc9, 0xe7, 0x3e,
	0xf1, 0xfe, 0x25, 0x96, 0x14, 0xc4, 0xf0, 0x7c, 0x88, 0xa9, 0x01, 0xc8, 0x54, 0x6a, 0x1b, 0xee,
	0xd4, 0x43, 0xf4, 0x53, 0x82, 0x1d, 0x6a, 0x53, 0x4e, 0xb1, 0xa9, 0x4d, 0xd7, 0x2e, 0x25, 0x28,
	0x9e, 0xcd, 0x87, 0x62, 0x2d, 0x80, 0x8a, 0x97, 0x70, 0x9b, 0xc6, 0x67, 0xa0, 0x5f, 0x12, 0xec,
	0x46, 0x64, 0xe3, 0x14, 0x5d, 0x14, 0x74, 0x7b, 0x73, 0xa5, 0x7b, 0x9d, 0xb0, 0x91, 0x7a, 0xd3,
	0x73, 0x72, 0x03, 0xd8, 0x8a, 0x1d, 0x16, 0x65, 0x20, 0xf9, 0x89, 0x0c, 0xc4, 0x1d, 0x5e, 0x54,
	0xfd, 0x25, 0x3a, 0x82, 0xc5, 0x73, 0x6c, 0xf6, 0x49, 0x36, 0x91, 0x97, 0xf6, 0xd2, 0xc5, 0x67,
	0xd3, 0xc7, 0x08, 0xdf, 0x92, 0x21, 0xfd, 0xba, 0xd8, 0x57, 0x30, 0xc7, 0x6a, 0x50, 0xfe, 0x32,
	0x71, 0x20, 0xe5, 0xbe, 0xc0, 0xf6, 0x35, 0xc4, 0x6f, 0x19, 0xfc, 0xab, 0x04, 0xf9, 0x59, 0x0f,
	0xfb, 0x96, 0x29, 0x5c, 0x4a, 0x50, 0x98, 0xfd, 0x00, 0x6f, 0x97, 0x44, 0xe1, 0x77, 0x02, 0x56,
	0x4b, 0xc2, 0xfa, 0x5a, 0xcc, 0x68, 0x31, 0x93, 0xea, 0x03, 0x24, 0x43, 0xca, 0xb7, 0x5d, 0x01,
	0x99, 0x2e, 0xe6, 0xe4, 0xc0, 0x93, 0xe5, 0xc8, 0x93, 0xe5, 0x76, 0xe4, 0xc9, 0xaa, 0xc8, 0x43,
	0x67, 0x70, 0xcf, 0x66, 0x06, 0xd1, 0x1c, 0x97, 0x32, 0x97, 0xf2, 0x41, 0xc8, 0xeb, 0x5d, 0xdc,
	0x1d, 0x0f, 0x5c, 0x58, 0x76, 0x98, 0x21, 0x87, 0xae, 0x3b, 0x24, 0xd9, 0x60, 0x06, 0x69, 0x85,
	0x6d, 0x8e, 0x17, 0xd4, 0xbb, 0xf6, 0x95, 0x3d, 0xea, 0x84, 0x38, 0x1e, 0x31, 0x89, 0xce, 0x99,
	0x9b, 0x4d, 0x0a, 0x9c, 0x57, 0xff, 0x81, 0x73, 0x1a, 0xb6, 0x88, 0x30, 0xa2, 0x3d, 0xda, 0x81,
	0x15, 0x81, 0x21, 0xcc, 0x3b, 0xe5, 0x9b, 0xf7, 0xf1, 0x82, 0xba, 0xec, 0x87, 0x1a, 0xd8, 0x22,
	0x87, 0xcb, 0xb0, 0xe4, 0x08, 0x91, 0x0a, 0x3f, 0x92, 0xb0, 0xd6, 0x62, 0xc6, 0x84, 0xed, 0xf7,
	0x60, 0xd5, 0xaf, 0xf4, 0x1c, 0xac, 0x13, 0x43, 0x1b, 0x7e, 0x01, 0x66, 0x8b, 0x31, 0xfc, 0x66,
	0x8d, 0xa4, 0x18, 0xf6, 0xf1, 0x57, 0xea, 0x7d, 0x7b, 0x6c, 0x8f, 0x0e, 0x20, 0x8b, 0x1d, 0xc7,
	0x1c, 0x4c, 0x18, 0x8c, 0x66, 0xb3, 0x0b, 0xa1, 0xff, 0xb2, 0xba, 0x21, 0xce, 0xc7, 0x09, 0x36,
	0xd8, 0x05, 0xa2, 0xb0, 0x16, 0xc8, 0xa1, 0x39, 0xcc, 0xd0, 0x82, 0x71, 0x42, 0x29, 0xdf, 0xdc,
	0xcc, 0x96, 0x26, 0x2e, 0x8e, 0xba, 0x8a, 0x27, 0x6e, 0xd2, 0xa5, 0x04, 0x5b, 0xc3, 0x8e, 0x31,
	0xce, 0x5d, 0x9d, 0x8b, 0x15, 0xaa, 0x59, 0x7d, 0xfa, 0x81, 0xf7, 0xf8, 0x14, 0xd6, 0xc7, 0x43,
	0x21, 0xb9, 0x02, 0x3c, 0x54, 0xab, 0xe5, 0x66, 0xbd, 0x5e, 0x6d, 0x54, 0x4a, 0xed, 0x5a, 0xb3,
	0xd1, 0x6a, 0x9e, 0xd4, 0xca, 0x1f, 0xb4, 0xf7, 0x8d, 0x4a, 0xf5, 0xa8, 0xd6, 0xa8, 0x56, 0x32,
	0x0b, 0x08, 0x60, 0xe9, 0xb4, 0x5d, 0x3a, 0x3c, 0xa9, 0x66, 0x24, 0x94, 0x86, 0x3b, 0xe5, 0x66,
	0xbd, 0x55, 0x2a, 0xb7, 0x33, 0x89, 0xc3, 0xb7, 0x1f, 0x5f, 0x77, 0x29, 0xf7, 0x89, 0xea, 0xcc,
	0x52, 0x46, 0x13, 0x3c, 0xc5, 0x54, 0xc1, 0x0e, 0x55, 0x66, 0xfc, 0x2b, 0x75, 0x96, 0xc4, 0xdb,
	0xf4, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa3, 0xac, 0x27, 0x55, 0x09, 0x00, 0x00,
}
