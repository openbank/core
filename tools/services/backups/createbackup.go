package backups

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/backups"
)

func (s *Server) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*pb.Backup, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.backup.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
