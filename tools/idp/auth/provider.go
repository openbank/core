package auth

import (
	"context"
	"errors"
)

// ErrInvalidCredentials or its wrapped equivalent should be returned by
// Provider if the provided username or password is incorrect.
var ErrInvalidCredentials = errors.New("invalid credentials provided")

// Challenge is the details that the user is using to authenticate to the service.
type Challenge struct {
	// Username is the username identifying the user.
	Username string
	// Password is the secret phrase set by the user.
	Password string
}

// Identity is the validated identity of the user by the provider.
type Identity struct {
	UserID string
}

// Provider is a source for the IdP service to query from.
type Provider interface {
	// Authenticate tries to authenticate with given username and password in parameters.
	Authenticate(ctx context.Context, challenge Challenge) (*Identity, error)

	// Lookup requests the provider to check if the specified ID is still valid
	// and return its corresponding identity. It is used to check for the
	// validity of cached successes by Hydra.
	Lookup(ctx context.Context, id string) (*Identity, error)
}
