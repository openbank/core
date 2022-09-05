// Package core provides the openbank core server.
package core

import (
	"context"

	"bnk.to/core/tools/cli"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// Run runs the command.
func Run(ctx context.Context, name, version string) error {
	var args Args
	c := cli.Command{
		SetupFlags: func(c *cobra.Command) {
			c.Flags().BoolVar(&args.Debug, "debug", false, "enable debug")
			c.Flags().StringVarP(&args.Addr, "listen", "l", ":9090", "listen address")
			c.Flags().StringVar(&args.Database, "db", "pg://localhost:5432", "dburl for the database")
			c.Flags().StringVar(&args.AuthType, "auth", "dummy", "backend to use for authentication")
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
	Debug      bool
	Addr       string
	Database   string
	AuthType   string
	HydraAdmin string
}
