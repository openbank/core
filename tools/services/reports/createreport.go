package reports

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/reports"
)

func (s *Server) CreateReport(ctx context.Context, req *pb.CreateReportRequest) (*pb.Report, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.reports.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
