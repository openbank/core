package auth

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// Dummy is an authentication strategy that considers all requests authenticated.
type Dummy struct{}

// CheckPerm returns a nil error, simulating a successful permission check.
func (Dummy) CheckPerm(ctx context.Context, req proto.Message, permName string) error {
	return nil
}

// InstanceName returns an empty string.
func (Dummy) InstanceName(ctx context.Context) (string, error) {
	return "", nil
}
