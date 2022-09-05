package clients

import (
	"context"

	pb "bnk.to/core/api/v1/clients"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListClientDocuments(ctx context.Context, req *pb.ListClientDocumentsRequest) (*pb.ListClientDocumentsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
