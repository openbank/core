package branches

import (
	"context"

	pb "bnk.to/core/api/v1/branches"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBranchHolidays(ctx context.Context, req *pb.ListBranchHolidaysRequest) (*pb.ListBranchHolidaysResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
