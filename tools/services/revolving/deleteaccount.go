// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package revolving

import (
	"context"

	pb "bnk.to/core/api/v1/revolving"
	"bnk.to/core/tools/db/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *AccountsServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.revolving.delete"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	if err := storage.DeleteRevolvingAccountByAccountID(ctx, req.AccountID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}