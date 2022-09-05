package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *DepositServer) CreateBulkDeposit(ctx context.Context, req *pb.CreateBulkDepositRequest) (*emptypb.Empty, error) {
	// TODO: Check for permission "v1.deposits.bulk"

	return nil, fmt.Errorf("unimplemented")
}
