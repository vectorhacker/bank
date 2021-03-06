// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

package accounts

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

type CreateAccountRequest struct {
	InitialDeposit       int64    `protobuf:"varint,1,opt,name=initialDeposit,proto3" json:"initialDeposit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{0}
}
func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(dst, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetInitialDeposit() int64 {
	if m != nil {
		return m.InitialDeposit
	}
	return 0
}

type CreateAccountResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{1}
}
func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (dst *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(dst, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type DebitAccountRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CorrelationID        string   `protobuf:"bytes,5,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebitAccountRequest) Reset()         { *m = DebitAccountRequest{} }
func (m *DebitAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DebitAccountRequest) ProtoMessage()    {}
func (*DebitAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{2}
}
func (m *DebitAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebitAccountRequest.Unmarshal(m, b)
}
func (m *DebitAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebitAccountRequest.Marshal(b, m, deterministic)
}
func (dst *DebitAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebitAccountRequest.Merge(dst, src)
}
func (m *DebitAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DebitAccountRequest.Size(m)
}
func (m *DebitAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebitAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebitAccountRequest proto.InternalMessageInfo

func (m *DebitAccountRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DebitAccountRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DebitAccountRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DebitAccountRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DebitAccountRequest) GetCorrelationID() string {
	if m != nil {
		return m.CorrelationID
	}
	return ""
}

type DebitAccountResponse struct {
	TransactionID        string   `protobuf:"bytes,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebitAccountResponse) Reset()         { *m = DebitAccountResponse{} }
func (m *DebitAccountResponse) String() string { return proto.CompactTextString(m) }
func (*DebitAccountResponse) ProtoMessage()    {}
func (*DebitAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{3}
}
func (m *DebitAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebitAccountResponse.Unmarshal(m, b)
}
func (m *DebitAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebitAccountResponse.Marshal(b, m, deterministic)
}
func (dst *DebitAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebitAccountResponse.Merge(dst, src)
}
func (m *DebitAccountResponse) XXX_Size() int {
	return xxx_messageInfo_DebitAccountResponse.Size(m)
}
func (m *DebitAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DebitAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DebitAccountResponse proto.InternalMessageInfo

func (m *DebitAccountResponse) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

type CreditAccountRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CorrelationID        string   `protobuf:"bytes,5,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreditAccountRequest) Reset()         { *m = CreditAccountRequest{} }
func (m *CreditAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreditAccountRequest) ProtoMessage()    {}
func (*CreditAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{4}
}
func (m *CreditAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditAccountRequest.Unmarshal(m, b)
}
func (m *CreditAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreditAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditAccountRequest.Merge(dst, src)
}
func (m *CreditAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreditAccountRequest.Size(m)
}
func (m *CreditAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreditAccountRequest proto.InternalMessageInfo

func (m *CreditAccountRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CreditAccountRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreditAccountRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreditAccountRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CreditAccountRequest) GetCorrelationID() string {
	if m != nil {
		return m.CorrelationID
	}
	return ""
}

type CreditAccountResponse struct {
	TransactionID        string   `protobuf:"bytes,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreditAccountResponse) Reset()         { *m = CreditAccountResponse{} }
func (m *CreditAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreditAccountResponse) ProtoMessage()    {}
func (*CreditAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_344e315ca8dd90c7, []int{5}
}
func (m *CreditAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditAccountResponse.Unmarshal(m, b)
}
func (m *CreditAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditAccountResponse.Marshal(b, m, deterministic)
}
func (dst *CreditAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditAccountResponse.Merge(dst, src)
}
func (m *CreditAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreditAccountResponse.Size(m)
}
func (m *CreditAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreditAccountResponse proto.InternalMessageInfo

func (m *CreditAccountResponse) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "accounts.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "accounts.CreateAccountResponse")
	proto.RegisterType((*DebitAccountRequest)(nil), "accounts.DebitAccountRequest")
	proto.RegisterType((*DebitAccountResponse)(nil), "accounts.DebitAccountResponse")
	proto.RegisterType((*CreditAccountRequest)(nil), "accounts.CreditAccountRequest")
	proto.RegisterType((*CreditAccountResponse)(nil), "accounts.CreditAccountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsCommandClient is the client API for AccountsCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsCommandClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	DebitAccount(ctx context.Context, in *DebitAccountRequest, opts ...grpc.CallOption) (*DebitAccountResponse, error)
	CreditAccount(ctx context.Context, in *CreditAccountRequest, opts ...grpc.CallOption) (*CreditAccountResponse, error)
}

type accountsCommandClient struct {
	cc *grpc.ClientConn
}

func NewAccountsCommandClient(cc *grpc.ClientConn) AccountsCommandClient {
	return &accountsCommandClient{cc}
}

func (c *accountsCommandClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountsCommand/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsCommandClient) DebitAccount(ctx context.Context, in *DebitAccountRequest, opts ...grpc.CallOption) (*DebitAccountResponse, error) {
	out := new(DebitAccountResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountsCommand/DebitAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsCommandClient) CreditAccount(ctx context.Context, in *CreditAccountRequest, opts ...grpc.CallOption) (*CreditAccountResponse, error) {
	out := new(CreditAccountResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountsCommand/CreditAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsCommandServer is the server API for AccountsCommand service.
type AccountsCommandServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	DebitAccount(context.Context, *DebitAccountRequest) (*DebitAccountResponse, error)
	CreditAccount(context.Context, *CreditAccountRequest) (*CreditAccountResponse, error)
}

func RegisterAccountsCommandServer(s *grpc.Server, srv AccountsCommandServer) {
	s.RegisterService(&_AccountsCommand_serviceDesc, srv)
}

func _AccountsCommand_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsCommandServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountsCommand/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsCommandServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsCommand_DebitAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsCommandServer).DebitAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountsCommand/DebitAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsCommandServer).DebitAccount(ctx, req.(*DebitAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsCommand_CreditAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsCommandServer).CreditAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountsCommand/CreditAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsCommandServer).CreditAccount(ctx, req.(*CreditAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountsCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.AccountsCommand",
	HandlerType: (*AccountsCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountsCommand_CreateAccount_Handler,
		},
		{
			MethodName: "DebitAccount",
			Handler:    _AccountsCommand_DebitAccount_Handler,
		},
		{
			MethodName: "CreditAccount",
			Handler:    _AccountsCommand_CreditAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command.proto",
}

func init() { proto.RegisterFile("command.proto", fileDescriptor_command_344e315ca8dd90c7) }

var fileDescriptor_command_344e315ca8dd90c7 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x4e, 0xeb, 0x30,
	0x10, 0x54, 0xd2, 0xf7, 0x2a, 0xba, 0x90, 0x22, 0x99, 0x82, 0xa2, 0x0a, 0x4a, 0x15, 0x55, 0xc0,
	0x29, 0x07, 0xb8, 0x02, 0x12, 0x6a, 0x2e, 0x39, 0x20, 0xa1, 0xfc, 0x81, 0x9b, 0xf8, 0x60, 0xa9,
	0xb1, 0x8d, 0xbd, 0xfd, 0x03, 0xfe, 0x04, 0x3e, 0x14, 0xc5, 0x31, 0x4a, 0xe3, 0xb6, 0x17, 0x4e,
	0x1c, 0x33, 0xde, 0xd9, 0x9d, 0xd9, 0x9d, 0x40, 0x54, 0xca, 0xba, 0xa6, 0xa2, 0x4a, 0x95, 0x96,
	0x28, 0xc9, 0x11, 0x2d, 0x4b, 0xb9, 0x11, 0x68, 0x92, 0x67, 0x98, 0x2c, 0x35, 0xa3, 0xc8, 0x5e,
	0x5a, 0xa4, 0x60, 0xef, 0x1b, 0x66, 0x90, 0xdc, 0xc0, 0x98, 0x0b, 0x8e, 0x9c, 0xae, 0x33, 0xa6,
	0xa4, 0xe1, 0x18, 0x07, 0xf3, 0xe0, 0x6e, 0x50, 0x78, 0x68, 0x72, 0x0b, 0xe7, 0x1e, 0xdf, 0x28,
	0x29, 0x0c, 0x23, 0x63, 0x08, 0xf3, 0xcc, 0x92, 0x46, 0x45, 0x98, 0x67, 0xc9, 0x67, 0x00, 0x67,
	0x19, 0x5b, 0x71, 0xf4, 0x06, 0x79, 0x75, 0xe4, 0x02, 0x86, 0xb4, 0x6e, 0x0a, 0xe2, 0xd0, 0x0e,
	0x74, 0x5f, 0x64, 0x0e, 0xc7, 0x15, 0x33, 0xa5, 0xe6, 0x0a, 0xb9, 0x14, 0xf1, 0xc0, 0x12, 0xb6,
	0x21, 0x72, 0x09, 0x23, 0xe4, 0x35, 0x33, 0x48, 0x6b, 0x15, 0xff, 0xb3, 0xe4, 0x0e, 0x20, 0x8b,
	0x66, 0x07, 0x5a, 0xb3, 0x35, 0x6d, 0x8a, 0xf3, 0x2c, 0xfe, 0x6f, 0x3b, 0xf4, 0xc1, 0xe4, 0x11,
	0x26, 0x7d, 0x91, 0xce, 0xcd, 0x02, 0x22, 0xd4, 0x54, 0x18, 0x5a, 0x3a, 0x76, 0x2b, 0xb8, 0x0f,
	0x26, 0x5f, 0x81, 0xdd, 0x66, 0xf5, 0xc7, 0x4d, 0x3e, 0xd9, 0x9b, 0x55, 0xbf, 0x74, 0x79, 0xff,
	0x11, 0xc2, 0xa9, 0x63, 0x9a, 0x65, 0x1b, 0x2b, 0xf2, 0x06, 0x51, 0x2f, 0x06, 0x64, 0x96, 0xfe,
	0x44, 0x2c, 0xdd, 0x97, 0xaf, 0xe9, 0xf5, 0xc1, 0x77, 0xa7, 0xe5, 0x15, 0x4e, 0xb6, 0x2f, 0x41,
	0xae, 0x3a, 0xc2, 0x9e, 0x18, 0x4d, 0x67, 0x87, 0x9e, 0x5d, 0xbb, 0x56, 0x60, 0xe7, 0xd9, 0x13,
	0xb8, 0x73, 0x32, 0x4f, 0xe0, 0xee, 0xb2, 0x56, 0x43, 0xfb, 0x2b, 0x3d, 0x7c, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xbf, 0x2a, 0xf6, 0x5b, 0x03, 0x00, 0x00,
}
