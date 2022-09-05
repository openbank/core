package loans

import (
	"context"
	"fmt"

	cardspb "bnk.to/core/api/v1/cards"
	pb "bnk.to/core/api/v1/loans"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/services"
)

func (s *Server) ListCards(ctx context.Context, req *pb.ListCardsRequest) (*pb.ListCardsResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.cards.list"); err != nil {
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
	// TODO:
	// v, err := s.Storage.ListLoanCards(req.Filter, req.PageSize, req.OrderBy, afterID)
	// Stub:
	_ = after
	var v []db.Card
	var newAfter db.ListPosition
	var err error
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, fmt.Errorf("no results found")
	}
	results := make([]*cardspb.Card, 0, len(v))
	for _, w := range v {
		pb, err := w.PB()
		if err != nil {
			return nil, err
		}
		results = append(results, pb)
	}
	return &pb.ListCardsResponse{
		Cards: results,
		NextPageToken: services.EncodePageToken(
			req.Filter, req.PageSize, req.OrderBy, newAfter,
		),
	}, nil
}
