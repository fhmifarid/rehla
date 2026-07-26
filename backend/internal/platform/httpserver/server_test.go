package httpserver

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rehla-platform/rehla/backend/internal/platform/apierror"
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
