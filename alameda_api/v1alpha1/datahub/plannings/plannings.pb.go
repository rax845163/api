// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/plannings/plannings.proto

package plannings

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
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

type Planning struct {
	LimitPlannings          []*common.MetricData `protobuf:"bytes,1,rep,name=limit_plannings,json=limitPlannings,proto3" json:"limit_plannings,omitempty"`
	RequestPlannings        []*common.MetricData `protobuf:"bytes,2,rep,name=request_plannings,json=requestPlannings,proto3" json:"request_plannings,omitempty"`
	InitialLimitPlannings   []*common.MetricData `protobuf:"bytes,3,rep,name=initial_limit_plannings,json=initialLimitPlannings,proto3" json:"initial_limit_plannings,omitempty"`
	InitialRequestPlannings []*common.MetricData `protobuf:"bytes,4,rep,name=initial_request_plannings,json=initialRequestPlannings,proto3" json:"initial_request_plannings,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *Planning) Reset()         { *m = Planning{} }
func (m *Planning) String() string { return proto.CompactTextString(m) }
func (*Planning) ProtoMessage()    {}
func (*Planning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{0}
}

func (m *Planning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Planning.Unmarshal(m, b)
}
func (m *Planning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Planning.Marshal(b, m, deterministic)
}
func (m *Planning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Planning.Merge(m, src)
}
func (m *Planning) XXX_Size() int {
	return xxx_messageInfo_Planning.Size(m)
}
func (m *Planning) XXX_DiscardUnknown() {
	xxx_messageInfo_Planning.DiscardUnknown(m)
}

var xxx_messageInfo_Planning proto.InternalMessageInfo

func (m *Planning) GetLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.LimitPlannings
	}
	return nil
}

func (m *Planning) GetRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.RequestPlannings
	}
	return nil
}

func (m *Planning) GetInitialLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialLimitPlannings
	}
	return nil
}

func (m *Planning) GetInitialRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialRequestPlannings
	}
	return nil
}

type ContainerPlanning struct {
	Name                    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitPlannings          []*common.MetricData `protobuf:"bytes,2,rep,name=limit_plannings,json=limitPlannings,proto3" json:"limit_plannings,omitempty"`
	RequestPlannings        []*common.MetricData `protobuf:"bytes,3,rep,name=request_plannings,json=requestPlannings,proto3" json:"request_plannings,omitempty"`
	InitialLimitPlannings   []*common.MetricData `protobuf:"bytes,4,rep,name=initial_limit_plannings,json=initialLimitPlannings,proto3" json:"initial_limit_plannings,omitempty"`
	InitialRequestPlannings []*common.MetricData `protobuf:"bytes,5,rep,name=initial_request_plannings,json=initialRequestPlannings,proto3" json:"initial_request_plannings,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *ContainerPlanning) Reset()         { *m = ContainerPlanning{} }
func (m *ContainerPlanning) String() string { return proto.CompactTextString(m) }
func (*ContainerPlanning) ProtoMessage()    {}
func (*ContainerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{1}
}

func (m *ContainerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPlanning.Unmarshal(m, b)
}
func (m *ContainerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPlanning.Marshal(b, m, deterministic)
}
func (m *ContainerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPlanning.Merge(m, src)
}
func (m *ContainerPlanning) XXX_Size() int {
	return xxx_messageInfo_ContainerPlanning.Size(m)
}
func (m *ContainerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPlanning proto.InternalMessageInfo

func (m *ContainerPlanning) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPlanning) GetLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.LimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.RequestPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialLimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialRequestPlannings
	}
	return nil
}

