package schema_test

import (
	"github.com/mark3labs/mcp-go/mcp"

	"github.com/selesy/ethereum-mcp/pkg/schema"
)

func ExampleEthGetBlockByNumberSchema() { //nolint:govet
	tool := mcp.NewToolWithRawSchema(
		"eth_getBlockByNumber",
		"Returns information about a block by number.",
		schema.EthGetBlockByNumberSchema,
	)

	// Do something with the created mcp.Tool.
	_ = tool
}
