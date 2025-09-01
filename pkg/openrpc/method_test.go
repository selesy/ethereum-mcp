package openrpc_test

import (
	"encoding/json"
	"testing"

	"github.com/invopop/yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/selesy/ethereum-mcp/pkg/openrpc"
)

func TestMethod_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("with execution-api example", func(t *testing.T) {
		t.Parallel()
		data := golden.Get(t, "transaction.yaml")
		assert.NotNil(t, data)

		data, err := yaml.YAMLToJSON(data)
		require.NoError(t, err)

		var methods []openrpc.Method

		require.NoError(t, json.Unmarshal(data, &methods))
		assert.Len(t, methods, 4)

		method := methods[0]
		assert.Equal(t, "eth_getTransactionByHash", method.Name())
		assert.Equal(t, "Returns the information about a transaction requested by transaction hash.", method.Description())
		assert.Equal(t, []string{"hash32"}, method.Refs())

		param := method.Params()[0]
		assert.Equal(t, "transaction_hash", param.Name())
		assert.Equal(t, "Transaction hash", param.Description())
		assert.True(t, param.Required())
		assert.False(t, param.Deprecated())

		schema := param.Schema()
		assert.Equal(t, "#/$defs/hash32", schema.Ref)
	})

	t.Run("with only summary", func(t *testing.T) {
		t.Parallel()

		const in = `{
			"name": "method_name",
			"summary": "summary text"
		}`

		var method openrpc.Method

		require.NoError(t, json.Unmarshal([]byte(in), &method))
		assert.Equal(t, "method_name", method.Name())
		assert.Equal(t, "summary text", method.Description())
	})

	t.Run("with only description", func(t *testing.T) {
		t.Parallel()

		const in = `{
			"name": "method_name",
			"description": "description text\nwhich can span multiple lines"
		}`

		var method openrpc.Method

		require.NoError(t, json.Unmarshal([]byte(in), &method))
		assert.Equal(t, "method_name", method.Name())
		assert.Equal(t, "description text\nwhich can span multiple lines", method.Description())
	})
}
