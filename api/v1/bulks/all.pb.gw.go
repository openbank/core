// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: bnk.to/core/api/v1/bulks/all.proto

/*
Package bulks is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package bulks

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code

var (
	_ io.Reader
	_ status.Status
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_BulksService_GetBulks_0(ctx context.Context, marshaler runtime.Marshaler, client BulksServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBulksRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["BulksID"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "BulksID")
	}

	protoReq.BulksID, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "BulksID", err)
	}

	msg, err := client.GetBulks(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_BulksService_GetBulks_0(ctx context.Context, marshaler runtime.Marshaler, server BulksServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBulksRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["BulksID"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "BulksID")
	}

	protoReq.BulksID, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "BulksID", err)
	}

	msg, err := server.GetBulks(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterBulksServiceHandlerServer registers the http handlers for service BulksService to "mux".
// UnaryRPC     :call BulksServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBulksServiceHandlerFromEndpoint instead.
func RegisterBulksServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BulksServiceServer) error {
	mux.Handle("GET", pattern_BulksService_GetBulks_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/openbank.core.v1.bulks.BulksService/GetBulks", runtime.WithHTTPPathPattern("/v1/bulks/{BulksID}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BulksService_GetBulks_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BulksService_GetBulks_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterBulksServiceHandlerFromEndpoint is same as RegisterBulksServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBulksServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBulksServiceHandler(ctx, mux, conn)
}

// RegisterBulksServiceHandler registers the http handlers for service BulksService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBulksServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBulksServiceHandlerClient(ctx, mux, NewBulksServiceClient(conn))
}

// RegisterBulksServiceHandlerClient registers the http handlers for service BulksService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BulksServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BulksServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BulksServiceClient" to call the correct interceptors.
func RegisterBulksServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BulksServiceClient) error {
	mux.Handle("GET", pattern_BulksService_GetBulks_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/openbank.core.v1.bulks.BulksService/GetBulks", runtime.WithHTTPPathPattern("/v1/bulks/{BulksID}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BulksService_GetBulks_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BulksService_GetBulks_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

var pattern_BulksService_GetBulks_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"v1", "bulks", "BulksID"}, ""))

var forward_BulksService_GetBulks_0 = runtime.ForwardResponseMessage