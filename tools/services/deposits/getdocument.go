package deposits

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/deposits"

	corepb "bnk.to/core/api/v1"
)

func (s *Server) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*corepb.File, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.doc.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
