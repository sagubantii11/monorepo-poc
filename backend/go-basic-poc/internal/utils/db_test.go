package utils

import (
	"os"
	"testing"

	"go-basic-poc/internal/config"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInitDB(t *testing.T) {
	// Set up mock config
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "dbname")
	os.Setenv("DB_SSLMODE", "disable")

	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	// Set up mock logger
	Logger, err = zap.NewDevelopment()
	assert.NoError(t, err)

	// Test InitDB
	err = InitDB(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, DB)

	// Test QueryRowDB
	row := QueryRowDB("SELECT 1")
	assert.NotNil(t, row)

	// Test QueryDB
	rows, err := QueryDB("SELECT 1")
	assert.NoError(t, err)
	assert.NotNil(t, rows)

	// Test ExecDB
	result, err := ExecDB("SELECT 1")
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Test closing the DB
	err = DB.Close()
	assert.NoError(t, err)
}
