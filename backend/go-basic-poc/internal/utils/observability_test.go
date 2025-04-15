package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestObserveHTTP(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	observedHandler := ObserveHTTP(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()

	observedHandler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "OK", recorder.Body.String())

	// Basic check for prometheus metrics
	expected := `# HELP http_requests_total Total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="GET",path="/test",status="OK"} 1
`
	err := testutil.CollectAndCompare(RequestsTotal, strings.NewReader(expected)) // Corrected line
	assert.NoError(t, err)
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	recorder := httptest.NewRecorder()

	healthCheck(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "OK", recorder.Body.String())
}
