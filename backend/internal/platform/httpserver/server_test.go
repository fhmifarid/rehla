package httpserver

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fhmifarid/rehla/backend/internal/config"
	"github.com/fhmifarid/rehla/backend/internal/platform/apierror"
	"github.com/fhmifarid/rehla/backend/openapi"
	"github.com/go-chi/chi/v5"
)

func TestHealthEndpoint(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	router := chi.NewRouter()
	router.Use(requestContext(logger))
	router.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if requestID := response.Header().Get("X-Request-ID"); requestID == "" {
		t.Fatal("expected X-Request-ID response header")
	}
	if !strings.Contains(response.Body.String(), `"status":"ok"`) {
		t.Fatalf("body = %q", response.Body.String())
	}
}

func TestAPIErrorDoesNotLeakCause(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request = request.WithContext(apierror.ContextWithRequestID(request.Context(), "request-123"))
	response := httptest.NewRecorder()

	apierror.Write(response, request, logger, &apierror.Error{
		Status:     http.StatusInternalServerError,
		Code:       "internal_error",
		Message:    "An unexpected error occurred.",
		Underlying: io.ErrUnexpectedEOF,
	})

	if strings.Contains(response.Body.String(), "unexpected EOF") {
		t.Fatal("response leaked internal error")
	}
	if !strings.Contains(response.Body.String(), `"request_id":"request-123"`) {
		t.Fatal("response did not include request ID")
	}
}

func TestOpenAPIEndpointServesCanonicalDocument(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/openapi.yaml", nil)
	response := httptest.NewRecorder()

	serveOpenAPI().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if !bytes.Equal(response.Body.Bytes(), openapi.Document()) {
		t.Fatal("served OpenAPI document differs from the canonical document")
	}
}

func TestCORSAllowsConfiguredPreflight(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := New(Dependencies{
		Config: config.Config{AllowedOrigins: []string{"https://admin.example"}},
		Logger: logger,
	})
	request := httptest.NewRequest(http.MethodOptions, "/v1/system/info", nil)
	request.Header.Set("Origin", "https://admin.example")
	request.Header.Set("Access-Control-Request-Method", http.MethodGet)
	request.Header.Set("Access-Control-Request-Headers", "Authorization, X-Request-ID")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusNoContent)
	}
	if origin := response.Header().Get("Access-Control-Allow-Origin"); origin != "https://admin.example" {
		t.Fatalf("Access-Control-Allow-Origin = %q", origin)
	}
	if response.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Fatal("expected credentialed CORS response")
	}
}

func TestCORSRejectsUnknownOrigin(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := New(Dependencies{
		Config: config.Config{AllowedOrigins: []string{"https://admin.example"}},
		Logger: logger,
	})
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	request.Header.Set("Origin", "https://attacker.example")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusForbidden)
	}
	if response.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Fatal("disallowed origin must not receive an allow-origin header")
	}
	if !strings.Contains(response.Body.String(), `"code":"cors_origin_denied"`) {
		t.Fatalf("body = %q", response.Body.String())
	}
}

func TestCORSRejectsUnknownHeaders(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := New(Dependencies{
		Config: config.Config{AllowedOrigins: []string{"https://admin.example"}},
		Logger: logger,
	})
	request := httptest.NewRequest(http.MethodOptions, "/v1/system/info", nil)
	request.Header.Set("Origin", "https://admin.example")
	request.Header.Set("Access-Control-Request-Method", http.MethodGet)
	request.Header.Set("Access-Control-Request-Headers", "X-Unapproved")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusForbidden)
	}
	if !strings.Contains(response.Body.String(), `"code":"cors_headers_denied"`) {
		t.Fatalf("body = %q", response.Body.String())
	}
}
