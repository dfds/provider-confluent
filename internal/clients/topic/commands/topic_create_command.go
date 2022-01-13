package commands

import (
	"fmt"
	"os/exec"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// NewTopicCreateCommand is a factory method for Topic create command
func NewTopicCreateCommand(tp v1alpha1.TopicParameters) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "topic", "create", tp.Topic.Name, "--cluster", tp.Cluster, "--environment", tp.Environment, "--partitions", fmt.Sprintf("%d", tp.Topic.Partitions), "--config", fmt.Sprintf("'retention.ms=%d'", tp.Topic.Config.Retention)},
	}

	return command
}
