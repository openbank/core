//go:build tools

package core

import (
	_ "github.com/chanbakjsd/protoc-gen-doc"
	_ "github.com/fullstorydev/grpcurl/cmd/grpcurl"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/grpc-ecosystem/grpc-health-probe"
	_ "github.com/gunk/gunk"
	_ "github.com/kenshaw/git-buildnumber"
	_ "github.com/xo/ecosystem/cmd/protoc-gen-xo"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "mvdan.cc/gofumpt"
)
