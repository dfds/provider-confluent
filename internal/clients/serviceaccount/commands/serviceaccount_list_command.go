package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ServiceAccountListCommand is a struct for ServiceAccount list command
type ServiceAccountListCommand exec.Cmd

// NewServiceAccountListCommand is a factory method for ServiceAccount list command
func NewServiceAccountListCommand() ServiceAccountListCommand {
	var command = ServiceAccountListCommand{
		Path: clients.CliName,
		Args: []string{"iam", "service-accont", "list", "-o", "json"},
	}

	return command
}
