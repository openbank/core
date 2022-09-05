// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/cards/all.proto

package cards

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

// CardsServiceClient is the client API for CardsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardsServiceClient interface {
	// GetAccountBalances retrieves the current account balances of the
	// specified card.
	GetAccountBalances(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*AccountBalance, error)
	// CreateHold creates an authorization hold.
	CreateHold(ctx context.Context, in *CreateHoldRequest, opts ...grpc.CallOption) (*Hold, error)
	// GetHold retrieves an authorization hold.
	GetHold(ctx context.Context, in *GetHoldRequest, opts ...grpc.CallOption) (*Hold, error)
	// UpdateHold updates an authorization hold.
	UpdateHold(ctx context.Context, in *UpdateHoldRequest, opts ...grpc.CallOption) (*Hold, error)
	// AdjustHold adjusts the amount of an authorization hold.
	// The amount represents the difference, which can be positive or negative.
	AdjustHold(ctx context.Context, in *AdjustHoldRequest, opts ...grpc.CallOption) (*Hold, error)
	// DeleteHold revokes an authorization hold.
	DeleteHold(ctx context.Context, in *DeleteHoldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CreateTransaction creates a financial transaction.
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	// ReverseTransaction reverses a card transaction.
	ReverseTransaction(ctx context.Context, in *ReverseTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
}

type cardsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardsServiceClient(cc grpc.ClientConnInterface) CardsServiceClient {
	return &cardsServiceClient{cc}
}

func (c *cardsServiceClient) GetAccountBalances(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*AccountBalance, error) {
	out := new(AccountBalance)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/GetAccountBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) CreateHold(ctx context.Context, in *CreateHoldRequest, opts ...grpc.CallOption) (*Hold, error) {
	out := new(Hold)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/CreateHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) GetHold(ctx context.Context, in *GetHoldRequest, opts ...grpc.CallOption) (*Hold, error) {
	out := new(Hold)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/GetHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) UpdateHold(ctx context.Context, in *UpdateHoldRequest, opts ...grpc.CallOption) (*Hold, error) {
	out := new(Hold)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/UpdateHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) AdjustHold(ctx context.Context, in *AdjustHoldRequest, opts ...grpc.CallOption) (*Hold, error) {
	out := new(Hold)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/AdjustHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) DeleteHold(ctx context.Context, in *DeleteHoldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/DeleteHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) ReverseTransaction(ctx context.Context, in *ReverseTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.CardsService/ReverseTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardsServiceServer is the server API for CardsService service.
// All implementations must embed UnimplementedCardsServiceServer
// for forward compatibility
type CardsServiceServer interface {
	// GetAccountBalances retrieves the current account balances of the
	// specified card.
	GetAccountBalances(context.Context, *GetAccountBalanceRequest) (*AccountBalance, error)
	// CreateHold creates an authorization hold.
	CreateHold(context.Context, *CreateHoldRequest) (*Hold, error)
	// GetHold retrieves an authorization hold.
	GetHold(context.Context, *GetHoldRequest) (*Hold, error)
	// UpdateHold updates an authorization hold.
	UpdateHold(context.Context, *UpdateHoldRequest) (*Hold, error)
	// AdjustHold adjusts the amount of an authorization hold.
	// The amount represents the difference, which can be positive or negative.
	AdjustHold(context.Context, *AdjustHoldRequest) (*Hold, error)
	// DeleteHold revokes an authorization hold.
	DeleteHold(context.Context, *DeleteHoldRequest) (*emptypb.Empty, error)
	// CreateTransaction creates a financial transaction.
	CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error)
	// ReverseTransaction reverses a card transaction.
	ReverseTransaction(context.Context, *ReverseTransactionRequest) (*Transaction, error)
	mustEmbedUnimplementedCardsServiceServer()
}

// UnimplementedCardsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardsServiceServer struct{}

func (UnimplementedCardsServiceServer) GetAccountBalances(context.Context, *GetAccountBalanceRequest) (*AccountBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalances not implemented")
}

func (UnimplementedCardsServiceServer) CreateHold(context.Context, *CreateHoldRequest) (*Hold, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHold not implemented")
}

func (UnimplementedCardsServiceServer) GetHold(context.Context, *GetHoldRequest) (*Hold, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHold not implemented")
}

func (UnimplementedCardsServiceServer) UpdateHold(context.Context, *UpdateHoldRequest) (*Hold, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHold not implemented")
}

func (UnimplementedCardsServiceServer) AdjustHold(context.Context, *AdjustHoldRequest) (*Hold, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustHold not implemented")
}

func (UnimplementedCardsServiceServer) DeleteHold(context.Context, *DeleteHoldRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHold not implemented")
}

func (UnimplementedCardsServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}

func (UnimplementedCardsServiceServer) ReverseTransaction(context.Context, *ReverseTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseTransaction not implemented")
}
func (UnimplementedCardsServiceServer) mustEmbedUnimplementedCardsServiceServer() {}

// UnsafeCardsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardsServiceServer will
// result in compilation errors.
type UnsafeCardsServiceServer interface {
	mustEmbedUnimplementedCardsServiceServer()
}

func RegisterCardsServiceServer(s grpc.ServiceRegistrar, srv CardsServiceServer) {
	s.RegisterService(&CardsService_ServiceDesc, srv)
}

func _CardsService_GetAccountBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).GetAccountBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/GetAccountBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).GetAccountBalances(ctx, req.(*GetAccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_CreateHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).CreateHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/CreateHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).CreateHold(ctx, req.(*CreateHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_GetHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).GetHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/GetHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).GetHold(ctx, req.(*GetHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_UpdateHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).UpdateHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/UpdateHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).UpdateHold(ctx, req.(*UpdateHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_AdjustHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).AdjustHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/AdjustHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).AdjustHold(ctx, req.(*AdjustHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_DeleteHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).DeleteHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/DeleteHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).DeleteHold(ctx, req.(*DeleteHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_ReverseTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).ReverseTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.CardsService/ReverseTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).ReverseTransaction(ctx, req.(*ReverseTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardsService_ServiceDesc is the grpc.ServiceDesc for CardsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.cards.CardsService",
	HandlerType: (*CardsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBalances",
			Handler:    _CardsService_GetAccountBalances_Handler,
		},
		{
			MethodName: "CreateHold",
			Handler:    _CardsService_CreateHold_Handler,
		},
		{
			MethodName: "GetHold",
			Handler:    _CardsService_GetHold_Handler,
		},
		{
			MethodName: "UpdateHold",
			Handler:    _CardsService_UpdateHold_Handler,
		},
		{
			MethodName: "AdjustHold",
			Handler:    _CardsService_AdjustHold_Handler,
		},
		{
			MethodName: "DeleteHold",
			Handler:    _CardsService_DeleteHold_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _CardsService_CreateTransaction_Handler,
		},
		{
			MethodName: "ReverseTransaction",
			Handler:    _CardsService_ReverseTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/cards/all.proto",
}

// HoldConfigServiceClient is the client API for HoldConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoldConfigServiceClient interface {
	// CreateHoldConfig creates a new merchant-specific authorization hold.
	CreateHoldConfig(ctx context.Context, in *CreateHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error)
	// ListHoldConfigs lists merchant-specific authorization holds.
	ListHoldConfigs(ctx context.Context, in *ListHoldConfigsRequest, opts ...grpc.CallOption) (*ListHoldConfigsResponse, error)
	// GetHoldConfig retrieves the specified merchant-specific authorization
	// hold. To retrieve the default authorization hold, use "default" as the
	// merchant code.
	GetHoldConfig(ctx context.Context, in *GetHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error)
	// UpdateHoldConfig updates a merchant-specific authorization hold. To
	// update the default authorization hold, use "default" as the merchant
	// code.
	UpdateHoldConfig(ctx context.Context, in *UpdateHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error)
	// DeleteHoldConfig deletes the specified merchant-specific authorization
	// hold.
	DeleteHoldConfig(ctx context.Context, in *DeleteHoldConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type holdConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHoldConfigServiceClient(cc grpc.ClientConnInterface) HoldConfigServiceClient {
	return &holdConfigServiceClient{cc}
}

func (c *holdConfigServiceClient) CreateHoldConfig(ctx context.Context, in *CreateHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error) {
	out := new(HoldConfig)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.HoldConfigService/CreateHoldConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdConfigServiceClient) ListHoldConfigs(ctx context.Context, in *ListHoldConfigsRequest, opts ...grpc.CallOption) (*ListHoldConfigsResponse, error) {
	out := new(ListHoldConfigsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.HoldConfigService/ListHoldConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdConfigServiceClient) GetHoldConfig(ctx context.Context, in *GetHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error) {
	out := new(HoldConfig)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.HoldConfigService/GetHoldConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdConfigServiceClient) UpdateHoldConfig(ctx context.Context, in *UpdateHoldConfigRequest, opts ...grpc.CallOption) (*HoldConfig, error) {
	out := new(HoldConfig)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.HoldConfigService/UpdateHoldConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdConfigServiceClient) DeleteHoldConfig(ctx context.Context, in *DeleteHoldConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.cards.HoldConfigService/DeleteHoldConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoldConfigServiceServer is the server API for HoldConfigService service.
// All implementations must embed UnimplementedHoldConfigServiceServer
// for forward compatibility
type HoldConfigServiceServer interface {
	// CreateHoldConfig creates a new merchant-specific authorization hold.
	CreateHoldConfig(context.Context, *CreateHoldConfigRequest) (*HoldConfig, error)
	// ListHoldConfigs lists merchant-specific authorization holds.
	ListHoldConfigs(context.Context, *ListHoldConfigsRequest) (*ListHoldConfigsResponse, error)
	// GetHoldConfig retrieves the specified merchant-specific authorization
	// hold. To retrieve the default authorization hold, use "default" as the
	// merchant code.
	GetHoldConfig(context.Context, *GetHoldConfigRequest) (*HoldConfig, error)
	// UpdateHoldConfig updates a merchant-specific authorization hold. To
	// update the default authorization hold, use "default" as the merchant
	// code.
	UpdateHoldConfig(context.Context, *UpdateHoldConfigRequest) (*HoldConfig, error)
	// DeleteHoldConfig deletes the specified merchant-specific authorization
	// hold.
	DeleteHoldConfig(context.Context, *DeleteHoldConfigRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHoldConfigServiceServer()
}

// UnimplementedHoldConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHoldConfigServiceServer struct{}

func (UnimplementedHoldConfigServiceServer) CreateHoldConfig(context.Context, *CreateHoldConfigRequest) (*HoldConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHoldConfig not implemented")
}

func (UnimplementedHoldConfigServiceServer) ListHoldConfigs(context.Context, *ListHoldConfigsRequest) (*ListHoldConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHoldConfigs not implemented")
}

func (UnimplementedHoldConfigServiceServer) GetHoldConfig(context.Context, *GetHoldConfigRequest) (*HoldConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHoldConfig not implemented")
}

func (UnimplementedHoldConfigServiceServer) UpdateHoldConfig(context.Context, *UpdateHoldConfigRequest) (*HoldConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHoldConfig not implemented")
}

func (UnimplementedHoldConfigServiceServer) DeleteHoldConfig(context.Context, *DeleteHoldConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHoldConfig not implemented")
}
func (UnimplementedHoldConfigServiceServer) mustEmbedUnimplementedHoldConfigServiceServer() {}

// UnsafeHoldConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoldConfigServiceServer will
// result in compilation errors.
type UnsafeHoldConfigServiceServer interface {
	mustEmbedUnimplementedHoldConfigServiceServer()
}

func RegisterHoldConfigServiceServer(s grpc.ServiceRegistrar, srv HoldConfigServiceServer) {
	s.RegisterService(&HoldConfigService_ServiceDesc, srv)
}

func _HoldConfigService_CreateHoldConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHoldConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldConfigServiceServer).CreateHoldConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.HoldConfigService/CreateHoldConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldConfigServiceServer).CreateHoldConfig(ctx, req.(*CreateHoldConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldConfigService_ListHoldConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHoldConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldConfigServiceServer).ListHoldConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.HoldConfigService/ListHoldConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldConfigServiceServer).ListHoldConfigs(ctx, req.(*ListHoldConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldConfigService_GetHoldConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHoldConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldConfigServiceServer).GetHoldConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.HoldConfigService/GetHoldConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldConfigServiceServer).GetHoldConfig(ctx, req.(*GetHoldConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldConfigService_UpdateHoldConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHoldConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldConfigServiceServer).UpdateHoldConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.HoldConfigService/UpdateHoldConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldConfigServiceServer).UpdateHoldConfig(ctx, req.(*UpdateHoldConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldConfigService_DeleteHoldConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHoldConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldConfigServiceServer).DeleteHoldConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.cards.HoldConfigService/DeleteHoldConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldConfigServiceServer).DeleteHoldConfig(ctx, req.(*DeleteHoldConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HoldConfigService_ServiceDesc is the grpc.ServiceDesc for HoldConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HoldConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.cards.HoldConfigService",
	HandlerType: (*HoldConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHoldConfig",
			Handler:    _HoldConfigService_CreateHoldConfig_Handler,
		},
		{
			MethodName: "ListHoldConfigs",
			Handler:    _HoldConfigService_ListHoldConfigs_Handler,
		},
		{
			MethodName: "GetHoldConfig",
			Handler:    _HoldConfigService_GetHoldConfig_Handler,
		},
		{
			MethodName: "UpdateHoldConfig",
			Handler:    _HoldConfigService_UpdateHoldConfig_Handler,
		},
		{
			MethodName: "DeleteHoldConfig",
			Handler:    _HoldConfigService_DeleteHoldConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/cards/all.proto",
}