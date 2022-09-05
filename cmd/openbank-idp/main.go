// Command openbank-idp is the identity provider used during OpenBank authentication.
package main

import (
	"context"
	"fmt"
	"os"

	"bnk.to/core/tools/idp"
)

// version is the version.
var version = "0.0.0-dev"

func main() {
	if err := idp.Run(context.Background(), "openbank-idp", version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
