package consumers

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/consumers"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/services"
)

func (s *Server) ListAPIKeys(ctx context.Context, req *pb.ListAPIKeysRequest) (*pb.ListAPIKeysResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.consumers.apikeys.list"); err != nil {
		return nil, err
	}

	var after *db.ListPosition
	if req.PageToken != "" {
		if req.Filter != "" {
			return nil, services.ErrFilterOnPagination
		}
		if req.PageSize != 0 {
			return nil, services.ErrSizeOnPagination
		}
		if req.OrderBy != "" {
			return nil, services.ErrOrderOnPagination
		}
		var err error
		req.Filter, req.PageSize, req.OrderBy, after, err = services.DecodePageToken(req.PageToken)
		if err != nil {
			return nil, err
		}
	}
	if req.PageSize == 0 {
		req.PageSize = services.DefaultPageSize
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		return nil, services.ErrInvalidPageSize
	}

	// Stub:
	_ = after
	return nil, fmt.Errorf("unimplemented")
	// return &pb.ListAPIKeysResponse{
	// 	Keys: results,
	// 	NextPageToken: services.EncodePageToken(
	// 		req.Filter, req.PageSize, req.OrderBy, v[len(v)-1].ID,
	// 	),
	// }, nil
}
