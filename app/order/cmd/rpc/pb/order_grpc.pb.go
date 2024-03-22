// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: order.proto

package pb

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

const (
	Order_CreateHomestayOrder_FullMethodName           = "/pb.order/createHomestayOrder"
	Order_HomestayOrderDetail_FullMethodName           = "/pb.order/homestayOrderDetail"
	Order_UpdateHomestayOrderTradeState_FullMethodName = "/pb.order/updateHomestayOrderTradeState"
	Order_UserHomestayOrderList_FullMethodName         = "/pb.order/userHomestayOrderList"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// 民宿下订单
	CreateHomestayOrder(ctx context.Context, in *CreateHomestayOrderReq, opts ...grpc.CallOption) (*CreateHomestayOrderResp, error)
	// 民宿订单详情
	HomestayOrderDetail(ctx context.Context, in *HomestayOrderDetailReq, opts ...grpc.CallOption) (*HomestayOrderDetailResp, error)
	// 更新民宿订单状态
	UpdateHomestayOrderTradeState(ctx context.Context, in *UpdateHomestayOrderTradeStateReq, opts ...grpc.CallOption) (*UpdateHomestayOrderTradeStateResp, error)
	// 用户民宿订单
	UserHomestayOrderList(ctx context.Context, in *UserHomestayOrderListReq, opts ...grpc.CallOption) (*UserHomestayOrderListResp, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateHomestayOrder(ctx context.Context, in *CreateHomestayOrderReq, opts ...grpc.CallOption) (*CreateHomestayOrderResp, error) {
	out := new(CreateHomestayOrderResp)
	err := c.cc.Invoke(ctx, Order_CreateHomestayOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) HomestayOrderDetail(ctx context.Context, in *HomestayOrderDetailReq, opts ...grpc.CallOption) (*HomestayOrderDetailResp, error) {
	out := new(HomestayOrderDetailResp)
	err := c.cc.Invoke(ctx, Order_HomestayOrderDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateHomestayOrderTradeState(ctx context.Context, in *UpdateHomestayOrderTradeStateReq, opts ...grpc.CallOption) (*UpdateHomestayOrderTradeStateResp, error) {
	out := new(UpdateHomestayOrderTradeStateResp)
	err := c.cc.Invoke(ctx, Order_UpdateHomestayOrderTradeState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UserHomestayOrderList(ctx context.Context, in *UserHomestayOrderListReq, opts ...grpc.CallOption) (*UserHomestayOrderListResp, error) {
	out := new(UserHomestayOrderListResp)
	err := c.cc.Invoke(ctx, Order_UserHomestayOrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// 民宿下订单
	CreateHomestayOrder(context.Context, *CreateHomestayOrderReq) (*CreateHomestayOrderResp, error)
	// 民宿订单详情
	HomestayOrderDetail(context.Context, *HomestayOrderDetailReq) (*HomestayOrderDetailResp, error)
	// 更新民宿订单状态
	UpdateHomestayOrderTradeState(context.Context, *UpdateHomestayOrderTradeStateReq) (*UpdateHomestayOrderTradeStateResp, error)
	// 用户民宿订单
	UserHomestayOrderList(context.Context, *UserHomestayOrderListReq) (*UserHomestayOrderListResp, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateHomestayOrder(context.Context, *CreateHomestayOrderReq) (*CreateHomestayOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomestayOrder not implemented")
}
func (UnimplementedOrderServer) HomestayOrderDetail(context.Context, *HomestayOrderDetailReq) (*HomestayOrderDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayOrderDetail not implemented")
}
func (UnimplementedOrderServer) UpdateHomestayOrderTradeState(context.Context, *UpdateHomestayOrderTradeStateReq) (*UpdateHomestayOrderTradeStateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomestayOrderTradeState not implemented")
}
func (UnimplementedOrderServer) UserHomestayOrderList(context.Context, *UserHomestayOrderListReq) (*UserHomestayOrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHomestayOrderList not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateHomestayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomestayOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateHomestayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateHomestayOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateHomestayOrder(ctx, req.(*CreateHomestayOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_HomestayOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayOrderDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).HomestayOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_HomestayOrderDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).HomestayOrderDetail(ctx, req.(*HomestayOrderDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateHomestayOrderTradeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomestayOrderTradeStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateHomestayOrderTradeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateHomestayOrderTradeState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateHomestayOrderTradeState(ctx, req.(*UpdateHomestayOrderTradeStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UserHomestayOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHomestayOrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UserHomestayOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UserHomestayOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UserHomestayOrderList(ctx, req.(*UserHomestayOrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHomestayOrder",
			Handler:    _Order_CreateHomestayOrder_Handler,
		},
		{
			MethodName: "homestayOrderDetail",
			Handler:    _Order_HomestayOrderDetail_Handler,
		},
		{
			MethodName: "updateHomestayOrderTradeState",
			Handler:    _Order_UpdateHomestayOrderTradeState_Handler,
		},
		{
			MethodName: "userHomestayOrderList",
			Handler:    _Order_UserHomestayOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}