package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// SchemaDescribeCommand is a struct for schema describe command
type SchemaDescribeCommand exec.Cmd

// NewSchemaDescribeCommand is a factory method for schema describe command
func NewSchemaDescribeCommand(subject string, version string, environment string, apiKey string, apiSecret string) SchemaDescribeCommand {
	var command = SchemaDescribeCommand{
		Path: clients.CliName,
		Args: []string{"schema-registry", "schema", "describe", "--subject", subject, "--version", version, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
