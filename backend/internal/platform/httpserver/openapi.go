package httpserver

import (
	"net/http"

	"github.com/fhmifarid/rehla/backend/openapi"
)

func serveOpenAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Header().Set("Cache-Control", "public, max-age=300")
		_, _ = w.Write(openapi.Document())
	}
}
