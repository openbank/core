// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package cards

import (
	"context"

	pb "bnk.to/core/api/v1/cards"
	"bnk.to/core/tools/db/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteHold(ctx context.Context, req *pb.DeleteHoldRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.cards.holds.delete"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	if err := storage.DeleteCardHoldByHoldID(ctx, req.HoldID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}