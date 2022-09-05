package org

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/org"
)

func (s *HolidaysServer) UpdateNonWorkingDays(ctx context.Context, req *pb.UpdateNonWorkingDaysRequest) (*pb.NonWorkingDays, error) {
	return nil, fmt.Errorf("unimplemented")
}
