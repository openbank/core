package deposits

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/deposits"
)

func (s *Server) GetFundedAccounts(ctx context.Context, req *pb.GetLoanAccountsRequest) (*pb.GetLoanAccountsResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.funding.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
