package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ServiceAccountDeleteCommand is a struct for ServiceAccount delete command
type ServiceAccountDeleteCommand exec.Cmd

// NewServiceAccountDeleteCommand is a factory method for ServiceAccount delete command
func NewServiceAccountDeleteCommand(id string) ServiceAccountDeleteCommand {
	var command = ServiceAccountDeleteCommand{
		Path: clients.CliName,
		Args: []string{"iam", "service-account", "delete", id},
	}

	return command
}
