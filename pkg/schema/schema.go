//go:generate go tool schema

// Package schema contains the "raw" JSONSchema source for each OpenRPC
// method defined by Ethereum's execution APIs.
package schema

// Schema returns the JSONSchema source for the Ethereum method with the
// matching name.
func Schema(name string) (string, bool) {
	schema, ok := Schemas()[name]

	return schema, ok
}
