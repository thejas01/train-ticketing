// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/train.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrainServiceClient is the client API for TrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainServiceClient interface {
	PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*Receipt, error)
	AllocateSeat(ctx context.Context, in *SeatAllocationRequest, opts ...grpc.CallOption) (*SeatAllocationResponse, error)
	GetReceiptDetails(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*Receipt, error)
	ViewUsersBySection(ctx context.Context, in *ViewUsersRequest, opts ...grpc.CallOption) (TrainService_ViewUsersBySectionClient, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
}

type trainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainServiceClient(cc grpc.ClientConnInterface) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/TrainService/PurchaseTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) AllocateSeat(ctx context.Context, in *SeatAllocationRequest, opts ...grpc.CallOption) (*SeatAllocationResponse, error) {
	out := new(SeatAllocationResponse)
	err := c.cc.Invoke(ctx, "/TrainService/AllocateSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) GetReceiptDetails(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/TrainService/GetReceiptDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewUsersBySection(ctx context.Context, in *ViewUsersRequest, opts ...grpc.CallOption) (TrainService_ViewUsersBySectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrainService_ServiceDesc.Streams[0], "/TrainService/ViewUsersBySection", opts...)
	if err != nil {
		return nil, err
	}
	x := &trainServiceViewUsersBySectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrainService_ViewUsersBySectionClient interface {
	Recv() (*SeatDetails, error)
	grpc.ClientStream
}

type trainServiceViewUsersBySectionClient struct {
	grpc.ClientStream
}

func (x *trainServiceViewUsersBySectionClient) Recv() (*SeatDetails, error) {
	m := new(SeatDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trainServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/TrainService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, "/TrainService/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainServiceServer is the server API for TrainService service.
// All implementations must embed UnimplementedTrainServiceServer
// for forward compatibility
type TrainServiceServer interface {
	PurchaseTicket(context.Context, *TicketRequest) (*Receipt, error)
	AllocateSeat(context.Context, *SeatAllocationRequest) (*SeatAllocationResponse, error)
	GetReceiptDetails(context.Context, *ReceiptRequest) (*Receipt, error)
	ViewUsersBySection(*ViewUsersRequest, TrainService_ViewUsersBySectionServer) error
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	mustEmbedUnimplementedTrainServiceServer()
}

// UnimplementedTrainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainServiceServer struct {
}

func (UnimplementedTrainServiceServer) PurchaseTicket(context.Context, *TicketRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainServiceServer) AllocateSeat(context.Context, *SeatAllocationRequest) (*SeatAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateSeat not implemented")
}
func (UnimplementedTrainServiceServer) GetReceiptDetails(context.Context, *ReceiptRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiptDetails not implemented")
}
func (UnimplementedTrainServiceServer) ViewUsersBySection(*ViewUsersRequest, TrainService_ViewUsersBySectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ViewUsersBySection not implemented")
}
func (UnimplementedTrainServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainServiceServer) mustEmbedUnimplementedTrainServiceServer() {}

// UnsafeTrainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainServiceServer will
// result in compilation errors.
type UnsafeTrainServiceServer interface {
	mustEmbedUnimplementedTrainServiceServer()
}

func RegisterTrainServiceServer(s grpc.ServiceRegistrar, srv TrainServiceServer) {
	s.RegisterService(&TrainService_ServiceDesc, srv)
}

func _TrainService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/PurchaseTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_AllocateSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).AllocateSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/AllocateSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).AllocateSeat(ctx, req.(*SeatAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_GetReceiptDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).GetReceiptDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/GetReceiptDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).GetReceiptDetails(ctx, req.(*ReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewUsersBySection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ViewUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrainServiceServer).ViewUsersBySection(m, &trainServiceViewUsersBySectionServer{stream})
}

type TrainService_ViewUsersBySectionServer interface {
	Send(*SeatDetails) error
	grpc.ServerStream
}

type trainServiceViewUsersBySectionServer struct {
	grpc.ServerStream
}

func (x *trainServiceViewUsersBySectionServer) Send(m *SeatDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _TrainService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainService_ServiceDesc is the grpc.ServiceDesc for TrainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainService_PurchaseTicket_Handler,
		},
		{
			MethodName: "AllocateSeat",
			Handler:    _TrainService_AllocateSeat_Handler,
		},
		{
			MethodName: "GetReceiptDetails",
			Handler:    _TrainService_GetReceiptDetails_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainService_ModifySeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ViewUsersBySection",
			Handler:       _TrainService_ViewUsersBySection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/train.proto",
}
