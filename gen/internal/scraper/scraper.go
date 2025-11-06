// Package scraper scrapes methods and merges definition schemas to return
// "complete" JSONScehma methods.
package scraper

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/go-github/v74/github"

	"github.com/selesy/ethereum-mcp/gen/internal/openrpc"
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
	tkn, ok := os.LookupEnv("GITHUB_TOKEN")
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

	localMethodSrc, err := s.LocalSource(ctx, "_schemas.yaml")
	if err != nil {
		return nil, err
	}
	s.log.InfoContext(ctx, "Local method files retrieved", "methods", len(localMethodSrc))
	methodSrc = append(methodSrc, localMethodSrc...)

	schemaSrc, err := s.SchemaSource(ctx)
	if err != nil {
		return nil, err
	}

	localSchemaSrc, err := s.LocalSource(ctx, "trace.yaml")
	if err != nil {
		return nil, err
	}
	s.log.InfoContext(ctx, "Local schema files retrieved", "schemas", len(localSchemaSrc))
	schemaSrc = append(schemaSrc, localSchemaSrc...)

	return s.MergeMethodsAndDefinitions(ctx, methodSrc, schemaSrc)
}

// LocalSource returns the OpenRPC source for all methods defined by the
// local OpenRPC specifications.
func (s *Scraper) LocalSource(ctx context.Context, filter string) ([]string, error) {
	root, err := s.findProjectRoot()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(root, "gen", "internal", "openrpc", "data")
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var sources []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}

		if filter == "trace.yaml" && file.Name() == "trace.yaml" {
			continue
		}

		if filter == "_schemas.yaml" && file.Name() != "trace.yaml" {
			continue
		}

		path := filepath.Join(dir, file.Name())
		path = filepath.Clean(path)
		s.log.DebugContext(ctx, "Reading local method file", "path", path)

		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		s.log.DebugContext(ctx, "Read local method file", "path", path)
		sources = append(sources, string(data))
	}

	return sources, nil
}

func (s *Scraper) findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		if dir == filepath.Dir(dir) {
			return "", os.ErrNotExist
		}

		dir = filepath.Dir(dir)
	}
}
