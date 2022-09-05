// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package revolving

import (
	"context"
	"time"

	pb "bnk.to/core/api/v1/revolving"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AccountsServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.revolving.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	req.Body.CreateTime = timestamppb.New(time.Now())
	req.Body.UpdateTime = timestamppb.New(time.Now())
	v, err := db.NewRevolvingAccount(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertRevolvingAccount(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}