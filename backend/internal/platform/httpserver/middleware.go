package httpserver

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/rehla-platform/rehla/backend/internal/platform/apierror"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *statusRecorder) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusRecorder) Write(body []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	n, err := w.ResponseWriter.Write(body)
	w.bytes += n
	return n, err
}

func requestContext(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := sanitizedRequestID(r.Header.Get("X-Request-ID"))
			if requestID == "" {
				requestID = newRequestID()
			}
			w.Header().Set("X-Request-ID", requestID)

			started := time.Now()
			recorder := &statusRecorder{ResponseWriter: w}
			ctx := apierror.ContextWithRequestID(r.Context(), requestID)
			next.ServeHTTP(recorder, r.WithContext(ctx))

			logger.Info("request complete",
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
				"status", recorder.status,
				"bytes", recorder.bytes,
				"duration_ms", time.Since(started).Milliseconds(),
			)
		})
	}
}

func recoverer(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					logger.Error("panic recovered",
						"request_id", apierror.RequestID(r.Context()),
						"panic", recovered,
						"stack", string(debug.Stack()),
					)
					apierror.Write(w, r, logger, &apierror.Error{
						Status:  http.StatusInternalServerError,
						Code:    "internal_error",
						Message: "An unexpected error occurred.",
					})
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
		next.ServeHTTP(w, r)
	})
}

func newRequestID() string {
	var data [16]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "req-" + time.Now().UTC().Format("20060102150405.000000000")
	}
	return hex.EncodeToString(data[:])
}

func sanitizedRequestID(value string) string {
	if len(value) < 8 || len(value) > 128 {
		return ""
	}
	for _, char := range value {
		if !(char == '-' || char == '_' || char == '.' ||
			char >= 'a' && char <= 'z' ||
			char >= 'A' && char <= 'Z' ||
			char >= '0' && char <= '9') {
			return ""
		}
	}
	return strings.TrimSpace(value)
}
