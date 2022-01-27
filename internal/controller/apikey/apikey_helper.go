package apikey

import (
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/apikey"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
)

func observeCreateResource(err error) (bool, error) {
	if err != nil {
		if err.Error() == apikey.ErrNotExists {
			return true, nil
		} else {
			return true, err
		}
	}

	return false, nil
}

func observeUpdateResource(ak *v1alpha1.ApiKey, akm apikey.ApiKeyMetadata) bool {
	compare := updateStrategy(ak, akm)

	return diffDetect(compare)
}

func externalNameHelper(ak *v1alpha1.ApiKey) (string, bool) {
	extName := meta.GetExternalName(ak)
	if extName != "" {
		return extName, true
	}
	return ak.Status.AtProvider.Key, false
}

func createResourceIsImport(err error) (bool, error) {
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, err
}

func updateResourceDestructive(ak *v1alpha1.ApiKey, akm apikey.ApiKeyMetadata) bool {
	compare := updateStrategy(ak, akm)

	return compare.IsDestructive()
}

func diffDetect(compare ApiKeyCompare) bool {
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
