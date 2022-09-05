package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *DepositServer) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.DepositTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.deposits.transfer"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
