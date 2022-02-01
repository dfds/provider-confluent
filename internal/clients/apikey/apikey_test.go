package apikey

import (
	"testing"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/stretchr/testify/assert"
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

// Asses and assert
func TestAPIKeyLifecycle(t *testing.T) {
	assert := assert.New(t)

	_, err := client.GetAPIKeyByKey("")
	if err != nil {
		assert.Equal(err.Error(), ErrNotExists, "empty key should should return not exists")
	}

	out, err := client.APIKeyCreate(resource, description, serviceAccount, environment)
	if err != nil {
		t.Errorf("api-key creation not working")
	}

	_, err = client.GetAPIKeyByKey(out.Key)
	if err != nil {
		t.Errorf("api-key get by key not working")
	}

	err = client.APIKeyUpdate(out.Key, "crossplane-test0")
	if err != nil {
		t.Errorf("api-key update not working")
	}

	err = client.APIKeyDelete(out.Key)
	if err != nil {
		t.Errorf("api-key delete not working manual OBS: clean up required, please run the following command \"confluent api-key list | grep \"crossplane-test\" | awk '{ print $1 }' | xargs -I {} confluent api-key delete {}\"")
	}

	_, err = client.GetAPIKeyByKey(out.Key)
	if err != nil {
		assert.Equal(err.Error(), ErrNotExists, "deleted key should should return not exists")
	}
}
