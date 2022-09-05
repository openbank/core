package groups

import (
	"context"

	pb "bnk.to/core/api/v1/groups"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.Member, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
