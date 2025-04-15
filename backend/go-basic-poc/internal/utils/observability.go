package utils

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request latency in seconds.",
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(RequestsTotal, RequestDuration)
}

func ObserveHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(ww, r)
		duration := time.Since(start)

		RequestsTotal.WithLabelValues(r.Method, r.URL.Path, http.StatusText(ww.statusCode)).Inc()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration.Seconds())
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ww *responseWriter) WriteHeader(statusCode int) {
	ww.statusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}

func SetupObservability(r *http.ServeMux) {
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/health", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	Logger.Info("Health check passed")
}
