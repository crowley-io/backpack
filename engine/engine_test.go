package engine

import (
	"os"
	"testing"
)

func setEnv(t *testing.T, key, value string) {
	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func unsetEnv(t *testing.T, key string) {
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
