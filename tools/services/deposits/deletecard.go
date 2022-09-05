// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package deposits

import (
	"context"

	pb "bnk.to/core/api/v1/deposits"
	"bnk.to/core/tools/db/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.cards.delete"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	if err := storage.DeleteCardByCardID(ctx, req.CardID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}