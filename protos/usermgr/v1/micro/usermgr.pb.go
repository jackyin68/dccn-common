// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

package usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Ankr-network/dccn-common/protos/common"

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
	// uuid
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// email user's email, unique.
	Email                string          `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Attributes           *UserAttributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{0}
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

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAttributes() *UserAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type UserAttribute struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*UserAttribute_IntValue
	//	*UserAttribute_DoubleValue
	//	*UserAttribute_StringValue
	//	*UserAttribute_BoolValue
	Value                isUserAttribute_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserAttribute) Reset()         { *m = UserAttribute{} }
func (m *UserAttribute) String() string { return proto.CompactTextString(m) }
func (*UserAttribute) ProtoMessage()    {}
func (*UserAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{1}
}
func (m *UserAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAttribute.Unmarshal(m, b)
}
func (m *UserAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAttribute.Marshal(b, m, deterministic)
}
func (dst *UserAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAttribute.Merge(dst, src)
}
func (m *UserAttribute) XXX_Size() int {
	return xxx_messageInfo_UserAttribute.Size(m)
}
func (m *UserAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_UserAttribute proto.InternalMessageInfo

func (m *UserAttribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isUserAttribute_Value interface {
	isUserAttribute_Value()
}

type UserAttribute_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type UserAttribute_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type UserAttribute_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type UserAttribute_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

func (*UserAttribute_IntValue) isUserAttribute_Value() {}

func (*UserAttribute_DoubleValue) isUserAttribute_Value() {}

func (*UserAttribute_StringValue) isUserAttribute_Value() {}

func (*UserAttribute_BoolValue) isUserAttribute_Value() {}

func (m *UserAttribute) GetValue() isUserAttribute_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UserAttribute) GetIntValue() int64 {
	if x, ok := m.GetValue().(*UserAttribute_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *UserAttribute) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*UserAttribute_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *UserAttribute) GetStringValue() string {
	if x, ok := m.GetValue().(*UserAttribute_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *UserAttribute) GetBoolValue() bool {
	if x, ok := m.GetValue().(*UserAttribute_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserAttribute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserAttribute_OneofMarshaler, _UserAttribute_OneofUnmarshaler, _UserAttribute_OneofSizer, []interface{}{
		(*UserAttribute_IntValue)(nil),
		(*UserAttribute_DoubleValue)(nil),
		(*UserAttribute_StringValue)(nil),
		(*UserAttribute_BoolValue)(nil),
	}
}

func _UserAttribute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UserAttribute)
	// value
	switch x := m.Value.(type) {
	case *UserAttribute_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *UserAttribute_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *UserAttribute_StringValue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *UserAttribute_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("UserAttribute.Value has unexpected type %T", x)
	}
	return nil
}

func _UserAttribute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UserAttribute)
	switch tag {
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &UserAttribute_IntValue{int64(x)}
		return true, err
	case 3: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &UserAttribute_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 4: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &UserAttribute_StringValue{x}
		return true, err
	case 5: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &UserAttribute_BoolValue{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _UserAttribute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UserAttribute)
	// value
	switch x := m.Value.(type) {
	case *UserAttribute_IntValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntValue))
	case *UserAttribute_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *UserAttribute_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *UserAttribute_BoolValue:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UserAttributes struct {
	// name of a user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// public key of tendermint wallet
	PubKey               string   `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	CreationDate         uint64   `protobuf:"varint,3,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,4,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAttributes) Reset()         { *m = UserAttributes{} }
func (m *UserAttributes) String() string { return proto.CompactTextString(m) }
func (*UserAttributes) ProtoMessage()    {}
func (*UserAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{2}
}
func (m *UserAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAttributes.Unmarshal(m, b)
}
func (m *UserAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAttributes.Marshal(b, m, deterministic)
}
func (dst *UserAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAttributes.Merge(dst, src)
}
func (m *UserAttributes) XXX_Size() int {
	return xxx_messageInfo_UserAttributes.Size(m)
}
func (m *UserAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_UserAttributes proto.InternalMessageInfo

func (m *UserAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserAttributes) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *UserAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *UserAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type RegisterRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{3}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
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
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{4}
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
	AuthenticationResult *AuthenticationResult `protobuf:"bytes,1,opt,name=authentication_result,json=authenticationResult,proto3" json:"authentication_result,omitempty"`
	User                 *User                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{5}
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

func (m *LoginResponse) GetAuthenticationResult() *AuthenticationResult {
	if m != nil {
		return m.AuthenticationResult
	}
	return nil
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type AuthenticationResult struct {
	Expiration           uint64   `protobuf:"varint,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	IssuedAt             uint64   `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	AccessToken          string   `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationResult) Reset()         { *m = AuthenticationResult{} }