type PodPlanning struct {
	ObjectMeta           *resources.ObjectMeta      `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType               `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                     `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64                    `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                       `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp       `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp       `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	AssignPodPolicy      *resources.AssignPodPolicy `protobuf:"bytes,8,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	TopController        *resources.Controller      `protobuf:"bytes,9,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	ContainerPlannings   []*ContainerPlanning       `protobuf:"bytes,10,rep,name=container_plannings,json=containerPlannings,proto3" json:"container_plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PodPlanning) Reset()         { *m = PodPlanning{} }
func (m *PodPlanning) String() string { return proto.CompactTextString(m) }
func (*PodPlanning) ProtoMessage()    {}
func (*PodPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{2}
}

func (m *PodPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodPlanning.Unmarshal(m, b)
}
func (m *PodPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodPlanning.Marshal(b, m, deterministic)
}
func (m *PodPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodPlanning.Merge(m, src)
}
func (m *PodPlanning) XXX_Size() int {
	return xxx_messageInfo_PodPlanning.Size(m)
}
func (m *PodPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_PodPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_PodPlanning proto.InternalMessageInfo

func (m *PodPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *PodPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *PodPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *PodPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *PodPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *PodPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *PodPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *PodPlanning) GetAssignPodPolicy() *resources.AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodPlanning) GetTopController() *resources.Controller {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *PodPlanning) GetContainerPlannings() []*ContainerPlanning {
	if m != nil {
		return m.ContainerPlannings
	}
	return nil
}

type ControllerPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                 resources.Kind        `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,3,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,4,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,5,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,6,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,9,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ControllerPlanning) Reset()         { *m = ControllerPlanning{} }
func (m *ControllerPlanning) String() string { return proto.CompactTextString(m) }
func (*ControllerPlanning) ProtoMessage()    {}
func (*ControllerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{3}
}

