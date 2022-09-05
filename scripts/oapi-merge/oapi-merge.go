package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chanbakjsd/protoc-gen-doc/config"
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"

	"bnk.to/core/api"
)

var (
	configFile = flag.String("config", "", "path to file with groups")
	outDir     = flag.String("out", "specs", "path to write config files")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.LoadFile(*configFile)
	if err != nil {
		return err
	}

	specs, err := api.AllSpecs("v1")
	if err != nil {
		return err
	}

	fi, err := os.Stat(*outDir)
	switch {
	case errors.Is(err, os.ErrNotExist):
		os.Mkdir(*outDir, 0o755)
	case err != nil:
		return err
	case !fi.IsDir():
		return fmt.Errorf("path %s exists and is not a directory", *outDir)
	}

	for name, section := range cfg.Sections {
		var swagger *spec.Swagger
		for _, pkg := range section.Packages {
			spec, ok := specs[pkg]
			switch {
			case !ok:
				return fmt.Errorf("could not find package %s", pkg)
			case swagger == nil:
				swagger = spec
			default:
				analysis.Mixin(swagger, spec)
			}
		}
		if swagger == nil {
			continue
		}

		swagger.BasePath = "https://api.corebank.brankas.dev"
		swagger.Info.Title = section.DisplayName
		swagger.Info.Description = section.PreambleContent

		buf, err := json.MarshalIndent(swagger, "", "  ")
		if err != nil {
			return err
		}
		out := filepath.Join(*outDir, name+".json")
		if err := os.WriteFile(out, buf, 0o644); err != nil {
			return err
		}
	}

	return nil
}
