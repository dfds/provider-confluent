package serviceaccount

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount/commands"
)

// Errors
const (
	ErrAlreadyInUse       = "service account already in use"
	ErrDescriptionTooLong = "service account description exceed 128 characters"
	ErrNameTooLong        = "service account name exceed 64 characters"
	ErrNotExists          = "service account does not exists"
)

const (
	nameMaxLength        = 64
	descriptionMaxLength = 128
)

// NewClient is a factory method for schemaregistry client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

func (c *Client) ServiceAccountCreate(name string, description string) (ServiceAccount, error) {
	var resp ServiceAccount

	// TODO: consider hitting the API and then handling the error
	if len(name) > nameMaxLength {
		return resp, errors.New(ErrNameTooLong)
	}
	// TODO: consider hitting the API and then handling the error
	if isDescriptionValid(description) {
		return resp, errors.New(ErrDescriptionTooLong)
	}

	var cmd = commands.NewServiceAccountCreateCommand(name, description)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		if strings.Contains(string(out), "Service name is already in use") {
			return resp, errors.New(ErrAlreadyInUse)
		}
		return resp, errors.Wrap(err, string(out))
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) ServiceAccountList(name string) (ServiceAccount, error) {
	var cmd = commands.NewServiceAccountListCommand()
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return ServiceAccount{}, errors.Wrap(err, string(out))
	}

	var resp ServiceAccountList
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return ServiceAccount{}, err
	}

	for _, v := range resp {
		if v.Name == name {
			return v, nil
		}
	}

	return ServiceAccount{}, errors.New(ErrNotExists)
}

func (c *Client) ServiceAccountUpdate(id string, description string) error {
	// TODO: consider hitting the API and then handling the error
	if isDescriptionValid(description) {
		return errors.New(ErrDescriptionTooLong)
	}

	var cmd = commands.NewServiceAccountUpdateCommand(id, description)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		if strings.Contains(string(out), "Service Account Not Found") {
			return errors.New(ErrNotExists)
		}
		return errors.Wrap(err, string(out))
	}

	return nil
}

func (c *Client) ServiceAccountDelete(id string) error {
	var cmd = commands.NewServiceAccountDeleteCommand(id)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		if strings.Contains(string(out), "error deleting service account: Forbidden") {
			return errors.New(ErrNotExists)
		}
		return errors.Wrap(err, string(out))
	}
	return nil
}

func isDescriptionValid(description string) bool {
	return len(description) > descriptionMaxLength
}
