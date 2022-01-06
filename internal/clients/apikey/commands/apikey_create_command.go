package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ApiKeyCreateCommand is a struct for ApiKey create command
type ApiKeyCreateCommand exec.Cmd

// NewApiKeyCreateCommand is a factory method for ApiKey create command
func NewApiKeyCreateCommand(resource string, description string, serviceAccount string, environment string) ApiKeyCreateCommand {
	var command = ApiKeyCreateCommand{
		Path: clients.CliName,
		Args: []string{"api-key", "create", "--resource", resource, "--description", description, "--service-account", serviceAccount, "--environment", environment, "-o", "json"},
	}

	return command
}
