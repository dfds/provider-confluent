package schemaregistry

import (
	"github.com/dfds/provider-confluent/internal/clients"
	"log"
	"testing"

	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
)

var (
	testSchema = `
{
  "type" : "record",
  "namespace" : "Example",
  "name" : "Employee",
  "fields" : [
    { "name" : "Name" , "type" : "string" },
    { "name" : "Age" , "type" : "int" }
  ]
}`
	
)

func TestDescribeSchemaRegistryCommand(t *testing.T) {
	var describeCommand = commands.NewDescribeSchemaRegistryCommand("subject", "version", "environment", "key", "secret")

	if (describeCommand.Args[4] != "subject") {
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "version") {
		t.Errorf("Version is not in correct index")
	}
	
	if(describeCommand.Args[8] != "environment") {
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[10] != "key") {
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[12] != "secret") {
		t.Errorf("Secret is not in correct index")
	}
}

func TestDeleteSchemaRegistryCommand(t *testing.T) {
	var describeCommand = commands.NewDeleteSchemaRegistryCommand("subject", "version", true, "environment", "key", "secret")

	if(describeCommand.Args[4] != "subject") {
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "version") {
		t.Errorf("Version is not in correct index")
	}
	
	if(describeCommand.Args[8] != "environment") {
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[10] != "key") {
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[12] != "secret") {
		t.Errorf("Secret is not in correct index")
	}
	
	if(describeCommand.Args[13] != "--permanent") {
		t.Errorf("Permanent flag is not set at correct index")
	}
}

func TestCreateSchemaRegistryCommand(t *testing.T) {
	var describeCommand = commands.NewCreateSchemaRegistryCommand("subject", "schema", "schemaType", "environment", "key", "secret")

	if(describeCommand.Args[4] != "subject") {
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "schema") {
		t.Errorf("Schema is not in correct index")
	}
	
	if(describeCommand.Args[8] != "schemaType") {
		t.Errorf("SchemaType is not in correct index")
	}

	if(describeCommand.Args[10] != "environment") {
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[12] != "key") {
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[14] != "secret") {
		t.Errorf("Secret is not in correct index")
	}
}

func TestClientCreate(t *testing.T) {
	client := NewClient(clients.Config{
		ApiKey:    "",
		ApiSecret: "",
	})

	resp, err := client.Create("provider-confluent-testclientcreate", testSchema, "AVRO", "env-zvzz7")
	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	// Teardown
	_, err = client.Delete("provider-confluent-testclientcreate", "all", false, "env-zvzz7")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(resp)
}

func TestClientDelete(t *testing.T) {
	client := NewClient(clients.Config{
		ApiKey:    "",
		ApiSecret: "",
	})

	respCreate, err := client.Create("provider-confluent-testclientdelete", testSchema, "AVRO", "env-zvzz7")
	if err != nil {
		log.Println(respCreate)
		t.Errorf(err.Error())
	}

	resp, err := client.Delete("provider-confluent-testclientdelete", "all", false, "env-zvzz7")
	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	t.Log(resp)
}

func TestClientDescribe(t *testing.T) {
	client := NewClient(clients.Config{
		ApiKey:    "",
		ApiSecret: "",
	})

	respCreate, err := client.Create("provider-confluent-testclientdescribe", testSchema, "AVRO", "env-zvzz7")
	if err != nil {
		log.Println(respCreate)
		t.Errorf(err.Error())
	}

	resp, err := client.Describe("provider-confluent-testclientdescribe", "latest", "env-zvzz7")
	if err != nil {
		log.Println(resp)
		t.Errorf(err.Error())
	}

	// Teardown
	_, err = client.Delete("provider-confluent-testclientdescribe", "all", false, "env-zvzz7")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(resp)
}