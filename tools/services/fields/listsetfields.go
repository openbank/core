package fields

import (
	"context"

	pb "bnk.to/core/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListSetFields(ctx context.Context, req *pb.ListSetFieldsRequest) (*pb.ListSetFieldsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
