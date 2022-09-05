package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreateRepayment(ctx context.Context, req *pb.CreateRepaymentRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.repayment"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
