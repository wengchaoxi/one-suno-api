package logger

import (
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	Init()
	if Logger == nil {
		t.Fatal("Logger was not initialized")
	}
}
