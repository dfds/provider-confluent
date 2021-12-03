package clients

import (
	"fmt"
	"os/exec"
)

type Client interface {
	Authenticate(email string, password string) error
}

func NewClient() Client {
	return &ConfluentClient{}
}

type ConfluentClient struct {
}

const CONFLUENT_EMAIL, CONFLUENT_PASSWORD = "CONFLUENT_PLATFORM_USERNAME", "CONFLUENT_PLATFORM_PASSWORD"
const CLI_NAME = "confluent"

func (c *ConfluentClient) Authenticate(email string, password string) error {
	cmd := exec.Command(CLI_NAME, "login")
	cmd.Env = []string{fmt.Sprintf("%v=%v", CONFLUENT_EMAIL, email), fmt.Sprintf("%v=%v", CONFLUENT_PASSWORD, password)}
	cmdOutput, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Not logged in:", err)
		fmt.Println(cmdOutput)
	}

	return err
}
