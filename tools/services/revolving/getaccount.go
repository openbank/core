// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package revolving

import (
	"context"

	pb "bnk.to/core/api/v1/revolving"
	"bnk.to/core/tools/db/mux"
)

func (s *AccountsServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.revolving.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.RevolvingAccountByAccountID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
