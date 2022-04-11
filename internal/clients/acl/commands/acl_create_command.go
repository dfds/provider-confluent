package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// Errors
const (
	errPatternInvalid                               = "pattern type must be either LITERAL or PREFIXED"
	errPermissionInvalid                            = "permission type must be either ALLOW or DENY"
	errPrincipalInvalid                             = "principal does only allow User:sa-55555 type input"
	errResourceTypeInvalid                          = "resource type must be either TOPIC, CONSUMER_GROUP or CLUSTER"
	errResourceNameSpecifiedWithResourceTypeCluster = "resource name can't be specified when resource type is CLUSTER"
)

// NewACLCreateCommand is a factory method for ACL create command
func NewACLCreateCommand(aclP v1alpha1.ACLParameters) (exec.Cmd, error) {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "create", "--environment", aclP.Environment, "--cluster", aclP.Cluster, "-o", "json"},
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
