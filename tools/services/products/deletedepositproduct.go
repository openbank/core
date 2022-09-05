// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package products

import (
	"context"

	pb "bnk.to/core/api/v1/products"
	"bnk.to/core/tools/db/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *DepositsServer) DeleteDepositProduct(ctx context.Context, req *pb.DeleteDepositProductRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.products.deposit.delete"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	if err := storage.DeleteProductDepositProductByProductID(ctx, req.ProductID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}