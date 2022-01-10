package acl

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/apikey/commands"
)

// Errors
const (
	errUnknown                              = "unknown error"
	errEnvironmentNotFound                  = "environment not found"
	errServiceAccountNotFoundOrLimitReached = "service not found or limit reached"
	errResourceNotFoundOrAccessForbidden    = "resource not found or access forbidden"
	ErrNotExists                            = "api key does not exists"
	errUnknownApiKey                        = "unknow apikey"
)

// NewClient is a factory method for apikey client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

func (c *Client) ApiKeyCreate(resource string, description string, serviceAccount string, environment string) (ApiKey, error) {
	var resp ApiKey

	var cmd = commands.NewApiKeyCreateCommand(resource, description, serviceAccount, environment)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return resp, errorParser(out)
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) GetApiKeyByKey(key string) (ApiKeyMetadata, error) {
	var resp ApiKeyList
	var akm ApiKeyMetadata

	var cmd = commands.NewApiKeyListCommand()
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return akm, errorParser(out)
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return akm, err
	}

	for _, v := range resp {
		if v.Key == key {
			akm = v
			return akm, nil
		}
	}

	return akm, errors.New(ErrNotExists)
}

func (c *Client) ApiKeyUpdate(key string, description string) error {
	var cmd = commands.NewApiKeyUpdateCommand(key, description)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return errorParser(out)
	}

	return nil
}

func (c *Client) ApiKeyDelete(id string) error {
	var cmd = commands.NewApiKeyDeleteCommand(id)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return errorParser(out)
	}
	return nil
}

func errorParser(cmdout []byte) error {
	str := string(cmdout)
	if strings.Contains(str, "Error: environment") && strings.Contains(str, "not found") {
		return errors.New(errEnvironmentNotFound)
	} else if strings.Contains(str, "Your Api Keys per User is currently limited to 10") {
		return errors.New(errServiceAccountNotFoundOrLimitReached)
	} else if strings.Contains(str, "Error: Kafka cluster not found or access forbidden") {
		return errors.New(errResourceNotFoundOrAccessForbidden)
	} else if strings.Contains(str, "Error: Unknown API key") {
		return errors.New(errUnknownApiKey)
	}
	return errors.Wrap(errors.New(errUnknown), string(str))
}
