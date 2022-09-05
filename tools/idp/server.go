package idp

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"bnk.to/core/tools/idp/auth"
	hydra "github.com/ory/hydra-client-go"
	"github.com/rs/zerolog"
)

// Server is the implementation of the OpenBank identity provider service.
type Server struct {
	logger   zerolog.Logger
	addr     string
	hydra    *hydra.APIClient
	provider auth.Provider
}

// NewServer initializes values in the identity provider server.
func NewServer(logger zerolog.Logger, args Args) (*Server, error) {
	cfg := hydra.NewConfiguration()
	adminURL, err := url.Parse(args.HydraAdmin)
	if err != nil {
		return nil, fmt.Errorf("error parsing hydra admin URL %q: %w", args.HydraAdmin, err)
	}
	cfg.Scheme = adminURL.Scheme
	cfg.Host = adminURL.Host
	cli := hydra.NewAPIClient(cfg)
	srv := &Server{
		logger:   logger,
		addr:     args.Addr,
		hydra:    cli,
		provider: auth.NewDummy(),
	}
	return srv, nil
}

// Run starts the identity provider service.
func (s *Server) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/login", handlers(s.getLoginHandler, s.postLoginHandler))
	mux.Handle("/logout", handlers(s.getLogoutHandler, nil))
	mux.Handle("/consent", handlers(s.getConsentHandler, nil))
	return http.ListenAndServe(s.addr, mux)
}

// handlers returns a handler function that maps requests to the corresponding
// handler for the method.
func handlers(getHandler http.HandlerFunc, postHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var handler http.HandlerFunc
		switch r.Method {
		case http.MethodGet:
			handler = getHandler
		case http.MethodPost:
			handler = postHandler
		}
		if handler == nil {
			http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
