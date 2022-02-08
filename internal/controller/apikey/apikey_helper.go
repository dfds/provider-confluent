package apikey

import (
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/apikey"
)

func observeCreateResource(ak *v1alpha1.APIKey, exists bool, err error) (bool, error) {
	if err != nil {
		if err.Error() == apikey.ErrNotExists {
			if exists {
				return true, nil
			}
			return true, nil
		}
		return false, err
	}

	if ak.Status.AtProvider.Key == "" {
		return true, nil
	}
	return false, nil
}

func observeUpdateResource(ak *v1alpha1.APIKey, akm apikey.Metadata) bool {
	compare := updateStrategy(ak, akm)
	if !compare.DescriptionMatch {
		return true
	}

	if !compare.EnvironmentMatch {
		return true
	}

	if !compare.ResourceMatch {
		return true
	}

	if !compare.ServiceAccountMatch {
		return true
	}

	return false
}

func externalNameHelper(ak *v1alpha1.APIKey) (string, bool) {
	extName := meta.GetExternalName(ak)
	if extName != "" {
		return extName, true
	}
	return ak.Status.AtProvider.Key, false
}

func createResourceIsImport(err error) (bool, error) {
	if err != nil {
		if err.Error() == apikey.ErrNotExists {
			return false, err
		}
		return false, err
	}
	return true, err
}

func updateResourceDestructive(ak *v1alpha1.APIKey, akm apikey.Metadata) bool {
	compare := updateStrategy(ak, akm)

	return compare.isDestructive()
}
