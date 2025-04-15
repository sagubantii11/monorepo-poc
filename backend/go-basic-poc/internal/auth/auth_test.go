package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := "123"
	secretKey := "test_secret_key"

	tokenString, err := GenerateJWT(userID, secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := ValidateJWT(tokenString, secretKey)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.True(t, claims.ExpiresAt.Unix() > time.Now().Unix())

	_, err = ValidateJWT("invalid_token", secretKey)
	assert.Error(t, err)
}
