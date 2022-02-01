package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewAPIKeyDeleteCommand is a factory method for ApiKey delete command
func NewAPIKeyDeleteCommand(key string) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"api-key", "delete", key},
	}

	return command
}
