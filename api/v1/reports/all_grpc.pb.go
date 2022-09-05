// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/reports/all.proto

package reports

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

// ReportsServiceClient is the client API for ReportsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportsServiceClient interface {
	// CreateReport creates an accounting report.
	CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*Report, error)
	// ListReports lists the reports.
	ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error)
	// GetReport retrieves an accounting report.
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error)
}

type reportsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportsServiceClient(cc grpc.ClientConnInterface) ReportsServiceClient {
	return &reportsServiceClient{cc}
}

func (c *reportsServiceClient) CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.reports.ReportsService/CreateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error) {
	out := new(ListReportsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.reports.ReportsService/ListReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.reports.ReportsService/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsServiceServer is the server API for ReportsService service.
// All implementations must embed UnimplementedReportsServiceServer
// for forward compatibility
type ReportsServiceServer interface {
	// CreateReport creates an accounting report.
	CreateReport(context.Context, *CreateReportRequest) (*Report, error)
	// ListReports lists the reports.
	ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error)
	// GetReport retrieves an accounting report.
	GetReport(context.Context, *GetReportRequest) (*Report, error)
	mustEmbedUnimplementedReportsServiceServer()
}

// UnimplementedReportsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportsServiceServer struct{}

func (UnimplementedReportsServiceServer) CreateReport(context.Context, *CreateReportRequest) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReport not implemented")
}

func (UnimplementedReportsServiceServer) ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReports not implemented")
}

func (UnimplementedReportsServiceServer) GetReport(context.Context, *GetReportRequest) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedReportsServiceServer) mustEmbedUnimplementedReportsServiceServer() {}

// UnsafeReportsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportsServiceServer will
// result in compilation errors.
type UnsafeReportsServiceServer interface {
	mustEmbedUnimplementedReportsServiceServer()
}

func RegisterReportsServiceServer(s grpc.ServiceRegistrar, srv ReportsServiceServer) {
	s.RegisterService(&ReportsService_ServiceDesc, srv)
}

func _ReportsService_CreateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).CreateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.reports.ReportsService/CreateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).CreateReport(ctx, req.(*CreateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_ListReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).ListReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.reports.ReportsService/ListReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).ListReports(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.reports.ReportsService/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportsService_ServiceDesc is the grpc.ServiceDesc for ReportsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.reports.ReportsService",
	HandlerType: (*ReportsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReport",
			Handler:    _ReportsService_CreateReport_Handler,
		},
		{
			MethodName: "ListReports",
			Handler:    _ReportsService_ListReports_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _ReportsService_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/reports/all.proto",
}