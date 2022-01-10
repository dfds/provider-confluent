package commands

import (
	"errors"
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// ACLDeleteCommand is a struct for ACL delete command
type ACLDeleteCommand exec.Cmd

// NewACLDeleteCommand is a factory method for ACL delete command
func NewACLDeleteCommand(action string, clusterScope string, consumerGroup string, operations []string, prefix string, serviceAccount string, topic string, environment string, cluster string) (ACLDeleteCommand, error) {
	var command = ACLDeleteCommand{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "create", "--action", action, "--cluster-scope", clusterScope, "--service-account", serviceAccount, "--environment", environment, "--prefix", prefix, "-o", "json"},
	}

	for _, v := range operations {
		command.Args = append(command.Args, "--operation", v)
	}

	if topic != "" && consumerGroup == "" {
		command.Args = append(command.Args, "--topic", topic)
	} else if topic == "" && consumerGroup != "" {
		command.Args = append(command.Args, "--consumer-group", consumerGroup)
	} else {
		return ACLDeleteCommand{}, errors.New(ErrTopicOrConsumerGroupAllowed)
	}

	return command, nil
}
