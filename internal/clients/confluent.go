package clients

import (
	"fmt"
	"os/exec"
)

// IClient interface for confluent client
type IClient interface {
	Authenticate(email string, password string) error
}

// NewClient is a factory method for confluent client
func NewClient() IClient {
	return &Client{}
}

// Client is a struct for confluent client
type Client struct {
}

// ConflientUsernameEnvKey is the environment key used to assign the username used by the confluent CLI
const ConflientUsernameEnvKey = "CONFLUENT_PLATFORM_USERNAME"

// ConfluentPasswordEnvKey is the environment key used to assign the password used by the confluent CLI
const ConfluentPasswordEnvKey = "CONFLUENT_PLATFORM_PASSWORD"

// CliName is the name of the confluent CLI application
const CliName = "confluent"

// Authenticate a user via the confluent client
func (c *Client) Authenticate(email string, password string) error {
	cmd := exec.Command(CliName, "login")
	cmd.Env = []string{fmt.Sprintf("%v=%v", ConflientUsernameEnvKey, email), fmt.Sprintf("%v=%v", ConfluentPasswordEnvKey, password)}
	cmdOutput, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Not logged in:", err)
		fmt.Println(cmdOutput)
	}

	return err
}
