package cards

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/cards"
)

func (s *Server) GetAccountBalances(ctx context.Context, req *pb.GetAccountBalanceRequest) (*pb.AccountBalance, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.cards.balance.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
