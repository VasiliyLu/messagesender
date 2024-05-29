package testingex

import (
	"os"
	"strings"
	"testing"
)

func CheckIntegrationEnv(t *testing.T) (skip bool) {
	if intg := os.Getenv("INTEGRATION"); intg == "" || strings.ToLower(intg) == "false" {
		t.Skip("skipping integration tests: set INTEGRATION environment variable")
		return true
	}
	return false
}
