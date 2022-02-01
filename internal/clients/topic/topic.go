package topic

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/topic/commands"
)

// Errors
const (
	errUnknown      = "unknown error"
	ErrUnknownTopic = "unknown topic"
	ErrInvalidInput = "input given may be invalid like empty topic name or so"
)

// NewClient is a factory method for apikey client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

// TopicCreate Executes Confluent CLI command to create a Topic in Confluent Cloud
func (c *Client) TopicCreate(tp v1alpha1.TopicParameters) error {

	var cmd = commands.NewTopicCreateCommand(tp)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}

	return nil
}

// TopicDescribe Executes Confluent CLI command to retrieve metadata about a Topic from Confluent Cloud
func (c *Client) TopicDescribe(to v1alpha1.TopicObservation) (DescribeResponse, error) {
	var resp DescribeResponse

	cmd := commands.NewTopicDescribeCommand(to)
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

// TopicUpdate Executes Confluent CLI command, and with its given TopicParameters, attempts to update a Topic in Confluent Cloud
func (c *Client) TopicUpdate(tp v1alpha1.TopicParameters) error {

	cmd := commands.NewTopicUpdateCommand(tp)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}

	return nil
}

// TopicDelete Executes Confluent CLI command, and with its given TopicParameters, attempts to delete a Topic in Confluent Cloud
func (c *Client) TopicDelete(tp v1alpha1.TopicParameters) error {
	cmd := commands.NewTopicDeleteCommand(tp)
	out, err := clients.ExecuteCommand(cmd)

	if err != nil {
		return errorParser(out)
	}

	return nil
}

func errorParser(cmdout []byte) error {
	str := string(cmdout)
	if strings.Contains(str, "Error: unknown topic") {
		return errors.New(ErrUnknownTopic)
	} else if strings.Contains(str, "Error: REST request failed") {
		return errors.New(ErrInvalidInput)
	}
	return errors.Wrap(errors.New(errUnknown), str)
}
