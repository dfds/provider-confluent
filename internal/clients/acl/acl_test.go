package acl

import (
	"fmt"
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/stretchr/testify/assert"
	"go.dfds.cloud/utils/config"
	"testing"
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
	aclParam       = v1alpha1.ACLParameters{
		ACLRule: v1alpha1.ACLRule{
			Operation:    "CREATE",
			PatternType:  "LITERAL",
			Permission:   "ALLOW",
			Principal:    fmt.Sprintf("User:%s", serviceAccount),
			ResourceName: "acltest_testacllifecycle",
			ResourceType: "TOPIC",
		},
		Environment: environment,
		Cluster:     resource,
	}
)

// Asses and assert
func TestAclLifecycle(t *testing.T) {
	assert := assert.New(t)

	_, err := client.ACLList("sa-00000", environment, resource)
	if err != nil {
		assert.Equal(err.Error(), ErrACLNotExistsOrInvalidServiceAccount, "empty acl should should return not exists")
	}

	_, err = client.ACLCreate(aclParam)
	if err != nil {
		t.Errorf("acl creation not working")
	}

	resp, err := client.ACLList(serviceAccount, environment, resource)
	if err != nil {
		t.Errorf("acl list not working")
	}

	if len(resp) != 1 {
		t.Errorf("Expected amount of ACLS after creation is not 1. Could be affected by external factors")
	}

	err = client.ACLDelete(aclParam)
	if err != nil {
		t.Errorf("acl delete not working manual OBS: clean up required, please run the following command \"confluent kafka acl list | grep \"acltest_testacllifecycle\" | awk '{ print $1 }' | xargs -I {} confluent kafka acl delete {}\"")
	}

	_, err = client.ACLList(serviceAccount, environment, resource)
	if err == nil {
		t.Errorf("acl deletion didn't work. 1 or more ACLS are attached to the specified service account, cluster & environment")
	}
}
