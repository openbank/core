// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package imports

import (
	"context"

	pb "bnk.to/core/api/v1/imports"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) UpdateImport(ctx context.Context, req *pb.UpdateImportRequest) (*pb.Import, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.imports.update"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewImport(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.UpdateImportByImportID(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}