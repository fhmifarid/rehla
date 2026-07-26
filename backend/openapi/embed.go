// Package openapi owns the canonical HTTP API contract embedded in the server.
package openapi

import _ "embed"

//go:embed openapi.yaml
var document []byte

// Document returns the canonical OpenAPI document.
func Document() []byte {
	return document
}
