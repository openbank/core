package consumers

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/consumers"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteAPIKey(ctx context.Context, req *pb.DeleteAPIKeyRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.consumers.apikeys.delete"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
