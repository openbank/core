// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/org/all.proto

package org

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

// AccountingServiceClient is the client API for AccountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingServiceClient interface {
	// GetConfig gets the current accounting rules configuration.
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccountingConfig, error)
	// UpdateConfig updates the accounting configuration.
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*AccountingConfig, error)
}

type accountingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingServiceClient(cc grpc.ClientConnInterface) AccountingServiceClient {
	return &accountingServiceClient{cc}
}

func (c *accountingServiceClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccountingConfig, error) {
	out := new(AccountingConfig)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.AccountingService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*AccountingConfig, error) {
	out := new(AccountingConfig)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.AccountingService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServiceServer is the server API for AccountingService service.
// All implementations must embed UnimplementedAccountingServiceServer
// for forward compatibility
type AccountingServiceServer interface {
	// GetConfig gets the current accounting rules configuration.
	GetConfig(context.Context, *emptypb.Empty) (*AccountingConfig, error)
	// UpdateConfig updates the accounting configuration.
	UpdateConfig(context.Context, *UpdateConfigRequest) (*AccountingConfig, error)
	mustEmbedUnimplementedAccountingServiceServer()
}

// UnimplementedAccountingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountingServiceServer struct{}

func (UnimplementedAccountingServiceServer) GetConfig(context.Context, *emptypb.Empty) (*AccountingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}

func (UnimplementedAccountingServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*AccountingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedAccountingServiceServer) mustEmbedUnimplementedAccountingServiceServer() {}

// UnsafeAccountingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingServiceServer will
// result in compilation errors.
type UnsafeAccountingServiceServer interface {
	mustEmbedUnimplementedAccountingServiceServer()
}

func RegisterAccountingServiceServer(s grpc.ServiceRegistrar, srv AccountingServiceServer) {
	s.RegisterService(&AccountingService_ServiceDesc, srv)
}

func _AccountingService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.AccountingService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServiceServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountingService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.AccountingService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountingService_ServiceDesc is the grpc.ServiceDesc for AccountingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.org.AccountingService",
	HandlerType: (*AccountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _AccountingService_GetConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _AccountingService_UpdateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/org/all.proto",
}

// ChannelsServiceClient is the client API for ChannelsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelsServiceClient interface {
	// CreateChannel creates a new transaction channel.
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// ListChannels lists all transaction channels.
	ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error)
	// GetChannel retrieves a transaction channel.
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// UpdateChannel updates an existing transaction channel.
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// DeleteChannel deletes a transaction channel by ID.
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type channelsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelsServiceClient(cc grpc.ClientConnInterface) ChannelsServiceClient {
	return &channelsServiceClient{cc}
}

func (c *channelsServiceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.ChannelsService/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.ChannelsService/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.ChannelsService/GetChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.ChannelsService/UpdateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.ChannelsService/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelsServiceServer is the server API for ChannelsService service.
// All implementations must embed UnimplementedChannelsServiceServer
// for forward compatibility
type ChannelsServiceServer interface {
	// CreateChannel creates a new transaction channel.
	CreateChannel(context.Context, *CreateChannelRequest) (*Channel, error)
	// ListChannels lists all transaction channels.
	ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error)
	// GetChannel retrieves a transaction channel.
	GetChannel(context.Context, *GetChannelRequest) (*Channel, error)
	// UpdateChannel updates an existing transaction channel.
	UpdateChannel(context.Context, *UpdateChannelRequest) (*Channel, error)
	// DeleteChannel deletes a transaction channel by ID.
	DeleteChannel(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChannelsServiceServer()
}

// UnimplementedChannelsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelsServiceServer struct{}

func (UnimplementedChannelsServiceServer) CreateChannel(context.Context, *CreateChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}

func (UnimplementedChannelsServiceServer) ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}

func (UnimplementedChannelsServiceServer) GetChannel(context.Context, *GetChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}

func (UnimplementedChannelsServiceServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}

func (UnimplementedChannelsServiceServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChannelsServiceServer) mustEmbedUnimplementedChannelsServiceServer() {}

// UnsafeChannelsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelsServiceServer will
// result in compilation errors.
type UnsafeChannelsServiceServer interface {
	mustEmbedUnimplementedChannelsServiceServer()
}

