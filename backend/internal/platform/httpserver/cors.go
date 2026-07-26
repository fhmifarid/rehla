package httpserver

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/fhmifarid/rehla/backend/internal/platform/apierror"
)

var corsAllowedMethods = map[string]struct{}{
	http.MethodGet:    {},
	http.MethodPost:   {},
	http.MethodPut:    {},
	http.MethodPatch:  {},
	http.MethodDelete: {},
}

var corsAllowedHeaders = map[string]struct{}{
	"accept":          {},
	"authorization":   {},
	"content-type":    {},
	"idempotency-key": {},
	"if-match":        {},
	"x-csrf-token":    {},
	"x-request-id":    {},
}

func corsPolicy(allowedOrigins []string, logger *slog.Logger) func(http.Handler) http.Handler {
	origins := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		origins[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin == "" {
				next.ServeHTTP(w, r)
				return
			}

			addVary(w.Header(), "Origin")
			if _, allowed := origins[origin]; !allowed {
				writeCORSError(w, r, logger, "cors_origin_denied", "The request origin is not allowed.")
				return
			}

			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "ETag, X-Request-ID")

			if r.Method != http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			addVary(w.Header(), "Access-Control-Request-Method")
			addVary(w.Header(), "Access-Control-Request-Headers")

			requestedMethod := r.Header.Get("Access-Control-Request-Method")
			if _, allowed := corsAllowedMethods[requestedMethod]; !allowed {
				writeCORSError(w, r, logger, "cors_method_denied", "The requested method is not allowed.")
				return
			}
			if !corsHeadersAllowed(r.Header.Get("Access-Control-Request-Headers")) {
				writeCORSError(w, r, logger, "cors_headers_denied", "One or more requested headers are not allowed.")
				return
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			w.Header().Set(
				"Access-Control-Allow-Headers",
				"Accept, Authorization, Content-Type, Idempotency-Key, If-Match, X-CSRF-Token, X-Request-ID",
			)
			w.Header().Set("Access-Control-Max-Age", "600")
			w.WriteHeader(http.StatusNoContent)
		})
	}
}

func corsHeadersAllowed(value string) bool {
	if strings.TrimSpace(value) == "" {
		return true
	}
	for _, header := range strings.Split(value, ",") {
		if _, allowed := corsAllowedHeaders[strings.ToLower(strings.TrimSpace(header))]; !allowed {
			return false
		}
	}
	return true
}

func writeCORSError(
	w http.ResponseWriter,
	r *http.Request,
	logger *slog.Logger,
	code string,
	message string,
) {
	apierror.Write(w, r, logger, &apierror.Error{
		Status:  http.StatusForbidden,
		Code:    code,
		Message: message,
	})
}

func addVary(header http.Header, value string) {
	for _, current := range header.Values("Vary") {
		for _, item := range strings.Split(current, ",") {
			if strings.EqualFold(strings.TrimSpace(item), value) {
				return
			}
		}
	}
	header.Add("Vary", value)
}
