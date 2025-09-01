// Package scraper scrapes methods and merges definition schemas to return
// "complete" JSONScehma methods.
package scraper

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/go-github/v74/github"

	"github.com/selesy/ethereum-mcp/pkg/openrpc"
)

// Scraper downloads and parses OpenRPC files from the Ethereum APIs
// repository
type Scraper struct {
	// cl is the GitHub client used to retrieve the OpenRPC documents.
	// TODO: make this an interface so that it can be mocked for testing.
	cl *github.Client

	// log is the structured logger.
	log *slog.Logger
}

// New creates a new Scraper.
func New(log *slog.Logger) *Scraper {
	cl := github.NewClient(nil)

	// TODO: get this env var from config
	tkn, ok := os.LookupEnv("RELEASE_PLEASE_PAT")
	if ok {
		cl = cl.WithAuthToken(tkn)
	}

	return &Scraper{
		cl:  cl,
		log: log,
	}
}

// Run executes the scraper with the provided context.
func (s *Scraper) Run(ctx context.Context) ([]openrpc.Method, error) {
	methodSrc, err := s.MethodSource(ctx)
	if err != nil {
		return nil, err
	}

	schemaSrc, err := s.SchemaSource(ctx)
	if err != nil {
		return nil, err
	}

	return s.MergeMethodsAndDefinitions(ctx, methodSrc, schemaSrc)
}
