package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// NewTopicDeleteCommand is a factory method for Topic Delete command
func NewTopicDeleteCommand(tp v1alpha1.TopicParameters) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "topic", "delete", tp.Topic.Name, "--cluster", tp.Cluster, "--environment", tp.Environment},
	}

	return command
}
