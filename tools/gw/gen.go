//go:build ignore
// +build ignore

package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
	"sort"
	"text/template"

	"github.com/gunk/gunk/loader"
)

func main() {
	out := flag.String("o", "svc.go", "output file")
	pkg := flag.String("p", "gw", "package name")
	flag.Parse()
	if err := run(*out, *pkg, flag.Args()...); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func run(out, pkg string, args ...string) error {
	// load gunk files
	loader := loader.Loader{
		Dir:   "",
		Fset:  token.NewFileSet(),
		Types: true,
	}
	pkgs, err := loader.Load(args...)
	if err != nil {
		return fmt.Errorf("could not load gunk files: %w", err)
	}
	data := Data{
		Package: pkg,
	}
	known := make(map[string]bool)
	for _, g := range pkgs {
		// prevent duplicate package names
		packageName := g.Name
		if known[packageName] {
			fmt.Println("skipping", g.Dir, "because package name", packageName, "is already used")
			continue
		}
		known[g.Name] = true
		// add services
		var hasService bool
		for _, f := range g.GunkSyntax {
			ast.Inspect(f, func(n ast.Node) bool {
				switch v := n.(type) {
				default:
					return false
				case *ast.GenDecl, *ast.File:
					return true
				case *ast.TypeSpec:
					if _, ok := v.Type.(*ast.InterfaceType); !ok {
						return false
					}
					data.Services = append(data.Services, Service{
						PackageName: packageName,
						Name:        v.Name.Name,
					})
					hasService = true
					return false
				}
			})
		}
		// add package
		if hasService {
			data.Packages = append(data.Packages, Package{
				Name: packageName,
				Path: g.PkgPath,
			})
		}
	}
	sort.Slice(data.Packages, func(i, j int) bool {
		return data.Packages[i].Path < data.Packages[j].Path
	})
	sort.Slice(data.Services, func(i, j int) bool {
		return data.Services[i].Name < data.Services[j].Name
	})
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, data); err != nil {
		return err
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(out, src, 0o644)
}

type Data struct {
	Package  string
	Packages []Package
	Services []Service
}

type Package struct {
	Name string
	Path string
}

type Service struct {
	Name        string
	PackageName string
}

var tpl = template.Must(template.New("").Parse(svcGoTpl))

//go:embed svc.go.tpl
var svcGoTpl string
