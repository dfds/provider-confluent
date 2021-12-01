package schemaregistry

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.dfds.cloud/utils/config"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
)

const (
	SCHEMA_REGISTRY_API_KEY    = "PROVIDER_CONFLUENT_SCHEMA_REGISTRY_API_KEY"
	SCHEMA_REGISTRY_API_SECRET = "PROVIDER_CONFLUENT_SCHEMA_REGISTRY_API_SECRET"
)

func NewClient(c clients.Config) Client {
	if len(c.ApiKey) == 0 {
		c.ApiKey = config.GetEnvValue(SCHEMA_REGISTRY_API_KEY, "")
	}
	if len(c.ApiSecret) == 0 {
		c.ApiSecret = config.GetEnvValue(SCHEMA_REGISTRY_API_SECRET, "")
	}
	return &SchemaRegistryClient{Config: c}
}

func (c *SchemaRegistryClient) SchemaCreate(subject string, schema string, schemaType string, environment string) (string, error) {
	schemaGuid := uuid.New().String()
	path, err := CreateFile([]byte(schema), schemaGuid, "")
	if err != nil {
		return "", err
	}

	var cmd = commands.NewSchemaCreateCommand(subject, path, schemaType, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	err = RemoveFile(path)
	if err != nil {
		return string(cmdOutput), err
	}

	return string(cmdOutput), cmdErr
}

func (c *SchemaRegistryClient) SchemaDelete(subject string, version string, permanent bool, environment string) (string, error) {
	var cmd = commands.NewSchemaDeleteCommand(subject, version, permanent, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	return string(cmdOutput), cmdErr
}

func (c *SchemaRegistryClient) SchemaDescribe(subject string, version string, environment string) (SchemaDescribeResponse, error) {
	var cmd = commands.NewSchemaDescribeCommand(subject, version, environment, c.Config.ApiKey, c.Config.ApiSecret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))
	var schema SchemaDescribeResponse

	if cmdErr != nil {
		return schema, cmdErr
	}

	cmdOutputString := string(cmdOutput)
	split := strings.SplitN(cmdOutputString, ":", 2)

	if len(split) > 2 {
		return schema, errors.New("Invalid response from Describe")
	}

	err := json.Unmarshal([]byte(split[1]), &schema)

	return schema, err
}

func executeCommand(cmd exec.Cmd) ([]byte, error) {
	execCmd := exec.Command(cmd.Path, cmd.Args...)

	execCmd.Env = append(execCmd.Env, config.GetEnvValue("PATH", ""))

	out, err := execCmd.CombinedOutput()

	if err != nil {
		return out, err
	}

	return out, err
}
