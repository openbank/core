package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	hydra "github.com/ory/hydra-client-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Hydra is an authentication strategy that authenticates against Ory Hydra.
type Hydra struct {
	Client *hydra.APIClient
	Logger zerolog.Logger
}

// NewHydra creates an instance of Hydra based on the provided parameters.
func NewHydra(adminURL string, logger zerolog.Logger) (*Hydra, error) {
	cfg := hydra.NewConfiguration()
	parsedURL, err := url.Parse(adminURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing hydra admin URL %q: %w", adminURL, err)
	}
	cfg.Scheme = parsedURL.Scheme
	cfg.Host = parsedURL.Host
	cli := hydra.NewAPIClient(cfg)
	return &Hydra{
		Client: cli,
		Logger: logger,
	}, nil
}

// CheckPerm checks the permission of the user to perform the specified
// request. The permission name of the type of the request must be provided.
func (h Hydra) CheckPerm(ctx context.Context, req proto.Message, permName string) error {
	resp, _, err := h.introspectCtx(ctx)
	if err != nil {
		return err
	}
	if !resp.Active {
		return status.Error(codes.PermissionDenied, "INVALID_TOKEN")
	}

	// TODO: Permission check unimplemented.
	return nil
}

// InstanceName returns the instance the token is associated with.
func (h Hydra) InstanceName(ctx context.Context) (string, error) {
	resp, _, err := h.introspectCtx(ctx)
	if err != nil {
		return "", err
	}
	if !resp.Active {
		return "", status.Error(codes.PermissionDenied, "INVALID_TOKEN")
	}
	// TODO: Retrieve instance name from resp.Ext.
	return "placeholder", nil
}

// introspectCtx introspects the token provided in the context.
func (h Hydra) introspectCtx(ctx context.Context) (*hydra.OAuth2TokenIntrospection, *http.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, status.Error(codes.Unauthenticated, "INTERNAL_SERVER_ERROR")
	}
	if len(meta["authorization"]) != 1 {
		return nil, nil, status.Error(codes.Unauthenticated, "AUTHORIZATION_NOT_PROVIDED")
	}
	header := strings.SplitN(meta["authorization"][0], " ", 2)
	if len(header) != 2 || header[0] != "Bearer" {
		return nil, nil, status.Error(codes.PermissionDenied, "INVALID_TOKEN")
	}

	introspectReq := h.Client.AdminApi.IntrospectOAuth2Token(ctx)
	parsedResp, rawResp, err := introspectReq.Token(header[1]).Execute()
	if err != nil {
		h.Logger.Error().Err(err).Str("token", header[1]).Msg("error introspecting token")
		return nil, nil, status.Error(codes.PermissionDenied, "INTERNAL_SERVER_ERROR")
	}
	return parsedResp, rawResp, nil
}
