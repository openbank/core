package consumers

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/consumers"
)

func (s *Server) GetAPIKey(ctx context.Context, req *pb.GetAPIKeyRequest) (*pb.APIKey, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.consumers.apikeys.get"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
