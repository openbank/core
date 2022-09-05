// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package deposits

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/deposits"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UndoMaturity(ctx context.Context, req *pb.UndoMaturityRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.undomaturity"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}