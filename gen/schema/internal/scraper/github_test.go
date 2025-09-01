package scraper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selesy/ethereum-mcp/gen/schema/internal/scraper"
)

// TODO: this tests the code but is NOT a unit test, requires a connection
// to GitHub and takes a while - it's also rate-limited pretty severely if
// a GitHub token is not present.  Refactor!
func TestClient(t *testing.T) {
	t.Parallel()

	s := scraper.New(newTestLogger(t))

	ctx := context.Background()

	methods, err := s.MethodSource(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, methods)

	schemas, err := s.SchemaSource(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, schemas)
}
