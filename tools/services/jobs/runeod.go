package jobs

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/jobs"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RunEOD(ctx context.Context, _ *emptypb.Empty) (*pb.Job, error) {
	if err := s.Auth.CheckPerm(ctx, nil, "v1.jobs.eod"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
