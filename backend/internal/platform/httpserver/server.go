package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/fhmifarid/rehla/backend/internal/config"
	"github.com/fhmifarid/rehla/backend/internal/platform/apierror"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	Config config.Config
	Logger *slog.Logger
	DB     *pgxpool.Pool
}

func New(deps Dependencies) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(requestContext(deps.Logger))
	router.Use(recoverer(deps.Logger))
	router.Use(securityHeaders)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		apierror.Write(w, r, deps.Logger, &apierror.Error{
			Status:  http.StatusNotFound,
			Code:    "route_not_found",
			Message: "The requested route does not exist.",
		})
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		apierror.Write(w, r, deps.Logger, &apierror.Error{
			Status:  http.StatusMethodNotAllowed,
			Code:    "method_not_allowed",
			Message: "The HTTP method is not allowed for this route.",
		})
	})

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	router.Get("/readyz", readinessHandler(deps))
	router.Get("/openapi.yaml", serveOpenAPI())
	router.Route("/v1", func(v1 chi.Router) {
		v1.Get("/system/info", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{
				"name":        "Rehla API",
				"environment": deps.Config.Environment,
				"version":     "0.1.0",
			})
		})
	})
	return router
}

func readinessHandler(deps Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := contextWithTimeout(r, 2*time.Second)
		defer cancel()
		if err := deps.DB.Ping(ctx); err != nil {
			apierror.Write(w, r, deps.Logger, &apierror.Error{
				Status:     http.StatusServiceUnavailable,
				Code:       "dependency_unavailable",
				Message:    "A required dependency is unavailable.",
				Retryable:  true,
				Underlying: err,
			})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
