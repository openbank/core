// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package backups

import (
	"context"

	pb "bnk.to/core/api/v1/backups"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.backup.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.BackupByBackupID(ctx, req.BackupID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
