package auth

import "context"

var _ Provider = Dummy{}

// Dummy is a simple provider intended for testing which will accept all logins
// as long as the password is set to 'password1'.
type Dummy struct{}

func NewDummy() Dummy {
	return Dummy{}
}

// Authenticate authenticates the user by checking if the password is set to
// 'password1'.
func (Dummy) Authenticate(ctx context.Context, challenge Challenge) (*Identity, error) {
	if challenge.Password != "password1" {
		return nil, ErrInvalidCredentials
	}
	return &Identity{
		UserID: challenge.Username,
	}, nil
}

// Lookup always reports that all cached logins are valid as there are no
// invalidation mechanics.
func (Dummy) Lookup(ctx context.Context, id string) (*Identity, error) {
	return &Identity{
		UserID: id,
	}, nil
}
