package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Command is a wrapper around cobra.Command that sets up the logger and parses
// arguments based on environment variable and the config file.
type Command struct {
	// SetupFlags is a callback that sets up the flags to parse for.
	SetupFlags func(*cobra.Command)
	// Debug returns if the debug flag has been set. It is used to setup the
	// logger.
	Debug func() bool
	// Run runs the actual command after the config has been parsed.
	Run func(zerolog.Logger) error
}

// Execute executes the command.
func (c Command) Execute(ctx context.Context, name, version string) error {
	var configFlag string
	var configFileUsed string
	cobraCmd := &cobra.Command{
		Use:     name,
		Version: version,
		Short:   name + " server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			configFileUsed, err = parseConfig(cmd, name, configFlag)
			return err
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			var w io.Writer = os.Stderr
			if c.Debug() {
				w = zerolog.NewConsoleWriter(func(cw *zerolog.ConsoleWriter) {
					cw.Out = w
					cw.TimeFormat = "2006-01-02 15:04:05"
				})
			}
			logger := zerolog.New(w).With().Timestamp().Logger()
			if configFileUsed != "" {
				logger.Info().Str("file", configFileUsed).Msg("config")
			}
			return c.Run(logger)
		},
	}
	envConfig := strings.TrimSpace(os.Getenv(strings.ToUpper(name) + "_CONFIG"))
	cobraCmd.Flags().StringVarP(&configFlag, "config", "c", envConfig, "config file")
	c.SetupFlags(cobraCmd)
	cobraCmd.SetVersionTemplate("{{ .Name }} {{ .Version }}\n")
	cobraCmd.InitDefaultHelpCmd()
	cobraCmd.SetArgs(os.Args[1:])
	cobraCmd.SilenceErrors = true
	cobraCmd.SilenceUsage = true
	if err := cobraCmd.Execute(); err != nil {
		return fmt.Errorf("error executing command: %w", err)
	}
	return nil
}

// parseConfig parses the config. NAME_CONFIG environment variables overrides
// command flags and command flags overrides the config file.
func parseConfig(cmd *cobra.Command, name string, configFile string) (string, error) {
	v := viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
	)
	// Use the command name as the default config file name but allow
	// overriding with the configFile flag.
	configFile = strings.TrimSpace(configFile)
	if configFile == "" {
		v.AddConfigPath(".")
		v.SetConfigName(name)
	} else {
		v.SetConfigFile(configFile)
	}
	v.AutomaticEnv()
	err := v.ReadInConfig()
	switch {
	case errors.As(err, &viper.ConfigFileNotFoundError{}):
		fmt.Fprintln(os.Stderr, "error locating config file:", err)
	case err != nil:
		return "", fmt.Errorf("error loading config file: %w", err)
	}
	// Allow overriding flags using NAME_CONFIG environment flag.
	prefix := toEnvName(cmd.Name()) + "_"
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Name == "config" {
			// config flag is used to determine the config file name.
			return
		}
		_ = v.BindEnv(f.Name, prefix+toEnvName(f.Name))
		if !f.Changed && v.IsSet(f.Name) {
			_ = cmd.Flags().Set(f.Name, v.GetString(f.Name))
		}
	})
	return v.ConfigFileUsed(), nil
}

// toEnvName converts the provided string to an environment variable by
// uppercasing the name and replacing dashes with underscores.
func toEnvName(key string) string {
	return strings.ToUpper(strings.ReplaceAll(key, "-", "_"))
}
