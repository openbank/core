// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package clients

import (
	"context"

	pb "bnk.to/core/api/v1/clients"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetClient(ctx context.Context, req *pb.GetClientRequest) (*pb.Client, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.clients.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.ClientByClientID(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
