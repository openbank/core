package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"

	corepb "bnk.to/core/api/v1"
)

func (s *DepositServer) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*corepb.File, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.deposits.transactions.documents.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
