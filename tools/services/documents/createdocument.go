package documents

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/documents"
)

func (s *Server) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.Document, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.docs.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
