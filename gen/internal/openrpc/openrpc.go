// Package openrpc provides the minimal OpenRPC parser needed to process
// the OpenRPC documents for the Ethereum JSON-RPC methods.  The OpenRPC
// team provides a [full-featured parser] if needed in the future but the
// API is overly cumbersome for this use case.
//
// [full-featured parser]: https://pkg.go.dev/github.com/open-rpc/meta-schema
package openrpc

import (
	"encoding/json"
	"fmt"
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

	params := walkJSON("", v, func(k string, v any) bool {
		return k == "/params"
	})

	if len(params) != 1 {
		return nil, nil
	}

	var refs []string

	for _, v := range walkJSON("", params[0], referenceFilter) {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert reference to string: %v", v)
		}

		refs = append(refs, s[strings.LastIndex(s, "/")+1:])
	}

	return refs, nil
}

func referenceFilter(path string, v any) bool {
	return strings.HasSuffix(path, "/$ref")
}

func rewriteReferences(data []byte) []byte {
	return []byte(strings.ReplaceAll(string(data), ethereumReference, localReference))
}

func walkJSON(path string, v any, filter func(path string, v any) bool) []any {
	var out []any

	if v == nil {
		return out
	}

	switch t := v.(type) {
	case map[string]any:
		for k, v := range t {
			p := path + "/" + k

			if filter(p, v) {
				out = append(out, v)
			}

			out = append(out, walkJSON(p, v, filter)...)
		}
	case []any:
		for i, v := range t {
			p := path + "[" + fmt.Sprint(i) + "]"

			out = append(out, walkJSON(p, v, filter)...)
		}
	}

	return out
}
