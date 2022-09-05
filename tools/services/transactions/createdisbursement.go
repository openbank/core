package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreateDisbursement(ctx context.Context, req *pb.CreateDisbursementRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.disburse"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
