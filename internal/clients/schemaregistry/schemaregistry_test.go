package schemaregistry

import (
	"testing"

	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
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