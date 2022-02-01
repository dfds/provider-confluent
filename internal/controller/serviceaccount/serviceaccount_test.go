package serviceaccount

import (
	"testing"

	"github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestExternalNameHelper(t *testing.T) {
	assert := assert.New(t)

	// No external name
	sa := v1alpha1.ServiceAccount{}
	sa.Name = "name"
	name, exists := ExternalNameHelper(&sa)
	assert.Equal(sa.Name, name)
	assert.False(exists)

	// With external name
	extName := make(map[string]string)
	extName["crossplane.io/external-name"] = "extname"
	sa.SetAnnotations(extName)
	name, exists = ExternalNameHelper(&sa)
	assert.NotEqual(sa.Name, name, "external name not used")
	assert.Equal("extname", name, "external name not used")
	assert.True(exists)
}

func TestCreateResource(t *testing.T) {
	assert := assert.New(t)

	// Resource do not exists
	sa := v1alpha1.ServiceAccount{}
	create, err := ObserveCreateResource(&sa, errors.New(serviceaccount.ErrNotExists))
	if err != nil {
		t.Errorf("no error expected when ErrorNotExists is passed to function")
	} else {
		assert.True(create, "resource do not exists so it should create")
	}

	// Resource has no status
	create, err = ObserveCreateResource(&sa, nil)
	if err != nil {
		t.Errorf("no error expected when non given")
	} else {
		assert.True(create, "resource has not status set so it should create")
	}

	// Resource has status
	sa.Status.AtProvider.ID = "sa-55555"
	create, err = ObserveCreateResource(&sa, nil)
	if err != nil {
		t.Errorf("no error expected when non given")
	} else {
		assert.False(create, "resource has status set so it should not create")
	}

	// Unknow error
	const uErr = "unknown"
	_, err = ObserveCreateResource(&sa, errors.New(uErr))
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
	sa := v1alpha1.ServiceAccount{}
	sa.Spec.ForProvider.Description = description
	sac := serviceaccount.ServiceAccount{}
	sac.Description = description
	assert.False(ObserveUpdateResource(&sa, sac), "no update required when descriptions match")

	// Descriptions do not match
	sa.Spec.ForProvider.Description = "almost my description"
	assert.True(ObserveUpdateResource(&sa, sac), "update required when descriptions do not match")
}

func TestCreateResourceIsImport(t *testing.T) {
	assert := assert.New(t)

	// ErrNotExists
	isImport, err := CreateResourceIsImport(errors.New(serviceaccount.ErrNotExists))
	assert.False(isImport)
	assert.NoError(err)

	// unknow error
	const uErr = "unknow error"
	isImport, err = CreateResourceIsImport(errors.New(uErr))
	assert.False(isImport)
	assert.Equal(err.Error(), uErr)

	// nil error
	isImport, err = CreateResourceIsImport(nil)
	assert.True(isImport)
	assert.NoError(err)
}
