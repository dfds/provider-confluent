/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package apis contains Kubernetes API for the Template provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	aclv1alpha1 "github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	apikeyv1alpha1 "github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	serviceaccountv1alpha1 "github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	topicv1alpha1 "github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	confluentv1alpha1 "github.com/dfds/provider-confluent/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		confluentv1alpha1.SchemeBuilder.AddToScheme,
		serviceaccountv1alpha1.SchemeBuilder.AddToScheme,
		apikeyv1alpha1.SchemeBuilder.AddToScheme,
		aclv1alpha1.SchemeBuilder.AddToScheme,
		topicv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
