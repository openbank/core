package imports

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/imports"
)

func (s *Server) CreateImport(ctx context.Context, req *pb.CreateImportRequest) (*pb.Import, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.imports.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
