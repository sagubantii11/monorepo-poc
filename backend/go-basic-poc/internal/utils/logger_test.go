package utils

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	if Logger == nil {
		t.Error("Logger not initialized")
	}
}
