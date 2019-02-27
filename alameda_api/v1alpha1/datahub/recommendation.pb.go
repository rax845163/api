// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/recommendation.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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
// Represents a resource configuration recommendation
//
// It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial stage
//
type ContainerRecommendation struct {
	Name                          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitRecommendations          []*MetricData `protobuf:"bytes,2,rep,name=limit_recommendations,json=limitRecommendations,proto3" json:"limit_recommendations,omitempty"`
	RequestRecommendations        []*MetricData `protobuf:"bytes,3,rep,name=request_recommendations,json=requestRecommendations,proto3" json:"request_recommendations,omitempty"`
	InitialLimitRecommendations   []*MetricData `protobuf:"bytes,4,rep,name=initial_limit_recommendations,json=initialLimitRecommendations,proto3" json:"initial_limit_recommendations,omitempty"`
	InitialRequestRecommendations []*MetricData `protobuf:"bytes,5,rep,name=initial_request_recommendations,json=initialRequestRecommendations,proto3" json:"initial_request_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}      `json:"-"`
	XXX_unrecognized              []byte        `json:"-"`
	XXX_sizecache                 int32         `json:"-"`
}

func (m *ContainerRecommendation) Reset()         { *m = ContainerRecommendation{} }
func (m *ContainerRecommendation) String() string { return proto.CompactTextString(m) }
func (*ContainerRecommendation) ProtoMessage()    {}
func (*ContainerRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_095cd91eb31a5f85, []int{0}
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

func (m *ContainerRecommendation) GetLimitRecommendations() []*MetricData {
	if m != nil {
		return m.LimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetRequestRecommendations() []*MetricData {
	if m != nil {
		return m.RequestRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialLimitRecommendations() []*MetricData {
	if m != nil {
		return m.InitialLimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialRequestRecommendations() []*MetricData {
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
	return fileDescriptor_recommendation_095cd91eb31a5f85, []int{1}
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
	NodePriority *NodePriority `protobuf:"bytes,2,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type AssignPodPolicy_NodeSelector struct {
	NodeSelector *Selector `protobuf:"bytes,3,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
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

func (m *AssignPodPolicy) GetNodePriority() *NodePriority {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeSelector() *Selector {
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
		msg := new(NodePriority)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodePriority{msg}
		return true, err
	case 3: // policy.node_selector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Selector)
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
	NamespacedName           *NamespacedName            `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ApplyRecommendationNow   bool                       `protobuf:"varint,2,opt,name=apply_recommendation_now,json=applyRecommendationNow,proto3" json:"apply_recommendation_now,omitempty"`
	AssignPodPolicy          *AssignPodPolicy           `protobuf:"bytes,3,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	ContainerRecommendations []*ContainerRecommendation `protobuf:"bytes,4,rep,name=container_recommendations,json=containerRecommendations,proto3" json:"container_recommendations,omitempty"`
	StartTime                *timestamp.Timestamp       `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                  *timestamp.Timestamp       `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TopController            *TopController             `protobuf:"bytes,7,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	RecommendationId         string                     `protobuf:"bytes,8,opt,name=recommendation_id,json=recommendationId,proto3" json:"recommendation_id,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *PodRecommendation) Reset()         { *m = PodRecommendation{} }
func (m *PodRecommendation) String() string { return proto.CompactTextString(m) }
func (*PodRecommendation) ProtoMessage()    {}
func (*PodRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_095cd91eb31a5f85, []int{2}
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

func (m *PodRecommendation) GetNamespacedName() *NamespacedName {
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

func (m *PodRecommendation) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *PodRecommendation) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *PodRecommendation) GetTopController() *TopController {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *PodRecommendation) GetRecommendationId() string {
	if m != nil {
		return m.RecommendationId
	}
	return ""
}

