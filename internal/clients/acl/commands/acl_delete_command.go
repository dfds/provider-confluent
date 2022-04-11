package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewACLDeleteCommand is a factory method for ACL delete command
func NewACLDeleteCommand(aclP v1alpha1.ACLParameters) (exec.Cmd, error) {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "delete", "--environment", aclP.Environment, "--cluster", aclP.Cluster},
	}
	err := parsePatternType(&command, aclP.ACLRule.PatternType)
	if err != nil {
		return command, err
	}

	err = parsePermission(&command, aclP.ACLRule.Permission)
	if err != nil {
		return command, err
	}

	err = parseServiceAccount(&command, aclP.ACLRule.Principal)
	if err != nil {
		return command, err
	}

	err = parseResource(&command, aclP.ACLRule.ResourceName, aclP.ACLRule.ResourceType)
	if err != nil {
		return command, err
	}
	command.Args = append(command.Args, "--operation", aclP.ACLRule.Operation)

	return command, nil
}
