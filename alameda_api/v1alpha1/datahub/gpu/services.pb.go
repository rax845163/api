// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/gpu/services.proto

package gpu

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type ListGpusRequest struct {
	QueryCondition       *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	Host                 string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber          string                 `protobuf:"bytes,3,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListGpusRequest) Reset()         { *m = ListGpusRequest{} }
func (m *ListGpusRequest) String() string { return proto.CompactTextString(m) }
func (*ListGpusRequest) ProtoMessage()    {}
func (*ListGpusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{0}
}

func (m *ListGpusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpusRequest.Unmarshal(m, b)
}
func (m *ListGpusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpusRequest.Marshal(b, m, deterministic)
}
func (m *ListGpusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpusRequest.Merge(m, src)
}
func (m *ListGpusRequest) XXX_Size() int {
	return xxx_messageInfo_ListGpusRequest.Size(m)
}
func (m *ListGpusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpusRequest proto.InternalMessageInfo

func (m *ListGpusRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListGpusRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ListGpusRequest) GetMinorNumber() string {
	if m != nil {
		return m.MinorNumber
	}
	return ""
}

type ListGpusResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Gpus                 []*Gpu         `protobuf:"bytes,2,rep,name=gpus,proto3" json:"gpus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListGpusResponse) Reset()         { *m = ListGpusResponse{} }
func (m *ListGpusResponse) String() string { return proto.CompactTextString(m) }
func (*ListGpusResponse) ProtoMessage()    {}
func (*ListGpusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{1}
}

func (m *ListGpusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpusResponse.Unmarshal(m, b)
}
func (m *ListGpusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpusResponse.Marshal(b, m, deterministic)
}
func (m *ListGpusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpusResponse.Merge(m, src)
}
func (m *ListGpusResponse) XXX_Size() int {
	return xxx_messageInfo_ListGpusResponse.Size(m)
}
func (m *ListGpusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpusResponse proto.InternalMessageInfo

func (m *ListGpusResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListGpusResponse) GetGpus() []*Gpu {
	if m != nil {
		return m.Gpus
	}
	return nil
}

type ListGpuMetricsRequest struct {
	QueryCondition       *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	Host                 string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber          string                 `protobuf:"bytes,3,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListGpuMetricsRequest) Reset()         { *m = ListGpuMetricsRequest{} }
func (m *ListGpuMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGpuMetricsRequest) ProtoMessage()    {}
func (*ListGpuMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{2}
}

func (m *ListGpuMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpuMetricsRequest.Unmarshal(m, b)
}
func (m *ListGpuMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpuMetricsRequest.Marshal(b, m, deterministic)
}
func (m *ListGpuMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpuMetricsRequest.Merge(m, src)
}
func (m *ListGpuMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGpuMetricsRequest.Size(m)
}
func (m *ListGpuMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpuMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpuMetricsRequest proto.InternalMessageInfo

func (m *ListGpuMetricsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListGpuMetricsRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ListGpuMetricsRequest) GetMinorNumber() string {
	if m != nil {
		return m.MinorNumber
	}
	return ""
}

type ListGpuMetricsResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	GpuMetrics           []*GpuMetric   `protobuf:"bytes,2,rep,name=gpu_metrics,json=gpuMetrics,proto3" json:"gpu_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListGpuMetricsResponse) Reset()         { *m = ListGpuMetricsResponse{} }
func (m *ListGpuMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGpuMetricsResponse) ProtoMessage()    {}
func (*ListGpuMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{3}
}

func (m *ListGpuMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpuMetricsResponse.Unmarshal(m, b)
}
func (m *ListGpuMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpuMetricsResponse.Marshal(b, m, deterministic)
}
func (m *ListGpuMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpuMetricsResponse.Merge(m, src)
}
func (m *ListGpuMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGpuMetricsResponse.Size(m)
}
func (m *ListGpuMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpuMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpuMetricsResponse proto.InternalMessageInfo

func (m *ListGpuMetricsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListGpuMetricsResponse) GetGpuMetrics() []*GpuMetric {
	if m != nil {
		return m.GpuMetrics
	}
	return nil
}

type CreateGpuPredictionsRequest struct {
	GpuPredictions       []*GpuPrediction `protobuf:"bytes,1,rep,name=gpu_predictions,json=gpuPredictions,proto3" json:"gpu_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateGpuPredictionsRequest) Reset()         { *m = CreateGpuPredictionsRequest{} }
func (m *CreateGpuPredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGpuPredictionsRequest) ProtoMessage()    {}
func (*CreateGpuPredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{4}
}

func (m *CreateGpuPredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGpuPredictionsRequest.Unmarshal(m, b)
}
func (m *CreateGpuPredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGpuPredictionsRequest.Marshal(b, m, deterministic)
}
func (m *CreateGpuPredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGpuPredictionsRequest.Merge(m, src)
}
func (m *CreateGpuPredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGpuPredictionsRequest.Size(m)
}
func (m *CreateGpuPredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGpuPredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGpuPredictionsRequest proto.InternalMessageInfo

func (m *CreateGpuPredictionsRequest) GetGpuPredictions() []*GpuPrediction {
	if m != nil {
		return m.GpuPredictions
	}
	return nil
}

type ListGpuPredictionsRequest struct {
	QueryCondition       *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	Host                 string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber          string                 `protobuf:"bytes,3,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Granularity          int64                  `protobuf:"varint,4,opt,name=granularity,proto3" json:"granularity,omitempty"`
	ModelId              string                 `protobuf:"bytes,5,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PredictionId         string                 `protobuf:"bytes,6,opt,name=prediction_id,json=predictionId,proto3" json:"prediction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListGpuPredictionsRequest) Reset()         { *m = ListGpuPredictionsRequest{} }
func (m *ListGpuPredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGpuPredictionsRequest) ProtoMessage()    {}
func (*ListGpuPredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{5}
}

func (m *ListGpuPredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpuPredictionsRequest.Unmarshal(m, b)
}
func (m *ListGpuPredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpuPredictionsRequest.Marshal(b, m, deterministic)
}
func (m *ListGpuPredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpuPredictionsRequest.Merge(m, src)
}
func (m *ListGpuPredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGpuPredictionsRequest.Size(m)
}
func (m *ListGpuPredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpuPredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpuPredictionsRequest proto.InternalMessageInfo

func (m *ListGpuPredictionsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListGpuPredictionsRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ListGpuPredictionsRequest) GetMinorNumber() string {
	if m != nil {
		return m.MinorNumber
	}
	return ""
}

func (m *ListGpuPredictionsRequest) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *ListGpuPredictionsRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *ListGpuPredictionsRequest) GetPredictionId() string {
	if m != nil {
		return m.PredictionId
	}
	return ""
}

type ListGpuPredictionsResponse struct {
	Status               *status.Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	GpuPredictions       []*GpuPrediction `protobuf:"bytes,2,rep,name=gpu_predictions,json=gpuPredictions,proto3" json:"gpu_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListGpuPredictionsResponse) Reset()         { *m = ListGpuPredictionsResponse{} }
func (m *ListGpuPredictionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGpuPredictionsResponse) ProtoMessage()    {}
func (*ListGpuPredictionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e24c7d763b1f22, []int{6}
}

func (m *ListGpuPredictionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGpuPredictionsResponse.Unmarshal(m, b)
}
func (m *ListGpuPredictionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGpuPredictionsResponse.Marshal(b, m, deterministic)
}
func (m *ListGpuPredictionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGpuPredictionsResponse.Merge(m, src)
}
func (m *ListGpuPredictionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGpuPredictionsResponse.Size(m)
}
func (m *ListGpuPredictionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGpuPredictionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGpuPredictionsResponse proto.InternalMessageInfo

func (m *ListGpuPredictionsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListGpuPredictionsResponse) GetGpuPredictions() []*GpuPrediction {
	if m != nil {
		return m.GpuPredictions
	}
	return nil
}

func init() {
	proto.RegisterType((*ListGpusRequest)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest")
	proto.RegisterType((*ListGpusResponse)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse")
	proto.RegisterType((*ListGpuMetricsRequest)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest")
	proto.RegisterType((*ListGpuMetricsResponse)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse")
	proto.RegisterType((*CreateGpuPredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest")
	proto.RegisterType((*ListGpuPredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest")
	proto.RegisterType((*ListGpuPredictionsResponse)(nil), "containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/gpu/services.proto", fileDescriptor_62e24c7d763b1f22)
}

var fileDescriptor_62e24c7d763b1f22 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x14, 0x94, 0xb7, 0xcb, 0x02, 0x4e, 0xe9, 0x22, 0x4b, 0x40, 0xba, 0x5c, 0x42, 0xb8, 0x2c, 0x95,
	0xb0, 0xb5, 0x85, 0x03, 0x07, 0x10, 0x52, 0x7b, 0x58, 0x55, 0x02, 0x04, 0x41, 0x5c, 0xb8, 0x44,
	0xde, 0xc4, 0xf2, 0x5a, 0x4a, 0x62, 0xaf, 0x3f, 0x2a, 0xf5, 0x86, 0xf8, 0x1b, 0xfd, 0x09, 0x1c,
	0x38, 0xf1, 0xff, 0x50, 0x1c, 0x6f, 0x43, 0xd1, 0x8a, 0x7e, 0x88, 0x03, 0xdc, 0xe2, 0x37, 0x7e,
	0x33, 0xf3, 0xe6, 0x45, 0x86, 0x84, 0x56, 0xb4, 0x66, 0x25, 0xcd, 0xa9, 0x12, 0xe4, 0x78, 0x46,
	0x2b, 0xb5, 0xa4, 0x33, 0x52, 0x52, 0x4b, 0x97, 0x6e, 0x41, 0xb8, 0x72, 0xc4, 0x30, 0x7d, 0x2c,
	0x0a, 0x66, 0xb0, 0xd2, 0xd2, 0x4a, 0xf4, 0xa4, 0x90, 0x8d, 0xa5, 0xa2, 0x61, 0xda, 0x50, 0x81,
	0x43, 0x37, 0x5e, 0x77, 0xe2, 0xd0, 0x89, 0xb9, 0x72, 0x93, 0xd9, 0x1f, 0xb9, 0x0b, 0x59, 0xd7,
	0xb2, 0x21, 0x2b, 0xc7, 0xb4, 0x58, 0xb3, 0x4f, 0xf6, 0x2e, 0xb4, 0xc3, 0x95, 0x0b, 0x77, 0x1f,
	0x70, 0x29, 0x79, 0xc5, 0x88, 0x56, 0x05, 0x31, 0x96, 0x5a, 0x17, 0x48, 0xd2, 0xef, 0x00, 0x8e,
	0xdf, 0x08, 0x63, 0xe7, 0xca, 0x99, 0x8c, 0xad, 0x1c, 0x33, 0x16, 0x31, 0x38, 0x6e, 0x95, 0x4e,
	0xf2, 0x42, 0x36, 0xa5, 0xb0, 0x42, 0x36, 0x31, 0x48, 0xc0, 0x34, 0xda, 0x7f, 0x89, 0x2f, 0x37,
	0x50, 0x67, 0x17, 0x7f, 0x68, 0x49, 0x0e, 0xd7, 0x1c, 0xd9, 0xce, 0xea, 0xdc, 0x19, 0x21, 0x38,
	0x5c, 0x4a, 0x63, 0xe3, 0x41, 0x02, 0xa6, 0xb7, 0x33, 0xff, 0x8d, 0x1e, 0xc1, 0xed, 0x5a, 0x34,
	0x52, 0xe7, 0x8d, 0xab, 0x17, 0x4c, 0xc7, 0x5b, 0x1e, 0x8b, 0x7c, 0xed, 0x9d, 0x2f, 0xa5, 0x5f,
	0x01, 0xbc, 0xdb, 0x3b, 0x36, 0x4a, 0x36, 0x86, 0xa1, 0x3d, 0x38, 0xea, 0xc6, 0x0a, 0x4e, 0x11,
	0xee, 0x06, 0xc6, 0x5a, 0x15, 0xf8, 0xa3, 0x47, 0xb2, 0x70, 0x03, 0x1d, 0xc0, 0x21, 0x57, 0xce,
	0xc4, 0x83, 0x64, 0x6b, 0x1a, 0xed, 0x63, 0x7c, 0xe9, 0x25, 0xe1, 0xb9, 0x72, 0x99, 0xef, 0x4d,
	0x7f, 0x00, 0x78, 0x2f, 0x98, 0x78, 0xcb, 0xac, 0x16, 0xc5, 0x7f, 0x12, 0xde, 0x29, 0x80, 0xf7,
	0x7f, 0xf7, 0x7d, 0x8d, 0x08, 0x3f, 0xc1, 0x88, 0x2b, 0x97, 0xd7, 0x1d, 0x45, 0x48, 0xf2, 0xf9,
	0xd5, 0x92, 0xec, 0xf4, 0x33, 0xc8, 0xcf, 0xac, 0xa4, 0x5f, 0x00, 0x7c, 0x78, 0xa8, 0x19, 0xb5,
	0x6c, 0xae, 0xdc, 0x7b, 0xcd, 0x4a, 0x51, 0xb4, 0xc3, 0x9e, 0x65, 0x4b, 0xe1, 0xb8, 0x95, 0x55,
	0x3d, 0x12, 0x03, 0x2f, 0xfd, 0xe2, 0x6a, 0xd2, 0x3d, 0x75, 0xb6, 0xc3, 0xcf, 0x29, 0xa5, 0xa7,
	0x03, 0xb8, 0x1b, 0x02, 0xda, 0x60, 0xe0, 0x9f, 0x5e, 0x2e, 0x4a, 0x60, 0xc4, 0x35, 0x6d, 0x5c,
	0x45, 0xb5, 0xb0, 0x27, 0xf1, 0x30, 0x01, 0xd3, 0xad, 0xec, 0xd7, 0x12, 0xda, 0x85, 0xb7, 0x6a,
	0x59, 0xb2, 0x2a, 0x17, 0x65, 0x7c, 0xc3, 0x13, 0xdc, 0xf4, 0xe7, 0xa3, 0x12, 0x3d, 0x86, 0x77,
	0xfa, 0x5c, 0x5b, 0x7c, 0xe4, 0xf1, 0xed, 0xbe, 0x78, 0x54, 0xa6, 0xdf, 0x00, 0x9c, 0x6c, 0x4a,
	0xe7, 0x1a, 0xbf, 0xd0, 0x86, 0x5d, 0x0e, 0xfe, 0xee, 0x2e, 0x0f, 0x5e, 0x7f, 0x7e, 0xc5, 0x85,
	0x0d, 0xb1, 0x93, 0x9e, 0xf5, 0x29, 0x15, 0xa4, 0x7d, 0x33, 0x2f, 0x7a, 0x3f, 0x17, 0x23, 0xff,
	0x46, 0x3e, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x9d, 0x1b, 0x99, 0xf9, 0x05, 0x00, 0x00,
}
