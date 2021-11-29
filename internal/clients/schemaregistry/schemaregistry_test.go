package schemaregistry

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/go-cmp/cmp"

	"github.com/dfds/provider-confluent/apis/schemaregistry/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/schemaregistry"
	"github.com/dfds/provider-confluent/internal/clients/schemaregistry/commands"
)

func TestDescribeSchemaRegistryCommand(t *testing.T) {
	var describeCommand = NewDescribeSchemaRegistryCommand("subject", "version", "environment", "key", "secret")

	if(describeCommand.Args[4] != "subject")
	{
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "version")
	{
		t.Errorf("Version is not in correct index")
	}
	
	if(describeCommand.Args[8] != "environment")
	{
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[10] != "key")
	{
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[12] != "secret")
	{
		t.Errorf("Secret is not in correct index")
	}
}

func TestDeleteSchemaRegistryCommand(t *testing.T) {
	var describeCommand = NewDeleteSchemaRegistryCommand("subject", "version", true, "environment", "key", "secret")

	if(describeCommand.Args[4] != "subject")
	{
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "version")
	{
		t.Errorf("Version is not in correct index")
	}
	
	if(describeCommand.Args[8] != "environment")
	{
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[10] != "key")
	{
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[12] != "secret")
	{
		t.Errorf("Secret is not in correct index")
	}
	
	if(describeCommand.Args[13] != "--permanent")
	{
		t.Errorf("Permanent flag is not set at correct index")
	}
}

func TestDeleteSchemaRegistryCommand(t *testing.T) {
	var describeCommand = NewCreateSchemaRegistryCommand("subject", "schema", "schemaType", "environment", "key", "secret")

	if(describeCommand.Args[4] != "subject")
	{
		t.Errorf("Subject is not in correct index")
	}
	
	if(describeCommand.Args[6] != "schema")
	{
		t.Errorf("Schema is not in correct index")
	}
	
	if(describeCommand.Args[8] != "schemaType")
	{
		t.Errorf("SchemaType is not in correct index")
	}

	if(describeCommand.Args[10] != "environment")
	{
		t.Errorf("Environment is not in correct index")
	}
	
	if(describeCommand.Args[12] != "key")
	{
		t.Errorf("Key is not in correct index")
	}
	
	if(describeCommand.Args[14] != "secret")
	{
		t.Errorf("Secret is not in correct index")
	}
}