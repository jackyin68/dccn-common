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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{0}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{1}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{2}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{3}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{4}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{5}
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
	return fileDescriptor_taskmgr_a7a0f0242b09c0e3, []int{6}
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
	TaskList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskListResponse, error)
	CancelTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*common.Error, error)
	PurgeTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*common.Error, error)
	TaskDetail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Error, error)
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

func (c *taskMgrClient) TaskList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) CancelTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) PurgeTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/PurgeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskDetail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, "/taskmgr.TaskMgr/TaskDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
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
	TaskList(context.Context, *ID) (*TaskListResponse, error)
	CancelTask(context.Context, *Request) (*common.Error, error)
	PurgeTask(context.Context, *Request) (*common.Error, error)
	TaskDetail(context.Context, *Request) (*TaskDetailResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Error, error)
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
	in := new(ID)
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
		return srv.(TaskMgrServer).TaskList(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
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
		return srv.(TaskMgrServer).CancelTask(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_PurgeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
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
		return srv.(TaskMgrServer).PurgeTask(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskmgr.TaskMgr/TaskDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskDetail(ctx, req.(*Request))
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
			MethodName: "TaskDetail",
			Handler:    _TaskMgr_TaskDetail_Handler,
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
	proto.RegisterFile("taskmgr/v1/grpc/taskmgr.proto", fileDescriptor_taskmgr_a7a0f0242b09c0e3)
}

var fileDescriptor_taskmgr_a7a0f0242b09c0e3 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcf, 0x6b, 0xf2, 0x40,
	0x10, 0x35, 0xf1, 0xd3, 0x7c, 0x8e, 0x17, 0x1d, 0x0f, 0xb5, 0x11, 0x41, 0xf6, 0x50, 0xec, 0x25,
	0xa1, 0xb6, 0x3d, 0x95, 0xd2, 0x83, 0x96, 0x12, 0x68, 0xa1, 0x88, 0x85, 0xde, 0x4a, 0x6a, 0x96,
	0x20, 0xfe, 0xd8, 0x74, 0x77, 0xed, 0x7f, 0xde, 0x7b, 0xd9, 0xc4, 0x24, 0x6b, 0x35, 0x15, 0xa1,
	0xa7, 0xe5, 0xf1, 0x66, 0xde, 0xbc, 0x37, 0xb3, 0xd0, 0x95, 0xbe, 0x98, 0x2f, 0x43, 0xee, 0x7e,
	0x5e, 0xb8, 0x21, 0x8f, 0xa6, 0xee, 0x06, 0x3b, 0x11, 0x67, 0x92, 0xa1, 0xb5, 0x81, 0x76, 0x6b,
	0xca, 0x96, 0x4b, 0xb6, 0x72, 0x93, 0x27, 0x61, 0xc9, 0x04, 0x9a, 0x43, 0x4e, 0x7d, 0x49, 0x27,
	0xbe, 0x98, 0x8f, 0xe9, 0xc7, 0x9a, 0x0a, 0x89, 0x27, 0x60, 0xad, 0x05, 0xe5, 0x6f, 0xb3, 0xa0,
	0x6d, 0xf4, 0x8c, 0x7e, 0x6d, 0x5c, 0x55, 0xd0, 0x0b, 0xf0, 0x0c, 0xfe, 0x29, 0xb5, 0xb6, 0xd9,
	0x33, 0xfa, 0xf5, 0x01, 0x3a, 0xba, 0x94, 0x13, 0x2b, 0xc4, 0x3c, 0x79, 0x05, 0xd4, 0x55, 0x45,
	0xc4, 0x56, 0x82, 0x2a, 0x59, 0xc5, 0x6a, 0xb2, 0x0a, 0x7a, 0x01, 0x9e, 0x43, 0x85, 0x72, 0xce,
	0xf8, 0x46, 0xb7, 0xb5, 0xad, 0x7b, 0xaf, 0xa8, 0x71, 0x52, 0x41, 0xba, 0x60, 0x7a, 0xa3, 0x42,
	0x83, 0xe4, 0x06, 0xac, 0x83, 0x21, 0x34, 0x1b, 0xa6, 0x6e, 0x83, 0x84, 0xd0, 0x50, 0x7e, 0x1f,
	0x67, 0x42, 0x66, 0x9e, 0xfb, 0x50, 0x51, 0xac, 0x68, 0x1b, 0xbd, 0x72, 0x41, 0xe4, 0xa4, 0x20,
	0x0f, 0x51, 0x3e, 0x18, 0x62, 0x02, 0xcd, 0x97, 0x28, 0xf8, 0xeb, 0xa5, 0x87, 0x80, 0x0a, 0x8d,
	0xa8, 0xf4, 0x67, 0x8b, 0x2c, 0x40, 0xda, 0x6d, 0xfc, 0xde, 0x7d, 0xc4, 0x0d, 0x06, 0x5f, 0x26,
	0x58, 0xaa, 0xf3, 0x29, 0xe4, 0xf8, 0x00, 0x90, 0x5f, 0x1a, 0x6d, 0x27, 0xfd, 0x7b, 0x3b, 0x9f,
	0xca, 0xee, 0xec, 0xe5, 0x12, 0x97, 0xa4, 0x84, 0x57, 0xf0, 0x3f, 0x5d, 0x3e, 0xd6, 0xb3, 0x52,
	0x6f, 0x64, 0x9f, 0x66, 0xe0, 0xe7, 0x71, 0x48, 0x09, 0xaf, 0x01, 0x86, 0xfe, 0x6a, 0x4a, 0x17,
	0xf1, 0xf8, 0x46, 0x56, 0x9a, 0x0e, 0xdd, 0x17, 0x23, 0x1e, 0x56, 0x7b, 0x5e, 0xf3, 0x90, 0x1e,
	0xd7, 0x75, 0x0b, 0x90, 0x2f, 0x78, 0x4f, 0x5b, 0x67, 0xcb, 0xe9, 0xf6, 0x1d, 0x48, 0x09, 0xef,
	0x00, 0xf2, 0xab, 0x6b, 0xab, 0xda, 0xf9, 0x0a, 0x05, 0xf3, 0xdf, 0xab, 0x31, 0xbc, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0xe7, 0x3e, 0x2b, 0xf1, 0x03, 0x00, 0x00,
}
