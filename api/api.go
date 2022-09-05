package api

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// CombinedSpec produces the combined swagger spec definition for the specified
// API version, marshaling as JSON.
func CombinedSpec(version string) ([]byte, error) {
	swagger, err := combineSpecs(version)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(swagger, "", "  ")
}

// AllSpecs produces a map of each individual parsed swagger spec definition.
// The key of the map is the package name as read from the parent directory of
// the spec file.
func AllSpecs(version string) (map[string]*spec.Swagger, error) {
	paths, docs, err := loadDocs(version)
	if err != nil {
		return nil, err
	}
	specs := make(map[string]*spec.Swagger, len(paths))
	for i, path := range paths {
		pkg := filepath.Base(filepath.Dir(path))
		if spec, ok := specs[pkg]; ok {
			analysis.Mixin(spec, docs[i].Spec())
			analysis.FixEmptyResponseDescriptions(spec)
			continue
		}
		specs[pkg] = docs[i].Spec()
	}
	return specs, nil
}

// combineSpecs produces the combined swagger spec definition for the specified
// API version.
func combineSpecs(version string) (*spec.Swagger, error) {
	_, docs, err := loadDocs(version)
	if err != nil {
		return nil, err
	}
	if len(docs) == 0 {
		return nil, fmt.Errorf("swagger: no specs found (%s)", version)
	}
	swagger := docs[0].Spec()
	for _, doc := range docs[1:] {
		analysis.Mixin(swagger, doc.Spec())
	}
	analysis.FixEmptyResponseDescriptions(swagger)
	return swagger, nil
}

// loadDocs loads all API specifications for the specified version.
func loadDocs(version string) ([]string, []*loads.Document, error) {
	var paths []string
	var docs []*loads.Document
	err := fs.WalkDir(specs, ".", func(path string, e fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if e.IsDir() {
			return nil
		}
		doc, err := load(version, path)
		if err != nil {
			return err
		}
		paths = append(paths, path)
		docs = append(docs, doc)
		return nil
	})
	return paths, docs, err
}

// load loads the specified API specification.
func load(version, path string) (*loads.Document, error) {
	buf, err := specs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return loads.Analyzed(buf, "")
}

// specs are the api's embedded swagger definitions.
//
//go:embed v1/_gunk/gen/json/*/all.swagger.json
var specs embed.FS
