// Package api exposes the control plane's HTTP transport.
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Dependencies makes operational dependencies explicit and testable.
type Dependencies struct {
	Logger *slog.Logger
}

// NewHandler builds the HTTP routing boundary.
func NewHandler(deps Dependencies) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", statusHandler("healthy"))
	mux.HandleFunc("GET /readyz", statusHandler("ready"))

	return requestLogger(deps.Logger, mux)
}

func statusHandler(status string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding failure is not actionable after headers are committed.
		_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
	}
}

func requestLogger(logger *slog.Logger, next http.Handler) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
