package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ApiKeyDeleteCommand is a struct for ApiKey delete command
type ApiKeyDeleteCommand exec.Cmd

// NewApiKeyDeleteCommand is a factory method for ApiKey delete command
func NewApiKeyDeleteCommand(key string) ApiKeyDeleteCommand {
	var command = ApiKeyDeleteCommand{
		Path: clients.CliName,
		Args: []string{"api-key", "delete", key},
	}

	return command
}