func (m *ControllerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerPlanning.Unmarshal(m, b)
}
func (m *ControllerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerPlanning.Marshal(b, m, deterministic)
}
func (m *ControllerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerPlanning.Merge(m, src)
}
func (m *ControllerPlanning) XXX_Size() int {
	return xxx_messageInfo_ControllerPlanning.Size(m)
}
func (m *ControllerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerPlanning proto.InternalMessageInfo

func (m *ControllerPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ControllerPlanning) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (m *ControllerPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ControllerPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ControllerPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ControllerPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ControllerPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ControllerPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ControllerPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type ApplicationPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApplicationPlanning) Reset()         { *m = ApplicationPlanning{} }
func (m *ApplicationPlanning) String() string { return proto.CompactTextString(m) }
func (*ApplicationPlanning) ProtoMessage()    {}
func (*ApplicationPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{4}
}

func (m *ApplicationPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationPlanning.Unmarshal(m, b)
}
func (m *ApplicationPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationPlanning.Marshal(b, m, deterministic)
}
func (m *ApplicationPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationPlanning.Merge(m, src)
}
func (m *ApplicationPlanning) XXX_Size() int {
	return xxx_messageInfo_ApplicationPlanning.Size(m)
}
func (m *ApplicationPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationPlanning proto.InternalMessageInfo

func (m *ApplicationPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ApplicationPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ApplicationPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ApplicationPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ApplicationPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ApplicationPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ApplicationPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ApplicationPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type NamespacePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NamespacePlanning) Reset()         { *m = NamespacePlanning{} }
func (m *NamespacePlanning) String() string { return proto.CompactTextString(m) }
func (*NamespacePlanning) ProtoMessage()    {}
func (*NamespacePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{5}
}

func (m *NamespacePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacePlanning.Unmarshal(m, b)
}
func (m *NamespacePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacePlanning.Marshal(b, m, deterministic)
}
func (m *NamespacePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacePlanning.Merge(m, src)
}
func (m *NamespacePlanning) XXX_Size() int {
	return xxx_messageInfo_NamespacePlanning.Size(m)
}
func (m *NamespacePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacePlanning proto.InternalMessageInfo

func (m *NamespacePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NamespacePlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *NamespacePlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *NamespacePlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *NamespacePlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *NamespacePlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *NamespacePlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *NamespacePlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type NodePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NodePlanning) Reset()         { *m = NodePlanning{} }
func (m *NodePlanning) String() string { return proto.CompactTextString(m) }
func (*NodePlanning) ProtoMessage()    {}
func (*NodePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{6}
}

func (m *NodePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePlanning.Unmarshal(m, b)
}
func (m *NodePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePlanning.Marshal(b, m, deterministic)
}
func (m *NodePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePlanning.Merge(m, src)
}
func (m *NodePlanning) XXX_Size() int {
	return xxx_messageInfo_NodePlanning.Size(m)
}
func (m *NodePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NodePlanning proto.InternalMessageInfo

func (m *NodePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NodePlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *NodePlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *NodePlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *NodePlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *NodePlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *NodePlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *NodePlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type ClusterPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterPlanning) Reset()         { *m = ClusterPlanning{} }
func (m *ClusterPlanning) String() string { return proto.CompactTextString(m) }
func (*ClusterPlanning) ProtoMessage()    {}
func (*ClusterPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{7}
}

func (m *ClusterPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPlanning.Unmarshal(m, b)
}
func (m *ClusterPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPlanning.Marshal(b, m, deterministic)
}
func (m *ClusterPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPlanning.Merge(m, src)
}
func (m *ClusterPlanning) XXX_Size() int {
	return xxx_messageInfo_ClusterPlanning.Size(m)
}
func (m *ClusterPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPlanning proto.InternalMessageInfo

func (m *ClusterPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ClusterPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ClusterPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ClusterPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ClusterPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ClusterPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ClusterPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ClusterPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

func init() {
	proto.RegisterType((*Planning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.Planning")
	proto.RegisterType((*ContainerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning")
	proto.RegisterType((*PodPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.PodPlanning")
	proto.RegisterType((*ControllerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning")
	proto.RegisterType((*ApplicationPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning")
	proto.RegisterType((*NamespacePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning")
	proto.RegisterType((*NodePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NodePlanning")
	proto.RegisterType((*ClusterPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/plannings/plannings.proto", fileDescriptor_36c6a4155f49dbb5)
}

var fileDescriptor_36c6a4155f49dbb5 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x5d, 0x6b, 0x13, 0x4d,
	0x14, 0x26, 0xcd, 0xa6, 0x4d, 0x4e, 0xfa, 0x95, 0x29, 0x2f, 0xef, 0xbe, 0x85, 0x97, 0x86, 0x5c,
	0xe5, 0x42, 0x77, 0x69, 0xa4, 0xc5, 0x22, 0x82, 0x6d, 0xf4, 0xa2, 0x68, 0x6b, 0x59, 0x0a, 0x82,
	0x08, 0xcb, 0x64, 0x77, 0x4c, 0x47, 0x77, 0x67, 0xc6, 0x9d, 0x89, 0x25, 0x3f, 0x41, 0x10, 0xf1,
	0x07, 0xf9, 0x47, 0xf4, 0x87, 0x78, 0x2d, 0x3b, 0xd9, 0x8f, 0xb4, 0x95, 0x76, 0xd3, 0xda, 0xa8,
	0x90, 0xbb, 0x99, 0xc9, 0x9e, 0xe7, 0x99, 0x73, 0xce, 0x73, 0x1e, 0xb2, 0x0b, 0xdb, 0x38, 0xc0,
	0x21, 0xf1, 0xb1, 0x8b, 0x05, 0xb5, 0xdf, 0x6f, 0xe2, 0x40, 0x9c, 0xe0, 0x4d, 0xdb, 0xc7, 0x0a,
	0x9f, 0x0c, 0x7a, 0xb6, 0x08, 0x30, 0x63, 0x94, 0xf5, 0x65, 0xbe, 0xb2, 0x44, 0xc4, 0x15, 0x47,
	0xb6, 0xc7, 0x99, 0xc2, 0x94, 0x91, 0x48, 0x62, 0x6a, 0x25, 0x20, 0x56, 0x0a, 0x60, 0x25, 0x00,
	0x56, 0x16, 0xb6, 0xbe, 0x79, 0x29, 0x91, 0xc7, 0xc3, 0x90, 0x33, 0x3b, 0x24, 0x2a, 0xa2, 0x5e,
	0xc2, 0xb1, 0xde, 0x29, 0x78, 0x37, 0x35, 0x14, 0x24, 0x8d, 0xd9, 0xba, 0x34, 0x26, 0x22, 0x92,
	0x0f, 0x22, 0x8f, 0xc8, 0x98, 0x09, 0xc7, 0xa7, 0x13, 0x86, 0x09, 0x1e, 0x50, 0x8f, 0x66, 0x6c,
	0xdb, 0x05, 0xc3, 0xb2, 0x55, 0x12, 0xb7, 0xd1, 0xe7, 0xbc, 0x1f, 0x10, 0x5b, 0xef, 0x7a, 0x83,
	0xd7, 0xb6, 0xa2, 0x21, 0x91, 0x0a, 0x87, 0x62, 0xf4, 0x40, 0xeb, 0x4b, 0x19, 0xaa, 0x47, 0x49,
	0x82, 0x08, 0xc3, 0x4a, 0x40, 0x43, 0xaa, 0xdc, 0x2c, 0x65, 0xb3, 0xd4, 0x2c, 0xb7, 0xeb, 0x9d,
	0xfb, 0x56, 0xb1, 0x2e, 0x8c, 0xaa, 0x6b, 0x1d, 0xe8, 0xea, 0x3e, 0xc6, 0x0a, 0x3b, 0xcb, 0x1a,
	0x30, 0x65, 0x90, 0x88, 0x40, 0x23, 0x22, 0xef, 0x06, 0x44, 0x8e, 0x93, 0xcc, 0xdd, 0x90, 0x64,
	0x35, 0x81, 0xcc, 0x69, 0x04, 0xfc, 0x4b, 0x19, 0x55, 0x14, 0x07, 0xee, 0xf9, 0x8c, 0xca, 0x37,
	0x24, 0xfb, 0x27, 0x01, 0x7e, 0x76, 0x36, 0x31, 0x05, 0xff, 0xa5, 0x8c, 0x17, 0x13, 0x34, 0x6e,
	0xc8, 0x99, 0x26, 0xe3, 0x9c, 0xcb, 0xb3, 0xf5, 0xb5, 0x0c, 0x8d, 0x6e, 0x0a, 0x9a, 0xf5, 0x11,
	0x81, 0xc1, 0x70, 0x48, 0xcc, 0x52, 0xb3, 0xd4, 0xae, 0x39, 0x7a, 0xfd, 0xb3, 0xde, 0xce, 0x4d,
	0xa3, 0xb7, 0xe5, 0x69, 0xf6, 0xd6, 0xf8, 0x0d, 0xbd, 0xad, 0xdc, 0x56, 0x6f, 0xbf, 0x57, 0xa0,
	0x7e, 0xc4, 0xfd, 0xac, 0xab, 0xaf, 0xa0, 0xce, 0x7b, 0x6f, 0x88, 0xa7, 0xdc, 0xd8, 0x53, 0x74,
	0x73, 0xeb, 0x9d, 0x07, 0x05, 0x79, 0x73, 0x63, 0x78, 0xae, 0x31, 0x0e, 0x88, 0xc2, 0x0e, 0xf0,
	0x6c, 0x8d, 0x7a, 0xb0, 0x94, 0xe6, 0xe4, 0xc6, 0x3e, 0x67, 0xce, 0x35, 0x4b, 0xed, 0xe5, 0xce,
	0x43, 0x6b, 0x42, 0xff, 0xb5, 0xd2, 0xfb, 0x1e, 0x0f, 0x05, 0x71, 0x16, 0xc5, 0xd8, 0x0e, 0x6d,
	0x40, 0x3d, 0xe3, 0xa0, 0xbe, 0x59, 0xd6, 0xf2, 0x84, 0xf4, 0x68, 0xdf, 0x47, 0xff, 0x03, 0x28,
	0xae, 0x70, 0xe0, 0x7a, 0x5c, 0x2a, 0xd3, 0x68, 0x96, 0xda, 0x25, 0xa7, 0xa6, 0x4f, 0xba, 0x5c,
	0x2a, 0x74, 0x07, 0x10, 0x16, 0x22, 0x18, 0x66, 0xd5, 0x77, 0x19, 0x3f, 0x35, 0x2b, 0xcd, 0x52,
	0xbb, 0xea, 0xac, 0xea, 0x5f, 0x52, 0xf2, 0x43, 0x7e, 0x8a, 0x76, 0x00, 0xa4, 0xc2, 0x91, 0x72,
	0x63, 0xcf, 0x33, 0xe7, 0x75, 0xb9, 0xd6, 0xad, 0x91, 0x21, 0x5a, 0xa9, 0x21, 0x5a, 0xc7, 0xa9,
	0x21, 0x3a, 0x35, 0xfd, 0x74, 0xbc, 0x47, 0x5b, 0x50, 0x25, 0xcc, 0x1f, 0x05, 0x2e, 0x5c, 0x19,
	0xb8, 0x40, 0x98, 0xaf, 0xc3, 0x02, 0x68, 0x60, 0x29, 0x69, 0x9f, 0xb9, 0x82, 0xfb, 0xae, 0xb6,
	0xf0, 0xa1, 0x59, 0xd5, 0xf1, 0x8f, 0x26, 0xee, 0xd3, 0xae, 0x46, 0x8a, 0x05, 0xa0, 0x71, 0x9c,
	0x15, 0x7c, 0xf6, 0x00, 0xf5, 0x60, 0x59, 0x71, 0xe1, 0xc6, 0xb8, 0x11, 0x0f, 0x02, 0x12, 0x99,
	0xb5, 0x6b, 0x4a, 0xa2, 0x9b, 0x41, 0x38, 0x4b, 0x8a, 0x8b, 0x7c, 0x8b, 0x24, 0xac, 0x65, 0x60,
	0x63, 0x9a, 0x07, 0xad, 0xf9, 0xbd, 0x89, 0xb5, 0x71, 0xc1, 0xaa, 0x1c, 0xe4, 0x9d, 0x3f, 0x92,
	0xad, 0x6f, 0x06, 0xa0, 0xfc, 0x0e, 0x53, 0xd2, 0xff, 0x3e, 0x18, 0x6f, 0x29, 0xf3, 0x13, 0xd9,
	0x6f, 0x4d, 0x0c, 0xfb, 0x94, 0x32, 0xdf, 0xd1, 0x10, 0x17, 0x47, 0xa9, 0x7c, 0xeb, 0xa3, 0x64,
	0x5c, 0x31, 0x4a, 0x95, 0x62, 0xa3, 0x34, 0x5f, 0x68, 0x94, 0x16, 0xae, 0x3b, 0x4a, 0xd5, 0xe2,
	0xa3, 0xf4, 0x02, 0x6a, 0xb9, 0xdc, 0x6a, 0x5a, 0x6e, 0x3b, 0xd7, 0xae, 0x9f, 0x93, 0x63, 0xb5,
	0x3e, 0x1b, 0xb0, 0xb6, 0x2b, 0x44, 0x40, 0x3d, 0xac, 0x28, 0x67, 0x33, 0x77, 0xfd, 0x6b, 0xdc,
	0xf5, 0x8c, 0x24, 0xaa, 0xbf, 0x50, 0x12, 0x9f, 0x0c, 0x68, 0x1c, 0xe2, 0x90, 0x48, 0x81, 0x3d,
	0x32, 0x13, 0xc4, 0x4c, 0x10, 0x1f, 0x0c, 0x58, 0x3c, 0xe4, 0xfe, 0x4c, 0x0b, 0x33, 0x2d, 0xb4,
	0x3e, 0x1a, 0xb0, 0xd2, 0x0d, 0x06, 0x52, 0x4d, 0xed, 0x9f, 0xc8, 0x4c, 0x0e, 0x7f, 0xae, 0x1c,
	0xf6, 0x9e, 0xbc, 0xec, 0xf6, 0xa9, 0x4a, 0xde, 0xe4, 0xc6, 0xbe, 0x4d, 0xdd, 0xc5, 0xd4, 0xc6,
	0x82, 0xda, 0xc5, 0xbe, 0x24, 0xf5, 0xe6, 0xf5, 0xe5, 0xef, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0x70, 0xfd, 0x63, 0x16, 0x13, 0x00, 0x00,
}
