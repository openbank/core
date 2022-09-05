package consumers

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/consumers"
)

func (s *Server) CreateAPIKey(ctx context.Context, req *pb.CreateAPIKeyRequest) (*pb.APIKey, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.consumers.apikeys.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
