package cards

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/cards"
)

func (s *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.cards.transactions.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
