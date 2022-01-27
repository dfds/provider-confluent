package serviceaccount

import (
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
)

func ObserveCreateResource(sa *v1alpha1.ServiceAccount, err error) (bool, error) {
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return true, nil
		} else {
			return true, err
		}
	}

	// Check status
	if sa.Status.AtProvider.Id == "" {
		return true, nil
	}

	return false, nil
}

func ObserveUpdateResource(sa *v1alpha1.ServiceAccount, sac serviceaccount.ServiceAccount) bool {
	// Diff
	return sa.Spec.ForProvider.Description != sac.Description
}

func ExternalNameHelper(sa *v1alpha1.ServiceAccount) (string, bool) {
	extName := meta.GetExternalName(sa)
	if extName != "" {
		return extName, true
	}
	return sa.Name, false
}

func CreateResourceIsImport(err error) (bool, error) {
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, err
}
