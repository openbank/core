// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package loans

import (
	"context"

	pb "bnk.to/core/api/v1/loans"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	"bnk.to/core/tools/services"
)

func (s *Server) ListFundings(ctx context.Context, req *pb.ListFundingsRequest) (*pb.ListFundingsResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.fundings.list"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
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
	stat, v, newAfter, err := storage.ListLoanInvestorFunds(ctx, req.Filter, req.PageSize, req.OrderBy, after)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.InvestorFund, 0, len(v))
	for _, w := range v {
		pb, err := w.PB()
		if err != nil {
			return nil, err
		}
		results = append(results, pb)
	}
	return &pb.ListFundingsResponse{
		Total:     stat.Total,
		Remaining: stat.Remaining,
		Funds:     results,
		NextPageToken: services.EncodePageToken(
			req.Filter, req.PageSize, req.OrderBy, *newAfter,
		),
	}, nil
}