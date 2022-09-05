package fields

import (
	"context"

	pb "bnk.to/core/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteField(ctx context.Context, req *pb.DeleteFieldRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
