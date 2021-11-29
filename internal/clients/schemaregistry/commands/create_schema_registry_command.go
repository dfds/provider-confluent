package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

type CreateSchemaRegistryCommand exec.Cmd

func NewCreateSchemaRegistryCommand(subject string, schemaPath string, schemaType string, environment string, apiKey string, apiSecret string) CreateSchemaRegistryCommand {
	var command = CreateSchemaRegistryCommand{
		Path: clients.CLI_NAME,
		Args: []string{"schema-registry", "schema", "create", "--subject", subject, "--schema", schemaPath, "--type", schemaType, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
