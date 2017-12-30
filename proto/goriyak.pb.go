// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goriyak.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	goriyak.proto

It has these top-level messages:
	Transaction
	TransactionRequest
	Node
	BlockRequest
	Block
	MiningResult
	Status
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// confirm between server
type Transaction struct {
	Name          string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SendNodeId    int64                      `protobuf:"varint,2,opt,name=send_node_id,json=sendNodeId" json:"send_node_id,omitempty"`
	ReceiveNodeId int64                      `protobuf:"varint,3,opt,name=receive_node_id,json=receiveNodeId" json:"receive_node_id,omitempty"`
	Amount        float64                    `protobuf:"fixed64,4,opt,name=amount" json:"amount,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto1.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transaction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Transaction) GetSendNodeId() int64 {
	if m != nil {
		return m.SendNodeId
	}
	return 0
}

func (m *Transaction) GetReceiveNodeId() int64 {
	if m != nil {
		return m.ReceiveNodeId
	}
	return 0
}

func (m *Transaction) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// post from client to server
type TransactionRequest struct {
	NodeName      string                     `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	Password      string                     `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Name          string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	SendNodeId    int64                      `protobuf:"varint,4,opt,name=send_node_id,json=sendNodeId" json:"send_node_id,omitempty"`
	ReceiveNodeId int64                      `protobuf:"varint,5,opt,name=receive_node_id,json=receiveNodeId" json:"receive_node_id,omitempty"`
	Amount        float64                    `protobuf:"fixed64,6,opt,name=amount" json:"amount,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *TransactionRequest) Reset()                    { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string            { return proto1.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()               {}
func (*TransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TransactionRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *TransactionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TransactionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransactionRequest) GetSendNodeId() int64 {
	if m != nil {
		return m.SendNodeId
	}
	return 0
}

func (m *TransactionRequest) GetReceiveNodeId() int64 {
	if m != nil {
		return m.ReceiveNodeId
	}
	return 0
}

func (m *TransactionRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionRequest) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// for confirm node
type Node struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto1.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// request for start mining
type BlockRequest struct {
}

func (m *BlockRequest) Reset()                    { *m = BlockRequest{} }
func (m *BlockRequest) String() string            { return proto1.CompactTextString(m) }
func (*BlockRequest) ProtoMessage()               {}
func (*BlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// for mining
type Block struct {
	Id           int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Transactions []*Block_Transaction       `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	StartedAt    *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	FinishedAt   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=finished_at,json=finishedAt" json:"finished_at,omitempty"`
	Sign         []string                   `protobuf:"bytes,5,rep,name=sign" json:"sign,omitempty"`
	PreviousHash string                     `protobuf:"bytes,6,opt,name=previous_hash,json=previousHash" json:"previous_hash,omitempty"`
	Nonce        string                     `protobuf:"bytes,7,opt,name=nonce" json:"nonce,omitempty"`
	CreatedAt    *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Difficulty   string                     `protobuf:"bytes,9,opt,name=difficulty" json:"difficulty,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto1.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Block) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetTransactions() []*Block_Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetStartedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Block) GetFinishedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Block) GetSign() []string {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Block) GetPreviousHash() string {
	if m != nil {
		return m.PreviousHash
	}
	return ""
}

func (m *Block) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *Block) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Block) GetDifficulty() string {
	if m != nil {
		return m.Difficulty
	}
	return ""
}

type Block_Transaction struct {
	Id            int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SendNodeId    int64                      `protobuf:"varint,2,opt,name=send_node_id,json=sendNodeId" json:"send_node_id,omitempty"`
	ReceiveNodeId int64                      `protobuf:"varint,3,opt,name=receive_node_id,json=receiveNodeId" json:"receive_node_id,omitempty"`
	Amount        float64                    `protobuf:"fixed64,4,opt,name=amount" json:"amount,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Block_Transaction) Reset()                    { *m = Block_Transaction{} }
func (m *Block_Transaction) String() string            { return proto1.CompactTextString(m) }
func (*Block_Transaction) ProtoMessage()               {}
func (*Block_Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Block_Transaction) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block_Transaction) GetSendNodeId() int64 {
	if m != nil {
		return m.SendNodeId
	}
	return 0
}

func (m *Block_Transaction) GetReceiveNodeId() int64 {
	if m != nil {
		return m.ReceiveNodeId
	}
	return 0
}

func (m *Block_Transaction) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Block_Transaction) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// post from client to server
type MiningResult struct {
	BlockId int64  `protobuf:"varint,1,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
	Hash    string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	Nonce   string `protobuf:"bytes,3,opt,name=nonce" json:"nonce,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *MiningResult) Reset()                    { *m = MiningResult{} }
func (m *MiningResult) String() string            { return proto1.CompactTextString(m) }
func (*MiningResult) ProtoMessage()               {}
func (*MiningResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MiningResult) GetBlockId() int64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

func (m *MiningResult) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *MiningResult) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *MiningResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// uses only in response
type Status struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto1.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*Transaction)(nil), "proto.Transaction")
	proto1.RegisterType((*TransactionRequest)(nil), "proto.TransactionRequest")
	proto1.RegisterType((*Node)(nil), "proto.Node")
	proto1.RegisterType((*BlockRequest)(nil), "proto.BlockRequest")
	proto1.RegisterType((*Block)(nil), "proto.Block")
	proto1.RegisterType((*Block_Transaction)(nil), "proto.Block.Transaction")
	proto1.RegisterType((*MiningResult)(nil), "proto.MiningResult")
	proto1.RegisterType((*Status)(nil), "proto.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Goriyak service

type GoriyakClient interface {
	RegisterNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error)
	DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error)
	Login(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error)
	PostTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Status, error)
	GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Block, error)
	PostMiningResult(ctx context.Context, in *MiningResult, opts ...grpc.CallOption) (*Status, error)
}

