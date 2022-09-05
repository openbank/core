package groups

import (
	"context"

	pb "bnk.to/core/api/v1/groups"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListMemberRoles(ctx context.Context, req *pb.ListMemberRolesRequest) (*pb.ListMemberRolesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
