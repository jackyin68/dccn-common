// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskmgr/v1/micro/taskmgr.proto

package taskmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The dccn client request message containing the user's token
type CreateTaskRequest struct {
	Task                 *common.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{0}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{1}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskListRequest struct {
	TaskFilter           *TaskFilter `protobuf:"bytes,1,opt,name=task_filter,json=taskFilter,proto3" json:"task_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskListRequest) Reset()         { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()    {}
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{2}
}
func (m *TaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListRequest.Unmarshal(m, b)
}
func (m *TaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListRequest.Marshal(b, m, deterministic)
}
func (dst *TaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListRequest.Merge(dst, src)
}
func (m *TaskListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskListRequest.Size(m)
}
func (m *TaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListRequest proto.InternalMessageInfo

func (m *TaskListRequest) GetTaskFilter() *TaskFilter {
	if m != nil {
		return m.TaskFilter
	}
	return nil
}

type TaskListResponse struct {
	Tasks                []*common.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskListResponse) Reset()         { *m = TaskListResponse{} }
func (m *TaskListResponse) String() string { return proto.CompactTextString(m) }
func (*TaskListResponse) ProtoMessage()    {}
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{3}
}
func (m *TaskListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListResponse.Unmarshal(m, b)
}
func (m *TaskListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListResponse.Marshal(b, m, deterministic)
}
func (dst *TaskListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListResponse.Merge(dst, src)
}
func (m *TaskListResponse) XXX_Size() int {
	return xxx_messageInfo_TaskListResponse.Size(m)
}
func (m *TaskListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListResponse proto.InternalMessageInfo

func (m *TaskListResponse) GetTasks() []*common.Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskFilter struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskFilter) Reset()         { *m = TaskFilter{} }
func (m *TaskFilter) String() string { return proto.CompactTextString(m) }
func (*TaskFilter) ProtoMessage()    {}
func (*TaskFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{4}
}
func (m *TaskFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskFilter.Unmarshal(m, b)
}
func (m *TaskFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskFilter.Marshal(b, m, deterministic)
}
func (dst *TaskFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskFilter.Merge(dst, src)
}
func (m *TaskFilter) XXX_Size() int {
	return xxx_messageInfo_TaskFilter.Size(m)
}
func (m *TaskFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TaskFilter proto.InternalMessageInfo

func (m *TaskFilter) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskID struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskID) Reset()         { *m = TaskID{} }
func (m *TaskID) String() string { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()    {}
func (*TaskID) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{5}
}
func (m *TaskID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskID.Unmarshal(m, b)
}
func (m *TaskID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskID.Marshal(b, m, deterministic)
}
func (dst *TaskID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskID.Merge(dst, src)
}
func (m *TaskID) XXX_Size() int {
	return xxx_messageInfo_TaskID.Size(m)
}
func (m *TaskID) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskID.DiscardUnknown(m)
}

var xxx_messageInfo_TaskID proto.InternalMessageInfo

func (m *TaskID) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type UpdateTaskRequest struct {
	Task                 *common.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{6}
}
func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(dst, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskOverviewResponse struct {
	ClusterCount         int32    `protobuf:"varint,1,opt,name=cluster_count,json=clusterCount,proto3" json:"cluster_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,2,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	RegionCount          int32    `protobuf:"varint,3,opt,name=region_count,json=regionCount,proto3" json:"region_count,omitempty"`
	TotalTaskCount       int32    `protobuf:"varint,4,opt,name=total_task_count,json=totalTaskCount,proto3" json:"total_task_count,omitempty"`
	HealthTaskCount      int32    `protobuf:"varint,5,opt,name=health_task_count,json=healthTaskCount,proto3" json:"health_task_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskOverviewResponse) Reset()         { *m = TaskOverviewResponse{} }
func (m *TaskOverviewResponse) String() string { return proto.CompactTextString(m) }
func (*TaskOverviewResponse) ProtoMessage()    {}
func (*TaskOverviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{7}
}
func (m *TaskOverviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskOverviewResponse.Unmarshal(m, b)
}
func (m *TaskOverviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskOverviewResponse.Marshal(b, m, deterministic)
}
func (dst *TaskOverviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskOverviewResponse.Merge(dst, src)
}
func (m *TaskOverviewResponse) XXX_Size() int {
	return xxx_messageInfo_TaskOverviewResponse.Size(m)
}
func (m *TaskOverviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskOverviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskOverviewResponse proto.InternalMessageInfo

func (m *TaskOverviewResponse) GetClusterCount() int32 {
	if m != nil {
		return m.ClusterCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetRegionCount() int32 {
	if m != nil {
		return m.RegionCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetTotalTaskCount() int32 {
	if m != nil {
		return m.TotalTaskCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetHealthTaskCount() int32 {
	if m != nil {
		return m.HealthTaskCount
	}
	return 0
}

type TaskLeaderBoardResponse struct {
	List                 []*TaskLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskLeaderBoardResponse) Reset()         { *m = TaskLeaderBoardResponse{} }
func (m *TaskLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardResponse) ProtoMessage()    {}
func (*TaskLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{8}
}
func (m *TaskLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardResponse.Unmarshal(m, b)
}
func (m *TaskLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (dst *TaskLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardResponse.Merge(dst, src)
}
func (m *TaskLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardResponse.Size(m)
}
func (m *TaskLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardResponse proto.InternalMessageInfo

func (m *TaskLeaderBoardResponse) GetList() []*TaskLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type TaskLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskLeaderBoardDetail) Reset()         { *m = TaskLeaderBoardDetail{} }
func (m *TaskLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardDetail) ProtoMessage()    {}
func (*TaskLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_7e528ef3f13a0b99, []int{9}
}
func (m *TaskLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardDetail.Unmarshal(m, b)
}
func (m *TaskLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (dst *TaskLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardDetail.Merge(dst, src)
}
func (m *TaskLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardDetail.Size(m)
}
func (m *TaskLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardDetail proto.InternalMessageInfo

func (m *TaskLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "taskmgr.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "taskmgr.CreateTaskResponse")
	proto.RegisterType((*TaskListRequest)(nil), "taskmgr.TaskListRequest")
	proto.RegisterType((*TaskListResponse)(nil), "taskmgr.TaskListResponse")
	proto.RegisterType((*TaskFilter)(nil), "taskmgr.TaskFilter")
	proto.RegisterType((*TaskID)(nil), "taskmgr.TaskID")
	proto.RegisterType((*UpdateTaskRequest)(nil), "taskmgr.UpdateTaskRequest")
	proto.RegisterType((*TaskOverviewResponse)(nil), "taskmgr.TaskOverviewResponse")
	proto.RegisterType((*TaskLeaderBoardResponse)(nil), "taskmgr.TaskLeaderBoardResponse")
	proto.RegisterType((*TaskLeaderBoardDetail)(nil), "taskmgr.TaskLeaderBoardDetail")
}

func init() {
	proto.RegisterFile("taskmgr/v1/micro/taskmgr.proto", fileDescriptor_taskmgr_7e528ef3f13a0b99)
}

var fileDescriptor_taskmgr_7e528ef3f13a0b99 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x6d, 0x6f, 0x12, 0x41,
	0x10, 0xc7, 0x41, 0x1e, 0x6a, 0x07, 0x14, 0x58, 0xd4, 0xe2, 0x19, 0x1b, 0xba, 0x46, 0x43, 0x34,
	0x42, 0xa4, 0x7d, 0x57, 0x13, 0xa3, 0xa0, 0x0d, 0xc6, 0x46, 0x73, 0xd1, 0xd7, 0x64, 0x0b, 0x23,
	0xbd, 0xf4, 0x6e, 0x17, 0xf7, 0x16, 0x8c, 0x9f, 0xd3, 0x4f, 0xe0, 0x37, 0x31, 0xfb, 0xc0, 0xdd,
	0x61, 0xb9, 0x17, 0xf6, 0x55, 0x6f, 0x67, 0x7e, 0xf3, 0xdf, 0x99, 0x9d, 0x7f, 0x81, 0x43, 0xc5,
	0xe2, 0xab, 0x68, 0x21, 0x07, 0xeb, 0x57, 0x83, 0x28, 0x98, 0x49, 0x31, 0x70, 0x81, 0xfe, 0x52,
	0x0a, 0x25, 0xc8, 0x9e, 0x3b, 0x7a, 0xed, 0x99, 0x88, 0x22, 0xc1, 0x07, 0xf6, 0x8f, 0xcd, 0xd2,
	0x53, 0x68, 0x8d, 0x24, 0x32, 0x85, 0x5f, 0x59, 0x7c, 0xe5, 0xe3, 0x8f, 0x15, 0xc6, 0x8a, 0x3c,
	0x83, 0xb2, 0x2e, 0xea, 0x14, 0xbb, 0xc5, 0x5e, 0x6d, 0x48, 0xfa, 0xd9, 0x8a, 0xbe, 0x01, 0x4d,
	0x9e, 0xbe, 0x04, 0x92, 0x2d, 0x8e, 0x97, 0x82, 0xc7, 0x48, 0x0e, 0xc0, 0x5c, 0x39, 0x0d, 0xe6,
	0x46, 0x60, 0xdf, 0xaf, 0xea, 0xe3, 0x64, 0x4e, 0xcf, 0xa0, 0xa1, 0xc1, 0x4f, 0x41, 0xac, 0x36,
	0x37, 0x9d, 0x40, 0xcd, 0xb0, 0xdf, 0x83, 0x50, 0xa1, 0x74, 0x17, 0xb6, 0xfb, 0x9b, 0x09, 0x34,
	0xfe, 0xc1, 0xa4, 0x7c, 0x50, 0xc9, 0x37, 0x7d, 0x0d, 0xcd, 0x54, 0xc8, 0xdd, 0xda, 0x83, 0x8a,
	0x26, 0xe2, 0x4e, 0xb1, 0x5b, 0xca, 0x69, 0xda, 0x02, 0xf4, 0x29, 0x40, 0xaa, 0x9b, 0xdf, 0xed,
	0x11, 0x54, 0x35, 0x36, 0x19, 0xe7, 0x23, 0xa7, 0xd0, 0xfa, 0xb6, 0x9c, 0xdf, 0xf0, 0xf1, 0xfe,
	0x14, 0xe1, 0x9e, 0x3e, 0x7e, 0x5e, 0xa3, 0x5c, 0x07, 0xf8, 0x33, 0x99, 0xe4, 0x09, 0xdc, 0x99,
	0x85, 0xab, 0x58, 0xa1, 0x9c, 0xce, 0xc4, 0x8a, 0x2b, 0xa3, 0x54, 0xf1, 0xeb, 0x2e, 0x38, 0xd2,
	0x31, 0xf2, 0x02, 0x5a, 0xc8, 0xd7, 0x81, 0x14, 0x3c, 0x42, 0xae, 0x1c, 0x78, 0xcb, 0x80, 0xcd,
	0x4c, 0xc2, 0xc2, 0x47, 0x50, 0x97, 0xb8, 0x08, 0x04, 0x77, 0x5c, 0xc9, 0x70, 0x35, 0x1b, 0xb3,
	0x48, 0x0f, 0x9a, 0x4a, 0x28, 0x16, 0x4e, 0xcd, 0xa4, 0x16, 0x2b, 0x1b, 0xec, 0xae, 0x89, 0xeb,
	0x4e, 0x2d, 0xf9, 0x1c, 0x5a, 0x97, 0xc8, 0x42, 0x75, 0x99, 0x45, 0x2b, 0x06, 0x6d, 0xd8, 0x44,
	0xc2, 0xd2, 0x73, 0x38, 0x30, 0x8b, 0x42, 0x36, 0x47, 0xf9, 0x4e, 0x30, 0x39, 0x4f, 0xa6, 0x1c,
	0x42, 0x39, 0x0c, 0x62, 0xe5, 0xd6, 0x75, 0xb8, 0xb5, 0xf2, 0x0c, 0x3f, 0x46, 0xc5, 0x82, 0xd0,
	0x37, 0x2c, 0x1d, 0xc1, 0xfd, 0x9d, 0x69, 0x42, 0xa0, 0xcc, 0x59, 0x84, 0x6e, 0x3d, 0xe6, 0x9b,
	0x3c, 0x80, 0x2a, 0x5f, 0x45, 0x17, 0x28, 0xcd, 0xb3, 0x14, 0x7d, 0x77, 0x1a, 0xfe, 0x2e, 0xc1,
	0x9e, 0x56, 0x39, 0x5f, 0x48, 0x72, 0x06, 0x90, 0x1a, 0x98, 0x78, 0x49, 0x13, 0xd7, 0xfe, 0x25,
	0xbc, 0x47, 0x3b, 0x73, 0x76, 0x16, 0x5a, 0x20, 0x6f, 0xe1, 0xf6, 0xc6, 0x91, 0xa4, 0xb3, 0x3d,
	0x4b, 0xea, 0x76, 0xef, 0xe1, 0x8e, 0x4c, 0x22, 0x71, 0x02, 0x30, 0x62, 0x7c, 0x86, 0xe6, 0xa9,
	0x49, 0x63, 0x0b, 0x9d, 0x8c, 0xbd, 0xf6, 0xb6, 0x91, 0xde, 0x47, 0x4b, 0xf5, 0x8b, 0x16, 0xc8,
	0x31, 0xec, 0x7f, 0x59, 0xc9, 0x05, 0xfe, 0x57, 0xd1, 0x1b, 0x80, 0xd4, 0xb7, 0x99, 0xb1, 0xaf,
	0x99, 0x39, 0x4f, 0x60, 0x0c, 0xf5, 0xac, 0x75, 0xc9, 0x2e, 0xcc, 0x7b, 0xbc, 0xd5, 0xcd, 0xbf,
	0x36, 0xa7, 0x05, 0xf2, 0xd1, 0xfd, 0x1e, 0xa4, 0xeb, 0xdc, 0x2d, 0xd4, 0xcd, 0x33, 0x47, 0xaa,
	0x75, 0x51, 0x35, 0x05, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x08, 0xeb, 0xde, 0x0e,
	0x05, 0x00, 0x00,
}
