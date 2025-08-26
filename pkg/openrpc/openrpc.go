// Package openrpc provides the minimal OpenRPC parser needed to process
// the OpenRPC documents for the Ethereum JSON-RPC methods.  The OpenRPC
// team provides a [full-featured parser] if needed in the future but the
// API is overly cumbersome for this use case.
//
// [full-featured parser]: https://pkg.go.dev/github.com/open-rpc/meta-schema
package openrpc

import (
	"encoding/json"
	"strings"
)

const (
	ethereumReference = "#/components/schemas/"
	localReference    = "#/$defs/"
)

func findParamReferences(data []byte) ([]string, error) {
	var v any

	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	return walkReferences(v, false), nil
}

func rewriteReferences(data []byte) []byte {
	return []byte(strings.ReplaceAll(string(data), ethereumReference, localReference))
}

func walkReferences(v any, inParams bool) []string {
	if v == nil {
		return nil
	}

	var refs []string

	switch t := v.(type) {
	case map[string]any:
		if _, ok := t["params"]; ok {
			return append(refs, walkReferences(t["params"], true)...)
		}

		ref, ok := t["$ref"]
		if ok && inParams {
			refs = append(refs, strings.TrimPrefix(ref.(string), localReference))
		}

		for _, v := range t {
			refs = append(refs, walkReferences(v, inParams)...)
		}
	case []any:
		for _, v := range t {
			refs = append(refs, walkReferences(v, inParams)...)
		}
	}

	return refs
}
