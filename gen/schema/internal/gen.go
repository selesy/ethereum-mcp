package internal

import (
	"context"
	"io"
	"log/slog"

	"github.com/lmittmann/tint"

	"github.com/selesy/ethereum-mcp/gen/schema/internal/generator"
	"github.com/selesy/ethereum-mcp/gen/schema/internal/scraper"
)

// Run scrapes Ethereum execution APIs and generates complete JSONSchemas
// for each OpenRPC method found.  The word "complete" indicates that the
// referenced OpenRPC definitions are merged into each JSONSchema method.
func Run(stderr io.Writer) int {
	ctx := context.Background()

	log := slog.New(tint.NewHandler(stderr, &tint.Options{
		Level: slog.LevelDebug,
	}))
	log.InfoContext(ctx, "Started Ethereum execution API schema generator")
	log.DebugContext(ctx, "Debug logging enabled")

	if err := run(ctx, log); err != nil {
		log.ErrorContext(ctx, "Finished Ethereum execution API schema generator", tint.Err(err))

		return 1
	}

	log.InfoContext(ctx, "Finished Ethereum execution API schema generator")

	return 0
}

func run(ctx context.Context, log *slog.Logger) error {
	methods, err := scraper.New(log).Run(ctx)
	if err != nil {
		return err
	}

	err = generator.New(log).Run(ctx, methods)
	if err != nil {
		return err
	}

	return nil
}
