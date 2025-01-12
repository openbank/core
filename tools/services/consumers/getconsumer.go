// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package consumers

import (
	"context"

	pb "bnk.to/core/api/v1/consumers"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetConsumer(ctx context.Context, req *pb.GetConsumerRequest) (*pb.Consumer, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.consumers.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.ConsumerByConsumerID(ctx, req.ConsumerID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
