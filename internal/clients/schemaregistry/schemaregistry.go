package schemaregistry

import (
	"encoding/json"
	"fmt"
	"os/exec"

	cli "github.com/dfds/provider-confluent/internal/clients"
)

type Client interface {
	Create(subject string, schema string, schemaType string, environment string)
	Delete(subject string, version string, permanent bool)
	Describe(subject string, version string) (SchemaDescribeResponse, error)
}

func NewClient(c cli.Config) Client {
	return &ConfluentCLIClient{Config: c}
}

type ConfluentCLIClient struct {
	Config cli.Config
}

func (c *ConfluentCLIClient) Create(subject string, schema string, schemaType string, environment string) {
	//TODO: Implement
}

func (c *ConfluentCLIClient) Delete(subject string, version string, permanent bool) {
	//TODO: Implement
}

func (c *ConfluentCLIClient) Describe(subject string, version string) (SchemaDescribeResponse, error) {
	// TODO: Abstract command into seperate package
	out, err := exec.Command("ccloud", "schema-registry", "schema", "describe", "--subject", subject, "--version", version).CombinedOutput()
	if err != nil {
		fmt.Println("FML", err)
	}
	var schema SchemaDescribeResponse
	err = json.Unmarshal([]byte(out), &schema)
	if err != nil {
		fmt.Println("FML unmarshalling fucked up", err)
	}
	return schema, err
}
