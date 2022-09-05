package org

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/org"
)

func (s *AccountingServer) UpdateConfig(ctx context.Context, req *pb.UpdateConfigRequest) (*pb.AccountingConfig, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.organization.accounting.update"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
