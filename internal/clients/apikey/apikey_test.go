package apikey

import (
	"testing"

	"github.com/dfds/provider-confluent/internal/clients"
	"go.dfds.cloud/utils/config"
)

// Assign
var (
	testConfig = Config{
		APICredentials: clients.APICredentials{
			Identifier: "FOO",
			Key:        config.GetEnvValue("CONFLUENT_PROVIDER_API_KEY", ""),
			Secret:     config.GetEnvValue("CONFLUENT_PROVIDER_API_SECRET", ""),
		},
	}
	resource       = config.GetEnvValue("CONFLUENT_CLUSTER_ID", "")
	environment    = config.GetEnvValue("CONFLUENT_ENVIRONMENT", "")
	serviceAccount = config.GetEnvValue("CONFLUENT_SERVICEACCOUNT", "")
	description    = "crossplane-test"
	client         = NewClient(testConfig)
)

// Asses
func TestApiKeyCreate(t *testing.T) {
	_, err := client.ApiKeyCreate(resource, description, serviceAccount, environment)
	if err != nil {
		t.Errorf("api-key creatio not working")
	}
}

// Assert
