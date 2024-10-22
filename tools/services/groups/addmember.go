// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package groups

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/groups"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddMember(ctx context.Context, req *pb.AddMemberRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.groups.members"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
