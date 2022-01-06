package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ApiKeyUpdateCommand is a struct for ApiKey update command
type ApiKeyUpdateCommand exec.Cmd

// NewApiKeyUpdateCommand is a factory method for ApiKey update command
func NewApiKeyUpdateCommand(key string, description string) ApiKeyUpdateCommand {
	var command = ApiKeyUpdateCommand{
		Path: clients.CliName,
		Args: []string{"api-key", "update", key, "--description", description},
	}

	return command
}
