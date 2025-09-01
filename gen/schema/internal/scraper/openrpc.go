package scraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/invopop/yaml"

	"github.com/selesy/ethereum-mcp/pkg/openrpc"
)

// MergeMethodsAndDefinitions adds a "def" key containing each definition
// that's referenced by schema "params" to the schema from the pool of
// definitions that are provided by Ethereum's execution APIs.
func (s *Scraper) MergeMethodsAndDefinitions(ctx context.Context, methodSource []string, schemaSource []string) ([]openrpc.Method, error) {
	var methods []openrpc.Method

	methods, err := s.decodeMethods(methodSource)
	if err != nil {
		return nil, err
	}

	for _, m := range methods {
		s.log.DebugContext(ctx, "Method processed", slog.String("name", m.Name()))
	}

	s.log.Info("Methods processed", slog.Int("count", len(methods)))

	schemas, err := s.decodeSchemas(schemaSource)
	if err != nil {
		return nil, err
	}

	for k := range schemas.GetAll() {
		s.log.DebugContext(ctx, "Definition (schema) processed", slog.String("name", k))
	}

	s.log.InfoContext(ctx, "Definitions (schemas) processed", slog.Int("count", schemas.Len()))

	for i, m := range methods {

		name := m.Name()
		if m.Name() == "" {
			return nil, errors.New("the MethodObject.Name in OpenRPC file is nil")
		}

		desc := m.Description()
		if desc == "" {
			return nil, fmt.Errorf("both the description and summary are missing for method %s", name)
		}

		if m.Params() == nil || len(m.Params()) == 0 {
			continue
		}

		for _, p := range m.Params() {
			name := p.Name()
			if name == "" {
				return nil, errors.New("the parameter name in OpenRPC file is missing")
			}
		}

		defs, err := schemas.Filter(m.Refs()...)
		if err != nil {
			return nil, err
		}

		methods[i] = *m.WithDefs(defs)
	}

	return methods, nil
}

func (s *Scraper) decodeMethods(methodSource []string) ([]openrpc.Method, error) {
	var methods []openrpc.Method

	for _, src := range methodSource {
		ms, err := s.decodeMethod(src)
		if err != nil {
			return methods, fmt.Errorf("%w: %s", err, src)
		}

		methods = append(methods, ms...)
	}

	return methods, nil
}

func (s *Scraper) decodeMethod(methodSource string) ([]openrpc.Method, error) {
	var methods []openrpc.Method

	data, err := yaml.YAMLToJSON([]byte(methodSource))
	if err != nil {
		return methods, err
	}

	return methods, json.Unmarshal(data, &methods)
}

func (s *Scraper) decodeSchemas(schemaSource []string) (openrpc.Definitions, error) {
	var schemas openrpc.Definitions

	for _, src := range schemaSource {
		ss, err := s.decodeSchema(src)
		if err != nil {
			return schemas, fmt.Errorf("%w: %s", err, src)
		}

		schemas.Merge(ss)
	}

	return schemas, nil
}

func (s *Scraper) decodeSchema(schemaSource string) (openrpc.Definitions, error) {
	var schemas openrpc.Definitions

	data, err := yaml.YAMLToJSON([]byte(schemaSource))
	if err != nil {
		return schemas, err
	}

	return schemas, json.Unmarshal(data, &schemas)
}
