// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/grpc/dcmgr.proto

package dcmgr

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
	return fileDescriptor_dcmgr_0f95de64ebe4fbb0, []int{0}
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

func init() {
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DCStreamerClient is the client API for DCStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCStreamerClient interface {
	ServerStream(ctx context.Context, opts ...grpc.CallOption) (DCStreamer_ServerStreamClient, error)
}

type dCStreamerClient struct {
	cc *grpc.ClientConn
}

func NewDCStreamerClient(cc *grpc.ClientConn) DCStreamerClient {
	return &dCStreamerClient{cc}
}

func (c *dCStreamerClient) ServerStream(ctx context.Context, opts ...grpc.CallOption) (DCStreamer_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DCStreamer_serviceDesc.Streams[0], "/dcmgr.DCStreamer/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dCStreamerServerStreamClient{stream}
	return x, nil
}

type DCStreamer_ServerStreamClient interface {
	Send(*common.Event) error
	Recv() (*common.Event, error)
	grpc.ClientStream
}

type dCStreamerServerStreamClient struct {
	grpc.ClientStream
}

func (x *dCStreamerServerStreamClient) Send(m *common.Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dCStreamerServerStreamClient) Recv() (*common.Event, error) {
	m := new(common.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DCStreamerServer is the server API for DCStreamer service.
type DCStreamerServer interface {
	ServerStream(DCStreamer_ServerStreamServer) error
}

func RegisterDCStreamerServer(s *grpc.Server, srv DCStreamerServer) {
	s.RegisterService(&_DCStreamer_serviceDesc, srv)
}

func _DCStreamer_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DCStreamerServer).ServerStream(&dCStreamerServerStreamServer{stream})
}

type DCStreamer_ServerStreamServer interface {
	Send(*common.Event) error
	Recv() (*common.Event, error)
	grpc.ServerStream
}

type dCStreamerServerStreamServer struct {
	grpc.ServerStream
}

func (x *dCStreamerServerStreamServer) Send(m *common.Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dCStreamerServerStreamServer) Recv() (*common.Event, error) {
	m := new(common.Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DCStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcmgr.DCStreamer",
	HandlerType: (*DCStreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _DCStreamer_ServerStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dcmgr/v1/grpc/dcmgr.proto",
}

// DCAPIClient is the client API for DCAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCAPIClient interface {
	DataCenterList(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterListResponse, error)
}

type dCAPIClient struct {
	cc *grpc.ClientConn
}

func NewDCAPIClient(cc *grpc.ClientConn) DCAPIClient {
	return &dCAPIClient{cc}
}

func (c *dCAPIClient) DataCenterList(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterListResponse, error) {
	out := new(DataCenterListResponse)
	err := c.cc.Invoke(ctx, "/dcmgr.DCAPI/DataCenterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DCAPIServer is the server API for DCAPI service.
type DCAPIServer interface {
	DataCenterList(context.Context, *common.Empty) (*DataCenterListResponse, error)
}

func RegisterDCAPIServer(s *grpc.Server, srv DCAPIServer) {
	s.RegisterService(&_DCAPI_serviceDesc, srv)
}

func _DCAPI_DataCenterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCAPIServer).DataCenterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DCAPI/DataCenterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCAPIServer).DataCenterList(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DCAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcmgr.DCAPI",
	HandlerType: (*DCAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataCenterList",
			Handler:    _DCAPI_DataCenterList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcmgr/v1/grpc/dcmgr.proto",
}

func init() { proto.RegisterFile("dcmgr/v1/grpc/dcmgr.proto", fileDescriptor_dcmgr_0f95de64ebe4fbb0) }

var fileDescriptor_dcmgr_0f95de64ebe4fbb0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0xce, 0x4d,
	0x2f, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x07, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x58, 0xc1, 0x1c, 0x29, 0xe1, 0xe4, 0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x7d, 0x08,
	0x05, 0x91, 0x53, 0xf2, 0xe2, 0x12, 0x73, 0x49, 0x2c, 0x49, 0x74, 0x4e, 0xcd, 0x2b, 0x49, 0x2d,
	0xf2, 0xc9, 0x2c, 0x2e, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe0, 0x62,
	0x4b, 0x49, 0x06, 0x89, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xe8, 0x21, 0x6b, 0xd4,
	0x43, 0xe8, 0x0a, 0x82, 0xaa, 0x33, 0xf2, 0xe1, 0xe2, 0x72, 0x71, 0x0e, 0x2e, 0x29, 0x4a, 0x4d,
	0xcc, 0x4d, 0x2d, 0x12, 0xb2, 0xe3, 0xe2, 0x09, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x82, 0x88, 0x08,
	0x09, 0xa3, 0xea, 0x77, 0x2d, 0x4b, 0xcd, 0x2b, 0x91, 0xc2, 0x26, 0xa8, 0xc4, 0xa0, 0xc1, 0x68,
	0xc0, 0x68, 0xe4, 0xcf, 0xc5, 0xea, 0xe2, 0xec, 0x18, 0xe0, 0x29, 0xe4, 0xc6, 0xc5, 0x87, 0xea,
	0x44, 0x0c, 0xa3, 0x72, 0x0b, 0x4a, 0x2a, 0xa5, 0x64, 0xf5, 0x20, 0x7e, 0xc6, 0xee, 0x1d, 0x25,
	0x86, 0x24, 0x36, 0xb0, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xac, 0x04, 0xe4,
	0x2a, 0x01, 0x00, 0x00,
}
