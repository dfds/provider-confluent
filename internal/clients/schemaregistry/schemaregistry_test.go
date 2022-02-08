// nolint
package schemaregistry

import (
	"log"
	"testing"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
	"go.dfds.cloud/utils/config"
)

const testSchema = `
{
  "type" : "record",
  "namespace" : "Example",
  "name" : "Employee",
  "fields" : [
    { "name" : "Name" , "type" : "string" },
    { "name" : "Age" , "type" : "int" }
  ]
}`

var testConfig = Config{
	APICredentials: clients.APICredentials{
		Identifier: "FOO",
		Key:        config.GetEnvValue("CONFLUENT_PROVIDER_API_KEY", ""),
		Secret:     config.GetEnvValue("CONFLUENT_PROVIDER_API_SECRET", ""),
	},
	SchemaPath: config.GetEnvValue("PROVIDER_CONFLUENT_SCHEMA_FILE_LOCATION", "/tmp"),
}

var (
	environment = config.GetEnvValue("CONFLUENT_ENVIRONMENT", "")
)

func TestSchemaDescribeCommand(t *testing.T) {
	var describeCommand = commands.NewSchemaDescribeCommand("subject", "version", "environment", "key", "secret")

	if describeCommand.Args[4] != "subject" {
		t.Errorf("Subject is not in correct index")
	}

	if describeCommand.Args[6] != "version" {
		t.Errorf("Version is not in correct index")
	}

	if describeCommand.Args[8] != "environment" {
		t.Errorf("Environment is not in correct index")
	}

	if describeCommand.Args[10] != "key" {
		t.Errorf("Key is not in correct index")
	}

	if describeCommand.Args[12] != "secret" {
		t.Errorf("Secret is not in correct index")
	}
}

func TestSchemaDeleteCommand(t *testing.T) {
	var describeCommand = commands.NewSchemaDeleteCommand("subject", "version", true, "environment", "key", "secret")

	if describeCommand.Args[4] != "subject" {
		t.Errorf("Subject is not in correct index")
	}

	if describeCommand.Args[6] != "version" {
		t.Errorf("Version is not in correct index")
	}

	if describeCommand.Args[8] != "environment" {
		t.Errorf("Environment is not in correct index")
	}

	if describeCommand.Args[10] != "key" {
		t.Errorf("Key is not in correct index")
	}

	if describeCommand.Args[12] != "secret" {
		t.Errorf("Secret is not in correct index")
	}

	if describeCommand.Args[13] != "--permanent" {
		t.Errorf("Permanent flag is not set at correct index")
	}
}

func TestSchemaCreateCommand(t *testing.T) {
	var describeCommand = commands.NewSchemaCreateCommand("subject", "schema", "schemaType", "environment", "key", "secret")

	if describeCommand.Args[4] != "subject" {
		t.Errorf("Subject is not in correct index")
	}

	if describeCommand.Args[6] != "schema" {
		t.Errorf("Schema is not in correct index")
	}

	if describeCommand.Args[8] != "schemaType" {
		t.Errorf("SchemaType is not in correct index")
	}

	if describeCommand.Args[10] != "environment" {
		t.Errorf("Environment is not in correct index")
	}

	if describeCommand.Args[12] != "key" {
		t.Errorf("Key is not in correct index")
	}

	if describeCommand.Args[14] != "secret" {
		t.Errorf("Secret is not in correct index")
	}
}

func TestClientSchemaCreate(t *testing.T) {
	clients.SkipCI(t)
	client := NewClient(testConfig)

	resp, err := client.SchemaCreate("provider-confluent-testclientcreate", testSchema, "AVRO", environment)

	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	// Teardown
	_, err = client.SchemaDelete("provider-confluent-testclientcreate", "all", false, environment)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(resp)
}

func TestClientSchemaDelete(t *testing.T) {
	clients.SkipCI(t)
	client := NewClient(testConfig)

	respCreate, err := client.SchemaCreate("provider-confluent-testclientdelete", testSchema, "AVRO", environment)
	if err != nil {
		log.Println(respCreate)
		t.Errorf(err.Error())
	}

	resp, err := client.SchemaDelete("provider-confluent-testclientdelete", "all", false, environment)
	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	t.Log(resp)
}

func TestClientSchemaDescribe(t *testing.T) {
	clients.SkipCI(t)
	client := NewClient(testConfig)

	respCreate, err := client.SchemaCreate("provider-confluent-testclientdescribe", testSchema, "AVRO", environment)
	if err != nil {
		log.Println(respCreate)
		t.Errorf(err.Error())
	}

	resp, err := client.SchemaDescribe("provider-confluent-testclientdescribe", "latest", environment)
	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	// Teardown
	_, err = client.SchemaDelete("provider-confluent-testclientdescribe", "all", false, environment)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(resp)
}
