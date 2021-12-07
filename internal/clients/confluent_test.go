package clients

import (
	"go.dfds.cloud/utils/config"
	"testing"
)


func TestClientAuthenticate(t *testing.T) {
	client := NewClient()
	err := client.Authenticate(config.GetEnvValue(ConflientUsernameEnvKey, ""), config.GetEnvValue(ConfluentPasswordEnvKey, ""))
	if err != nil {
		t.Error(err)
	}
}
