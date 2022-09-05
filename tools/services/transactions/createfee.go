package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreateFee(ctx context.Context, req *pb.CreateLoanFeeRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.applyfee"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
