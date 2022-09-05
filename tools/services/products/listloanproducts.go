// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package products

import (
	"context"

	pb "bnk.to/core/api/v1/products"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	"bnk.to/core/tools/services"
)

func (s *LoansServer) ListLoanProducts(ctx context.Context, req *pb.ListLoanProductsRequest) (*pb.ListLoanProductsResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.products.loan.list"); err != nil {
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
	stat, v, newAfter, err := storage.ListProductLoanProducts(ctx, req.Filter, req.PageSize, req.OrderBy, after)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.LoanProduct, 0, len(v))
	for _, w := range v {
		pb, err := w.PB()
		if err != nil {
			return nil, err
		}
		results = append(results, pb)
	}
	return &pb.ListLoanProductsResponse{
		Total:     stat.Total,
		Remaining: stat.Remaining,
		Products:  results,
		NextPageToken: services.EncodePageToken(
			req.Filter, req.PageSize, req.OrderBy, *newAfter,
		),
	}, nil
}