// Package gw provides the openbank core gateway server.
package gw

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"bnk.to/core/api"
	"bnk.to/core/tools/cli"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
)

// Run runs the openbank core gateway server.
func Run(ctx context.Context, name, version string) error {
	var args Args
	c := cli.Command{
		SetupFlags: func(c *cobra.Command) {
			c.Flags().BoolVar(&args.Debug, "debug", false, "enable debug")
			c.Flags().StringVarP(&args.Addr, "listen", "l", ":9091", "listen address")
			c.Flags().StringVar(&args.Endpoint, "endpoint", "127.0.0.1:9090", "grpc endpoint")
			c.Flags().StringSliceVar(&args.Origins, "origins", nil, "cors origins")
		},
		Debug: func() bool {
			return args.Debug
		},
		Run: func(logger zerolog.Logger) error {
			s, err := NewServer(logger, args)
			if err != nil {
				logger.Error().Err(err).Msg("init")
				return err
			}
			if err := s.Run(ctx); err != nil {
				logger.Error().Err(err).Msg("run")
				return err
			}
			return nil
		},
	}
	return c.Execute(ctx, name, version)
}

// Args are the command-line arguments.
type Args struct {
	Debug    bool
	Addr     string
	Endpoint string
	Origins  []string
}

// Server is the openbank core gateway server.
type Server struct {
	logger   zerolog.Logger
	addr     string
	endpoint string
	origins  []string
}

// NewServer creates a openbank core gateway server.
func NewServer(logger zerolog.Logger, args Args) (*Server, error) {
	s := &Server{
		logger:   logger,
		addr:     args.Addr,
		endpoint: args.Endpoint,
		origins:  args.Origins,
	}
	return s, nil
}

// Run runs the openbank core gateway server.
func (s *Server) Run(ctx context.Context) error {
	// listen
	s.logger.Info().Str("addr", s.addr).Msg("listen")
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", s.addr)
	if err != nil {
		return err
	}
	defer l.Close() //nolint:errcheck
	// build mux
	mux, err := s.mux(ctx)
	if err != nil {
		return err
	}
	// serve
	return http.Serve(l, s.logmw(mux))
}

// mux builds the serve mux.
func (s *Server) mux(ctx context.Context) (http.Handler, error) {
	// build gateway
	gw, err := s.gw(ctx)
	if err != nil {
		return nil, err
	}
	// create mux, handle gw, spec, docs, index
	spec, err := s.spec()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/v1/", gw)
	mux.Handle("/spec/", spec)
	return mux, nil
}

// gw builds the gateway handler.
func (s *Server) gw(ctx context.Context) (http.Handler, error) {
	s.logger.Info().Str("endpoint", s.endpoint).Msg("dial")
	// build gateway mux
	gw, opts := gwruntime.NewServeMux(), []grpc.DialOption{grpc.WithInsecure()}
	pkgs := make(map[string]bool)
	for _, handler := range Services {
		if err := handler(ctx, gw, s.endpoint, opts); err != nil {
			return nil, err
		}
		// determine the handler, package, name
		handlerName := strings.TrimPrefix(runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(), "bnk.to/core/api/v1/")
		i := strings.Index(handlerName, ".")
		pkgName := handlerName[:i]
		name := strings.TrimPrefix(strings.TrimSuffix(handlerName[i+1:], "HandlerFromEndpoint"), "Register")
		s.logger.Info().Str("pkg", pkgName).Str("name", name).Msg("service")
		pkgs[pkgName] = true
	}
	return s.cors(gw), nil
}

func (s *Server) cors(h http.Handler) http.Handler {
	origins := make(map[string]struct{})
	for _, v := range s.origins {
		origins[v] = struct{}{}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := origins[r.Header.Get("Origin")]
		if ok {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

// spec builds the spec handler.
func (s *Server) spec() (http.Handler, error) {
	spec, err := api.CombinedSpec("v1")
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		_, _ = res.Write(spec)
	}), nil
}

// logmw is the logging middleware.
func (s *Server) logmw(h http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		requestID := ksuid.New().String()
		logger := s.logger.With().
			Str("ip", req.RemoteAddr).
			Str("method", req.Method).
			Str("path", req.URL.Path).
			Str("request_id", requestID).
			Logger()
		logger.Info().Msg("req")
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				logger.Info().Bytes("stack", stack).Err(fmt.Errorf("%v", err)).Msg("internal server error")
				fmt.Fprintf(os.Stderr, "------------------\nRECOVERED: %v\n\n%s\n------------------\n", err, string(stack))
				http.Error(res, "internal server error", http.StatusInternalServerError)
			}
		}()
		ctx := context.WithValue(req.Context(), LoggerKey, logger)
		rw := negroni.NewResponseWriter(res)
		h.ServeHTTP(rw, req.WithContext(ctx))
		logger.Info().Int("code", rw.Status()).Msg("res")
	}
}

// ContextKey is a context key.
type ContextKey int

// Context keys.
const (
	LoggerKey ContextKey = iota
)
