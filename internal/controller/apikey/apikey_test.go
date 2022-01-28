package apikey

import (
	"testing"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
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

func TestObserveCreateResource(t *testing.T) {
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

func TestObserveUpdateResourceAndUpdateResourceDestrutive(t *testing.T) {
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
	assert.False(updateResourceDestructive(&ak, akm), "updates to description is not destructive")

	// Environment match
	environment := "env-vvvvv"
	ak = v1alpha1.ApiKey{}
	ak.Spec.ForProvider.Environment = environment
	akm = apikey.ApiKeyMetadata{}
	ak.Status.AtProvider.Environment = environment
	assert.False(observeUpdateResource(&ak, akm), "no update required when environment match")

	// Environment do not match
	ak.Spec.ForProvider.Environment = "env-vvvvy"
	assert.True(observeUpdateResource(&ak, akm), "update required when environment do not match")
	assert.True(updateResourceDestructive(&ak, akm), "updates to environment is destructive")

	// Resource match
	resource := "lkc-yyyyy"
	ak = v1alpha1.ApiKey{}
	ak.Spec.ForProvider.Resource = resource
	akm = apikey.ApiKeyMetadata{}
	ak.Status.AtProvider.Resource = resource
	assert.False(observeUpdateResource(&ak, akm), "no update required when Resource match")

	// Resource do not match
	ak.Spec.ForProvider.Resource = "lkc-yyyyv"
	assert.True(observeUpdateResource(&ak, akm), "update required when Resource do not match")
	assert.True(updateResourceDestructive(&ak, akm), "updates to resource is destructive")

	// Service account match
	sa := "sa-55555"
	ak = v1alpha1.ApiKey{}
	ak.Spec.ForProvider.ServiceAccount = sa
	akm = apikey.ApiKeyMetadata{}
	ak.Status.AtProvider.ServiceAccount = sa
	assert.False(observeUpdateResource(&ak, akm), "no update required when Resource match")

	// Service account do not match
	ak.Spec.ForProvider.ServiceAccount = "sa-55556"
	assert.True(observeUpdateResource(&ak, akm), "update required when Resource do not match")
	assert.True(updateResourceDestructive(&ak, akm), "updates to service account is destructive")
}

func TestCreateResourceIsImport(t *testing.T) {
	assert := assert.New(t)

	// Error is not exists
	isImport, err := createResourceIsImport(errors.New(apikey.ErrNotExists))
	assert.Equal(err, nil)
	assert.False(isImport)

	// Error is inl
	isImport, err = createResourceIsImport(nil)
	assert.Equal(err, nil)
	assert.True(isImport)

	// Unknown error
	const Uerr = "I have no idea of what I'm doing"
	isImport, err = createResourceIsImport(errors.New(Uerr))
	assert.Error(err)
	assert.False(isImport)
}

func TestDestructiveIsAllowed(t *testing.T) {
	assert := assert.New(t)

	var dp v1.DeletionPolicy = "Orphan"
	assert.False(destructiveActionsAllowed(dp))

	dp = "Delete"
	assert.True(destructiveActionsAllowed(dp))

	dp = "blabla"
	assert.False(destructiveActionsAllowed(dp))
}