func RegisterChannelsServiceServer(s grpc.ServiceRegistrar, srv ChannelsServiceServer) {
	s.RegisterService(&ChannelsService_ServiceDesc, srv)
}

func _ChannelsService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.ChannelsService/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.ChannelsService/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).ListChannels(ctx, req.(*ListChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.ChannelsService/GetChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.ChannelsService/UpdateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.ChannelsService/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelsService_ServiceDesc is the grpc.ServiceDesc for ChannelsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.org.ChannelsService",
	HandlerType: (*ChannelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _ChannelsService_CreateChannel_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _ChannelsService_ListChannels_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _ChannelsService_GetChannel_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _ChannelsService_UpdateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _ChannelsService_DeleteChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/org/all.proto",
}

// HolidaysServiceClient is the client API for HolidaysService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HolidaysServiceClient interface {
	// CreateHoliday creates a new holiday.
	CreateHoliday(ctx context.Context, in *CreateHolidayRequest, opts ...grpc.CallOption) (*Holiday, error)
	// ListHolidays lists holidays.
	ListHolidays(ctx context.Context, in *ListHolidaysRequest, opts ...grpc.CallOption) (*ListHolidaysResponse, error)
	// GetHoliday retrieves the specified holiday.
	GetHoliday(ctx context.Context, in *GetHolidayRequest, opts ...grpc.CallOption) (*Holiday, error)
	// UpdateHoliday updates the specified holiday.
	UpdateHoliday(ctx context.Context, in *UpdateHolidayRequest, opts ...grpc.CallOption) (*Holiday, error)
	// DeleteHoliday deletes the specified holiday.
	DeleteHoliday(ctx context.Context, in *DeleteHolidayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetNonWorkingDays retrieves the list of non-working days.
	GetNonWorkingDays(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NonWorkingDays, error)
	// UpdateNonWorkingDays updates the list of non-working days.
	UpdateNonWorkingDays(ctx context.Context, in *UpdateNonWorkingDaysRequest, opts ...grpc.CallOption) (*NonWorkingDays, error)
}

type holidaysServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHolidaysServiceClient(cc grpc.ClientConnInterface) HolidaysServiceClient {
	return &holidaysServiceClient{cc}
}

func (c *holidaysServiceClient) CreateHoliday(ctx context.Context, in *CreateHolidayRequest, opts ...grpc.CallOption) (*Holiday, error) {
	out := new(Holiday)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/CreateHoliday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) ListHolidays(ctx context.Context, in *ListHolidaysRequest, opts ...grpc.CallOption) (*ListHolidaysResponse, error) {
	out := new(ListHolidaysResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/ListHolidays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) GetHoliday(ctx context.Context, in *GetHolidayRequest, opts ...grpc.CallOption) (*Holiday, error) {
	out := new(Holiday)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/GetHoliday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) UpdateHoliday(ctx context.Context, in *UpdateHolidayRequest, opts ...grpc.CallOption) (*Holiday, error) {
	out := new(Holiday)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/UpdateHoliday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) DeleteHoliday(ctx context.Context, in *DeleteHolidayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/DeleteHoliday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) GetNonWorkingDays(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NonWorkingDays, error) {
	out := new(NonWorkingDays)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/GetNonWorkingDays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysServiceClient) UpdateNonWorkingDays(ctx context.Context, in *UpdateNonWorkingDaysRequest, opts ...grpc.CallOption) (*NonWorkingDays, error) {
	out := new(NonWorkingDays)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.HolidaysService/UpdateNonWorkingDays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HolidaysServiceServer is the server API for HolidaysService service.
// All implementations must embed UnimplementedHolidaysServiceServer
// for forward compatibility
type HolidaysServiceServer interface {
	// CreateHoliday creates a new holiday.
	CreateHoliday(context.Context, *CreateHolidayRequest) (*Holiday, error)
	// ListHolidays lists holidays.
	ListHolidays(context.Context, *ListHolidaysRequest) (*ListHolidaysResponse, error)
	// GetHoliday retrieves the specified holiday.
	GetHoliday(context.Context, *GetHolidayRequest) (*Holiday, error)
	// UpdateHoliday updates the specified holiday.
	UpdateHoliday(context.Context, *UpdateHolidayRequest) (*Holiday, error)
	// DeleteHoliday deletes the specified holiday.
	DeleteHoliday(context.Context, *DeleteHolidayRequest) (*emptypb.Empty, error)
	// GetNonWorkingDays retrieves the list of non-working days.
	GetNonWorkingDays(context.Context, *emptypb.Empty) (*NonWorkingDays, error)
	// UpdateNonWorkingDays updates the list of non-working days.
	UpdateNonWorkingDays(context.Context, *UpdateNonWorkingDaysRequest) (*NonWorkingDays, error)
	mustEmbedUnimplementedHolidaysServiceServer()
}

// UnimplementedHolidaysServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHolidaysServiceServer struct{}

func (UnimplementedHolidaysServiceServer) CreateHoliday(context.Context, *CreateHolidayRequest) (*Holiday, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHoliday not implemented")
}

func (UnimplementedHolidaysServiceServer) ListHolidays(context.Context, *ListHolidaysRequest) (*ListHolidaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolidays not implemented")
}

func (UnimplementedHolidaysServiceServer) GetHoliday(context.Context, *GetHolidayRequest) (*Holiday, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHoliday not implemented")
}

func (UnimplementedHolidaysServiceServer) UpdateHoliday(context.Context, *UpdateHolidayRequest) (*Holiday, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHoliday not implemented")
}

func (UnimplementedHolidaysServiceServer) DeleteHoliday(context.Context, *DeleteHolidayRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHoliday not implemented")
}

func (UnimplementedHolidaysServiceServer) GetNonWorkingDays(context.Context, *emptypb.Empty) (*NonWorkingDays, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNonWorkingDays not implemented")
}

func (UnimplementedHolidaysServiceServer) UpdateNonWorkingDays(context.Context, *UpdateNonWorkingDaysRequest) (*NonWorkingDays, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNonWorkingDays not implemented")
}
func (UnimplementedHolidaysServiceServer) mustEmbedUnimplementedHolidaysServiceServer() {}

// UnsafeHolidaysServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HolidaysServiceServer will
// result in compilation errors.
type UnsafeHolidaysServiceServer interface {
	mustEmbedUnimplementedHolidaysServiceServer()
}

func RegisterHolidaysServiceServer(s grpc.ServiceRegistrar, srv HolidaysServiceServer) {
	s.RegisterService(&HolidaysService_ServiceDesc, srv)
}

func _HolidaysService_CreateHoliday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHolidayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).CreateHoliday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/CreateHoliday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).CreateHoliday(ctx, req.(*CreateHolidayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_ListHolidays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHolidaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).ListHolidays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/ListHolidays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).ListHolidays(ctx, req.(*ListHolidaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_GetHoliday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHolidayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).GetHoliday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/GetHoliday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).GetHoliday(ctx, req.(*GetHolidayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_UpdateHoliday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHolidayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).UpdateHoliday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/UpdateHoliday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).UpdateHoliday(ctx, req.(*UpdateHolidayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_DeleteHoliday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHolidayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).DeleteHoliday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/DeleteHoliday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).DeleteHoliday(ctx, req.(*DeleteHolidayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_GetNonWorkingDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).GetNonWorkingDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/GetNonWorkingDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).GetNonWorkingDays(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HolidaysService_UpdateNonWorkingDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNonWorkingDaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolidaysServiceServer).UpdateNonWorkingDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.HolidaysService/UpdateNonWorkingDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolidaysServiceServer).UpdateNonWorkingDays(ctx, req.(*UpdateNonWorkingDaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HolidaysService_ServiceDesc is the grpc.ServiceDesc for HolidaysService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HolidaysService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.org.HolidaysService",
	HandlerType: (*HolidaysServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHoliday",
			Handler:    _HolidaysService_CreateHoliday_Handler,
		},
		{
			MethodName: "ListHolidays",
			Handler:    _HolidaysService_ListHolidays_Handler,
		},
		{
			MethodName: "GetHoliday",
			Handler:    _HolidaysService_GetHoliday_Handler,
		},
		{
			MethodName: "UpdateHoliday",
			Handler:    _HolidaysService_UpdateHoliday_Handler,
		},
		{
			MethodName: "DeleteHoliday",
			Handler:    _HolidaysService_DeleteHoliday_Handler,
		},
		{
			MethodName: "GetNonWorkingDays",
			Handler:    _HolidaysService_GetNonWorkingDays_Handler,
		},
		{
			MethodName: "UpdateNonWorkingDays",
			Handler:    _HolidaysService_UpdateNonWorkingDays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/org/all.proto",
}

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	// GetOrganization retrieves the details of the organization.
	GetOrganization(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Organization, error)
	// UpdateOrganization updates the details of the organization.
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) GetOrganization(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.OrganizationService/GetOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.OrganizationService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations must embed UnimplementedOrganizationServiceServer
// for forward compatibility
type OrganizationServiceServer interface {
	// GetOrganization retrieves the details of the organization.
	GetOrganization(context.Context, *emptypb.Empty) (*Organization, error)
	// UpdateOrganization updates the details of the organization.
	UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*Organization, error)
	mustEmbedUnimplementedOrganizationServiceServer()
}

// UnimplementedOrganizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationServiceServer struct{}

func (UnimplementedOrganizationServiceServer) GetOrganization(context.Context, *emptypb.Empty) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganization not implemented")
}

func (UnimplementedOrganizationServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) mustEmbedUnimplementedOrganizationServiceServer() {}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
	s.RegisterService(&OrganizationService_ServiceDesc, srv)
}

