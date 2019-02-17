// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

/*
Package usermgr is a generated protocol buffer package.

It is generated from these files:
	usermgr/v1/micro/usermgr.proto

It has these top-level messages:
	User
	UserAttribute
	UserAttributes
	RegisterRequest
	LoginRequest
	LoginResponse
	AuthenticationResult
	RefreshToken
	AccessToken
	ConfirmRegistrationRequst
	ForgotPasswordRequst
	ConfirmPasswordRequst
	ChangePasswordRequst
	ChangeEmailRequst
	VerifyEmailRequst
	UpdateAttributesRequest
*/
package usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_proto "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = common_proto.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserMgr service

type UserMgrService interface {
	// Register Create a new user
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*common_proto.Empty, error)
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// Logout need verify permission , disable RefreshToken , access_token still work for 2 hours.
	Logout(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*common_proto.Empty, error)
	// RefreshToken reset token last access token
	RefreshSession(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*AuthenticationResult, error)
	ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequst, opts ...client.CallOption) (*common_proto.Empty, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error)
	ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error)
	ChangePasword(ctx context.Context, in *ChangePasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error)
	UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, opts ...client.CallOption) (*User, error)
	ChangeEmail(ctx context.Context, in *ChangeEmailRequst, opts ...client.CallOption) (*common_proto.Empty, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequst, opts ...client.CallOption) (*User, error)
	VerifyAccessToken(ctx context.Context, in *AccessToken, opts ...client.CallOption) (*User, error)
}

type userMgrService struct {
	c    client.Client
	name string
}

func NewUserMgrService(name string, c client.Client) UserMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "usermgr"
	}
	return &userMgrService{
		c:    c,
		name: name,
	}
}

func (c *userMgrService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Register", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Logout(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Logout", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) RefreshSession(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*AuthenticationResult, error) {
	req := c.c.NewRequest(c.name, "UserMgr.RefreshSession", in)
	out := new(AuthenticationResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequst, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ConfirmRegistration", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ForgotPassword(ctx context.Context, in *ForgotPasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ForgotPassword", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ConfirmPassword", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ChangePasword(ctx context.Context, in *ChangePasswordRequst, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ChangePasword", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.UpdateAttributes", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ChangeEmail(ctx context.Context, in *ChangeEmailRequst, opts ...client.CallOption) (*common_proto.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ChangeEmail", in)
	out := new(common_proto.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyEmail(ctx context.Context, in *VerifyEmailRequst, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyEmail", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyAccessToken(ctx context.Context, in *AccessToken, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyAccessToken", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrHandler interface {
	// Register Create a new user
	Register(context.Context, *RegisterRequest, *common_proto.Empty) error
	// Login login
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// Logout need verify permission , disable RefreshToken , access_token still work for 2 hours.
	Logout(context.Context, *RefreshToken, *common_proto.Empty) error
	// RefreshToken reset token last access token
	RefreshSession(context.Context, *RefreshToken, *AuthenticationResult) error
	ConfirmRegistration(context.Context, *ConfirmRegistrationRequst, *common_proto.Empty) error
	ForgotPassword(context.Context, *ForgotPasswordRequst, *common_proto.Empty) error
	ConfirmPassword(context.Context, *ConfirmPasswordRequst, *common_proto.Empty) error
	ChangePasword(context.Context, *ChangePasswordRequst, *common_proto.Empty) error
	UpdateAttributes(context.Context, *UpdateAttributesRequest, *User) error
	ChangeEmail(context.Context, *ChangeEmailRequst, *common_proto.Empty) error
	VerifyEmail(context.Context, *VerifyEmailRequst, *User) error
	VerifyAccessToken(context.Context, *AccessToken, *User) error
}

func RegisterUserMgrHandler(s server.Server, hdlr UserMgrHandler, opts ...server.HandlerOption) error {
	type userMgr interface {
		Register(ctx context.Context, in *RegisterRequest, out *common_proto.Empty) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *RefreshToken, out *common_proto.Empty) error
		RefreshSession(ctx context.Context, in *RefreshToken, out *AuthenticationResult) error
		ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequst, out *common_proto.Empty) error
		ForgotPassword(ctx context.Context, in *ForgotPasswordRequst, out *common_proto.Empty) error
		ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequst, out *common_proto.Empty) error
		ChangePasword(ctx context.Context, in *ChangePasswordRequst, out *common_proto.Empty) error
		UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, out *User) error
		ChangeEmail(ctx context.Context, in *ChangeEmailRequst, out *common_proto.Empty) error
		VerifyEmail(ctx context.Context, in *VerifyEmailRequst, out *User) error
		VerifyAccessToken(ctx context.Context, in *AccessToken, out *User) error
	}
	type UserMgr struct {
		userMgr
	}
	h := &userMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&UserMgr{h}, opts...))
}

type userMgrHandler struct {
	UserMgrHandler
}

func (h *userMgrHandler) Register(ctx context.Context, in *RegisterRequest, out *common_proto.Empty) error {
	return h.UserMgrHandler.Register(ctx, in, out)
}

func (h *userMgrHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserMgrHandler.Login(ctx, in, out)
}

func (h *userMgrHandler) Logout(ctx context.Context, in *RefreshToken, out *common_proto.Empty) error {
	return h.UserMgrHandler.Logout(ctx, in, out)
}

func (h *userMgrHandler) RefreshSession(ctx context.Context, in *RefreshToken, out *AuthenticationResult) error {
	return h.UserMgrHandler.RefreshSession(ctx, in, out)
}

func (h *userMgrHandler) ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequst, out *common_proto.Empty) error {
	return h.UserMgrHandler.ConfirmRegistration(ctx, in, out)
}

func (h *userMgrHandler) ForgotPassword(ctx context.Context, in *ForgotPasswordRequst, out *common_proto.Empty) error {
	return h.UserMgrHandler.ForgotPassword(ctx, in, out)
}

func (h *userMgrHandler) ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequst, out *common_proto.Empty) error {
	return h.UserMgrHandler.ConfirmPassword(ctx, in, out)
}

func (h *userMgrHandler) ChangePasword(ctx context.Context, in *ChangePasswordRequst, out *common_proto.Empty) error {
	return h.UserMgrHandler.ChangePasword(ctx, in, out)
}

func (h *userMgrHandler) UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, out *User) error {
	return h.UserMgrHandler.UpdateAttributes(ctx, in, out)
}

func (h *userMgrHandler) ChangeEmail(ctx context.Context, in *ChangeEmailRequst, out *common_proto.Empty) error {
	return h.UserMgrHandler.ChangeEmail(ctx, in, out)
}

func (h *userMgrHandler) VerifyEmail(ctx context.Context, in *VerifyEmailRequst, out *User) error {
	return h.UserMgrHandler.VerifyEmail(ctx, in, out)
}

func (h *userMgrHandler) VerifyAccessToken(ctx context.Context, in *AccessToken, out *User) error {
	return h.UserMgrHandler.VerifyAccessToken(ctx, in, out)
}
