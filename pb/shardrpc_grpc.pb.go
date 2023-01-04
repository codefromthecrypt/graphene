// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: shardrpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShardRPCClient is the client API for ShardRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShardRPCClient interface {
	AnnounceProposal(ctx context.Context, in *ProposalAnnouncement, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBlockHashAtSlot(ctx context.Context, in *SlotRequest, opts ...grpc.CallOption) (*BlockHashResponse, error)
	GenerateBlockTemplate(ctx context.Context, in *BlockGenerationRequest, opts ...grpc.CallOption) (*ShardBlock, error)
	SubmitBlock(ctx context.Context, in *ShardBlockSubmission, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetActionStream(ctx context.Context, in *ShardActionStreamRequest, opts ...grpc.CallOption) (ShardRPC_GetActionStreamClient, error)
	GetListeningAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListeningAddressesResponse, error)
	Connect(ctx context.Context, in *ConnectMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSlotNumber(ctx context.Context, in *SlotNumberRequest, opts ...grpc.CallOption) (*SlotNumberResponse, error)
}

type shardRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewShardRPCClient(cc grpc.ClientConnInterface) ShardRPCClient {
	return &shardRPCClient{cc}
}

func (c *shardRPCClient) AnnounceProposal(ctx context.Context, in *ProposalAnnouncement, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/AnnounceProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) GetBlockHashAtSlot(ctx context.Context, in *SlotRequest, opts ...grpc.CallOption) (*BlockHashResponse, error) {
	out := new(BlockHashResponse)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/GetBlockHashAtSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) GenerateBlockTemplate(ctx context.Context, in *BlockGenerationRequest, opts ...grpc.CallOption) (*ShardBlock, error) {
	out := new(ShardBlock)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/GenerateBlockTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) SubmitBlock(ctx context.Context, in *ShardBlockSubmission, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/SubmitBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) GetActionStream(ctx context.Context, in *ShardActionStreamRequest, opts ...grpc.CallOption) (ShardRPC_GetActionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShardRPC_ServiceDesc.Streams[0], "/pb.ShardRPC/GetActionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &shardRPCGetActionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShardRPC_GetActionStreamClient interface {
	Recv() (*ShardChainAction, error)
	grpc.ClientStream
}

type shardRPCGetActionStreamClient struct {
	grpc.ClientStream
}

func (x *shardRPCGetActionStreamClient) Recv() (*ShardChainAction, error) {
	m := new(ShardChainAction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shardRPCClient) GetListeningAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListeningAddressesResponse, error) {
	out := new(ListeningAddressesResponse)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/GetListeningAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) Connect(ctx context.Context, in *ConnectMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardRPCClient) GetSlotNumber(ctx context.Context, in *SlotNumberRequest, opts ...grpc.CallOption) (*SlotNumberResponse, error) {
	out := new(SlotNumberResponse)
	err := c.cc.Invoke(ctx, "/pb.ShardRPC/GetSlotNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShardRPCServer is the server API for ShardRPC service.
// All implementations must embed UnimplementedShardRPCServer
// for forward compatibility
type ShardRPCServer interface {
	AnnounceProposal(context.Context, *ProposalAnnouncement) (*emptypb.Empty, error)
	GetBlockHashAtSlot(context.Context, *SlotRequest) (*BlockHashResponse, error)
	GenerateBlockTemplate(context.Context, *BlockGenerationRequest) (*ShardBlock, error)
	SubmitBlock(context.Context, *ShardBlockSubmission) (*emptypb.Empty, error)
	GetActionStream(*ShardActionStreamRequest, ShardRPC_GetActionStreamServer) error
	GetListeningAddresses(context.Context, *emptypb.Empty) (*ListeningAddressesResponse, error)
	Connect(context.Context, *ConnectMessage) (*emptypb.Empty, error)
	GetSlotNumber(context.Context, *SlotNumberRequest) (*SlotNumberResponse, error)
	mustEmbedUnimplementedShardRPCServer()
}

// UnimplementedShardRPCServer must be embedded to have forward compatible implementations.
type UnimplementedShardRPCServer struct {
}

func (UnimplementedShardRPCServer) AnnounceProposal(context.Context, *ProposalAnnouncement) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceProposal not implemented")
}
func (UnimplementedShardRPCServer) GetBlockHashAtSlot(context.Context, *SlotRequest) (*BlockHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHashAtSlot not implemented")
}
func (UnimplementedShardRPCServer) GenerateBlockTemplate(context.Context, *BlockGenerationRequest) (*ShardBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBlockTemplate not implemented")
}
func (UnimplementedShardRPCServer) SubmitBlock(context.Context, *ShardBlockSubmission) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBlock not implemented")
}
func (UnimplementedShardRPCServer) GetActionStream(*ShardActionStreamRequest, ShardRPC_GetActionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActionStream not implemented")
}
func (UnimplementedShardRPCServer) GetListeningAddresses(context.Context, *emptypb.Empty) (*ListeningAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListeningAddresses not implemented")
}
func (UnimplementedShardRPCServer) Connect(context.Context, *ConnectMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedShardRPCServer) GetSlotNumber(context.Context, *SlotNumberRequest) (*SlotNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlotNumber not implemented")
}
func (UnimplementedShardRPCServer) mustEmbedUnimplementedShardRPCServer() {}

// UnsafeShardRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShardRPCServer will
// result in compilation errors.
type UnsafeShardRPCServer interface {
	mustEmbedUnimplementedShardRPCServer()
}

func RegisterShardRPCServer(s grpc.ServiceRegistrar, srv ShardRPCServer) {
	s.RegisterService(&ShardRPC_ServiceDesc, srv)
}

func _ShardRPC_AnnounceProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalAnnouncement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).AnnounceProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/AnnounceProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).AnnounceProposal(ctx, req.(*ProposalAnnouncement))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_GetBlockHashAtSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).GetBlockHashAtSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/GetBlockHashAtSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).GetBlockHashAtSlot(ctx, req.(*SlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_GenerateBlockTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockGenerationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).GenerateBlockTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/GenerateBlockTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).GenerateBlockTemplate(ctx, req.(*BlockGenerationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_SubmitBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardBlockSubmission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).SubmitBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/SubmitBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).SubmitBlock(ctx, req.(*ShardBlockSubmission))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_GetActionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShardActionStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShardRPCServer).GetActionStream(m, &shardRPCGetActionStreamServer{stream})
}

