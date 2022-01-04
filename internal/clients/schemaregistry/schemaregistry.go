package schemaregistry

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
)

const (
	ErrNotFound             = "schema not found"
	errNotParsed            = "cannot be parsed"
	errGeneral              = "unknown error:"
	errInvalidResponse      = "invalid response from describe"
	ErrNotCompatible        = "schema not compatible"
	ErrInvalidCompatibility = "invalid compatibility level"
	errUnkownFormat         = "unknow error format"
)

// NewClient is a factory method for schemaregistry client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

// SchemaCreate creates a schema in the schemaregistry
func (c *Client) SchemaCreate(subject string, schema string, schemaType string, environment string) (string, error) {
	schemaGUID := uuid.New().String()

	path, err := CreateFile([]byte(schema), schemaGUID, c.Config.SchemaPath)

	if err != nil {
		return "", err
	}

	var cmd = commands.NewSchemaCreateCommand(subject, path, schemaType, environment, c.Config.APICredentials.Key, c.Config.APICredentials.Secret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	err = RemoveFile(path) // TODO: consider implementing with defer

	if err != nil {
		return string(cmdOutput), err
	}

	if cmdErr != nil {
		err = errorParser(cmdOutput)
		return string(cmdOutput), err
	}

	return string(cmdOutput), nil
}

// SchemaDelete deletes a schema in the schemaregistry
func (c *Client) SchemaDelete(subject string, version string, permanent bool, environment string) (string, error) {
	var cmd = commands.NewSchemaDeleteCommand(subject, version, permanent, environment, c.Config.APICredentials.Key, c.Config.APICredentials.Secret)
	var cmdOutput, cmdErr = executeCommand(exec.Cmd(cmd))

	return string(cmdOutput), cmdErr
}

// SchemaDescribe gets a schema in the schemaregistry
func (c *Client) SchemaDescribe(subject string, version string, environment string) (SchemaDescribeResponse, error) {
	var cmd = commands.NewSchemaDescribeCommand(subject, version, environment, c.Config.APICredentials.Key, c.Config.APICredentials.Secret)
	var cmdOutput, err = executeCommand(exec.Cmd(cmd))
	var schema SchemaDescribeResponse

	if err != nil {
		err = errorParser(cmdOutput)
		return schema, err
	}

	split, err := responseSanitiser(cmdOutput)

	if err != nil {
		return schema, err
	}

	err = json.Unmarshal([]byte(split[1]), &schema)

	return schema, err
}

func (c *Client) SchemaSubjectUpdateCommand(subject string, compatibility string, environment string) (string, error) {
	var cmd = commands.NewSchemaSubjectUpdateCommand(subject, compatibility, environment, c.Config.APICredentials.Key, c.Config.APICredentials.Secret)
	cmdOutput, err := executeCommand(exec.Cmd(cmd))

	if err != nil {
		return string(cmdOutput), errorParser(cmdOutput)
	}

	return string(cmdOutput), nil
}

func executeCommand(cmd exec.Cmd) ([]byte, error) {
	execCmd := exec.Command(cmd.Path, cmd.Args...) //nolint:gosec
	execCmd.Env = os.Environ()

	out, err := execCmd.CombinedOutput()

	if err != nil {
		return out, err
	}

	return out, err
}

func errorParser(cmdout []byte) error {
	split, err := responseSanitiser(cmdout)
	if err != nil {
		return err
	}

	var schema errorResponse

	if len(split) <= 1 {
		return errors.New(errUnkownFormat)
	}

	fmt.Println(split[1])

	err = json.Unmarshal([]byte(split[1]), &schema)
	if err != nil {
		message := fmt.Sprintf("%s%v", errGeneral, split)
		err = errors.Wrap(err, message)
		return err
	}

	switch schema.ErrorCode {
	case 409:
		return errors.New(ErrNotCompatible)
	case 40401:
		return errors.New(ErrNotFound)
	case 42203:
		return errors.New(ErrInvalidCompatibility)
	default:
		return errors.New(fmt.Sprintf("%s%d%s", errGeneral, schema.ErrorCode, schema.Message))
	}
}

func responseSanitiser(cmdoutput []byte) ([]string, error) {
	out := string(cmdoutput)
	split := strings.SplitN(out, ":", 2)
	if len(split) > 2 {
		return []string{}, errors.New(errInvalidResponse)
	}
	return split, nil
}

type errorResponse struct {
	ErrorCode int64  `json:"error_code"`
	Message   string `json:"message"`
}
