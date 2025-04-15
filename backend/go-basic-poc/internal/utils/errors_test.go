package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppError(t *testing.T) {
	err := errors.New("test error")
	appErr := NewAppError(500, "Internal Server Error", err)

	assert.Equal(t, 500, appErr.Code)
	assert.Equal(t, "Internal Server Error", appErr.Message)
	assert.Equal(t, err, appErr.Err)
	assert.Contains(t, appErr.Error(), "test error")

	appErrWithoutErr := NewAppError(400, "Bad Request", nil)
	assert.NotContains(t, appErrWithoutErr.Error(), "test error")
}
