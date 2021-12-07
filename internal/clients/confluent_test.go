// nolint
package clients

import (
	"testing"

	"go.dfds.cloud/utils/config"
)

func TestClientAuthenticate(t *testing.T) {
	client := NewClient()
	err := client.Authenticate(config.GetEnvValue(ConflientUsernameEnvKey, ""), config.GetEnvValue(ConfluentPasswordEnvKey, ""))
	if err != nil {
		t.Error(err)
	}
}
