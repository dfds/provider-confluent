package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// SchemaSubjectUpdateCommand is a struct for schema create command
type SchemaSubjectUpdateCommand exec.Cmd

// NewSchemaSubjectUpdateCommand is a factory method for schema create command
func NewSchemaSubjectUpdateCommand(subject string, compatibility string, environment string, apiKey string, apiSecret string) SchemaSubjectUpdateCommand {
	var command = SchemaSubjectUpdateCommand{
		Path: clients.CliName,
		Args: []string{"schema-registry", "subject", "update", subject, "--compatibility", compatibility, "--environment", environment, "--api-key", apiKey, "--api-secret", apiSecret},
	}

	return command
}
