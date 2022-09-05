package products

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/products"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *DepositsServer) UpdateDepositAccounts(ctx context.Context, req *pb.UpdateDepositAccountsRequest) (*emptypb.Empty, error) {
	return nil, fmt.Errorf("unimplemented")
}
