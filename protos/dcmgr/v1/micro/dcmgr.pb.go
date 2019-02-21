// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

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

type DataCenterListResponse struct {
	DcList               []*common.DataCenter `protobuf:"bytes,1,rep,name=dcList,proto3" json:"dcList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataCenterListResponse) Reset()         { *m = DataCenterListResponse{} }
func (m *DataCenterListResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterListResponse) ProtoMessage()    {}
func (*DataCenterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_c3a1b3a9887c3018, []int{0}
}
func (m *DataCenterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListResponse.Unmarshal(m, b)
}
func (m *DataCenterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListResponse.Marshal(b, m, deterministic)
}
func (dst *DataCenterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListResponse.Merge(dst, src)
}
func (m *DataCenterListResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterListResponse.Size(m)
}
func (m *DataCenterListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterListResponse proto.InternalMessageInfo

func (m *DataCenterListResponse) GetDcList() []*common.DataCenter {
	if m != nil {
		return m.DcList
	}
	return nil
}

type NetworkInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfoResponse) Reset()         { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()    {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_c3a1b3a9887c3018, []int{1}
}
func (m *NetworkInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfoResponse.Unmarshal(m, b)
}
func (m *NetworkInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfoResponse.Marshal(b, m, deterministic)
}
func (dst *NetworkInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfoResponse.Merge(dst, src)
}
func (m *NetworkInfoResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInfoResponse.Size(m)
}
func (m *NetworkInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfoResponse proto.InternalMessageInfo

func (m *NetworkInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/dcmgr.proto", fileDescriptor_dcmgr_c3a1b3a9887c3018) }

var fileDescriptor_dcmgr_c3a1b3a9887c3018 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xd1, 0x4a, 0xc3, 0x40,
	0x10, 0xcc, 0xa1, 0x2d, 0xb8, 0x15, 0x1f, 0x36, 0x50, 0xc2, 0x81, 0x50, 0xee, 0x29, 0xbe, 0x24,
	0x35, 0xfe, 0x80, 0x92, 0x28, 0x54, 0x44, 0x24, 0xfd, 0x82, 0x98, 0xae, 0x52, 0xe4, 0x6e, 0xc3,
	0xe6, 0xa8, 0xf8, 0x19, 0xfe, 0xb1, 0x98, 0x04, 0x6d, 0x25, 0x3e, 0xdd, 0xcd, 0xdc, 0xcc, 0x1c,
	0x33, 0xa0, 0x37, 0xb5, 0x7d, 0x95, 0x74, 0x77, 0x99, 0xda, 0x6d, 0x2d, 0x9c, 0x76, 0x30, 0x69,
	0x84, 0x3d, 0xe3, 0xa4, 0x03, 0x3a, 0xac, 0xd9, 0x5a, 0x76, 0x69, 0x7f, 0xf4, 0x6f, 0xe6, 0x1e,
	0xe6, 0x45, 0xe5, 0xab, 0x9c, 0x9c, 0x27, 0x79, 0xd8, 0xb6, 0xbe, 0xa4, 0xb6, 0x61, 0xd7, 0x12,
	0x2e, 0x61, 0xba, 0xa9, 0xbf, 0x99, 0x48, 0x2d, 0x8e, 0xe2, 0x59, 0x16, 0x25, 0xfb, 0xc6, 0xe4,
	0xd7, 0x55, 0x0e, 0x3a, 0x73, 0x01, 0xe1, 0x23, 0xf9, 0x77, 0x96, 0xb7, 0x95, 0x7b, 0xe1, 0x9f,
	0x20, 0x84, 0x63, 0x57, 0x59, 0x8a, 0xd4, 0x42, 0xc5, 0x27, 0x65, 0x77, 0xcf, 0x4a, 0x80, 0x22,
	0x5f, 0x7b, 0xa1, 0xca, 0x92, 0x60, 0x01, 0xa7, 0x6b, 0x92, 0x1d, 0x49, 0xcf, 0xe0, 0xfc, 0xcf,
	0x57, 0x83, 0x52, 0xff, 0xc3, 0x9b, 0x20, 0x56, 0x4b, 0x95, 0x7d, 0x2a, 0x98, 0x14, 0xf9, 0xcd,
	0xd3, 0x0a, 0xef, 0xe0, 0xec, 0xb0, 0x14, 0x86, 0x87, 0xce, 0x5b, 0xdb, 0xf8, 0x0f, 0x7d, 0x9e,
	0xf4, 0x2b, 0x8d, 0x0f, 0x60, 0x02, 0xbc, 0x86, 0xd9, 0x5e, 0xa1, 0xf1, 0x10, 0x3d, 0x84, 0x8c,
	0x34, 0x37, 0xc1, 0xf3, 0xb4, 0x93, 0x5e, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xfe, 0x48,
	0xdf, 0x9f, 0x01, 0x00, 0x00,
}
