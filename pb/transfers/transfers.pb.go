// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfers.proto

package transfers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type BeginTransferRequest struct {
	FromAccountID        string   `protobuf:"bytes,1,opt,name=fromAccountID,proto3" json:"fromAccountID,omitempty"`
	ToAccountID          string   `protobuf:"bytes,2,opt,name=toAccountID,proto3" json:"toAccountID,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginTransferRequest) Reset()         { *m = BeginTransferRequest{} }
func (m *BeginTransferRequest) String() string { return proto.CompactTextString(m) }
func (*BeginTransferRequest) ProtoMessage()    {}
func (*BeginTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_fcba74e1816b7f22, []int{0}
}
func (m *BeginTransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginTransferRequest.Unmarshal(m, b)
}
func (m *BeginTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginTransferRequest.Marshal(b, m, deterministic)
}
func (dst *BeginTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginTransferRequest.Merge(dst, src)
}
func (m *BeginTransferRequest) XXX_Size() int {
	return xxx_messageInfo_BeginTransferRequest.Size(m)
}
func (m *BeginTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeginTransferRequest proto.InternalMessageInfo

func (m *BeginTransferRequest) GetFromAccountID() string {
	if m != nil {
		return m.FromAccountID
	}
	return ""
}

func (m *BeginTransferRequest) GetToAccountID() string {
	if m != nil {
		return m.ToAccountID
	}
	return ""
}

func (m *BeginTransferRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BeginTransferRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type BeginTransferResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginTransferResponse) Reset()         { *m = BeginTransferResponse{} }
func (m *BeginTransferResponse) String() string { return proto.CompactTextString(m) }
func (*BeginTransferResponse) ProtoMessage()    {}
func (*BeginTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_fcba74e1816b7f22, []int{1}
}
func (m *BeginTransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginTransferResponse.Unmarshal(m, b)
}
func (m *BeginTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginTransferResponse.Marshal(b, m, deterministic)
}
func (dst *BeginTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginTransferResponse.Merge(dst, src)
}
func (m *BeginTransferResponse) XXX_Size() int {
	return xxx_messageInfo_BeginTransferResponse.Size(m)
}
func (m *BeginTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BeginTransferResponse proto.InternalMessageInfo

func (m *BeginTransferResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*BeginTransferRequest)(nil), "transfers.BeginTransferRequest")
	proto.RegisterType((*BeginTransferResponse)(nil), "transfers.BeginTransferResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransfersClient is the client API for Transfers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransfersClient interface {
	BeginTransfer(ctx context.Context, in *BeginTransferRequest, opts ...grpc.CallOption) (*BeginTransferResponse, error)
}

type transfersClient struct {
	cc *grpc.ClientConn
}

func NewTransfersClient(cc *grpc.ClientConn) TransfersClient {
	return &transfersClient{cc}
}

func (c *transfersClient) BeginTransfer(ctx context.Context, in *BeginTransferRequest, opts ...grpc.CallOption) (*BeginTransferResponse, error) {
	out := new(BeginTransferResponse)
	err := c.cc.Invoke(ctx, "/transfers.Transfers/BeginTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServer is the server API for Transfers service.
type TransfersServer interface {
	BeginTransfer(context.Context, *BeginTransferRequest) (*BeginTransferResponse, error)
}

func RegisterTransfersServer(s *grpc.Server, srv TransfersServer) {
	s.RegisterService(&_Transfers_serviceDesc, srv)
}

func _Transfers_BeginTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).BeginTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfers.Transfers/BeginTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).BeginTransfer(ctx, req.(*BeginTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transfers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transfers.Transfers",
	HandlerType: (*TransfersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeginTransfer",
			Handler:    _Transfers_BeginTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfers.proto",
}

func init() { proto.RegisterFile("transfers.proto", fileDescriptor_transfers_fcba74e1816b7f22) }

var fileDescriptor_transfers_fcba74e1816b7f22 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4b, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xcd, 0x60, 0xe4, 0x12, 0x71, 0x4a, 0x4d, 0xcf, 0xcc, 0x0b, 0x81, 0x0a, 0x05, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0xa9, 0x70, 0xf1, 0xa6, 0x15, 0xe5, 0xe7, 0x3a, 0x26, 0x27, 0xe7, 0x97,
	0xe6, 0x95, 0x78, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x0a, 0x29, 0x70,
	0x71, 0x97, 0xe4, 0x23, 0xd4, 0x30, 0x81, 0xd5, 0x20, 0x0b, 0x09, 0x89, 0x71, 0xb1, 0x25, 0xe6,
	0x82, 0xd8, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x1e, 0x48, 0x67, 0x4a, 0x6a, 0x71,
	0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x0b, 0x44, 0x27, 0x92, 0x90, 0x92, 0x3a,
	0x97, 0x28, 0x9a, 0xcb, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0xe0, 0xee,
	0x61, 0xf2, 0x74, 0x31, 0x8a, 0xe7, 0xe2, 0x84, 0xa9, 0x29, 0x16, 0x0a, 0xe2, 0xe2, 0x45, 0xd1,
	0x25, 0x24, 0xaf, 0x87, 0xf0, 0x3e, 0x36, 0x9f, 0x4a, 0x29, 0xe0, 0x56, 0x00, 0xb1, 0x30, 0x89,
	0x0d, 0x1c, 0x6c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x97, 0xbe, 0x8d, 0x49, 0x01,
	0x00, 0x00,
}