package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// NewTopicDescribeCommand is a factory method for Topic describe command
func NewTopicDescribeCommand(to v1alpha1.TopicObservation) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "topic", "describe", to.Name, "--cluster", to.Cluster, "--environment", to.Environment, "-o", "json"},
	}

	return command
}
