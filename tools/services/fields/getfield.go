package fields

import (
	"context"

	pb "bnk.to/core/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetField(ctx context.Context, req *pb.GetFieldRequest) (*pb.Field, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
