package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *DepositServer) CreateDeposit(ctx context.Context, req *pb.CreateDepositRequest) (*pb.DepositTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.deposits.deposit"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
