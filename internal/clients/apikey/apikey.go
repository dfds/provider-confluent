package apikey

import (
	"encoding/json"
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
	errUnknownAPIKey                        = "unknown apikey"
)

// NewClient is a factory method for apikey client
func NewClient() IClient {
	return &Client{}
}

// APIKeyCreate create API key
func (c *Client) APIKeyCreate(resource string, description string, serviceAccount string, environment string) (APIKey, error) {
	var resp APIKey

	var cmd = commands.NewAPIKeyCreateCommand(resource, description, serviceAccount, environment)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return resp, errorParser(out)
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetAPIKeyByKey get API key by key
func (c *Client) GetAPIKeyByKey(key string) (Metadata, error) {
	var resp List
	var akm Metadata

	var cmd = commands.NewAPIKeyListCommand()
	out, err := clients.ExecuteCommand(cmd)

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

// APIKeyUpdate update API key description by key
func (c *Client) APIKeyUpdate(key string, description string) error {
	var cmd = commands.NewAPIKeyUpdateCommand(key, description)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}

	return nil
}

// APIKeyDelete delete API key by key
func (c *Client) APIKeyDelete(key string) error {
	var cmd = commands.NewAPIKeyDeleteCommand(key)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}
	return nil
}

func errorParser(cmdout []byte) error {
	str := string(cmdout)

	switch {
	case strings.Contains(str, "Error: environment"):
		return errors.New(errEnvironmentNotFound)
	case strings.Contains(str, "Your Api Keys per User is currently limited to 10"):
		return errors.New(errServiceAccountNotFoundOrLimitReached)
	case strings.Contains(str, "Error: Kafka cluster not found or access forbidden"):
		return errors.New(errResourceNotFoundOrAccessForbidden)
	case strings.Contains(str, "Error: Unknown API key"):
		return errors.New(errUnknownAPIKey)
	default:
		return errors.Wrap(errors.New(errUnknown), str)
	}
}
