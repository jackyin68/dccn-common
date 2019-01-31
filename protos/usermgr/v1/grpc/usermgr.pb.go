// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgr/v1/grpc/usermgr.proto

package usermgr

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

type User struct {
	// id self-increase, de
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name should be unique
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// nickname show on UI
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// email user's email, unique.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// password string
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// balance user's balance in account
	Balance int32 `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
	// is_deleted user's status
	IsDeleted bool `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// token jwt token string
	Token                string   `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *User) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{1}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string        `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Error                *common.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{2}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LogoutRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{3}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{4}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type NewTokenResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error                *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NewTokenResponse) Reset()         { *m = NewTokenResponse{} }
func (m *NewTokenResponse) String() string { return proto.CompactTextString(m) }
func (*NewTokenResponse) ProtoMessage()    {}
func (*NewTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_231594396ab6444e, []int{5}
}
func (m *NewTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTokenResponse.Unmarshal(m, b)
}
func (m *NewTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTokenResponse.Marshal(b, m, deterministic)
}
func (dst *NewTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTokenResponse.Merge(dst, src)
}
func (m *NewTokenResponse) XXX_Size() int {
	return xxx_messageInfo_NewTokenResponse.Size(m)
}
func (m *NewTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewTokenResponse proto.InternalMessageInfo

func (m *NewTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NewTokenResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "usermgr.User")
	proto.RegisterType((*LoginRequest)(nil), "usermgr.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "usermgr.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "usermgr.LogoutRequest")
	proto.RegisterType((*Token)(nil), "usermgr.Token")
	proto.RegisterType((*NewTokenResponse)(nil), "usermgr.NewTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserMgrClient is the client API for UserMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserMgrClient interface {
	// Register Create a new user
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*common.Error, error)
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Logout need verify permission
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Error, error)
	// Auth  validates user
	NewToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewTokenResponse, error)
	// VerifyToken Validated Token
	VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error)
}

type userMgrClient struct {
	cc *grpc.ClientConn
}

func NewUserMgrClient(cc *grpc.ClientConn) UserMgrClient {
	return &userMgrClient{cc}
}

func (c *userMgrClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) NewToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewTokenResponse, error) {
	out := new(NewTokenResponse)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/NewToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMgrServer is the server API for UserMgr service.
type UserMgrServer interface {
	// Register Create a new user
	Register(context.Context, *User) (*common.Error, error)
	// Login login
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Logout need verify permission
	Logout(context.Context, *LogoutRequest) (*common.Error, error)
	// Auth  validates user
	NewToken(context.Context, *User) (*NewTokenResponse, error)
	// VerifyToken Validated Token
	VerifyToken(context.Context, *Token) (*common.Error, error)
}

func RegisterUserMgrServer(s *grpc.Server, srv UserMgrServer) {
	s.RegisterService(&_UserMgr_serviceDesc, srv)
}

func _UserMgr_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_NewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).NewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/NewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).NewToken(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).VerifyToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usermgr.UserMgr",
	HandlerType: (*UserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserMgr_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserMgr_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserMgr_Logout_Handler,
		},
		{
			MethodName: "NewToken",
			Handler:    _UserMgr_NewToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _UserMgr_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermgr/v1/grpc/usermgr.proto",
}

func init() {
	proto.RegisterFile("usermgr/v1/grpc/usermgr.proto", fileDescriptor_usermgr_231594396ab6444e)
}

