package main

import (
	"os"

	"github.com/selesy/ethereum-mcp/gen/internal"
)

func main() {
	os.Exit(internal.Run(os.Stderr))
}
