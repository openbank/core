package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreateWithdrawalRedraw(ctx context.Context, req *pb.CreateWithdrawalRedrawRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.redrawwithdrawal"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
