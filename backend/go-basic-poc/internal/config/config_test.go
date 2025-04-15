package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("PORT", "9000")
	os.Setenv("CLIENT_ID", "test_client_id")
	os.Setenv("CLIENT_SECRET", "test_client_secret")
	os.Setenv("DB_HOST", "test_db_host")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "test_db_user")
	os.Setenv("DB_PASSWORD", "test_db_password")
	os.Setenv("DB_NAME", "test_db_name")
	os.Setenv("DB_SSLMODE", "require")

	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "9000", cfg.Port)
	assert.Equal(t, "test_client_id", cfg.ClientID)
	assert.Equal(t, "test_client_secret", cfg.ClientSecret)
	assert.Equal(t, "test_db_host", cfg.DBHost)
	assert.Equal(t, "5433", cfg.DBPort)
	assert.Equal(t, "test_db_user", cfg.DBUser)
	assert.Equal(t, "test_db_password", cfg.DBPassword)
	assert.Equal(t, "test_db_name", cfg.DBName)
	assert.Equal(t, "require", cfg.DBSSLMode)

	// Test with default port
	os.Unsetenv("PORT")
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "8080", cfg.Port)

	// Test with default DB port
	os.Unsetenv("DB_PORT")
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "5432", cfg.DBPort)

	// Test with default ssl mode
	os.Unsetenv("DB_SSLMODE")
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "disable", cfg.DBSSLMode)
}
