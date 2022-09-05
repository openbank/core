package documents

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/documents"

	corepb "bnk.to/core/api/v1"
)

func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*corepb.File, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.docs.file.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
