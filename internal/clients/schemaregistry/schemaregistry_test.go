package schemaregistry

// import (
// 	"log"
// 	"testing"

// 	"github.com/dfds/provider-confluent/internal/clients"
// 	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
// 	"go.dfds.cloud/utils/config"
// )

// const testSchema = `
// {
//   "type" : "record",
//   "namespace" : "Example",
//   "name" : "Employee",
//   "fields" : [
//     { "name" : "Name" , "type" : "string" },
//     { "name" : "Age" , "type" : "int" }
//   ]
// }`

// var testConfig = SchemaRegistryConfig{
// 	ApiConfig: clients.ApiConfig{
// 		ApiKey:    config.GetEnvValue("PROVIDER_CONFLUENT_SCHEMA_REGISTRY_API_KEY", ""),
// 		ApiSecret: config.GetEnvValue("PROVIDER_CONFLUENT_SCHEMA_REGISTRY_API_SECRET", ""),
// 	},
// 	SchemaPath: config.GetEnvValue("PROVIDER_CONFLUENT_SCHEMA_FILE_LOCATION", "/tmp"),
// }

// func TestSchemaDescribeCommand(t *testing.T) {
// 	var describeCommand = commands.NewSchemaDescribeCommand("subject", "version", "environment", "key", "secret")

// 	if describeCommand.Args[4] != "subject" {
// 		t.Errorf("Subject is not in correct index")
// 	}

// 	if describeCommand.Args[6] != "version" {
// 		t.Errorf("Version is not in correct index")
// 	}

// 	if describeCommand.Args[8] != "environment" {
// 		t.Errorf("Environment is not in correct index")
// 	}

// 	if describeCommand.Args[10] != "key" {
// 		t.Errorf("Key is not in correct index")
// 	}

// 	if describeCommand.Args[12] != "secret" {
// 		t.Errorf("Secret is not in correct index")
// 	}
// }

// func TestSchemaDeleteCommand(t *testing.T) {
// 	var describeCommand = commands.NewSchemaDeleteCommand("subject", "version", true, "environment", "key", "secret")

// 	if describeCommand.Args[4] != "subject" {
// 		t.Errorf("Subject is not in correct index")
// 	}

// 	if describeCommand.Args[6] != "version" {
// 		t.Errorf("Version is not in correct index")
// 	}

// 	if describeCommand.Args[8] != "environment" {
// 		t.Errorf("Environment is not in correct index")
// 	}

// 	if describeCommand.Args[10] != "key" {
// 		t.Errorf("Key is not in correct index")
// 	}

// 	if describeCommand.Args[12] != "secret" {
// 		t.Errorf("Secret is not in correct index")
// 	}

// 	if describeCommand.Args[13] != "--permanent" {
// 		t.Errorf("Permanent flag is not set at correct index")
// 	}
// }

// func TestSchemaCreateCommand(t *testing.T) {
// 	var describeCommand = commands.NewSchemaCreateCommand("subject", "schema", "schemaType", "environment", "key", "secret")

// 	if describeCommand.Args[4] != "subject" {
// 		t.Errorf("Subject is not in correct index")
// 	}

// 	if describeCommand.Args[6] != "schema" {
// 		t.Errorf("Schema is not in correct index")
// 	}

// 	if describeCommand.Args[8] != "schemaType" {
// 		t.Errorf("SchemaType is not in correct index")
// 	}

// 	if describeCommand.Args[10] != "environment" {
// 		t.Errorf("Environment is not in correct index")
// 	}

// 	if describeCommand.Args[12] != "key" {
// 		t.Errorf("Key is not in correct index")
// 	}

// 	if describeCommand.Args[14] != "secret" {
// 		t.Errorf("Secret is not in correct index")
// 	}
// }

// func TestClientSchemaCreate(t *testing.T) {
// 	client := NewClient(testConfig)

// 	resp, err := client.SchemaCreate("provider-confluent-testclientcreate", testSchema, "AVRO", "env-zvzz7")

// 	if err != nil {
// 		log.Println(resp)
// 		t.Errorf(err.Error())
// 	}

// 	// Teardown
// 	_, err = client.SchemaDelete("provider-confluent-testclientcreate", "all", false, "env-zvzz7")

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	t.Log(resp)
// }

// func TestClientSchemaDelete(t *testing.T) {
// 	client := NewClient(testConfig)

// 	respCreate, err := client.SchemaCreate("provider-confluent-testclientdelete", testSchema, "AVRO", "env-zvzz7")
// 	if err != nil {
// 		log.Println(respCreate)
// 		t.Errorf(err.Error())
// 	}

// 	resp, err := client.SchemaDelete("provider-confluent-testclientdelete", "all", false, "env-zvzz7")
// 	if err != nil {
// 		log.Println(resp)
// 		t.Errorf(err.Error())
// 	}

// 	t.Log(resp)
// }

// func TestClientSchemaDescribe(t *testing.T) {
// 	client := NewClient(testConfig)

// 	respCreate, err := client.SchemaCreate("provider-confluent-testclientdescribe", testSchema, "AVRO", "env-zvzz7")
// 	if err != nil {
// 		log.Println(respCreate)
// 		t.Errorf(err.Error())
// 	}

// 	resp, err := client.SchemaDescribe("provider-confluent-testclientdescribe", "latest", "env-zvzz7")
// 	if err != nil {
// 		log.Println(resp)
// 		t.Errorf(err.Error())
// 	}

// 	// Teardown
// 	_, err = client.SchemaDelete("provider-confluent-testclientdescribe", "all", false, "env-zvzz7")
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	t.Log(resp)
// }
