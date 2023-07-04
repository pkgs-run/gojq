// gojq - Go implementation of jq
package main

import (
	"os"

	"github.com/pkgs-run/gojq/cli"
)

func main() {
	os.Exit(cli.Run())
}
