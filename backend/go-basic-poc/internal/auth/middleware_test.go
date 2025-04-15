package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	secretKey := "test_secret_key"
	os.Setenv("CLIENT_SECRET", secretKey)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middlewareHandler := AuthMiddleware(handler)

	// Test with valid token
	tokenString, _ := GenerateJWT("123", secretKey)
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", tokenString)
	recorder := httptest.NewRecorder()
	middlewareHandler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Test with missing token
	req = httptest.NewRequest("GET", "/test", nil)
	recorder = httptest.NewRecorder()
	middlewareHandler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)

	// Test with invalid token
	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "invalid_token")
	recorder = httptest.NewRecorder()
	middlewareHandler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}
