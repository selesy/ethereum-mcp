package openrpc

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	"github.com/invopop/jsonschema"
)

// Definitions is a managed set of jsonschema.Schema values referenced
// by string names.
type Definitions struct {
	defs jsonschema.Definitions
}

// Merges the key/value pairs of all the Definitions into a new returned
// Definition.
func Merge(defs ...Definitions) Definitions {
	merged := Definitions{defs: make(jsonschema.Definitions)}
	merged.Merge(defs...)

	return merged
}

// Contains returns true if the Definitions contains the given key.
func (d *Definitions) Contains(key string) bool {
	_, ok := d.defs[key]

	return ok
}

// Filter returns a new Definitions containing only the keys provided and
// using the underlying jsonschema.Definitions format so that it can be
// directly added to the raw schema when creating the tool.
func (d *Definitions) Filter(keys ...string) (Definitions, error) {
	// Add the references from the passed in keys (presumably found by
	// searching) for @ref in the parameters.
	filtered := Definitions{defs: make(jsonschema.Definitions)}

	for _, key := range keys {
		if d.Contains(key) {
			filtered.defs[key] = d.defs[key]
		}
	}

	// This is gross and perhaps recursion would be cleaner - finds
	// additional $ref within the defs section and adds them to the
	// filtered list.  Any $ref added might itself contain more $ref
	// elements, so this runs an arbitrary number of times until
	// the before and after counts match.
	var (
		beforeCount = 0
		afterCount  = len(filtered.defs)
	)

	for beforeCount != afterCount {
		beforeCount = afterCount

		data, err := json.Marshal(filtered.defs)
		if err != nil {
			return Definitions{}, err
		}

		var a any
		if err := json.Unmarshal(data, &a); err != nil {
			return Definitions{}, err
		}

		refs := walkJSON("", a, referenceFilter)

		for _, ref := range refs {
			key, ok := ref.(string)
			if !ok {
				return Definitions{}, fmt.Errorf("no references was found for key: %s", key)
			}

			key = key[strings.LastIndex(key, "/")+1:]

			filtered.defs[key] = d.defs[key]
		}

		afterCount = len(filtered.defs)
	}

	return filtered, nil
}

// Get returns the jsonschema.Schema for the given key.
func (d *Definitions) Get(key string) *jsonschema.Schema {
	return d.defs[key]
}

// GetAll returns the jsonschema.Schemas for all definitions.
// TODO: use a pointer receiver?
func (d Definitions) GetAll() jsonschema.Definitions {
	return d.defs
}

// Len returns the number of definitions.
func (d *Definitions) Len() int {
	return len(d.defs)
}

// Merges the key/value pairs into the other Definition(s) into the receiver.
func (d *Definitions) Merge(other ...Definitions) {
	if d.defs == nil {
		d.defs = make(jsonschema.Definitions)
	}

	for _, def := range other {
		maps.Copy(d.defs, def.defs)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (d Definitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.defs)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Definitions) UnmarshalJSON(data []byte) error {
	data = rewriteReferences(data)

	var defs jsonschema.Definitions

	if err := json.Unmarshal(data, &defs); err != nil {
		return err
	}

	d.defs = defs

	return nil
}