type goriyakClient struct {
	cc *grpc.ClientConn
}

func NewGoriyakClient(cc *grpc.ClientConn) GoriyakClient {
	return &goriyakClient{cc}
}

func (c *goriyakClient) RegisterNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Goriyak/RegisterNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goriyakClient) DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Goriyak/DeleteNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goriyakClient) Login(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Goriyak/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goriyakClient) PostTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Goriyak/PostTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goriyakClient) GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := grpc.Invoke(ctx, "/proto.Goriyak/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goriyakClient) PostMiningResult(ctx context.Context, in *MiningResult, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Goriyak/PostMiningResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goriyak service

type GoriyakServer interface {
	RegisterNode(context.Context, *Node) (*Status, error)
	DeleteNode(context.Context, *Node) (*Status, error)
	Login(context.Context, *Node) (*Status, error)
	PostTransaction(context.Context, *TransactionRequest) (*Status, error)
	GetBlock(context.Context, *BlockRequest) (*Block, error)
	PostMiningResult(context.Context, *MiningResult) (*Status, error)
}

func RegisterGoriyakServer(s *grpc.Server, srv GoriyakServer) {
	s.RegisterService(&_Goriyak_serviceDesc, srv)
}

func _Goriyak_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).RegisterNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goriyak_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).DeleteNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goriyak_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).Login(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goriyak_PostTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).PostTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/PostTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).PostTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goriyak_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).GetBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goriyak_PostMiningResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiningResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoriyakServer).PostMiningResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goriyak/PostMiningResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoriyakServer).PostMiningResult(ctx, req.(*MiningResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goriyak_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Goriyak",
	HandlerType: (*GoriyakServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Goriyak_RegisterNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _Goriyak_DeleteNode_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Goriyak_Login_Handler,
		},
		{
			MethodName: "PostTransaction",
			Handler:    _Goriyak_PostTransaction_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Goriyak_GetBlock_Handler,
		},
		{
			MethodName: "PostMiningResult",
			Handler:    _Goriyak_PostMiningResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goriyak.proto",
}

// Client API for Admin service

type AdminClient interface {
	PostBlock(ctx context.Context, in *MiningResult, opts ...grpc.CallOption) (*Status, error)
	PostTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Status, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) PostBlock(ctx context.Context, in *MiningResult, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Admin/PostBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PostTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Admin/PostTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	PostBlock(context.Context, *MiningResult) (*Status, error)
	PostTransaction(context.Context, *Transaction) (*Status, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_PostBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiningResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PostBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Admin/PostBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PostBlock(ctx, req.(*MiningResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PostTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PostTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Admin/PostTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PostTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostBlock",
			Handler:    _Admin_PostBlock_Handler,
		},
		{
			MethodName: "PostTransaction",
			Handler:    _Admin_PostTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goriyak.proto",
}

func init() { proto1.RegisterFile("goriyak.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x6e, 0x2e, 0xc9, 0xdd, 0x65, 0x2e, 0xd7, 0xca, 0x2a, 0x92, 0x9e, 0xa0, 0x21, 0xa2, 0x1c,
	0x22, 0x29, 0x9e, 0x50, 0x14, 0xf5, 0xe1, 0x44, 0xa8, 0x05, 0x2d, 0x12, 0xfb, 0x7e, 0x6c, 0x93,
	0xbd, 0xdc, 0xd2, 0xcb, 0xee, 0x99, 0xdd, 0x54, 0xfa, 0xbf, 0xfa, 0xe4, 0x83, 0x3f, 0x4d, 0x64,
	0x37, 0xc9, 0x35, 0x6d, 0xe4, 0xda, 0xbe, 0xf9, 0x94, 0xfd, 0x86, 0x6f, 0x76, 0xe7, 0xfb, 0x66,
	0x26, 0x30, 0x4c, 0x79, 0x4e, 0xcf, 0xf1, 0x69, 0xb8, 0xca, 0xb9, 0xe4, 0xc8, 0xd6, 0x9f, 0xd1,
	0x93, 0x94, 0xf3, 0x74, 0x49, 0xf6, 0x34, 0x3a, 0x29, 0xe6, 0x7b, 0x92, 0x66, 0x44, 0x48, 0x9c,
	0xad, 0x4a, 0x5e, 0xf0, 0xdb, 0x80, 0xc1, 0x71, 0x8e, 0x99, 0xc0, 0xb1, 0xa4, 0x9c, 0x21, 0x04,
	0x16, 0xc3, 0x19, 0xf1, 0x0c, 0xdf, 0x18, 0x3b, 0x91, 0x3e, 0x23, 0x1f, 0x5c, 0x41, 0x58, 0x32,
	0x63, 0x3c, 0x21, 0x33, 0x9a, 0x78, 0x1d, 0xdf, 0x18, 0x9b, 0x11, 0xa8, 0xd8, 0x11, 0x4f, 0xc8,
	0x61, 0x82, 0x9e, 0xc3, 0x4e, 0x4e, 0x62, 0x42, 0xcf, 0xc8, 0x9a, 0x64, 0x6a, 0xd2, 0xb0, 0x0a,
	0x57, 0xbc, 0x87, 0xd0, 0xc5, 0x19, 0x2f, 0x98, 0xf4, 0x2c, 0xdf, 0x18, 0x1b, 0x51, 0x85, 0xd0,
	0x5b, 0x80, 0x38, 0x27, 0x58, 0x92, 0x64, 0x86, 0xa5, 0x67, 0xfb, 0xc6, 0x78, 0x30, 0x19, 0x85,
	0x65, 0xed, 0x61, 0x5d, 0x7b, 0x78, 0x5c, 0xd7, 0x1e, 0x39, 0x15, 0x7b, 0x2a, 0x83, 0x3f, 0x06,
	0xa0, 0x86, 0x80, 0x88, 0xfc, 0x28, 0x88, 0x90, 0xe8, 0x11, 0x38, 0xba, 0x92, 0x86, 0x98, 0xbe,
	0x0a, 0x1c, 0x29, 0x41, 0x23, 0xe8, 0xaf, 0xb0, 0x10, 0x3f, 0x79, 0x5e, 0x8a, 0x71, 0xa2, 0x35,
	0x5e, 0x1b, 0x60, 0x6e, 0x30, 0xc0, 0xba, 0x8d, 0x01, 0xf6, 0x66, 0x03, 0xba, 0x1b, 0x0c, 0xe8,
	0xdd, 0xc5, 0x80, 0x7d, 0xb0, 0xd4, 0xe5, 0xff, 0xec, 0xdc, 0x06, 0xa1, 0xc1, 0x36, 0xb8, 0x1f,
	0x97, 0x3c, 0x3e, 0xad, 0x1c, 0x0b, 0x2e, 0x2c, 0xb0, 0x75, 0x00, 0x6d, 0x43, 0x87, 0x26, 0xfa,
	0x1e, 0x33, 0xea, 0xd0, 0x04, 0xbd, 0x07, 0x57, 0x5e, 0x3a, 0x2c, 0xbc, 0x8e, 0x6f, 0x8e, 0x07,
	0x13, 0xaf, 0xac, 0x2b, 0xd4, 0x39, 0x61, 0xb3, 0x05, 0x57, 0xd8, 0x4a, 0x9a, 0x90, 0x38, 0xaf,
	0xa4, 0x99, 0x37, 0x4b, 0xab, 0xd8, 0x53, 0x89, 0xde, 0xc1, 0x60, 0x4e, 0x19, 0x15, 0x8b, 0x32,
	0xd7, 0xba, 0x31, 0x17, 0x6a, 0xfa, 0x54, 0x2a, 0x3f, 0x04, 0x4d, 0x99, 0x67, 0xfb, 0xa6, 0xf2,
	0x43, 0x9d, 0xd1, 0x53, 0x18, 0xae, 0x72, 0x72, 0x46, 0x79, 0x21, 0x66, 0x0b, 0x2c, 0x16, 0xba,
	0x0b, 0x4e, 0xe4, 0xd6, 0xc1, 0xcf, 0x58, 0x2c, 0xd0, 0x03, 0xb0, 0x19, 0x67, 0x31, 0xd1, 0x6d,
	0x70, 0xa2, 0x12, 0x5c, 0xeb, 0x50, 0xff, 0x0e, 0x1d, 0x42, 0x8f, 0x01, 0x12, 0x3a, 0x9f, 0xd3,
	0xb8, 0x58, 0xca, 0x73, 0xcf, 0xd1, 0xb7, 0x36, 0x22, 0xa3, 0x5f, 0xd7, 0x76, 0xf0, 0xba, 0xff,
	0xff, 0xf5, 0xfe, 0xa5, 0xe0, 0x7e, 0xa5, 0x8c, 0xb2, 0x34, 0x22, 0xa2, 0x58, 0x4a, 0xb4, 0x0b,
	0xfd, 0x13, 0x35, 0x11, 0xb3, 0xb5, 0x84, 0x9e, 0xc6, 0x87, 0x7a, 0xb5, 0xb4, 0xe9, 0xe5, 0x24,
	0xea, 0xf3, 0xa5, 0xd9, 0x66, 0xd3, 0xec, 0x7a, 0x96, 0xad, 0xcb, 0x59, 0x0e, 0x02, 0xe8, 0x7e,
	0x97, 0x58, 0x16, 0x02, 0x79, 0xd0, 0xcb, 0x88, 0x10, 0x38, 0xad, 0x87, 0xbd, 0x86, 0x93, 0x8b,
	0x0e, 0xf4, 0x0e, 0xca, 0xff, 0x20, 0x7a, 0x09, 0x6e, 0x44, 0x52, 0x2a, 0x24, 0xc9, 0xf5, 0x7e,
	0x0c, 0xaa, 0x79, 0x55, 0x60, 0x34, 0xac, 0x40, 0x79, 0x63, 0xb0, 0x85, 0x5e, 0x00, 0x7c, 0x22,
	0x4b, 0x22, 0xc9, 0x2d, 0xb8, 0xcf, 0xc0, 0xfe, 0xc2, 0x53, 0xca, 0x6e, 0xa0, 0x7d, 0x80, 0x9d,
	0x6f, 0x5c, 0xc8, 0x66, 0x67, 0x77, 0x2b, 0x4e, 0xfb, 0x87, 0xd5, 0x4e, 0xdf, 0x83, 0xfe, 0x01,
	0x91, 0xe5, 0x46, 0xde, 0x6f, 0xee, 0x5a, 0x9d, 0xe1, 0x36, 0x83, 0xc1, 0x16, 0x7a, 0x03, 0xf7,
	0xd4, 0x7b, 0x57, 0xba, 0x51, 0x27, 0x36, 0x83, 0xad, 0xa7, 0x26, 0x39, 0xd8, 0xd3, 0x24, 0xa3,
	0x0c, 0xbd, 0x02, 0x47, 0x5d, 0x71, 0xf5, 0xd1, 0x8d, 0xb9, 0x68, 0xbf, 0xad, 0x12, 0xb5, 0x55,
	0xb6, 0xf2, 0x4e, 0xba, 0x1a, 0xbf, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x39, 0xee, 0xea, 0xaf,
	0xb8, 0x06, 0x00, 0x00,
}