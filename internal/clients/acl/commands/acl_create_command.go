package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// Errors
const (
	errPatternInvalid      = "pattern type must be either LITERAL or PREFIXED"
	errPermissionInvalid   = "permission type must be either ALLOW or DENY"
	errPrincipalInvalid    = "principal does only allow User:sa-55555 type input"
	errResourceTypeInvalid = "resource type must be either TOPIC or CONSUMER_GROUP"
)

// ACLCreateCommand is a struct for ACL create command
type ACLCreateCommand exec.Cmd

// NewACLCreateCommand is a factory method for ACL create command
// func NewACLCreateCommand(action string, clusterScope string, consumerGroup string, operations []string, prefix string, serviceAccount string, topic string, environment string, cluster string) (ACLCreateCommand, error) {
// 	var command = ACLCreateCommand{
// 		Path: clients.CliName,
// 		Args: []string{"kafka", "acl", "create", "--action", action, "--cluster-scope", clusterScope, "--service-account", serviceAccount, "--environment", environment, "--prefix", prefix, "-o", "json"},
// 	}

// 	for _, v := range operations {
// 		command.Args = append(command.Args, "--operation", v)
// 	}

// 	if topic != "" && consumerGroup == "" {
// 		command.Args = append(command.Args, "--topic", topic)
// 	} else if topic == "" && consumerGroup != "" {
// 		command.Args = append(command.Args, "--consumer-group", consumerGroup)
// 	} else {
// 		return ACLCreateCommand{}, errors.New(ErrTopicOrConsumerGroupAllowed)
// 	}

// 	return command, nil
// }

// NewACLCreateCommand is a factory method for ACL create command
func NewACLCreateCommand(aclP v1alpha1.ACLParameters) (ACLCreateCommand, error) {

	var command = ACLCreateCommand{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "create", "--cluster-scope", aclP.ClusterScope, "--environment", aclP.Environment, "-o", "json"},
	}
	for _, v := range aclP.ACLBlockList {
		err := parsePatternType(&command, v.PatternType)
		if err != nil {
			return command, err
		}

		err = parsePermission(&command, v.Permission)
		if err != nil {
			return command, err
		}

		err = parseServiceAccount(&command, v.Principal)
		if err != nil {
			return command, err
		}

		err = parseResource(&command, v.ResourceName, v.ResourceType)
		if err != nil {
			return command, nil
		}
		command.Args = append(command.Args, "--operation", v.Operation)
	}

	return command, nil
}
