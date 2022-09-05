package ledgers

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/ledgers"
)

func (s *EntriesServer) CreateEntries(ctx context.Context, req *pb.CreateEntriesRequest) (*pb.CreateEntriesResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.ledgers.entries.create"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
