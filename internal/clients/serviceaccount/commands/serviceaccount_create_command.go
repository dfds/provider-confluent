package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ServiceAccountCreateCommand is a struct for ServiceAccount create command
type ServiceAccountCreateCommand exec.Cmd

// NewServiceAccountCreateCommand is a factory method for ServiceAccount create command
func NewServiceAccountCreateCommand(name string, description string) ServiceAccountCreateCommand {
	var command = ServiceAccountCreateCommand{
		Path: clients.CliName,
		Args: []string{"iam", "service-account", "create", name, "--description", description, "-o", "json"},
	}

	return command
}
