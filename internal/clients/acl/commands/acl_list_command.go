package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewACLListCommand is a factory method for ACL list command
func NewACLListCommand(environment string, cluster string, serviceAccount string) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "list", "--environment", environment, "--cluster", cluster, "--service-account", serviceAccount, "-o", "json"},
	}

	return command
}
