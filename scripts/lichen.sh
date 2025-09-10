#!/usr/bin/env bash

go build -o ethereum-mcp.so ./main.go
go tool lichen --json sbom.tmp ethereum-mcp.so
cat sbom.tmp | jq 'del(.Modules[].Dir, .Modules[].Licenses[].Path, .Modules[].UsedBy)' > sbom.json

rm sbom.tmp
rm ethereum-mcp.so
