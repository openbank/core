// Command openbank-core-gw is the openbank core gateway server.
package main

import (
	"context"
	"fmt"
	"os"

	"bnk.to/core/tools/gw"
)

// version is the version.
var version = "0.0.0-dev"

func main() {
	if err := gw.Run(context.Background(), "openbank-core-gw", version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
