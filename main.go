package main

import (
	"context"
	_ "embed"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/selesy/ethereum-mcp/pkg/tool"
)

var instructions string

func main() {
	ctx := context.Background()

	log := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level: slog.LevelDebug,
	}))
	log.InfoContext(ctx, "Started Ethereum MCP server")

	srv := server.NewMCPServer(
		"Ethereum JSON-RPC server",
		"v1.0",
		server.WithInstructions(instructions),
	)

	srv.AddTool(mcp.NewTool("version", mcp.WithDescription("Returns the version of the server.")), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return &mcp.CallToolResult{}, nil
	})

	for k, v := range tool.Tools() {
		srv.AddTool(v, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return &mcp.CallToolResult{}, nil
		})

		log.DebugContext(ctx, "Registered tool", slog.String("name", k))
	}

	if err := server.NewStdioServer(srv).Listen(ctx, os.Stdin, os.Stdout); err != nil {
		log.Error("Failed to start the stdio server", tint.Err(err))
		os.Exit(1)
	}
}
