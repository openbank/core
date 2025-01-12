// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/imports/all.proto

package imports

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

// ImportsServiceClient is the client API for ImportsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImportsServiceClient interface {
	// CreateImport creates a new file to be processed.
	CreateImport(ctx context.Context, in *CreateImportRequest, opts ...grpc.CallOption) (*Import, error)
	// ListImports lists imports based on the query.
	ListImports(ctx context.Context, in *ListImportsRequest, opts ...grpc.CallOption) (*ListImportsResponse, error)
	// GetImport retrieves the status of an import, and any errors that occurred
	// during the import.
	GetImport(ctx context.Context, in *GetImportRequest, opts ...grpc.CallOption) (*Import, error)
	// UpdateImport updates the status of an import, preparing it for importing
	// or discarding it.
	UpdateImport(ctx context.Context, in *UpdateImportRequest, opts ...grpc.CallOption) (*Import, error)
}

type importsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImportsServiceClient(cc grpc.ClientConnInterface) ImportsServiceClient {
	return &importsServiceClient{cc}
}

func (c *importsServiceClient) CreateImport(ctx context.Context, in *CreateImportRequest, opts ...grpc.CallOption) (*Import, error) {
	out := new(Import)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.imports.ImportsService/CreateImport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importsServiceClient) ListImports(ctx context.Context, in *ListImportsRequest, opts ...grpc.CallOption) (*ListImportsResponse, error) {
	out := new(ListImportsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.imports.ImportsService/ListImports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importsServiceClient) GetImport(ctx context.Context, in *GetImportRequest, opts ...grpc.CallOption) (*Import, error) {
	out := new(Import)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.imports.ImportsService/GetImport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importsServiceClient) UpdateImport(ctx context.Context, in *UpdateImportRequest, opts ...grpc.CallOption) (*Import, error) {
	out := new(Import)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.imports.ImportsService/UpdateImport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImportsServiceServer is the server API for ImportsService service.
// All implementations must embed UnimplementedImportsServiceServer
// for forward compatibility
type ImportsServiceServer interface {
	// CreateImport creates a new file to be processed.
	CreateImport(context.Context, *CreateImportRequest) (*Import, error)
	// ListImports lists imports based on the query.
	ListImports(context.Context, *ListImportsRequest) (*ListImportsResponse, error)
	// GetImport retrieves the status of an import, and any errors that occurred
	// during the import.
	GetImport(context.Context, *GetImportRequest) (*Import, error)
	// UpdateImport updates the status of an import, preparing it for importing
	// or discarding it.
	UpdateImport(context.Context, *UpdateImportRequest) (*Import, error)
	mustEmbedUnimplementedImportsServiceServer()
}

// UnimplementedImportsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImportsServiceServer struct{}

func (UnimplementedImportsServiceServer) CreateImport(context.Context, *CreateImportRequest) (*Import, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImport not implemented")
}

func (UnimplementedImportsServiceServer) ListImports(context.Context, *ListImportsRequest) (*ListImportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImports not implemented")
}

func (UnimplementedImportsServiceServer) GetImport(context.Context, *GetImportRequest) (*Import, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImport not implemented")
}

func (UnimplementedImportsServiceServer) UpdateImport(context.Context, *UpdateImportRequest) (*Import, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImport not implemented")
}
func (UnimplementedImportsServiceServer) mustEmbedUnimplementedImportsServiceServer() {}

// UnsafeImportsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImportsServiceServer will
// result in compilation errors.
type UnsafeImportsServiceServer interface {
	mustEmbedUnimplementedImportsServiceServer()
}

func RegisterImportsServiceServer(s grpc.ServiceRegistrar, srv ImportsServiceServer) {
	s.RegisterService(&ImportsService_ServiceDesc, srv)
}

func _ImportsService_CreateImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportsServiceServer).CreateImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.imports.ImportsService/CreateImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportsServiceServer).CreateImport(ctx, req.(*CreateImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImportsService_ListImports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportsServiceServer).ListImports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.imports.ImportsService/ListImports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportsServiceServer).ListImports(ctx, req.(*ListImportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImportsService_GetImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportsServiceServer).GetImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.imports.ImportsService/GetImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportsServiceServer).GetImport(ctx, req.(*GetImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImportsService_UpdateImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportsServiceServer).UpdateImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.imports.ImportsService/UpdateImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportsServiceServer).UpdateImport(ctx, req.(*UpdateImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImportsService_ServiceDesc is the grpc.ServiceDesc for ImportsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImportsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.imports.ImportsService",
	HandlerType: (*ImportsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImport",
			Handler:    _ImportsService_CreateImport_Handler,
		},
		{
			MethodName: "ListImports",
			Handler:    _ImportsService_ListImports_Handler,
		},
		{
			MethodName: "GetImport",
			Handler:    _ImportsService_GetImport_Handler,
		},
		{
			MethodName: "UpdateImport",
			Handler:    _ImportsService_UpdateImport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/imports/all.proto",
}
