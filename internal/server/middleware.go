package server

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/telemetry"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

type metricsRecorder struct {
	metrics telemetry.Metrics
}

func (mr *metricsRecorder) sendHTTPRequestMetric(r *http.Request) {
	mr.metrics.Incr(r.Context(), telemetry.HTTPRequestsTotal, nil)
}

func MetricsMiddleware(metrics telemetry.Metrics) func(http.Handler) http.Handler {
	recorder := &metricsRecorder{metrics: metrics}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)

			recorder.sendHTTPRequestMetric(r)
		})
	}
}
