package util

import (
	"context"

	pb "bnk.to/core/api/v1/util"
)

func (s *Server) Echo(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	return req, nil
}