func (m *AuthenticationResult) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResult) ProtoMessage()    {}
func (*AuthenticationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{6}
}
func (m *AuthenticationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResult.Unmarshal(m, b)
}
func (m *AuthenticationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResult.Marshal(b, m, deterministic)
}
func (dst *AuthenticationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResult.Merge(dst, src)
}
func (m *AuthenticationResult) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResult.Size(m)
}
func (m *AuthenticationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResult.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResult proto.InternalMessageInfo

func (m *AuthenticationResult) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *AuthenticationResult) GetIssuedAt() uint64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *AuthenticationResult) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthenticationResult) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshToken struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshToken) Reset()         { *m = RefreshToken{} }
func (m *RefreshToken) String() string { return proto.CompactTextString(m) }
func (*RefreshToken) ProtoMessage()    {}
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{7}
}
func (m *RefreshToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshToken.Unmarshal(m, b)
}
func (m *RefreshToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshToken.Marshal(b, m, deterministic)
}
func (dst *RefreshToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshToken.Merge(dst, src)
}
func (m *RefreshToken) XXX_Size() int {
	return xxx_messageInfo_RefreshToken.Size(m)
}
func (m *RefreshToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshToken.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshToken proto.InternalMessageInfo

func (m *RefreshToken) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type ConfirmRegistrationRequst struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmationCode     string   `protobuf:"bytes,2,opt,name=confirmation_code,json=confirmationCode,proto3" json:"confirmation_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmRegistrationRequst) Reset()         { *m = ConfirmRegistrationRequst{} }
func (m *ConfirmRegistrationRequst) String() string { return proto.CompactTextString(m) }
func (*ConfirmRegistrationRequst) ProtoMessage()    {}
func (*ConfirmRegistrationRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{8}
}
func (m *ConfirmRegistrationRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmRegistrationRequst.Unmarshal(m, b)
}
func (m *ConfirmRegistrationRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmRegistrationRequst.Marshal(b, m, deterministic)
}
func (dst *ConfirmRegistrationRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmRegistrationRequst.Merge(dst, src)
}
func (m *ConfirmRegistrationRequst) XXX_Size() int {
	return xxx_messageInfo_ConfirmRegistrationRequst.Size(m)
}
func (m *ConfirmRegistrationRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmRegistrationRequst.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmRegistrationRequst proto.InternalMessageInfo

func (m *ConfirmRegistrationRequst) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ConfirmRegistrationRequst) GetConfirmationCode() string {
	if m != nil {
		return m.ConfirmationCode
	}
	return ""
}

type ForgotPasswordRequst struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordRequst) Reset()         { *m = ForgotPasswordRequst{} }
func (m *ForgotPasswordRequst) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordRequst) ProtoMessage()    {}
func (*ForgotPasswordRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{9}
}
func (m *ForgotPasswordRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordRequst.Unmarshal(m, b)
}
func (m *ForgotPasswordRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordRequst.Marshal(b, m, deterministic)
}
func (dst *ForgotPasswordRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordRequst.Merge(dst, src)
}
func (m *ForgotPasswordRequst) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordRequst.Size(m)
}
func (m *ForgotPasswordRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordRequst.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordRequst proto.InternalMessageInfo

func (m *ForgotPasswordRequst) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ConfirmPasswordRequst struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmationCode     string   `protobuf:"bytes,2,opt,name=confirmation_code,json=confirmationCode,proto3" json:"confirmation_code,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmPasswordRequst) Reset()         { *m = ConfirmPasswordRequst{} }
func (m *ConfirmPasswordRequst) String() string { return proto.CompactTextString(m) }
func (*ConfirmPasswordRequst) ProtoMessage()    {}
func (*ConfirmPasswordRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{10}
}
func (m *ConfirmPasswordRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmPasswordRequst.Unmarshal(m, b)
}
func (m *ConfirmPasswordRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmPasswordRequst.Marshal(b, m, deterministic)
}
func (dst *ConfirmPasswordRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmPasswordRequst.Merge(dst, src)
}
func (m *ConfirmPasswordRequst) XXX_Size() int {
	return xxx_messageInfo_ConfirmPasswordRequst.Size(m)
}
func (m *ConfirmPasswordRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmPasswordRequst.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmPasswordRequst proto.InternalMessageInfo

func (m *ConfirmPasswordRequst) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ConfirmPasswordRequst) GetConfirmationCode() string {
	if m != nil {
		return m.ConfirmationCode
	}
	return ""
}

func (m *ConfirmPasswordRequst) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangePasswordRequst struct {
	OldPassword          string   `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordRequst) Reset()         { *m = ChangePasswordRequst{} }
func (m *ChangePasswordRequst) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordRequst) ProtoMessage()    {}
func (*ChangePasswordRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{11}
}
func (m *ChangePasswordRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordRequst.Unmarshal(m, b)
}
func (m *ChangePasswordRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordRequst.Marshal(b, m, deterministic)
}
func (dst *ChangePasswordRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordRequst.Merge(dst, src)
}
func (m *ChangePasswordRequst) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordRequst.Size(m)
}
func (m *ChangePasswordRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordRequst.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordRequst proto.InternalMessageInfo

func (m *ChangePasswordRequst) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ChangePasswordRequst) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangeEmailRequst struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEmailRequst) Reset()         { *m = ChangeEmailRequst{} }
func (m *ChangeEmailRequst) String() string { return proto.CompactTextString(m) }
func (*ChangeEmailRequst) ProtoMessage()    {}
func (*ChangeEmailRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{12}
}
func (m *ChangeEmailRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEmailRequst.Unmarshal(m, b)
}
func (m *ChangeEmailRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEmailRequst.Marshal(b, m, deterministic)
}
func (dst *ChangeEmailRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEmailRequst.Merge(dst, src)
}
func (m *ChangeEmailRequst) XXX_Size() int {
	return xxx_messageInfo_ChangeEmailRequst.Size(m)
}
func (m *ChangeEmailRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEmailRequst.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEmailRequst proto.InternalMessageInfo

func (m *ChangeEmailRequst) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type VerifyEmailRequst struct {
	ConfirmationCode     string   `protobuf:"bytes,1,opt,name=confirmation_code,json=confirmationCode,proto3" json:"confirmation_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyEmailRequst) Reset()         { *m = VerifyEmailRequst{} }
func (m *VerifyEmailRequst) String() string { return proto.CompactTextString(m) }
func (*VerifyEmailRequst) ProtoMessage()    {}
func (*VerifyEmailRequst) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{13}
}
func (m *VerifyEmailRequst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyEmailRequst.Unmarshal(m, b)
}
func (m *VerifyEmailRequst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyEmailRequst.Marshal(b, m, deterministic)
}
func (dst *VerifyEmailRequst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyEmailRequst.Merge(dst, src)
}
func (m *VerifyEmailRequst) XXX_Size() int {
	return xxx_messageInfo_VerifyEmailRequst.Size(m)
}
func (m *VerifyEmailRequst) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyEmailRequst.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyEmailRequst proto.InternalMessageInfo

func (m *VerifyEmailRequst) GetConfirmationCode() string {
	if m != nil {
		return m.ConfirmationCode
	}
	return ""
}

type UpdateAttributesRequest struct {
	UserAttributes       []*UserAttribute `protobuf:"bytes,1,rep,name=user_attributes,json=userAttributes,proto3" json:"user_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateAttributesRequest) Reset()         { *m = UpdateAttributesRequest{} }
func (m *UpdateAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAttributesRequest) ProtoMessage()    {}
func (*UpdateAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_5312f8d80f9925d4, []int{14}
}
func (m *UpdateAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAttributesRequest.Unmarshal(m, b)
}
func (m *UpdateAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAttributesRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAttributesRequest.Merge(dst, src)
}
func (m *UpdateAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAttributesRequest.Size(m)
}
func (m *UpdateAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAttributesRequest proto.InternalMessageInfo

func (m *UpdateAttributesRequest) GetUserAttributes() []*UserAttribute {
	if m != nil {
		return m.UserAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "usermgr.User")
	proto.RegisterType((*UserAttribute)(nil), "usermgr.UserAttribute")
	proto.RegisterType((*UserAttributes)(nil), "usermgr.UserAttributes")
	proto.RegisterType((*RegisterRequest)(nil), "usermgr.RegisterRequest")
	proto.RegisterType((*LoginRequest)(nil), "usermgr.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "usermgr.LoginResponse")
	proto.RegisterType((*AuthenticationResult)(nil), "usermgr.AuthenticationResult")
	proto.RegisterType((*RefreshToken)(nil), "usermgr.RefreshToken")
	proto.RegisterType((*ConfirmRegistrationRequst)(nil), "usermgr.ConfirmRegistrationRequst")
	proto.RegisterType((*ForgotPasswordRequst)(nil), "usermgr.ForgotPasswordRequst")
	proto.RegisterType((*ConfirmPasswordRequst)(nil), "usermgr.ConfirmPasswordRequst")
	proto.RegisterType((*ChangePasswordRequst)(nil), "usermgr.ChangePasswordRequst")
	proto.RegisterType((*ChangeEmailRequst)(nil), "usermgr.ChangeEmailRequst")
	proto.RegisterType((*VerifyEmailRequst)(nil), "usermgr.VerifyEmailRequst")
	proto.RegisterType((*UpdateAttributesRequest)(nil), "usermgr.UpdateAttributesRequest")
}

func init() {
	proto.RegisterFile("usermgr/v1/micro/usermgr.proto", fileDescriptor_usermgr_5312f8d80f9925d4)
}

var fileDescriptor_usermgr_5312f8d80f9925d4 = []byte{
	// 878 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0xf6, 0x39, 0x4e, 0x62, 0x8f, 0x7f, 0x24, 0xd9, 0x38, 0x4d, 0x30, 0x6a, 0x70, 0xaf, 0x2f,
	0x41, 0x54, 0x89, 0x48, 0x25, 0x40, 0xf0, 0xd0, 0x84, 0x90, 0x2a, 0x12, 0xad, 0x14, 0x2d, 0xb4,
	0x0f, 0x08, 0x61, 0x9d, 0x7d, 0x13, 0x67, 0x55, 0xdf, 0xad, 0xbb, 0xbb, 0xd7, 0x90, 0x07, 0x5e,
	0x79, 0xe6, 0x9d, 0x3f, 0x03, 0xfe, 0x40, 0xb4, 0x3f, 0x7c, 0x59, 0xdb, 0x77, 0x56, 0x79, 0xf2,
	0xed, 0xcc, 0x37, 0x33, 0xdf, 0xec, 0xec, 0x37, 0x86, 0xc3, 0x4c, 0xa2, 0x48, 0xc6, 0xe2, 0xe4,
	0xc3, 0x97, 0x27, 0x09, 0x1b, 0x09, 0x7e, 0xe2, 0x0c, 0xc7, 0x53, 0xc1, 0x15, 0x27, 0x9b, 0xee,
	0xd8, 0xdb, 0x1d, 0xf1, 0x24, 0xe1, 0xe9, 0x89, 0xfd, 0xb1, 0xde, 0x10, 0xa1, 0xf6, 0x46, 0xa2,
	0x20, 0x1d, 0xa8, 0xb2, 0xf8, 0x20, 0xe8, 0x07, 0x47, 0x0d, 0x5a, 0x65, 0x31, 0xe9, 0xc2, 0x3a,
	0x26, 0x11, 0x9b, 0x1c, 0x54, 0x8d, 0xc9, 0x1e, 0xc8, 0xd7, 0x00, 0x91, 0x52, 0x82, 0x0d, 0x33,
	0x85, 0xf2, 0x60, 0xad, 0x1f, 0x1c, 0x35, 0x4f, 0xf7, 0x8f, 0x67, 0xf5, 0x74, 0xa2, 0xf3, 0xdc,
	0x4d, 0x3d, 0x68, 0xf8, 0x6f, 0x00, 0xed, 0x39, 0x37, 0xd9, 0x86, 0xb5, 0x77, 0x78, 0xef, 0x2a,
	0xea, 0x4f, 0xf2, 0x18, 0x1a, 0x2c, 0x55, 0x83, 0x0f, 0xd1, 0x24, 0x43, 0x53, 0x76, 0xed, 0xaa,
	0x42, 0xeb, 0x2c, 0x55, 0x6f, 0xb5, 0x85, 0x3c, 0x85, 0x56, 0xcc, 0xb3, 0xe1, 0x04, 0x1d, 0x42,
	0x57, 0x0f, 0xae, 0x2a, 0xb4, 0x69, 0xad, 0x39, 0x48, 0x2a, 0xc1, 0xd2, 0xb1, 0x03, 0xd5, 0x74,
	0x7a, 0x0d, 0xb2, 0x56, 0x0b, 0xfa, 0x0c, 0x60, 0xc8, 0xf9, 0xc4, 0x41, 0xd6, 0xfb, 0xc1, 0x51,
	0xfd, 0xaa, 0x42, 0x1b, 0xda, 0x66, 0x00, 0xdf, 0x6f, 0xc2, 0xba, 0xf1, 0x85, 0x7f, 0x05, 0xd0,
	0x99, 0xef, 0x8a, 0x10, 0xa8, 0xa5, 0x51, 0x82, 0x8e, 0xb8, 0xf9, 0x26, 0xfb, 0xb0, 0x39, 0xcd,
	0x86, 0x03, 0xdd, 0x8f, 0xbd, 0xae, 0x8d, 0x69, 0x36, 0xfc, 0x11, 0xef, 0xc9, 0x53, 0x68, 0x8f,
	0x04, 0x46, 0x8a, 0xf1, 0x74, 0x10, 0x47, 0xca, 0x92, 0xae, 0xd1, 0xd6, 0xcc, 0xf8, 0x43, 0xa4,
	0x90, 0x3c, 0x03, 0x32, 0x89, 0xa4, 0x1a, 0x24, 0x3c, 0x66, 0x37, 0x0c, 0x63, 0x8b, 0xac, 0x19,
	0xe4, 0xb6, 0xf6, 0xbc, 0x76, 0x0e, 0x8d, 0x0e, 0xaf, 0x61, 0x8b, 0xe2, 0x98, 0x49, 0x85, 0x82,
	0xe2, 0xfb, 0x0c, 0xa5, 0x22, 0x4f, 0xa0, 0xa6, 0x47, 0x60, 0x28, 0x35, 0x4f, 0xdb, 0x73, 0xf3,
	0xa0, 0xc6, 0x45, 0x7a, 0x50, 0x9f, 0x46, 0x52, 0xde, 0x71, 0x11, 0x3b, 0x8a, 0xf9, 0x39, 0x3c,
	0x83, 0xd6, 0x2b, 0x3e, 0x66, 0xe9, 0x2c, 0x5d, 0x3e, 0xfa, 0xc0, 0x1f, 0xfd, 0xaa, 0x0c, 0x7f,
	0x06, 0xd0, 0x76, 0x29, 0xe4, 0x94, 0xa7, 0x12, 0x09, 0x85, 0xbd, 0x28, 0x53, 0xb7, 0x98, 0x2a,
	0x36, 0xb2, 0xed, 0x0b, 0x94, 0xd9, 0x44, 0x39, 0x8e, 0x8f, 0x73, 0x8e, 0xe7, 0x73, 0x28, 0x6a,
	0x40, 0xb4, 0x1b, 0x15, 0x58, 0xf3, 0x36, 0xab, 0xa5, 0x6d, 0x86, 0x7f, 0x07, 0xd0, 0x2d, 0xca,
	0x48, 0x0e, 0x01, 0xf0, 0xf7, 0x29, 0x13, 0xc6, 0x66, 0x48, 0xd4, 0xa8, 0x67, 0x21, 0x9f, 0x42,
	0x83, 0x49, 0x99, 0x61, 0x3c, 0x88, 0x94, 0x29, 0x50, 0xa3, 0x75, 0x6b, 0x38, 0xd7, 0x85, 0x5b,
	0xd1, 0x68, 0x84, 0x52, 0x0e, 0x14, 0x7f, 0x87, 0xa9, 0x19, 0x62, 0x83, 0x36, 0xad, 0xed, 0x67,
	0x6d, 0xd2, 0x83, 0x16, 0x78, 0x23, 0x50, 0xde, 0x3a, 0x8c, 0x79, 0x78, 0xb4, 0xe5, 0x8c, 0x06,
	0x14, 0x3e, 0x87, 0x16, 0xf5, 0xce, 0xcb, 0x41, 0x41, 0x41, 0xd0, 0x6f, 0xf0, 0xc9, 0x05, 0x4f,
	0x6f, 0x98, 0x48, 0xec, 0xd8, 0x85, 0x6b, 0xeb, 0x7d, 0x56, 0x3a, 0xaa, 0x2f, 0x60, 0x67, 0x64,
	0x43, 0xec, 0xd5, 0x8f, 0x78, 0x8c, 0x6e, 0x66, 0xdb, 0xbe, 0xe3, 0x82, 0xc7, 0x18, 0x3e, 0x83,
	0xee, 0x4b, 0x2e, 0xc6, 0x5c, 0x5d, 0xbb, 0x69, 0xae, 0x4a, 0x1d, 0xfe, 0x01, 0x7b, 0x8e, 0xcd,
	0xc7, 0xc0, 0xff, 0x17, 0x13, 0x7d, 0xcd, 0x29, 0xde, 0x0d, 0xf2, 0x57, 0xe6, 0xae, 0x39, 0xc5,
	0xbb, 0x59, 0xad, 0xf0, 0x57, 0xe8, 0x5e, 0xdc, 0x46, 0xe9, 0x18, 0x17, 0xaa, 0x3f, 0x81, 0x16,
	0x9f, 0xc4, 0x0f, 0xa1, 0x96, 0x44, 0x93, 0x4f, 0xe2, 0x19, 0x70, 0x29, 0x7b, 0x75, 0x39, 0xfb,
	0xe7, 0xb0, 0x63, 0xb3, 0x5f, 0x6a, 0xf2, 0x2b, 0xef, 0xe1, 0x0c, 0x76, 0xde, 0xa2, 0x60, 0x37,
	0xf7, 0x3e, 0xb4, 0xb0, 0xdb, 0xa0, 0xe4, 0xde, 0x7f, 0x81, 0xfd, 0x37, 0x53, 0xad, 0x74, 0x6f,
	0x63, 0x3a, 0x01, 0xbe, 0x80, 0x2d, 0xfd, 0x9a, 0x07, 0xde, 0xaa, 0x0d, 0xfa, 0x6b, 0x47, 0xcd,
	0xd3, 0x47, 0xc5, 0xab, 0x96, 0x76, 0xb2, 0xb9, 0x1d, 0x75, 0xfa, 0xcf, 0x06, 0x6c, 0x6a, 0xc4,
	0xeb, 0xb1, 0x20, 0xdf, 0x41, 0x7d, 0xb6, 0x2f, 0xc8, 0x41, 0x1e, 0xbf, 0xb0, 0x42, 0x7a, 0xbb,
	0xc7, 0xfe, 0xbf, 0xc2, 0xf1, 0x65, 0x32, 0x55, 0xf7, 0x61, 0x85, 0x7c, 0x03, 0xeb, 0x46, 0xd7,
	0x64, 0x2f, 0x8f, 0xf4, 0x57, 0x45, 0xef, 0xd1, 0xa2, 0xd9, 0xca, 0x3f, 0xac, 0x90, 0xaf, 0x60,
	0xe3, 0x15, 0x1f, 0xf3, 0x4c, 0x79, 0xa1, 0xfe, 0xe3, 0x2f, 0xab, 0x78, 0x05, 0x1d, 0x07, 0xfb,
	0x09, 0xa5, 0xd4, 0xd2, 0x2c, 0x89, 0x5f, 0xbd, 0x42, 0xc2, 0x0a, 0xb9, 0x86, 0xdd, 0x02, 0xe1,
	0x90, 0x30, 0x8f, 0x2b, 0x95, 0x55, 0x19, 0xb7, 0x97, 0xd0, 0x99, 0x97, 0x0a, 0x79, 0x20, 0x51,
	0xa4, 0xa1, 0xf2, 0x1e, 0xb7, 0x16, 0x44, 0x44, 0x0e, 0x17, 0x59, 0x7d, 0x5c, 0xa6, 0x4b, 0x68,
	0xe7, 0x7a, 0x58, 0x20, 0x54, 0xa4, 0x93, 0xf2, 0x34, 0xdb, 0x8b, 0x6f, 0x91, 0xf4, 0x1f, 0xde,
	0x5a, 0xf1, 0x33, 0xed, 0xcd, 0x6f, 0xe0, 0xb0, 0x42, 0xce, 0xa0, 0xe9, 0xe9, 0x87, 0xf4, 0x16,
	0xb8, 0x78, 0x52, 0x29, 0x23, 0xf2, 0x2d, 0x34, 0x3d, 0x59, 0x79, 0x19, 0x96, 0xc4, 0xb6, 0x5c,
	0xfd, 0xc5, 0x4c, 0x92, 0xe7, 0xde, 0x5e, 0x2e, 0xaa, 0x53, 0x52, 0x7c, 0xb8, 0x61, 0x8e, 0xcf,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x56, 0x67, 0xb9, 0x51, 0x09, 0x00, 0x00,
}
