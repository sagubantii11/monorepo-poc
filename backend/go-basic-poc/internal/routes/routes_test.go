package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	SetupRoutes(r)

	// Test ping route
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())

	// Test user route
	req, _ = http.NewRequest("GET", "/users/1", nil)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code) // Because the database is mocked, and the id is valid, but the user is not found.
}
