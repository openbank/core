// Command openbank-core is the openbank core server.
package main

import (
	"context"
	"fmt"
	"os"

	"bnk.to/core/tools/core"
)

// version is the version.
var version = "0.0.0-dev"

func main() {
	if err := core.Run(context.Background(), "openbank-core", version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
