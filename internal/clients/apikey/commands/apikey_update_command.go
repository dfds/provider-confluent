package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewAPIKeyUpdateCommand is a factory method for ApiKey update command
func NewAPIKeyUpdateCommand(key string, description string) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"api-key", "update", key, "--description", description},
	}

	return command
}
