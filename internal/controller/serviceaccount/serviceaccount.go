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

package serviceaccount

import (
	"context"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	apisv1alpha1 "github.com/dfds/provider-confluent/apis/v1alpha1"

	"github.com/dfds/provider-confluent/internal/clients"
	confluentClient "github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
	serviceaccountClient "github.com/dfds/provider-confluent/internal/clients/serviceaccount"
)

const (
	errNotMyType       = "managed resource is not a Schema custom resource"
	errTrackPCUsage    = "cannot track ProviderConfig usage"
	errGetPC           = "cannot get ProviderConfig"
	errGetCreds        = "cannot get credentials"
	errNewClient       = "cannot create new Service"
	errAuthCredentials = "invalid client credentials"
	errUnmarshalState  = "kubernetes state mismatch with type"
)

var (
	createAndConvertClientFunc = func(clientCreds []byte, apiCreds clients.APICredentials) (interface{}, error) { //nolint
		credParts := strings.Split(string(clientCreds), ":")

		if len(credParts) != 2 {
			return nil, errors.New(errAuthCredentials)
		}

		cClient := confluentClient.NewClient()
		authErr := cClient.Authenticate(credParts[0], credParts[1])

		if authErr != nil {
			return nil, authErr
		}

		srConfig := serviceaccountClient.Config{
			APICredentials: apiCreds,
		}

		return serviceaccountClient.NewClient(srConfig).(interface{}), nil
	}
)

// Setup adds a controller that reconciles ServiceAccount managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.ServiceAccountGroupKind)

	o := controller.Options{
		RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.ServiceAccountGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			usage:        resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFn: createAndConvertClientFunc}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&v1alpha1.ServiceAccount{}).
		Complete(r)
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(creds []byte, apiCreds confluentClient.APICredentials) (interface{}, error)
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.ServiceAccount)
	if !ok {
		return nil, errors.New(errNotMyType)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	pc := &apisv1alpha1.ProviderConfig{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetPC)
	}

	clientCredentialData, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c.kube, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errGetCreds)
	}

	var apiCredentials confluentClient.APICredentials

	for _, value := range pc.Spec.APICredentials {
		if value.Identifier == v1alpha1.SchemeGroupVersion.Identifier() {
			apiCredentials = value

			break
		}
	}

	svc, err := c.newServiceFn(clientCredentialData, apiCredentials)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &external{service: svc, kube: c.kube}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	// A 'client' used to connect to the external resource API. In practice this
	// would be something like an AWS SDK client.
	service interface{}
	kube    client.Client
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.ServiceAccount)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotMyType)
	}

	// Confluent
	var client = c.service.(serviceaccount.IClient)
	ccsa, err := client.ServiceAccountList(cr.Spec.ForProvider.Name)

	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return managed.ExternalObservation{
				ResourceExists:    false,
				ConnectionDetails: managed.ConnectionDetails{},
			}, nil // returning nil because we want create on not found
		} else {
			return managed.ExternalObservation{
				ResourceExists:    false,
				ConnectionDetails: managed.ConnectionDetails{},
			}, err
		}
	}

	// Diff
	if cr.Spec.ForProvider.Description != ccsa.Description {
		return managed.ExternalObservation{
			ResourceExists:    true,
			ResourceUpToDate:  false,
			ConnectionDetails: managed.ConnectionDetails{},
		}, nil
	}

	return managed.ExternalObservation{
		ResourceExists:    true,
		ResourceUpToDate:  true,
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.ServiceAccount)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotMyType)
	}

	var client = c.service.(serviceaccount.IClient)
	out, err := client.ServiceAccountCreate(cr.Spec.ForProvider.Name, cr.Spec.ForProvider.Description)

	if err != nil {
		return managed.ExternalCreation{}, err
	}

	current := cr.Status.DeepCopy()
	cr.Status.AtProvider.Id = out.Id

	if !cmp.Equal(current, &cr.Status) {
		err = c.kube.Update(ctx, cr)
		if err != nil {
			return managed.ExternalCreation{
				ConnectionDetails: managed.ConnectionDetails{},
			}, err
		}
	}

	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.ServiceAccount)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotMyType)
	}

	var client = c.service.(serviceaccount.IClient)

	err := client.ServiceAccountUpdate(cr.Status.AtProvider.Id, cr.Spec.ForProvider.Description)

	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.ServiceAccount)
	if !ok {
		return errors.New(errNotMyType)
	}

	var client = c.service.(serviceaccount.IClient)

	err := client.ServiceAccountDelete(cr.Status.AtProvider.Id)
	if err != nil {
		return err
	}

	return nil
}
