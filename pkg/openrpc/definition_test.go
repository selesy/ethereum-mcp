package openrpc_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/invopop/yaml"
	"github.com/selesy/jsonschema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/selesy/ethereum-mcp/pkg/openrpc"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	const (
		inA = `{
			"A1": {},
			"A2": {},
			"A3": {}
		}`
		inB = `{
			"A3": {"deprecated":true},
			"B1": {},
			"B2": {},
			"B3": {}
		}`
	)

	var (
		a openrpc.Definitions
		b openrpc.Definitions
	)

	require.NoError(t, json.Unmarshal([]byte(inA), &a))
	require.NoError(t, json.Unmarshal([]byte(inB), &b))

	defs := openrpc.Merge(a, b)
	assert.Equal(t, 6, defs.Len())

	assert.True(t, defs.Contains("A1"))
	assert.True(t, defs.Contains("A2"))
	assert.True(t, defs.Contains("A3"))
	assert.True(t, defs.Contains("B1"))
	assert.True(t, defs.Contains("B2"))
	assert.True(t, defs.Contains("B3"))

	assert.False(t, defs.Get("A1").Deprecated)
	assert.False(t, defs.Get("A2").Deprecated)
	assert.True(t, defs.Get("A3").Deprecated)
	assert.False(t, defs.Get("B1").Deprecated)
	assert.False(t, defs.Get("B2").Deprecated)
	assert.False(t, defs.Get("B3").Deprecated)
}

func TestDefinitions_Filter(t *testing.T) {
	t.Parallel()

	const in = `{
		"A1": {},
		"A2": {},
		"A3": {"deprecated":true},
		"B1": {},
		"B2": {},
		"B3": {}
	}`

	var defs openrpc.Definitions

	require.NoError(t, json.Unmarshal([]byte(in), &defs))
	assert.Equal(t, 6, defs.Len())

	assert.True(t, defs.Contains("A1"))
	assert.True(t, defs.Contains("A2"))
	assert.True(t, defs.Contains("A3"))
	assert.True(t, defs.Contains("B1"))
	assert.True(t, defs.Contains("B2"))
	assert.True(t, defs.Contains("B3"))

	defs, err := defs.Filter("A3", "B2")
	require.NoError(t, err)
	assert.Equal(t, 2, defs.Len())

	assert.True(t, defs.Contains("A3"))
	assert.True(t, defs.Contains("B2"))

	assert.True(t, defs.Get("A3").Deprecated)
	assert.False(t, defs.Get("B2").Deprecated)
}

func TestDefinitions_Unmarshal(t *testing.T) {
	t.Parallel()

	t.Run("with execution-api example", func(t *testing.T) {
		t.Parallel()
		data := golden.Get(t, "base-types.yaml")
		assert.NotNil(t, data)

		data, err := yaml.YAMLToJSON(data)
		require.NoError(t, err)

		var defs openrpc.Definitions

		require.NoError(t, json.Unmarshal(data, &defs))
		assert.Equal(t, 17, defs.Len())
		assert.True(t, defs.Contains("address"))

		def := defs.Get("address")
		assert.Equal(t, "", def.Description)
		assert.Equal(t, "hex encoded address", def.Title)
		assert.Equal(t, "^0x[0-9a-fA-F]{40}$", def.Pattern)
		assert.Equal(t, jsonschema.TypeString, def.Type)
		fmt.Println("=====")
	})
}
