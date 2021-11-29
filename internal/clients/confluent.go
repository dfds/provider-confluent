package clients

import (
	"fmt"
	"os/exec"
)

type Client interface {
	Login(email string, password string) error
}

func NewClient() Client {
	return &ConfluentClient{}
}

type ConfluentClient struct {
}

const CONFLUENT_EMAIL, CONFLUENT_PASSWORD = "CONFLUENT_EMAIL", "CONFLUENT_PASSWORD"
const CLI_NAME = "confluent"


func (c *ConfluentClient) Login(email string, password string) error {
	cmd := exec.Command(CLI_NAME, "login")
	cmd.Env = []string{fmt.Sprintf("%v=%v", CONFLUENT_EMAIL, email), fmt.Sprintf("%v=%v", CONFLUENT_PASSWORD, password)}
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Not logged in:", err)
	}

	return err
}
