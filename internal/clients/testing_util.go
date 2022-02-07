package clients

import (
	"os"
	"testing"
)

// SkipCI Used for skipping tests in an environment where credentials for external providers can't be provided
func SkipCI(t *testing.T) {
	if os.Getenv("TESTING_CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}
