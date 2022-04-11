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

package apikey

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/dfds/provider-confluent/apis/apikey/v1alpha1"
	apisv1alpha1 "github.com/dfds/provider-confluent/apis/v1alpha1"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/apikey"
	"github.com/dfds/provider-confluent/internal/clients/serviceaccount"
)

const (
	errNotMyType                                 = "managed resource is not a APIKey custom resource"
	errTrackPCUsage                              = "cannot track ProviderConfig usage"
	errGetPC                                     = "cannot get ProviderConfig"
	errGetCreds                                  = "cannot get credentials"
	errNewClient                                 = "cannot create new Service"
	errAuthCredentials                           = "invalid client credentials"
	errBlockingCreationServiceAccountDoNotExists = "creation blocked service-account referenced do not exists"
	errExternalNameNotPresent                    = "external name is not present"
	errDestructiveUpdateNotAllowed               = "cannot update resource. An immutable field has been changed, underlying API doesn't support updating some fields"
)

var (
	createAndConvertClientFunc = func(clientCreds []byte) (interface{}, interface{}, error) { //nolint
		credParts := strings.Split(string(clientCreds), ":")

		if len(credParts) != 2 {
			return nil, nil, errors.New(errAuthCredentials)
		}

		cClient := clients.NewClient()
		authErr := cClient.Authenticate(credParts[0], credParts[1])

		if authErr != nil {
			return nil, nil, authErr
		}

		return apikey.NewClient().(interface{}), serviceaccount.NewClient().(interface{}), nil
	}
)

// Setup adds a controller that reconciles ServiceAccount managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.APIKeyGroupKind)

	o := controller.Options{
		RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.APIKeyGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			usage:        resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFn: createAndConvertClientFunc}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithInitializers(),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&v1alpha1.APIKey{}).
		Complete(r)
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(creds []byte) (interface{}, interface{}, error)
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.APIKey)
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

	svc, saSvc, err := c.newServiceFn(clientCredentialData)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &external{service: svc, saService: saSvc, kube: c.kube}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	// A 'client' used to connect to the external resource API. In practice this
	// would be something like an AWS SDK client.
	service   interface{}
	saService interface{}
	kube      client.Client
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.APIKey)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotMyType)
	}

	// Support for importing resource using external name
	key, exists := externalNameHelper(cr)

	// Confluent cloud
	var client = c.service.(apikey.IClient)
	observe, err := client.GetAPIKeyByKey(key)

	// Check if resource require creation
	create, err := observeCreateResource(cr, exists, err)
	if err != nil {
		return managed.ExternalObservation{
			ResourceExists:    false,
			ConnectionDetails: managed.ConnectionDetails{},
		}, err
	}

	fmt.Println("Observe if should create:", create)

	if create {
		return managed.ExternalObservation{
			ResourceExists:    false,
			ConnectionDetails: managed.ConnectionDetails{},
		}, nil
	}

	// Check if resource require update
	if observeUpdateResource(cr, observe) {
		return managed.ExternalObservation{
			ResourceExists:    true,
			ResourceUpToDate:  false,
			ConnectionDetails: managed.ConnectionDetails{},
		}, nil
	}

	cr.Status.SetConditions(xpv1.Available())
	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalObservation{}, err
	}

	return managed.ExternalObservation{
		ResourceExists:    true,
		ResourceUpToDate:  true,
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.APIKey)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotMyType)
	}

	cr.Status.SetConditions(xpv1.Creating())
	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalCreation{}, err
	}

	// Need to check if service account is valid otherwise it will return key pair with God like access (bug stems from confluent cli)
	var saClient = c.saService.(serviceaccount.IClient)
	_, err := saClient.ServiceAccountByID(cr.Spec.ForProvider.ServiceAccount)
	if err != nil {
		if err.Error() == serviceaccount.ErrNotExists {
			return managed.ExternalCreation{}, errors.New(errBlockingCreationServiceAccountDoNotExists)
		}
		return managed.ExternalCreation{}, err
	}

	key, exists := externalNameHelper(cr)

	var createIsImport bool

	conn := managed.ConnectionDetails{}

	var client = c.service.(apikey.IClient)

	if exists {
		observe, err := client.GetAPIKeyByKey(key)
		createIsImport, err = createResourceIsImport(err)
		if err != nil {
			return managed.ExternalCreation{}, err
		}
		if createIsImport {
			cr.Status.AtProvider.Key = observe.Key
			cr.Status.AtProvider.Environment = cr.Spec.ForProvider.Environment
			cr.Status.AtProvider.Resource = cr.Spec.ForProvider.Resource
			cr.Status.AtProvider.ServiceAccount = cr.Spec.ForProvider.ServiceAccount
			conn = managed.ConnectionDetails{xpv1.ResourceCredentialsSecretUserKey: []byte(observe.Key),
				xpv1.ResourceCredentialsSecretPasswordKey: []byte("YOU NEED TO SUPPLY YOUR OWN SECRET FOR IMPORTED RESOURCES"),
			}
		}
	}

	fmt.Println("createIsImport: ", createIsImport)
	if !createIsImport {
		out, err := client.APIKeyCreate(cr.Spec.ForProvider.Resource, cr.Spec.ForProvider.Description, cr.Spec.ForProvider.ServiceAccount, cr.Spec.ForProvider.Environment)
		if err != nil {
			return managed.ExternalCreation{}, err
		}
		meta.SetExternalName(cr, out.Key)
		if err := c.kube.Update(ctx, cr); err != nil {
			return managed.ExternalCreation{}, err
		}

		cr.Status.AtProvider.Key = out.Key
		cr.Status.AtProvider.Environment = cr.Spec.ForProvider.Environment
		cr.Status.AtProvider.Resource = cr.Spec.ForProvider.Resource
		cr.Status.AtProvider.ServiceAccount = cr.Spec.ForProvider.ServiceAccount
		conn = managed.ConnectionDetails{
			xpv1.ResourceCredentialsSecretUserKey:     []byte(out.Key),
			xpv1.ResourceCredentialsSecretPasswordKey: []byte(out.Secret),
		}
	}

	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalCreation{}, err
	}

	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: conn,
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.APIKey)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotMyType)
	}

	// Use external name since we set in create
	key, exists := externalNameHelper(cr)
	if !exists {
		return managed.ExternalUpdate{}, errors.New(errExternalNameNotPresent)
	}

	// Confluent cloud
	var client = c.service.(apikey.IClient)
	observed, err := client.GetAPIKeyByKey(key)
	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	// Is update destructive
	if updateResourceDestructive(cr, observed) {
		return managed.ExternalUpdate{}, errors.New(errDestructiveUpdateNotAllowed)
	}
	// Continue with non-destructive action
	err = client.APIKeyUpdate(key, cr.Spec.ForProvider.Description)
	if err != nil {
		return managed.ExternalUpdate{}, err
	}
	return managed.ExternalUpdate{}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.APIKey)
	if !ok {
		return errors.New(errNotMyType)
	}

	cr.Status.SetConditions(xpv1.Deleting())
	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return err
	}

	var client = c.service.(apikey.IClient)

	err := client.APIKeyDelete(cr.Status.AtProvider.Key)
	if err != nil {
		return err
	}

	return nil
}