func _OrganizationService_GetOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).GetOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.OrganizationService/GetOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).GetOrganization(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.OrganizationService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationService_ServiceDesc is the grpc.ServiceDesc for OrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.org.OrganizationService",
	HandlerType: (*OrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrganization",
			Handler:    _OrganizationService_GetOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _OrganizationService_UpdateOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/org/all.proto",
}

// SetupServiceClient is the client API for SetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SetupServiceClient interface {
	// GetSetup retrieves the general setup.
	GetSetup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Setup, error)
	// UpdateSetup updates the general setup.
	UpdateSetup(ctx context.Context, in *UpdateSetupRequest, opts ...grpc.CallOption) (*Setup, error)
}

type setupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSetupServiceClient(cc grpc.ClientConnInterface) SetupServiceClient {
	return &setupServiceClient{cc}
}

func (c *setupServiceClient) GetSetup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Setup, error) {
	out := new(Setup)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.SetupService/GetSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setupServiceClient) UpdateSetup(ctx context.Context, in *UpdateSetupRequest, opts ...grpc.CallOption) (*Setup, error) {
	out := new(Setup)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.org.SetupService/UpdateSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetupServiceServer is the server API for SetupService service.
// All implementations must embed UnimplementedSetupServiceServer
// for forward compatibility
type SetupServiceServer interface {
	// GetSetup retrieves the general setup.
	GetSetup(context.Context, *emptypb.Empty) (*Setup, error)
	// UpdateSetup updates the general setup.
	UpdateSetup(context.Context, *UpdateSetupRequest) (*Setup, error)
	mustEmbedUnimplementedSetupServiceServer()
}

// UnimplementedSetupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSetupServiceServer struct{}

func (UnimplementedSetupServiceServer) GetSetup(context.Context, *emptypb.Empty) (*Setup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetup not implemented")
}

func (UnimplementedSetupServiceServer) UpdateSetup(context.Context, *UpdateSetupRequest) (*Setup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSetup not implemented")
}
func (UnimplementedSetupServiceServer) mustEmbedUnimplementedSetupServiceServer() {}

// UnsafeSetupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SetupServiceServer will
// result in compilation errors.
type UnsafeSetupServiceServer interface {
	mustEmbedUnimplementedSetupServiceServer()
}

func RegisterSetupServiceServer(s grpc.ServiceRegistrar, srv SetupServiceServer) {
	s.RegisterService(&SetupService_ServiceDesc, srv)
}

func _SetupService_GetSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetupServiceServer).GetSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.SetupService/GetSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetupServiceServer).GetSetup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetupService_UpdateSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetupServiceServer).UpdateSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.org.SetupService/UpdateSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetupServiceServer).UpdateSetup(ctx, req.(*UpdateSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SetupService_ServiceDesc is the grpc.ServiceDesc for SetupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SetupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.org.SetupService",
	HandlerType: (*SetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSetup",
			Handler:    _SetupService_GetSetup_Handler,
		},
		{
			MethodName: "UpdateSetup",
			Handler:    _SetupService_UpdateSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/org/all.proto",
}
