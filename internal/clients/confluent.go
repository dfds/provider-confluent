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

func (c *ConfluentClient) Login(email string, password string) error {
	cmd := exec.Command("confluent", "login")
	cmd.Env = []string{"CONFLUENT_EMAIL=" + email, "CONFLUENT_PASSWORD=" + password}
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Not logged in:", err)
	}
	return err
}
