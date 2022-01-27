package apikey

import (
	"testing"

	"github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/apikey"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestExternalNameHelper(t *testing.T) {
	assert := assert.New(t)

	// No external name
	ak := v1alpha1.ApiKey{}
	ak.Status.AtProvider.Key = "name"
	name, exists := externalNameHelper(&ak)
	assert.Equal(ak.Status.AtProvider.Key, name)
	assert.False(exists)

	// With external name
	extName := make(map[string]string)
	extName["crossplane.io/external-name"] = "extname"
	ak.SetAnnotations(extName)
	name, exists = externalNameHelper(&ak)
	assert.NotEqual(ak.Name, name, "external name not used")
	assert.Equal("extname", name, "external name not used")
	assert.True(exists)
}

func TestCreateResource(t *testing.T) {
	assert := assert.New(t)

	// Resource do not exists
	create, err := observeCreateResource(errors.New(apikey.ErrNotExists))
	if err != nil {
		t.Errorf("no error expected when ErrorNotExists is passed to function")
	} else {
		assert.True(create, "resource do not exists so it should create")
	}

	// Resource exists
	create, err = observeCreateResource(nil)
	if err != nil {
		t.Errorf("no error expected when non given")
	} else {
		assert.False(create, "resource has not status set so it should create")
	}

	// Unknow error
	const uErr = "unknown"
	_, err = observeCreateResource(errors.New(uErr))
	if err == nil {
		t.Errorf("error expected when given")
	} else {
		assert.Equal(err.Error(), uErr)
	}
}

func TestObserveUpdateResource(t *testing.T) {
	assert := assert.New(t)

	// Descriptions match
	description := "my description"
	ak := v1alpha1.ApiKey{}
	ak.Spec.ForProvider.Description = description
	akm := apikey.ApiKeyMetadata{}
	akm.Description = description
	assert.False(observeUpdateResource(&ak, akm), "no update required when descriptions match")

	// Descriptions do not match
	ak.Spec.ForProvider.Description = "almost my description"
	assert.True(observeUpdateResource(&ak, akm), "update required when descriptions do not match")
}
