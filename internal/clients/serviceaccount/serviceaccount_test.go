package serviceaccount

import (
	"testing"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/stretchr/testify/assert"
)

var (
	testConfig = Config{
		APICredentials: clients.APICredentials{},
	}
	serviceAccount = "crossplane-testing"
	description    = "crossplane-testing"
	client         = NewClient(testConfig)
)

func TestServiceAccountLifecycle(t *testing.T) {
	clients.SkipCI(t)
	assert := assert.New(t)

	_, err := client.ServiceAccountByName("")
	if err != nil {
		assert.Equal(err.Error(), ErrNotExists)
	} else {
		t.Errorf("getting an empty service account should produce error")
	}

	resp, err := client.ServiceAccountCreate(serviceAccount, description)
	if err != nil {
		t.Errorf("service account creation not working with error: %s", err.Error())
	}

	_, err = client.ServiceAccountByName(resp.Name)
	if err != nil {
		t.Errorf("could not get already created service account by name")
	}

	_, err = client.ServiceAccountByID(resp.ID)
	if err != nil {
		t.Errorf("could not get already created service account by id")
	}

	err = client.ServiceAccountUpdate(resp.ID, "crossplane-test-update")
	if err != nil {
		t.Errorf("update does not work as indented")
	}

	err = client.ServiceAccountDelete(resp.ID)
	if err != nil {
		t.Errorf("delete does not work as indented")
	}
}