var fileDescriptor_usermgr_231594396ab6444e = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0x26, 0x4e, 0x9c, 0x53, 0x5a, 0xa1, 0xc3, 0x6d, 0x88, 0x14, 0x29, 0x9a, 0x95,
	0xd9, 0xd8, 0x50, 0x10, 0xb0, 0x64, 0x01, 0x0b, 0x24, 0x60, 0x61, 0x2e, 0xdb, 0xca, 0xb5, 0x0f,
	0xd6, 0xa8, 0xb1, 0xc7, 0xcc, 0x38, 0x54, 0xbc, 0x0b, 0xef, 0xc3, 0x6b, 0x55, 0x9e, 0xb1, 0x1d,
	0x37, 0x6a, 0x94, 0x9d, 0xbf, 0x73, 0xfb, 0x7f, 0xfd, 0x1e, 0x58, 0x6d, 0x0d, 0xe9, 0xb2, 0xd0,
	0xf1, 0x9f, 0x97, 0x71, 0xa1, 0xeb, 0x2c, 0xee, 0x38, 0xaa, 0xb5, 0x6a, 0x14, 0xce, 0x3b, 0x5c,
	0x62, 0xa6, 0xca, 0x52, 0x55, 0x31, 0x69, 0xad, 0xba, 0xa6, 0xf8, 0xcf, 0x60, 0xfa, 0xc3, 0x90,
	0xc6, 0x33, 0xf0, 0x64, 0xce, 0xd9, 0x9a, 0x85, 0x8b, 0xc4, 0x93, 0x39, 0x22, 0x4c, 0xab, 0xb4,
	0x24, 0xee, 0xd9, 0x8a, 0xfd, 0xc6, 0x25, 0x04, 0x95, 0xcc, 0xae, 0x6c, 0x7d, 0x6a, 0xeb, 0x03,
	0xe3, 0x23, 0xf0, 0xa9, 0x4c, 0xe5, 0x86, 0xfb, 0xb6, 0xe1, 0xa0, 0xdd, 0xa8, 0x53, 0x63, 0xae,
	0x95, 0xce, 0xf9, 0xcc, 0x6d, 0xf4, 0x8c, 0x1c, 0xe6, 0x97, 0xe9, 0x26, 0xad, 0x32, 0xe2, 0xf3,
	0x35, 0x0b, 0xfd, 0xa4, 0x47, 0x5c, 0x01, 0x48, 0x73, 0x91, 0xd3, 0x86, 0x1a, 0xca, 0x79, 0xb0,
	0x66, 0x61, 0x90, 0x2c, 0xa4, 0xf9, 0xe0, 0x0a, 0xad, 0x54, 0xa3, 0xae, 0xa8, 0xe2, 0x0b, 0x27,
	0x65, 0x41, 0xbc, 0x87, 0xfb, 0x9f, 0x55, 0x21, 0xab, 0x84, 0x7e, 0x6f, 0xc9, 0x34, 0x3b, 0x43,
	0xec, 0x90, 0x21, 0xef, 0xb6, 0x21, 0x21, 0xe1, 0xb4, 0xbb, 0x60, 0x6a, 0x55, 0x19, 0xda, 0x09,
	0xb1, 0x91, 0x10, 0x3e, 0x05, 0x9b, 0xe8, 0x85, 0xec, 0x2f, 0xcc, 0x5a, 0xfc, 0x94, 0xe3, 0x73,
	0xf0, 0x6d, 0xb4, 0xfc, 0xde, 0x9a, 0x85, 0x27, 0xe7, 0x0f, 0x23, 0x97, 0xb7, 0x4b, 0x3a, 0xfa,
	0xd8, 0xb6, 0x12, 0x37, 0x21, 0x42, 0x2b, 0xa5, 0xb6, 0x4d, 0xef, 0x76, 0x74, 0x94, 0x8d, 0x8f,
	0x8a, 0x15, 0xf8, 0xdf, 0xad, 0xec, 0x9d, 0x66, 0xc4, 0x37, 0x78, 0xf0, 0x95, 0xae, 0xed, 0xc4,
	0x11, 0xdb, 0x83, 0x3b, 0xef, 0x98, 0xbb, 0xf3, 0x7f, 0x1e, 0xcc, 0xdb, 0x47, 0xf1, 0xa5, 0xd0,
	0xf8, 0x02, 0x82, 0x84, 0x0a, 0x69, 0x1a, 0xd2, 0x78, 0x1a, 0xf5, 0x2f, 0xab, 0xed, 0x2e, 0xef,
	0x3a, 0x21, 0x26, 0xf8, 0x0e, 0x7c, 0x1b, 0x23, 0x3e, 0x1e, 0xc6, 0xc7, 0x3f, 0x66, 0xf9, 0x64,
	0xbf, 0xec, 0x6c, 0x8b, 0x09, 0xbe, 0x85, 0x99, 0x4b, 0x05, 0x6f, 0xcd, 0xec, 0x62, 0x3a, 0x24,
	0xf9, 0x06, 0x82, 0x3e, 0x85, 0x7d, 0x93, 0xcf, 0x06, 0xdc, 0xcf, 0x49, 0x4c, 0xf0, 0x35, 0x9c,
	0xfc, 0x24, 0x2d, 0x7f, 0xfd, 0x75, 0xab, 0x67, 0xc3, 0xac, 0xe5, 0x03, 0x6a, 0x97, 0x33, 0x8b,
	0xaf, 0x6e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe2, 0x18, 0xc4, 0x78, 0x03, 0x00, 0x00,
}