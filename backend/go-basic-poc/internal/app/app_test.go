package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-basic-poc/internal/config"
	"go-basic-poc/internal/routes"
	"go-basic-poc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRun(t *testing.T) {
	// Set up mock config
	os.Setenv("PORT", "9999")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "dbname")
	os.Setenv("DB_SSLMODE", "disable")

	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	// Set up mock logger
	utils.Logger, err = zap.NewDevelopment()
	assert.NoError(t, err)

	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Set up mock DB connection (InitDB is called in Run)
	err = utils.InitDB(cfg)
	assert.NoError(t, err)

	// Set up mock routes
	routes.SetupRoutes(r)

	// Set up test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make test request
	resp, err := http.Get(ts.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
