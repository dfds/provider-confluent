package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewAPIKeyListCommand is a factory method for ApiKey list command
func NewAPIKeyListCommand() exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"api-key", "list", "-o", "json"},
	}

	return command
}
