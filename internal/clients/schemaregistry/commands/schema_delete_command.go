package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// SchemaDeleteCommand is a struct for schema delete command
type SchemaDeleteCommand exec.Cmd

// NewSchemaDeleteCommand is a factory method for schema delete command
func NewSchemaDeleteCommand(subject string, version string, permanent bool, environment string, apiKey string, apiSecret string) SchemaDeleteCommand {
	var args = []string{"schema-registry", "schema", "delete", "--subject", subject, "--version", version, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret}

	if permanent {
		args = append(args, "--permanent")
	}

	var command = SchemaDeleteCommand{
		Path: clients.CliName,
		Args: args,
	}

	return command
}
