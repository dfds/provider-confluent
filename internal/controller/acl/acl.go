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

package acl

import (
	"context"
	"reflect"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
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
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	apisv1alpha1 "github.com/dfds/provider-confluent/apis/v1alpha1"

	"github.com/dfds/provider-confluent/internal/clients"
	confluentClient "github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/acl"
)

const (
	errNotMyType                      = "managed resource is not an ACL custom resource"
	errTrackPCUsage                   = "cannot track ProviderConfig usage"
	errGetPC                          = "cannot get ProviderConfig"
	errGetCreds                       = "cannot get credentials"
	errNewClient                      = "cannot create new Service"
	errAuthCredentials                = "invalid client credentials"
	errACLRuleInputDoesNotMatchOutput = "A single rule was not returned after creation. As only one rule is supposed to be created, this ain't right son."
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

		srConfig := acl.Config{
			APICredentials: apiCreds,
		}

		return acl.NewClient(srConfig).(interface{}), nil
	}
)

// Setup adds a controller that reconciles ServiceAccount managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.ACLGroupKind)

	o := controller.Options{
		RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.ACLGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			usage:        resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFn: createAndConvertClientFunc}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&v1alpha1.ACL{}).
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
	cr, ok := mg.(*v1alpha1.ACL)
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
	cr, ok := mg.(*v1alpha1.ACL)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotMyType)
	}

	// Confluent
	var client = c.service.(acl.IClient)
	aclResp, err := client.ACLList(cr.Status.AtProvider.ACLP.ACLRule.Principal, cr.Status.AtProvider.ACLP.Environment, cr.Status.AtProvider.ACLP.Cluster)

	if err != nil {
		if err.Error() == acl.ErrACLNotExistsOrInvalidServiceAccount {
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
	ruleStatusMatched := false
	ruleSpecMatched := false
	for _, rule := range aclResp {
		if ruleStatusMatched && ruleSpecMatched {
			break
		}

		if reflect.DeepEqual(rule, cr.Status.AtProvider.ACLP.ACLRule) {
			ruleStatusMatched = true
		}

		if reflect.DeepEqual(rule, cr.Spec.ForProvider.ACLRule) {
			ruleSpecMatched = true
		}
	}

	// Rule stored in Status matched, but rule in Spec doesn't. Delete rule specified in Status & create a new rule based from Spec. Update Status with rule from Spec.
	if ruleStatusMatched && !ruleSpecMatched {
		return managed.ExternalObservation{
			ResourceExists:    true,
			ResourceUpToDate:  false,
			ConnectionDetails: managed.ConnectionDetails{},
		}, nil
	}

	// Rule stored in Status & Spec didn't match. Create rule from Spec, update Status with rule from Spec.
	if !ruleStatusMatched && !ruleSpecMatched {
		return managed.ExternalObservation{
			ResourceExists:    false,
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
	cr, ok := mg.(*v1alpha1.ACL)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotMyType)
	}

	cr.Status.SetConditions(xpv1.Creating())
	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalCreation{}, err
	}

	var client = c.service.(acl.IClient)
	out, err := client.ACLCreate(cr.Spec.ForProvider)

	if err != nil {
		return managed.ExternalCreation{}, err
	}

	if len(out) != 1 {
		return managed.ExternalCreation{}, errors.New(errACLRuleInputDoesNotMatchOutput)
	}

	cr.Status.AtProvider.ACLP.ACLRule = out[0]
	cr.Status.AtProvider.ACLP.Cluster = cr.Spec.ForProvider.Cluster
	cr.Status.AtProvider.ACLP.Environment = cr.Spec.ForProvider.Environment

	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalCreation{}, err
	}

	conn := managed.ConnectionDetails{}

	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: conn,
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.ACL)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotMyType)
	}

	var client = c.service.(acl.IClient)

	// Update description
	err := client.ACLDelete(cr.Status.AtProvider.ACLP)
	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	out, err := client.ACLCreate(cr.Spec.ForProvider)
	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	if len(out) != 1 {
		return managed.ExternalUpdate{}, errors.New(errACLRuleInputDoesNotMatchOutput)
	}

	cr.Status.AtProvider.ACLP.ACLRule = out[0]
	cr.Status.AtProvider.ACLP.Cluster = cr.Spec.ForProvider.Cluster
	cr.Status.AtProvider.ACLP.Environment = cr.Spec.ForProvider.Environment

	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalUpdate{}, err
	}

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.ACL)
	if !ok {
		return errors.New(errNotMyType)
	}

	cr.Status.SetConditions(xpv1.Deleting())
	if err := c.kube.Status().Update(ctx, cr); err != nil {
		return err
	}

	var client = c.service.(acl.IClient)

	err := client.ACLDelete(cr.Spec.ForProvider)
	if err != nil {
		return err
	}

	return nil
}
