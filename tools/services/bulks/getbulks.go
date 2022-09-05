package bulks

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/bulks"
)

func (s *Server) GetBulks(ctx context.Context, req *pb.GetBulksRequest) (*pb.GetBulksResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.bulks.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