type ShardRPC_GetActionStreamServer interface {
	Send(*ShardChainAction) error
	grpc.ServerStream
}

type shardRPCGetActionStreamServer struct {
	grpc.ServerStream
}

func (x *shardRPCGetActionStreamServer) Send(m *ShardChainAction) error {
	return x.ServerStream.SendMsg(m)
}

func _ShardRPC_GetListeningAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).GetListeningAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/GetListeningAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).GetListeningAddresses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).Connect(ctx, req.(*ConnectMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardRPC_GetSlotNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlotNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardRPCServer).GetSlotNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ShardRPC/GetSlotNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardRPCServer).GetSlotNumber(ctx, req.(*SlotNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShardRPC_ServiceDesc is the grpc.ServiceDesc for ShardRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShardRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ShardRPC",
	HandlerType: (*ShardRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnnounceProposal",
			Handler:    _ShardRPC_AnnounceProposal_Handler,
		},
		{
			MethodName: "GetBlockHashAtSlot",
			Handler:    _ShardRPC_GetBlockHashAtSlot_Handler,
		},
		{
			MethodName: "GenerateBlockTemplate",
			Handler:    _ShardRPC_GenerateBlockTemplate_Handler,
		},
		{
			MethodName: "SubmitBlock",
			Handler:    _ShardRPC_SubmitBlock_Handler,
		},
		{
			MethodName: "GetListeningAddresses",
			Handler:    _ShardRPC_GetListeningAddresses_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _ShardRPC_Connect_Handler,
		},
		{
			MethodName: "GetSlotNumber",
			Handler:    _ShardRPC_GetSlotNumber_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActionStream",
			Handler:       _ShardRPC_GetActionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shardrpc.proto",
}
