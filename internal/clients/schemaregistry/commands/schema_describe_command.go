package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

type SchemaDescribeCommand exec.Cmd

func NewSchemaDescribeCommand(subject string, version string, environment string, apiKey string, apiSecret string) SchemaDescribeCommand {
	var command = SchemaDescribeCommand{
		Path: clients.CLI_NAME,
		Args: []string{"schema-registry", "schema", "describe", "--subject", subject, "--version", version, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
