package mcp_test

import (
	"testing"

	"github.com/selesy/ethereum-mcp/pkg/mcp"
)

func Test(t *testing.T) {
	t.Parallel()

	tool := mcp.EthGetBlockByNumberTool()

	t.Log(tool)
	t.Log(tool.GetName())
}
