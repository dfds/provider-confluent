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

package controller

import (
	"github.com/dfds/provider-confluent/internal/controller/acl"
	"github.com/dfds/provider-confluent/internal/controller/topic"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	"github.com/dfds/provider-confluent/internal/controller/apikey"
	"github.com/dfds/provider-confluent/internal/controller/config"
	"github.com/dfds/provider-confluent/internal/controller/serviceaccount"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter) error{
		config.Setup,
		serviceaccount.Setup,
		apikey.Setup,
		acl.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, l, wl); err != nil {
			return err
		}
	}
	return nil
}
