package scraper_test

import (
	"context"
	"log/slog"
	"testing"
)

var _ slog.Handler = (*testHandler)(nil)

type testHandler struct{}

func newTestHandler(t *testing.T) *testHandler {
	t.Helper()

	return &testHandler{}
}

func (h *testHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (h *testHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *testHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *testHandler) WithGroup(_ string) slog.Handler {
	return h
}

func newTestLogger(t *testing.T) *slog.Logger {
	t.Helper()

	return slog.New(newTestHandler(t))
}
