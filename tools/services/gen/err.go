package main

import (
	"strings"

	"golang.org/x/tools/go/packages"
)

// multiError is a list of errors combined as one error.
type multiError []error

// Error returns the errors combined in one.
func (m multiError) Error() string {
	msg := make([]string, len(m))
	for k, v := range m {
		msg[k] = v.Error()
	}
	return strings.Join(msg, "\n\t")
}

func pkgErrorToMultiError(p []packages.Error) multiError {
	err := make(multiError, len(p))
	for k, v := range p {
		err[k] = v
	}
	return err
}
