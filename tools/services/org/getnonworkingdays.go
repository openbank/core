package org

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/org"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *HolidaysServer) GetNonWorkingDays(ctx context.Context, _ *emptypb.Empty) (*pb.NonWorkingDays, error) {
	if err := s.Auth.CheckPerm(ctx, nil, "v1.organization.holidays.nonworkingdays.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
