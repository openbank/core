// Package idp provides the openbank identity provider server.
package idp

import (
	"context"
	"embed"
	"html/template"

	"bnk.to/core/tools/cli"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// Run runs the OpenBank IdP server.
func Run(ctx context.Context, name, version string) error {
	var args Args
	c := cli.Command{
		SetupFlags: func(c *cobra.Command) {
			c.Flags().BoolVar(&args.Debug, "debug", false, "enable debug")
			c.Flags().StringVarP(&args.Addr, "listen", "l", ":9090", "listen address")
			c.Flags().StringVar(&args.HydraAdmin, "hydra-admin", "", "address of Hydra admin endpoint")
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
	// Debug is a flag to enable debug logging.
	Debug bool
	// Addr is the address for the identity provider to listen to.
	Addr string
	// HydraAdmin is the domain Hydra's admin interface is exposed at.
	HydraAdmin string
}

var (
	tpl = template.Must(template.ParseFS(tplFS, "templates/*"))
	//go:embed templates/*
	tplFS embed.FS
)
