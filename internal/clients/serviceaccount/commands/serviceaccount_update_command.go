package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ServiceAccountUpdateCommand is a struct for ServiceAccount update command
type ServiceAccountUpdateCommand exec.Cmd

// NewServiceAccountUpdateCommand is a factory method for ServiceAccount update command
func NewServiceAccountUpdateCommand(id string, description string) ServiceAccountUpdateCommand {
	var command = ServiceAccountUpdateCommand{
		Path: clients.CliName,
		Args: []string{"iam", "service-account", "update", id, "--description", description},
	}

	return command
}
