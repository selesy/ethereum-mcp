package scraper

import (
	"context"
	"strings"
)

const (
	executionAPIsOwner = "ethereum"
	executionAPIsRepo  = "execution-apis"
	sourcePath         = "src"
	schemaPath         = "src/schemas"
)

// MethodSource returns the OpenRPC source for all methods defined by the
// Ethereum OpenRPC specifications.
func (s *Scraper) MethodSource(ctx context.Context) ([]string, error) {
	methodsSource, err := s.walkDirectories(ctx, sourcePath, func(path string) bool {
		return !strings.HasPrefix(path, "src/schemas") && !strings.HasPrefix(path, "src/engine/openrpc/schemas")
	})
	if err != nil {
		return nil, err
	}

	s.log.InfoContext(ctx, "Method files retrieved", "methods", len(methodsSource))

	return methodsSource, nil
}

// SchemaSource returns the OpenRPC source for all schemas defined by the
// Ethereum OpenRPC specifications.
func (s *Scraper) SchemaSource(ctx context.Context) ([]string, error) {
	schemaSource, err := s.walkDirectories(ctx, sourcePath, func(path string) bool {
		return strings.HasPrefix(path, "src/schemas") || strings.HasPrefix(path, "src/engine/openrpc/schemas")
	})
	if err != nil {
		return nil, err
	}

	s.log.InfoContext(ctx, "Schema files retrieved", "schemas", len(schemaSource))

	return schemaSource, nil
}

func (s *Scraper) walkDirectories(ctx context.Context, path string, filter func(string) bool) ([]string, error) {
	_, dcs, _, err := s.cl.Repositories.GetContents(ctx, executionAPIsOwner, executionAPIsRepo, path, nil)
	if err != nil {
		return nil, err
	}

	cont := []string{}

	for _, dc := range dcs {
		switch dc.GetType() {
		case "dir":
			s.log.DebugContext(ctx, "Reading directory", "path", dc.GetPath())

			sub, err := s.walkDirectories(ctx, dc.GetPath(), filter)
			if err != nil {
				return nil, err
			}

			cont = append(cont, sub...)
		case "file":
			if !filter(dc.GetPath()) {
				continue
			}

			if !strings.HasSuffix(dc.GetPath(), ".yaml") {
				continue
			}

			fc, _, _, err := s.cl.Repositories.GetContents(ctx, executionAPIsOwner, executionAPIsRepo, dc.GetPath(), nil)
			if err != nil {
				return nil, err
			}

			s.log.DebugContext(ctx, "Reading file", "path", dc.GetPath())

			yaml, err := fc.GetContent()
			if err != nil {
				return nil, err
			}

			cont = append(cont, yaml)
		default:
			continue
		}
	}

	return cont, nil
}
