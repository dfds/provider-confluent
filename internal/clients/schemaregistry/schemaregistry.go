package schemaregistry

import (
	"encoding/json"
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
)

type Client interface {
	Create(subject string, schema string, schemaType string, environment string) (string, error)
	Delete(subject string, version string, permanent bool, environment string) (string, error)
	Describe(subject string, version string, environment string) (SchemaDescribeResponse, error)
}

func NewClient(c clients.Config) Client {
	return &ConfluentCLIClient{Config: c}
}

type ConfluentCLIClient struct {
	Config clients.Config
}

func (c *ConfluentCLIClient) Create(subject string, schema string, schemaType string, environment string) (string, error) {
	var cmd = commands.NewCreateSchemaRegistryCommand(subject, schema, schemaType, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	return string(cmdOutput), cmdErr
}

func (c *ConfluentCLIClient) Delete(subject string, version string, permanent bool, environment string) (string, error) {
	var cmd = commands.NewDeleteSchemaRegistryCommand(subject, version, permanent, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	return string(cmdOutput), cmdErr
}

func (c *ConfluentCLIClient) Describe(subject string, version string, environment string) (SchemaDescribeResponse, error) {
	var cmd = commands.NewDescribeSchemaRegistryCommand(subject, version, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))
	var schema SchemaDescribeResponse

	if cmdErr != nil {
		return schema, cmdErr
	}

	err := json.Unmarshal([]byte(cmdOutput), &schema)

	return schema, err
}

func executeCommand(cmd exec.Cmd) ([]byte, error) {
	out, err := exec.Command(cmd.Path, cmd.Args...).CombinedOutput()

	if err != nil {
		return nil, err
	}

	return out, err
}
