// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package cards

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/cards"
)

func (s *Server) AdjustHold(ctx context.Context, req *pb.AdjustHoldRequest) (*pb.Hold, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.cards.holds.adjust"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
