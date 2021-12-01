package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

type SchemaDeleteCommand exec.Cmd

func NewSchemaDeleteCommand(subject string, version string, permanent bool, environment string, apiKey string, apiSecret string) SchemaDeleteCommand {
	var args = []string{"schema-registry", "schema", "delete", "--subject", subject, "--version", version, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret}

	if permanent {
		args = append(args, "--permanent")
	}

	var command = SchemaDeleteCommand{
		Path: clients.CLI_NAME,
		Args: args,
	}

	return command
}
