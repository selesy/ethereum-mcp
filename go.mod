module github.com/selesy/ethereum-mcp

go 1.24.4

tool (
	github.com/selesy/ethereum-mcp/gen
	golang.org/x/vuln/cmd/govulncheck
)

require github.com/mark3labs/mcp-go v0.39.0

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/dave/jennifer v1.7.1 // indirect
	github.com/google/go-github/v74 v74.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/invopop/jsonschema v0.13.0 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/lmittmann/tint v1.1.2 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/selesy/ethereum-mcp/gen v0.4.0 // indirect
	github.com/selesy/jsonschema v0.14.0-rc1 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7 // indirect
	golang.org/x/tools v0.29.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/selesy/ethereum-mcp/gen => ./gen
