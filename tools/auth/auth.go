package auth

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// Auth is an interface defining methods required for determining the user.
type Auth interface {
	// CheckPerm checks the permission of the user based on the context.
	CheckPerm(ctx context.Context, req proto.Message, permName string) error
	// InstanceName returns the instance name the user belongs to.
	InstanceName(ctx context.Context) (string, error)
}
