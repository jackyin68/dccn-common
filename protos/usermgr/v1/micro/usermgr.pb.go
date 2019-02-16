// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

package usermgr

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

type User struct {
	// id self-increase, de
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name should be unique
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// email user's email, unique.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// password string
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// balance user's balance in account
	Balance int32 `protobuf:"varint,5,opt,name=balance,proto3" json:"balance,omitempty"`
	// is_deleted user's status
	IsDeleted bool `protobuf:"varint,6,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// token jwt token string
	Token string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	// block chains public keys
	PublicKeys []string `protobuf:"bytes,8,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	// is_activation register if confirmed
	IsActivation bool `protobuf:"varint,9,opt,name=is_activation,json=isActivation,proto3" json:"is_activation,omitempty"`
	// user create date
	CreationDate         uint64   `protobuf:"varint,10,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{0}
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

func (m *User) GetPublicKeys() []string {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func (m *User) GetIsActivation() bool {
	if m != nil {
		return m.IsActivation
	}
	return false
}

func (m *User) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
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
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{1}
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
	User                 *User         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Error                *common.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{2}
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

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
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
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{3}
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

type AskResetPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskResetPasswordRequest) Reset()         { *m = AskResetPasswordRequest{} }
func (m *AskResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*AskResetPasswordRequest) ProtoMessage()    {}
func (*AskResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{4}
}
func (m *AskResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskResetPasswordRequest.Unmarshal(m, b)
}
func (m *AskResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskResetPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *AskResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskResetPasswordRequest.Merge(dst, src)
}
func (m *AskResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_AskResetPasswordRequest.Size(m)
}
func (m *AskResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskResetPasswordRequest proto.InternalMessageInfo

func (m *AskResetPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ResetPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordRequest) Reset()         { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordRequest) ProtoMessage()    {}
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{5}
}
func (m *ResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordRequest.Unmarshal(m, b)
}
func (m *ResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *ResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordRequest.Merge(dst, src)
}
func (m *ResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordRequest.Size(m)
}
func (m *ResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordRequest proto.InternalMessageInfo

func (m *ResetPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ActivateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateRequest) Reset()         { *m = ActivateRequest{} }
func (m *ActivateRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateRequest) ProtoMessage()    {}
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{6}
}
func (m *ActivateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateRequest.Unmarshal(m, b)
}
func (m *ActivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateRequest.Marshal(b, m, deterministic)
}
func (dst *ActivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateRequest.Merge(dst, src)
}
func (m *ActivateRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateRequest.Size(m)
}
func (m *ActivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateRequest proto.InternalMessageInfo

func (m *ActivateRequest) GetToken() string {
	if m != nil {
		return m.Token
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
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{7}
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
	return fileDescriptor_usermgr_dbe4cfba28e2f754, []int{8}
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
	proto.RegisterType((*AskResetPasswordRequest)(nil), "usermgr.AskResetPasswordRequest")
	proto.RegisterType((*ResetPasswordRequest)(nil), "usermgr.ResetPasswordRequest")
	proto.RegisterType((*ActivateRequest)(nil), "usermgr.ActivateRequest")
	proto.RegisterType((*Token)(nil), "usermgr.Token")
	proto.RegisterType((*NewTokenResponse)(nil), "usermgr.NewTokenResponse")
}

func init() {
	proto.RegisterFile("usermgr/v1/micro/usermgr.proto", fileDescriptor_usermgr_dbe4cfba28e2f754)
}

var fileDescriptor_usermgr_dbe4cfba28e2f754 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x5d, 0x3b, 0x97, 0x49, 0xda, 0xa2, 0x6d, 0x4b, 0x97, 0x48, 0x05, 0x63, 0x1e, 0x30,
	0x2f, 0x09, 0x94, 0xeb, 0x13, 0xa2, 0xa2, 0x14, 0x55, 0x5c, 0x84, 0xcc, 0xe5, 0x91, 0xc8, 0xb5,
	0x27, 0x61, 0x95, 0xd8, 0x1b, 0x76, 0x9d, 0x56, 0xf9, 0x11, 0x7e, 0x93, 0x5f, 0x40, 0xde, 0x75,
	0xec, 0x24, 0xd4, 0x2d, 0xe5, 0x29, 0x3e, 0x73, 0x39, 0xb3, 0x73, 0xce, 0x04, 0x6e, 0x4f, 0x25,
	0x8a, 0x78, 0x28, 0x7a, 0x67, 0x8f, 0x7a, 0x31, 0x0b, 0x05, 0xef, 0xe5, 0x81, 0xee, 0x44, 0xf0,
	0x94, 0x93, 0x7a, 0x0e, 0x3b, 0xdb, 0x21, 0x8f, 0x63, 0x9e, 0xf4, 0xf4, 0x8f, 0xce, 0xba, 0xbf,
	0x4c, 0xb0, 0xbe, 0x4a, 0x14, 0x64, 0x13, 0x4c, 0x16, 0x51, 0xc3, 0x31, 0xbc, 0xa6, 0x6f, 0xb2,
	0x88, 0x10, 0xb0, 0x92, 0x20, 0x46, 0x6a, 0xaa, 0x88, 0xfa, 0x26, 0x3b, 0x60, 0x63, 0x1c, 0xb0,
	0x31, 0x5d, 0x57, 0x41, 0x0d, 0x48, 0x07, 0x1a, 0x93, 0x40, 0xca, 0x73, 0x2e, 0x22, 0x6a, 0xa9,
	0x44, 0x81, 0x09, 0x85, 0xfa, 0x69, 0x30, 0x0e, 0x92, 0x10, 0xa9, 0xed, 0x18, 0x9e, 0xed, 0xcf,
	0x21, 0xd9, 0x07, 0x60, 0xb2, 0x1f, 0xe1, 0x18, 0x53, 0x8c, 0x68, 0xcd, 0x31, 0xbc, 0x86, 0xdf,
	0x64, 0xf2, 0x48, 0x07, 0xb2, 0x51, 0x29, 0x1f, 0x61, 0x42, 0xeb, 0x7a, 0x94, 0x02, 0xe4, 0x0e,
	0xb4, 0x26, 0xd3, 0xd3, 0x31, 0x0b, 0xfb, 0x23, 0x9c, 0x49, 0xda, 0x70, 0xd6, 0xbd, 0xa6, 0x0f,
	0x3a, 0xf4, 0x0e, 0x67, 0x92, 0xdc, 0x83, 0x0d, 0x26, 0xfb, 0x41, 0x98, 0xb2, 0xb3, 0x20, 0x65,
	0x3c, 0xa1, 0x4d, 0x45, 0xdc, 0x66, 0xf2, 0xb0, 0x88, 0x65, 0x45, 0xa1, 0x40, 0xf5, 0xdd, 0x8f,
	0x82, 0x14, 0x29, 0x38, 0x86, 0x67, 0xf9, 0xed, 0x79, 0xf0, 0x28, 0x48, 0xd1, 0x7d, 0x05, 0xed,
	0xf7, 0x7c, 0xc8, 0x12, 0x1f, 0x7f, 0x4e, 0x51, 0xa6, 0xe5, 0xee, 0x46, 0xd5, 0xee, 0xe6, 0xf2,
	0xee, 0xae, 0x84, 0x8d, 0x9c, 0x41, 0x4e, 0x78, 0x22, 0xb1, 0xdc, 0xc9, 0x58, 0xdc, 0xe9, 0x2e,
	0x58, 0x99, 0x43, 0xaa, 0xbd, 0x75, 0xb0, 0xd1, 0x9d, 0xbb, 0x97, 0xb9, 0xe2, 0xab, 0x14, 0x79,
	0x00, 0x36, 0x0a, 0xc1, 0x85, 0xd2, 0xbd, 0x75, 0xb0, 0xdd, 0x5d, 0xb4, 0xb0, 0xfb, 0x26, 0x4b,
	0xf9, 0xba, 0xc2, 0xf5, 0xd4, 0x50, 0x3e, 0x4d, 0xe7, 0xef, 0xde, 0x03, 0x75, 0x00, 0xfd, 0xc2,
	0xdc, 0x5a, 0x06, 0x4f, 0x22, 0xb7, 0x07, 0x7b, 0x87, 0x72, 0xe4, 0xa3, 0xc4, 0xf4, 0x53, 0xfe,
	0xe4, 0x4b, 0x77, 0x75, 0xbf, 0xc3, 0xce, 0xbf, 0x57, 0x5f, 0xa6, 0x4c, 0x29, 0xc4, 0xfa, 0x82,
	0x10, 0xee, 0x7d, 0xd8, 0xca, 0x4d, 0xc2, 0x05, 0xea, 0xbf, 0x15, 0x73, 0xf7, 0xc1, 0xfe, 0xa2,
	0xa4, 0xbb, 0x38, 0xfd, 0x19, 0x6e, 0x7c, 0xc4, 0x73, 0x55, 0x71, 0x85, 0xf4, 0x85, 0xae, 0xe6,
	0x55, 0xba, 0x1e, 0xfc, 0xb6, 0xa0, 0x9e, 0x39, 0xf2, 0x61, 0x28, 0xc8, 0x43, 0x68, 0xf8, 0x38,
	0x64, 0x32, 0x45, 0x41, 0x96, 0xfd, 0xea, 0x5c, 0x44, 0xe1, 0xae, 0x91, 0x17, 0x60, 0xab, 0x53,
	0x20, 0xbb, 0x45, 0xf9, 0xe2, 0x71, 0x75, 0x6e, 0xae, 0x86, 0xf5, 0xb3, 0xdd, 0x35, 0xf2, 0x1c,
	0x6a, 0xda, 0x4f, 0xb2, 0x54, 0x53, 0x1a, 0x5c, 0x35, 0xf2, 0x04, 0x36, 0x8f, 0xb9, 0x18, 0x96,
	0x76, 0x11, 0xa7, 0x20, 0xa8, 0xf0, 0xbd, 0x8a, 0xea, 0x2d, 0x6c, 0xbd, 0xe6, 0xc9, 0x80, 0x89,
	0xb8, 0xe0, 0xda, 0x2f, 0xb8, 0xae, 0x43, 0x74, 0x0c, 0xdb, 0x39, 0x91, 0xd6, 0x4f, 0xe8, 0xff,
	0x23, 0x2d, 0x1f, 0xb6, 0xec, 0x7f, 0x15, 0xcf, 0x33, 0x68, 0xcc, 0x1d, 0x5e, 0x35, 0xe0, 0x56,
	0x01, 0x57, 0x6f, 0xc0, 0x5d, 0x23, 0x4f, 0xa0, 0xf5, 0x0d, 0x05, 0x1b, 0xcc, 0x74, 0xeb, 0x66,
	0x51, 0xab, 0x70, 0xd5, 0xb4, 0x97, 0xb0, 0xab, 0xbb, 0x0e, 0x93, 0xc8, 0xc7, 0x81, 0x40, 0xf9,
	0xe3, 0x5a, 0xfd, 0x4f, 0xa1, 0xfd, 0x1f, 0x6d, 0xa7, 0x35, 0x05, 0x1f, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x77, 0xbe, 0xcd, 0x48, 0xe0, 0x05, 0x00, 0x00,
}
