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

// NewClient is a factory method for serviceaccount client
func NewClient() IClient {
	return &Client{}
}

// ServiceAccountCreate Executes Confluent CLI command to create ServiceAccount in Confluent Cloud & return a ServiceAccount object
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

// ServiceAccountList Executes Confluent CLI command to list all ServiceAccounts in Confluent Cloud & return a slice of ServiceAccount objects
func (c *Client) ServiceAccountList() ([]ServiceAccount, error) {
	var cmd = commands.NewServiceAccountListCommand()
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return []ServiceAccount{}, errors.Wrap(err, string(out))
	}

	var resp List
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return []ServiceAccount{}, err
	}

	return resp, errors.New(ErrNotExists)
}

// ServiceAccountByID Executes Confluent CLI command to list all ServiceAccounts in Confluent Cloud, filter by id & return a non-empty ServiceAccount object if found
func (c *Client) ServiceAccountByID(id string) (ServiceAccount, error) {
	var cmd = commands.NewServiceAccountListCommand()
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return ServiceAccount{}, errors.Wrap(err, string(out))
	}

	var resp List
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return ServiceAccount{}, err
	}

	for _, v := range resp {
		if v.ID == id {
			return v, nil
		}
	}

	return ServiceAccount{}, errors.New(ErrNotExists)
}

// ServiceAccountByName Executes Confluent CLI command to list all ServiceAccounts in Confluent Cloud, filter by name & return a non-empty ServiceAccount object if found
func (c *Client) ServiceAccountByName(name string) (ServiceAccount, error) {
	var cmd = commands.NewServiceAccountListCommand()
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return ServiceAccount{}, errors.Wrap(err, string(out))
	}

	var resp List
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return ServiceAccount{}, err
	}

	for _, v := range resp {
		if strings.EqualFold(v.Name, name) {
			return v, nil
		}
	}

	return ServiceAccount{}, errors.New(ErrNotExists)
}

// ServiceAccountUpdate Executes Confluent CLI command to update the description of a ServiceAccount in Confluent Cloud
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

// ServiceAccountDelete Executes Confluent CLI command to delete a ServiceAccount in Confluent Cloud
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
