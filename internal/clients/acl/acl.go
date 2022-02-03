package acl

import (
	"encoding/json"
	"strings"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/acl/commands"
)

// Errors
const (
	errUnknown = "unknown error"
	// errEnvironmentNotFound                  = "environment not found"
	// errServiceAccountNotFoundOrLimitReached = "service not found or limit reached"
	// errResourceNotFoundOrAccessForbidden    = "resource not found or access forbidden"
	ErrACLNotExistsOrInvalidServiceAccount = "acl for service account does not exists or invalid service account"
	// errUnknownApiKey                        = "unknow apikey"
)

// NewClient is a factory method for apikey client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

// ACLCreate create acl
func (c *Client) ACLCreate(aclP v1alpha1.ACLParameters) ([]v1alpha1.ACLRule, error) {
	var resp []v1alpha1.ACLRule

	cmd, err := commands.NewACLCreateCommand(aclP)
	if err != nil {
		return resp, err
	}

	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return resp, errorParser(out)
	}

	var aclBlocks []Block
	err = json.Unmarshal(out, &aclBlocks)
	if err != nil {
		return resp, err
	}

	for _, block := range aclBlocks {
		resp = append(resp, FromACLBlockToACLRule(block))
	}

	return resp, nil
}

// ACLDelete delete ACL
func (c *Client) ACLDelete(aclP v1alpha1.ACLParameters) error {
	cmd, err := commands.NewACLDeleteCommand(aclP)
	if err != nil {
		return err
	}

	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}

	return nil
}

// ACLList list ACL's
func (c *Client) ACLList(serviceAccount string, environment string, cluster string) ([]v1alpha1.ACLRule, error) {
	var resp []v1alpha1.ACLRule

	cmd := commands.NewACLListCommand(environment, cluster, serviceAccount)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return resp, errorParser(out)
	}

	var aclBlocks []Block
	err = json.Unmarshal(out, &aclBlocks)
	if err != nil {
		return resp, err
	}

	for _, block := range aclBlocks {
		resp = append(resp, FromACLBlockToACLRule(block))
	}

	if len(resp) == 0 {
		return resp, errors.New(ErrACLNotExistsOrInvalidServiceAccount)
	}

	return resp, nil
}

//TODO: Improve error parsing, e.g. don't always returning errUnknown.
func errorParser(cmdout []byte) error {
	str := string(cmdout)

	switch {
	case strings.Contains(str, "Error: service account") && strings.Contains(str, "not found"):
		return errors.New(ErrACLNotExistsOrInvalidServiceAccount)
	default:
		return errors.Wrap(errors.New(errUnknown), str)
	}
}
