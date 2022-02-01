package apikey

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/apikey"
)

type ApiKeyCompare struct {
	DescriptionMatch    bool
	EnvironmentMatch    bool
	ResourceMatch       bool
	ServiceAccountMatch bool
}

func updateStrategy(ak *v1alpha1.ApiKey, akm apikey.APIKeyMetadata) ApiKeyCompare {
	var compare ApiKeyCompare
	if ak.Spec.ForProvider.Description == akm.Description {
		compare.DescriptionMatch = true
	}

	if ak.Spec.ForProvider.Environment == ak.Status.AtProvider.Environment {
		compare.EnvironmentMatch = true
	}

	if ak.Spec.ForProvider.Resource == ak.Status.AtProvider.Resource {
		compare.ResourceMatch = true
	}

	if ak.Spec.ForProvider.ServiceAccount == ak.Status.AtProvider.ServiceAccount {
		compare.ServiceAccountMatch = true
	}

	return compare
}

func (ac *ApiKeyCompare) isDestructive() bool {
	var destructive bool

	if !ac.EnvironmentMatch {
		destructive = true
	}

	if !ac.ResourceMatch {
		destructive = true
	}

	if !ac.ServiceAccountMatch {
		destructive = true
	}

	return destructive
}

func destructiveActionsAllowed(dp v1.DeletionPolicy) bool {
	return dp == "Delete"
}
