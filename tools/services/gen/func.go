package main

import (
	"fmt"
	"go/types"
	"strings"
	"text/template"
	"unicode"

	"github.com/kenshaw/inflector"
)

var extraFuncs = template.FuncMap{
	"hasPrefix":           hasPrefix,
	"hasRequestBodyField": hasRequestBodyField,
	"shortName":           shortName,
	"pbName":              pbName,
	"uniqueField":         uniqueField,
}

func hasPrefix(text string, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}

func hasRequestBodyField(m MethodData, field string) bool {
	var body *types.Struct
	for i := 0; i < m.reqType.NumFields(); i++ {
		f := m.reqType.Field(i)
		if f.Name() == "Body" {
			var ok bool
			body, ok = f.Type().Underlying().(*types.Struct)
			if !ok {
				panic(fmt.Sprintf("Body is not a struct in %s", m.ModelName))
			}
		}
	}
	if body == nil {
		return false
	}
	for i := 0; i < body.NumFields(); i++ {
		f := body.Field(i)
		if f.Name() == field {
			return true
		}
	}
	return false
}

func shortName(fullName string) (string, error) {
	r := []rune(fullName)
	for i := len(r) - 1; i >= 0; i-- {
		if !unicode.IsUpper(r[i]) {
			continue
		}
		return string(r[i:]), nil
	}
	return "", fmt.Errorf(
		"cannot locate uppercase in %q while inferring short name", fullName,
	)
}

func pbName(m MethodData) string {
	pkgSingular := inflector.Singularize(string(m.PackageName))
	pkgTitle := strings.Title(pkgSingular)
	if pkgTitle == m.ModelName {
		return pkgTitle
	}
	return strings.TrimPrefix(m.ModelName, pkgTitle)
}

func uniqueField(m MethodData) (string, error) {
	if override, ok := uniqueOverride[m.ModelName]; ok {
		return override, nil
	}
	short, err := shortName(m.ModelName)
	if err != nil {
		return "", err
	}
	return short + "ID", nil
}
