package serviceaccount

import (
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
)

// ObserveCreateResource Checks if a ServiceAccount should be created
func ObserveCreateResource(sa *v1alpha1.ServiceAccount, err error) (bool, error) {
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return true, nil
		}

		return true, err
	}

	// Check status
	if sa.Status.AtProvider.ID == "" {
		return true, nil
	}

	return false, nil
}

// ObserveUpdateResource Checks if a ServiceAccount should be updated
func ObserveUpdateResource(sa *v1alpha1.ServiceAccount, sac serviceaccount.ServiceAccount) bool {
	// Diff
	return sa.Spec.ForProvider.Description != sac.Description
}

// ExternalNameHelper Checks if a ServiceAccount k8s object has an external-name attached. If it does, return that external-name, if it doesn't, return the name of the k8s object
func ExternalNameHelper(sa *v1alpha1.ServiceAccount) (string, bool) {
	extName := meta.GetExternalName(sa)
	if extName != "" {
		return extName, true
	}
	return sa.Name, false
}

// CreateResourceIsImport Checks if a ServiceAccount k8s object is considered an import
func CreateResourceIsImport(err error) (bool, error) {
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return false, nil
		}

		return false, err
	}
	return true, err
}
