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
	UserId               string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Task                 *common.Task `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{0}
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

func (m *CreateTaskRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               string        `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Error                *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{1}
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

func (m *CreateTaskResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ID struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{2}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Request struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{3}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Request) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskListResponse struct {
	Tasks                []*common.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Error                *common.Error  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskListResponse) Reset()         { *m = TaskListResponse{} }
func (m *TaskListResponse) String() string { return proto.CompactTextString(m) }
func (*TaskListResponse) ProtoMessage()    {}
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{4}
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

func (m *TaskListResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type UpdateTaskRequest struct {
	UserId               string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Task                 *common.Task `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{5}
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

func (m *UpdateTaskRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskDetailResponse struct {
	Task                 *common.Task  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Error                *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TaskDetailResponse) Reset()         { *m = TaskDetailResponse{} }
func (m *TaskDetailResponse) String() string { return proto.CompactTextString(m) }
func (*TaskDetailResponse) ProtoMessage()    {}
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_82f3ca707276e54d, []int{6}
}
func (m *TaskDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDetailResponse.Unmarshal(m, b)
}
func (m *TaskDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDetailResponse.Marshal(b, m, deterministic)
}
func (dst *TaskDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDetailResponse.Merge(dst, src)
}
func (m *TaskDetailResponse) XXX_Size() int {
	return xxx_messageInfo_TaskDetailResponse.Size(m)
}
func (m *TaskDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDetailResponse proto.InternalMessageInfo

func (m *TaskDetailResponse) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TaskDetailResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "taskmgr.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "taskmgr.CreateTaskResponse")
	proto.RegisterType((*ID)(nil), "taskmgr.ID")
	proto.RegisterType((*Request)(nil), "taskmgr.Request")
	proto.RegisterType((*TaskListResponse)(nil), "taskmgr.TaskListResponse")
	proto.RegisterType((*UpdateTaskRequest)(nil), "taskmgr.UpdateTaskRequest")
	proto.RegisterType((*TaskDetailResponse)(nil), "taskmgr.TaskDetailResponse")
}

func init() {
	proto.RegisterFile("taskmgr/v1/micro/taskmgr.proto", fileDescriptor_taskmgr_82f3ca707276e54d)
}

var fileDescriptor_taskmgr_82f3ca707276e54d = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xe5, 0x83, 0x7e, 0x0c, 0x17, 0x18, 0x0e, 0x62, 0x89, 0x86, 0xec, 0xc1, 0xe0, 0xa5,
	0x8d, 0xa8, 0x27, 0x63, 0x3c, 0x80, 0x31, 0x4d, 0x34, 0x31, 0x04, 0x13, 0x6f, 0xa6, 0xd2, 0x4d,
	0xd3, 0x40, 0x59, 0xdc, 0x5d, 0xfc, 0xe7, 0xde, 0xcd, 0xb6, 0xb4, 0x5d, 0x04, 0x24, 0x24, 0x9e,
	0x9a, 0xd7, 0x37, 0xf3, 0xe6, 0xbd, 0x99, 0x85, 0x53, 0xe9, 0x8b, 0x69, 0x1c, 0x72, 0xf7, 0xf3,
	0xc2, 0x8d, 0xa3, 0x09, 0x67, 0xee, 0xea, 0x87, 0xb3, 0xe0, 0x4c, 0x32, 0xb4, 0x56, 0xd0, 0x6e,
	0x4d, 0x58, 0x1c, 0xb3, 0xb9, 0x9b, 0x7e, 0x52, 0x96, 0x8c, 0xa1, 0x39, 0xe0, 0xd4, 0x97, 0x74,
	0xec, 0x8b, 0xe9, 0x88, 0x7e, 0x2c, 0xa9, 0x90, 0x78, 0x04, 0xd6, 0x52, 0x50, 0xfe, 0x16, 0x05,
	0x6d, 0xa3, 0x6b, 0xf4, 0x6a, 0xa3, 0xaa, 0x82, 0x5e, 0x80, 0x67, 0xf0, 0x4f, 0xa9, 0xb5, 0xcd,
	0xae, 0xd1, 0xab, 0xf7, 0xd1, 0xd1, 0xa5, 0x9c, 0x44, 0x21, 0xe1, 0xc9, 0x2b, 0xa0, 0xae, 0x2a,
	0x16, 0x6c, 0x2e, 0xa8, 0x92, 0x55, 0xac, 0x26, 0xab, 0xa0, 0x17, 0xe0, 0x39, 0x54, 0x28, 0xe7,
	0x8c, 0xaf, 0x74, 0x5b, 0xeb, 0xba, 0xf7, 0x8a, 0x1a, 0xa5, 0x15, 0xe4, 0x04, 0x4c, 0x6f, 0xb8,
	0xd3, 0x20, 0xb9, 0x01, 0x6b, 0x6f, 0x08, 0xcd, 0x86, 0xa9, 0xdb, 0x20, 0x21, 0x34, 0x94, 0xdf,
	0xc7, 0x48, 0xc8, 0xdc, 0x73, 0x0f, 0x2a, 0x8a, 0x15, 0x6d, 0xa3, 0x5b, 0xde, 0x11, 0x39, 0x2d,
	0x28, 0x42, 0x94, 0xf7, 0x86, 0x18, 0x43, 0xf3, 0x65, 0x11, 0xfc, 0xf5, 0xd2, 0x43, 0x40, 0x85,
	0x86, 0x54, 0xfa, 0xd1, 0x2c, 0x0f, 0x90, 0x75, 0x1b, 0xbf, 0x77, 0x1f, 0x70, 0x83, 0xfe, 0x97,
	0x09, 0x96, 0xea, 0x7c, 0x0a, 0x39, 0x3e, 0x00, 0x14, 0x97, 0x46, 0xdb, 0xc9, 0xde, 0xde, 0xc6,
	0xa3, 0xb2, 0x3b, 0x5b, 0xb9, 0xd4, 0x25, 0x29, 0xe1, 0x15, 0xfc, 0xcf, 0x96, 0x8f, 0xf5, 0xbc,
	0xd4, 0x1b, 0xda, 0xc7, 0x39, 0xf8, 0x79, 0x1c, 0x52, 0xc2, 0x6b, 0x80, 0x81, 0x3f, 0x9f, 0xd0,
	0x59, 0x32, 0xbe, 0x91, 0x97, 0x66, 0x43, 0xb7, 0xc5, 0x48, 0x86, 0xd5, 0x9e, 0x97, 0x3c, 0xa4,
	0x87, 0x75, 0xdd, 0x02, 0x14, 0x0b, 0xde, 0xd2, 0xd6, 0x59, 0x73, 0xba, 0x7e, 0x07, 0x52, 0xc2,
	0x3b, 0x80, 0xe2, 0xea, 0xda, 0xaa, 0x36, 0x9e, 0xc2, 0x8e, 0xf9, 0xef, 0xd5, 0x04, 0x5e, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x29, 0x16, 0xfa, 0xc2, 0xf2, 0x03, 0x00, 0x00,
}
