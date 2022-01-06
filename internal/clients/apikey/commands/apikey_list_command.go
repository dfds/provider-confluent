package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ApiKeyListCommand is a struct for ApiKey list command
type ApiKeyListCommand exec.Cmd

// NewApiKeyListCommand is a factory method for ApiKey list command
func NewApiKeyListCommand() ApiKeyListCommand {
	var command = ApiKeyListCommand{
		Path: clients.CliName,
		Args: []string{"api-key", "list", "-o", "json"},
	}

	return command
}
