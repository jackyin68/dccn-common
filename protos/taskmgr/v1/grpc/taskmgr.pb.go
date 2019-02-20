// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskmgr/v1/grpc/taskmgr.proto

package taskmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{0}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{1}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{2}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{3}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{4}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{5}
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
	return fileDescriptor_taskmgr_62a49eb60e864133, []int{6}
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

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "taskmgr.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "taskmgr.CreateTaskResponse")
	proto.RegisterType((*TaskListRequest)(nil), "taskmgr.TaskListRequest")
	proto.RegisterType((*TaskListResponse)(nil), "taskmgr.TaskListResponse")
	proto.RegisterType((*TaskFilter)(nil), "taskmgr.TaskFilter")
	proto.RegisterType((*TaskID)(nil), "taskmgr.TaskID")
	proto.RegisterType((*UpdateTaskRequest)(nil), "taskmgr.UpdateTaskRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskMgrClient is the client API for TaskMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskMgrClient interface {
	// Sends request to start a task and list task
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type taskMgrClient struct {
	cc *grpc.ClientConn
}

func NewTaskMgrClient(cc *grpc.ClientConn) TaskMgrClient {
	return &taskMgrClient{cc}
}

func (c *taskMgrClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/PurgeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskMgrServer is the server API for TaskMgr service.
type TaskMgrServer interface {
	// Sends request to start a task and list task
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
	CancelTask(context.Context, *TaskID) (*common.Empty, error)
	PurgeTask(context.Context, *TaskID) (*common.Empty, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error)
}

func RegisterTaskMgrServer(s *grpc.Server, srv TaskMgrServer) {
	s.RegisterService(&_TaskMgr_serviceDesc, srv)
}

func _TaskMgr_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CancelTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_PurgeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).PurgeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/PurgeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).PurgeTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "taskmgr.TaskMgr",
	HandlerType: (*TaskMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskMgr_CreateTask_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskMgr_TaskList_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _TaskMgr_CancelTask_Handler,
		},
		{
			MethodName: "PurgeTask",
			Handler:    _TaskMgr_PurgeTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskMgr_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskmgr/v1/grpc/taskmgr.proto",
}

func init() {
	proto.RegisterFile("taskmgr/v1/grpc/taskmgr.proto", fileDescriptor_taskmgr_62a49eb60e864133)
}

var fileDescriptor_taskmgr_62a49eb60e864133 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xdb, 0xff, 0x5f, 0x53, 0x3b, 0x3d, 0xd4, 0x6e, 0x0e, 0xc6, 0x88, 0x50, 0x17, 0x94,
	0x5e, 0x4c, 0xb0, 0xed, 0xad, 0x82, 0x48, 0xab, 0x25, 0xa0, 0x20, 0x41, 0xcf, 0x12, 0x93, 0x35,
	0x84, 0x36, 0x2f, 0xee, 0x6e, 0x05, 0x3f, 0x9e, 0xdf, 0x4c, 0x76, 0xf3, 0x4e, 0x9b, 0x83, 0x9e,
	0x92, 0xe1, 0xf9, 0xcd, 0x33, 0x33, 0xcf, 0xc2, 0x29, 0x77, 0xd8, 0x2a, 0xf4, 0xa9, 0xf9, 0x79,
	0x65, 0xfa, 0x34, 0x71, 0xcd, 0xac, 0x36, 0x12, 0x1a, 0xf3, 0x18, 0x75, 0xb2, 0x52, 0x57, 0xdd,
	0x38, 0x0c, 0xe3, 0xc8, 0x4c, 0x3f, 0xa9, 0x8a, 0x67, 0x30, 0x98, 0x53, 0xe2, 0x70, 0xf2, 0xec,
	0xb0, 0x95, 0x4d, 0x3e, 0x36, 0x84, 0x71, 0x74, 0x01, 0x7b, 0xa2, 0x49, 0x6b, 0x0f, 0xdb, 0xa3,
	0xde, 0x18, 0x19, 0xd5, 0x0e, 0x43, 0x82, 0x52, 0xc7, 0x97, 0x80, 0xaa, 0xcd, 0x2c, 0x89, 0x23,
	0x46, 0xd0, 0x11, 0xc8, 0x91, 0xaf, 0x81, 0x27, 0x0d, 0xba, 0xb6, 0x22, 0x4a, 0xcb, 0xc3, 0x4b,
	0xe8, 0x0b, 0xf0, 0x21, 0x60, 0x3c, 0x9f, 0x34, 0x85, 0x9e, 0x64, 0xdf, 0x83, 0x35, 0x27, 0x34,
	0x1b, 0xa8, 0x1a, 0xf9, 0x05, 0x02, 0xbf, 0x97, 0x92, 0x0d, 0xbc, 0xf8, 0xc7, 0xd7, 0x70, 0x58,
	0x1a, 0x65, 0x53, 0x47, 0xb0, 0x2f, 0x08, 0xa6, 0xb5, 0x87, 0xff, 0x1b, 0x96, 0x4e, 0x01, 0x7c,
	0x0e, 0x50, 0xfa, 0x36, 0x6f, 0x7b, 0x06, 0x8a, 0xc0, 0xac, 0x45, 0x33, 0x32, 0x83, 0xc1, 0x4b,
	0xe2, 0xfd, 0x2d, 0xbc, 0xf1, 0xf7, 0x3f, 0xe8, 0x88, 0xf2, 0xd1, 0xa7, 0x68, 0x09, 0x50, 0x06,
	0x89, 0xf4, 0xe2, 0xfe, 0xad, 0xa7, 0xd1, 0x4f, 0x76, 0x6a, 0x69, 0x06, 0xb8, 0x85, 0x6e, 0xe1,
	0x20, 0x4f, 0x06, 0x69, 0xb5, 0x18, 0x2b, 0xa9, 0xeb, 0xc7, 0x3b, 0x94, 0xc2, 0x62, 0x0a, 0x30,
	0x77, 0x22, 0x97, 0xac, 0xe5, 0x2e, 0xfd, 0x1a, 0x6a, 0x2d, 0x74, 0xb5, 0x7e, 0xd0, 0x5d, 0x98,
	0xf0, 0x2f, 0xdc, 0x42, 0x13, 0xe8, 0x3e, 0x6d, 0xa8, 0x4f, 0x7e, 0xd5, 0x74, 0x03, 0x50, 0xe6,
	0x57, 0x39, 0x7b, 0x2b, 0xd4, 0x06, 0x83, 0x37, 0x45, 0x96, 0x93, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x08, 0x84, 0x6d, 0x9e, 0x03, 0x03, 0x00, 0x00,
}
