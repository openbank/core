package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreateRedrawPayment(ctx context.Context, req *pb.CreateRedrawPaymentRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.redrawpayment"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
