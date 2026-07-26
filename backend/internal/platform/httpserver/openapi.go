package httpserver

import (
	_ "embed"
	"net/http"
)

//go:embed openapi.yaml
var openAPIDocument []byte

func serveOpenAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Header().Set("Cache-Control", "public, max-age=300")
		_, _ = w.Write(openAPIDocument)
	}
}