func init() {
	proto.RegisterType((*ContainerRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation")
	proto.RegisterType((*AssignPodPolicy)(nil), "containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy")
	proto.RegisterType((*PodRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.PodRecommendation")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/recommendation.proto", fileDescriptor_recommendation_095cd91eb31a5f85)
}

var fileDescriptor_recommendation_095cd91eb31a5f85 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0xdd, 0xd6, 0x6e, 0x6b, 0x5d, 0xb6, 0x52, 0x0b, 0xb6, 0x50, 0x34, 0x6d, 0xea, 0x05, 0x1a,
	0x9a, 0x94, 0xd2, 0xc2, 0xf8, 0xb9, 0x42, 0x30, 0x2e, 0x86, 0x04, 0x55, 0x65, 0x26, 0x71, 0x01,
	0x92, 0xe5, 0xc6, 0xa6, 0xb3, 0x48, 0x6c, 0x63, 0xbb, 0x9b, 0x8a, 0x78, 0x06, 0x9e, 0x81, 0xb7,
	0xe4, 0x16, 0xc5, 0x49, 0xaa, 0x25, 0xb4, 0x6b, 0xd4, 0x3b, 0x7f, 0xb1, 0xcf, 0x39, 0xdf, 0xf9,
	0x72, 0x6c, 0xd0, 0x23, 0x21, 0x89, 0x18, 0x25, 0x98, 0x28, 0xde, 0xbd, 0xea, 0x91, 0x50, 0x5d,
	0x92, 0x5e, 0x97, 0x12, 0x4b, 0x2e, 0x27, 0xa3, 0xae, 0x66, 0x81, 0x8c, 0x22, 0x26, 0x28, 0xb1,
	0x5c, 0x0a, 0x5f, 0x69, 0x69, 0x25, 0x7c, 0x14, 0x48, 0x61, 0x09, 0x17, 0x4c, 0x1b, 0x4c, 0xb8,
	0x9f, 0x12, 0xf8, 0x19, 0xd8, 0x4f, 0xc1, 0xed, 0xc3, 0xb1, 0x94, 0xe3, 0x90, 0x75, 0x1d, 0x6a,
	0x34, 0xf9, 0xd6, 0xb5, 0x3c, 0x62, 0xc6, 0x92, 0x48, 0x25, 0x44, 0xed, 0x93, 0x5b, 0xb5, 0x23,
	0x66, 0x49, 0xbc, 0x4e, 0x0f, 0xdf, 0xde, 0xa8, 0x92, 0x14, 0x13, 0x63, 0xf8, 0x58, 0x44, 0x4c,
	0xd8, 0x14, 0xf2, 0x78, 0x19, 0xbf, 0xe6, 0x41, 0xa9, 0x56, 0x34, 0x33, 0x72, 0xa2, 0x03, 0x96,
	0x1c, 0xee, 0xfc, 0xad, 0x80, 0xfd, 0xb3, 0x6c, 0x06, 0x28, 0x37, 0x22, 0x08, 0x41, 0x55, 0x90,
	0x88, 0x79, 0xeb, 0x47, 0xeb, 0xc7, 0x75, 0xe4, 0xd6, 0x70, 0x0c, 0xee, 0x87, 0x3c, 0xe2, 0x16,
	0xe7, 0xc7, 0x69, 0xbc, 0x8d, 0xa3, 0xca, 0x71, 0xa3, 0xdf, 0xf7, 0xcb, 0x0d, 0xd4, 0xff, 0xe8,
	0x3a, 0x7e, 0x47, 0x2c, 0x41, 0xf7, 0x1c, 0x61, 0x5e, 0xdb, 0xc0, 0xef, 0x60, 0x5f, 0xb3, 0x1f,
	0x13, 0x66, 0xfe, 0x97, 0xaa, 0xac, 0x2c, 0xb5, 0x97, 0x52, 0x16, 0xc5, 0xae, 0xc0, 0x01, 0x17,
	0xdc, 0x72, 0x12, 0xe2, 0xf9, 0xee, 0xaa, 0x2b, 0x4b, 0x3e, 0x4c, 0x89, 0x3f, 0xcc, 0x33, 0xf9,
	0x13, 0x1c, 0x66, 0xba, 0x8b, 0xcc, 0x6e, 0xae, 0xac, 0x9c, 0x59, 0x42, 0x73, 0x3d, 0x77, 0xfe,
	0x6c, 0x80, 0xe6, 0x1b, 0x17, 0xb3, 0xa1, 0xa4, 0x43, 0x19, 0xf2, 0x60, 0x0a, 0x7d, 0x50, 0x8d,
	0x83, 0xed, 0xfe, 0x78, 0xa3, 0xdf, 0xf6, 0x93, 0xd4, 0xfb, 0x59, 0xea, 0xfd, 0x8b, 0x2c, 0xf5,
	0xc8, 0x9d, 0x83, 0x5f, 0xc0, 0x8e, 0x90, 0x94, 0x61, 0xa5, 0xb9, 0xd4, 0xdc, 0x4e, 0xbd, 0x0d,
	0x07, 0x7c, 0x56, 0xb6, 0xdb, 0x81, 0xa4, 0x6c, 0x98, 0x62, 0xcf, 0xd7, 0xd0, 0x1d, 0x71, 0xa3,
	0x86, 0x9f, 0x53, 0x72, 0xc3, 0x42, 0x16, 0x58, 0xa9, 0xbd, 0x8a, 0x23, 0x7f, 0x52, 0x96, 0xfc,
	0x53, 0x8a, 0xcb, 0x88, 0xb3, 0x1a, 0x1e, 0x80, 0xba, 0x23, 0x76, 0xe1, 0xae, 0xc6, 0xe1, 0x3e,
	0x5f, 0x43, 0xb5, 0xf8, 0xd3, 0x80, 0x44, 0xec, 0x6d, 0x0d, 0x6c, 0x29, 0x37, 0x8e, 0xce, 0xef,
	0x4d, 0xd0, 0x1a, 0x4a, 0x5a, 0xb8, 0x16, 0x18, 0x34, 0x63, 0xa4, 0x51, 0x24, 0x60, 0x14, 0xcf,
	0x6e, 0x48, 0xa3, 0xff, 0xbc, 0xb4, 0xed, 0x19, 0x3c, 0x5e, 0xa1, 0x5d, 0x91, 0xab, 0xe1, 0x4b,
	0xe0, 0x11, 0xa5, 0xc2, 0x69, 0x21, 0x0b, 0x58, 0xc8, 0x6b, 0x37, 0xe0, 0x1a, 0xda, 0x73, 0xfb,
	0xf9, 0xbe, 0x06, 0xf2, 0x1a, 0x06, 0xa0, 0x95, 0xbc, 0x1c, 0x38, 0x7e, 0x44, 0x12, 0x17, 0xe9,
	0xd8, 0x5e, 0x94, 0x6d, 0xae, 0x90, 0x09, 0xd4, 0x24, 0x85, 0x90, 0xfc, 0x02, 0x0f, 0x66, 0x54,
	0x0b, 0x2e, 0xca, 0xeb, 0xb2, 0x62, 0x0b, 0x9e, 0x1e, 0xe4, 0x05, 0xf3, 0x37, 0x0c, 0x7c, 0x05,
	0x80, 0xb1, 0x44, 0x5b, 0xec, 0x82, 0xba, 0xb9, 0x34, 0xa8, 0x75, 0x77, 0x3a, 0xae, 0xe1, 0x29,
	0xa8, 0x31, 0x41, 0x13, 0xe0, 0xd6, 0x52, 0xe0, 0x36, 0x13, 0xd4, 0xc1, 0xbe, 0x82, 0x5d, 0x2b,
	0x15, 0x8e, 0x3b, 0xd2, 0x32, 0x0c, 0x99, 0xf6, 0xb6, 0x1d, 0xf8, 0xb4, 0xac, 0xc9, 0x0b, 0xa9,
	0xce, 0x66, 0x60, 0xb4, 0x63, 0x6f, 0x96, 0xf0, 0x04, 0xb4, 0x0a, 0xbf, 0x99, 0x53, 0xaf, 0xe6,
	0x5e, 0xdc, 0xbb, 0xf9, 0x8d, 0xf7, 0x74, 0xb4, 0xe5, 0xfa, 0x7c, 0xfa, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0x28, 0x08, 0xd3, 0xea, 0x06, 0x00, 0x00,
}
