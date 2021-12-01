package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

type SchemaCreateCommand exec.Cmd

func NewSchemaCreateCommand(subject string, schemaPath string, schemaType string, environment string, apiKey string, apiSecret string) SchemaCreateCommand {
	var command = SchemaCreateCommand{
		Path: clients.CLI_NAME,
		Args: []string{"schema-registry", "schema", "create", "--subject", subject, "--schema", schemaPath, "--type", schemaType, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
