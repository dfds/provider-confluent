package topic

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/topic/commands"
)

// Errors
const (
	errUnknown = "unknown error"
)

// NewClient is a factory method for apikey client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

func (c *Client) TopicCreate(tp v1alpha1.TopicParameters) (interface{}, error) {
	var resp interface{}

	var cmd = commands.NewTopicCreateCommand(tp)
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

func (c *Client) TopicDescribe(to v1alpha1.TopicObservation) (interface{}, error) {
	var resp interface{}

	cmd := commands.NewTopicDescribeCommand(to)
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

func (c *Client) TopicUpdate(tp v1alpha1.TopicParameters) error {

	cmd := commands.NewTopicUpdateCommand(tp)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return errorParser(out)
	}

	return nil
}

func (c *Client) TopicDelete(tp v1alpha1.TopicParameters) error {
	cmd := commands.NewTopicDeleteCommand(tp)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return errorParser(out)
	}

	return nil
}

func errorParser(cmdout []byte) error {
	str := string(cmdout)
	return errors.Wrap(errors.New(errUnknown), string(str))
}
