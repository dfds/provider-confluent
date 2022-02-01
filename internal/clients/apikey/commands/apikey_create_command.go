package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewAPIKeyCreateCommand is a factory method for ApiKey create command
func NewAPIKeyCreateCommand(resource string, description string, serviceAccount string, environment string) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"api-key", "create", "--resource", resource, "--description", description, "--service-account", serviceAccount, "--environment", environment, "-o", "json"},
	}

	return command
}
