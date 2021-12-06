package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// SchemaCreateCommand is a struct for schema create command
type SchemaCreateCommand exec.Cmd

// NewSchemaCreateCommand is a factory method for schema create command
func NewSchemaCreateCommand(subject string, schemaPath string, schemaType string, environment string, apiKey string, apiSecret string) SchemaCreateCommand {
	var command = SchemaCreateCommand{
		Path: clients.CliName,
		Args: []string{"schema-registry", "schema", "create", "--subject", subject, "--schema", schemaPath, "--type", schemaType, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
