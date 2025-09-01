package mcp

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/selesy/ethereum-mcp/pkg/schema"
)

func EthGetBlockByNumberTool() mcp.Tool {
	return mcp.NewToolWithRawSchema("eth_getBlockByNumber", "", json.RawMessage(schema.EthGetBlockByNumberJSON))
}
