package loans

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/loans"
)

func (s *Server) GetSchedule(ctx context.Context, req *pb.GetScheduleRequest) (*pb.Schedule, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.schedule.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
