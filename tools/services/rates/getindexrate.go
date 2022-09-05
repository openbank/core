// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package rates

import (
	"context"

	pb "bnk.to/core/api/v1/rates"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetIndexRate(ctx context.Context, req *pb.GetIndexRateRequest) (*pb.IndexRate, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.sources.indexrates.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.RateIndexRateByRateID(ctx, req.RateID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}