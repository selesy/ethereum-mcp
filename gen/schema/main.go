package main

import (
	"os"

	"github.com/selesy/ethereum-mcp/gen/schema/internal"
)

func main() {
	os.Exit(internal.Run(os.Stderr))
}
