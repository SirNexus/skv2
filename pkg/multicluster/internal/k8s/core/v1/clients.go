// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the /v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the /v1 APIs
type Clientset interface {
	// clienset for the v1/v1 APIs
	Secrets() SecretClient
	// clienset for the v1/v1 APIs
	ServiceAccounts() ServiceAccountClient
	// clienset for the v1/v1 APIs
	Namespaces() NamespaceClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := v1.SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the v1/v1 APIs
func (c *clientSet) Secrets() SecretClient {
	return NewSecretClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) ServiceAccounts() ServiceAccountClient {
	return NewServiceAccountClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) Namespaces() NamespaceClient {
	return NewNamespaceClient(c.client)
}

// Reader knows how to read and list Secrets.
type SecretReader interface {
	// Get retrieves a Secret for the given object key
	GetSecret(ctx context.Context, key client.ObjectKey) (*v1.Secret, error)

	// List retrieves list of Secrets for a given namespace and list options.
	ListSecret(ctx context.Context, opts ...client.ListOption) (*v1.SecretList, error)
}

// SecretTransitionFunction instructs the SecretWriter how to transition between an existing
// Secret object and a desired on an Upsert
type SecretTransitionFunction func(existing, desired *v1.Secret) error

// Writer knows how to create, delete, and update Secrets.
type SecretWriter interface {
	// Create saves the Secret object.
	CreateSecret(ctx context.Context, obj *v1.Secret, opts ...client.CreateOption) error

	// Delete deletes the Secret object.
	DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Secret object.
	UpdateSecret(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object.
	PatchSecret(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Secret objects matching the given options.
	DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Secret object.
	UpsertSecret(ctx context.Context, obj *v1.Secret, transitionFuncs ...SecretTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Secret object.
type SecretStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Secret object.
	UpdateSecretStatus(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object's subresource.
	PatchSecretStatus(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Secrets.
type SecretClient interface {
	SecretReader
	SecretWriter
	SecretStatusWriter
}

type secretClient struct {
	client client.Client
}

func NewSecretClient(client client.Client) *secretClient {
	return &secretClient{client: client}
}

func (c *secretClient) GetSecret(ctx context.Context, key client.ObjectKey) (*v1.Secret, error) {
	obj := &v1.Secret{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *secretClient) ListSecret(ctx context.Context, opts ...client.ListOption) (*v1.SecretList, error) {
	list := &v1.SecretList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *secretClient) CreateSecret(ctx context.Context, obj *v1.Secret, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *secretClient) DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &v1.Secret{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *secretClient) UpdateSecret(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecret(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *secretClient) DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &v1.Secret{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *secretClient) UpsertSecret(ctx context.Context, obj *v1.Secret, transitionFuncs ...SecretTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Secret), desired.(*v1.Secret)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *secretClient) UpdateSecretStatus(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecretStatus(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides SecretClients for multiple clusters.
type MulticlusterSecretClient interface {
	// Cluster returns a SecretClient for the given cluster
	Cluster(cluster string) (SecretClient, error)
}

type multiclusterSecretClient struct {
	client multicluster.Client
}

func NewMulticlusterSecretClient(client multicluster.Client) MulticlusterSecretClient {
	return &multiclusterSecretClient{client: client}
}

func (m *multiclusterSecretClient) Cluster(cluster string) (SecretClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewSecretClient(client), nil
}

// Reader knows how to read and list ServiceAccounts.
type ServiceAccountReader interface {
	// Get retrieves a ServiceAccount for the given object key
	GetServiceAccount(ctx context.Context, key client.ObjectKey) (*v1.ServiceAccount, error)

	// List retrieves list of ServiceAccounts for a given namespace and list options.
	ListServiceAccount(ctx context.Context, opts ...client.ListOption) (*v1.ServiceAccountList, error)
}

// ServiceAccountTransitionFunction instructs the ServiceAccountWriter how to transition between an existing
// ServiceAccount object and a desired on an Upsert
type ServiceAccountTransitionFunction func(existing, desired *v1.ServiceAccount) error

// Writer knows how to create, delete, and update ServiceAccounts.
type ServiceAccountWriter interface {
	// Create saves the ServiceAccount object.
	CreateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.CreateOption) error

	// Delete deletes the ServiceAccount object.
	DeleteServiceAccount(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ServiceAccount object.
	UpdateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error

	// Patch patches the given ServiceAccount object.
	PatchServiceAccount(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ServiceAccount objects matching the given options.
	DeleteAllOfServiceAccount(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ServiceAccount object.
	UpsertServiceAccount(ctx context.Context, obj *v1.ServiceAccount, transitionFuncs ...ServiceAccountTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ServiceAccount object.
type ServiceAccountStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ServiceAccount object.
	UpdateServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error

	// Patch patches the given ServiceAccount object's subresource.
	PatchServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ServiceAccounts.
type ServiceAccountClient interface {
	ServiceAccountReader
	ServiceAccountWriter
	ServiceAccountStatusWriter
}

type serviceAccountClient struct {
	client client.Client
}

func NewServiceAccountClient(client client.Client) *serviceAccountClient {
	return &serviceAccountClient{client: client}
}

func (c *serviceAccountClient) GetServiceAccount(ctx context.Context, key client.ObjectKey) (*v1.ServiceAccount, error) {
	obj := &v1.ServiceAccount{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *serviceAccountClient) ListServiceAccount(ctx context.Context, opts ...client.ListOption) (*v1.ServiceAccountList, error) {
	list := &v1.ServiceAccountList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *serviceAccountClient) CreateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *serviceAccountClient) DeleteServiceAccount(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &v1.ServiceAccount{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *serviceAccountClient) UpdateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *serviceAccountClient) PatchServiceAccount(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *serviceAccountClient) DeleteAllOfServiceAccount(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &v1.ServiceAccount{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *serviceAccountClient) UpsertServiceAccount(ctx context.Context, obj *v1.ServiceAccount, transitionFuncs ...ServiceAccountTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.ServiceAccount), desired.(*v1.ServiceAccount)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *serviceAccountClient) UpdateServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *serviceAccountClient) PatchServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ServiceAccountClients for multiple clusters.
type MulticlusterServiceAccountClient interface {
	// Cluster returns a ServiceAccountClient for the given cluster
	Cluster(cluster string) (ServiceAccountClient, error)
}

type multiclusterServiceAccountClient struct {
	client multicluster.Client
}

func NewMulticlusterServiceAccountClient(client multicluster.Client) MulticlusterServiceAccountClient {
	return &multiclusterServiceAccountClient{client: client}
}

func (m *multiclusterServiceAccountClient) Cluster(cluster string) (ServiceAccountClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewServiceAccountClient(client), nil
}

// Reader knows how to read and list Namespaces.
type NamespaceReader interface {
	// Get retrieves a Namespace for the given object key
	GetNamespace(ctx context.Context, name string) (*v1.Namespace, error)

	// List retrieves list of Namespaces for a given namespace and list options.
	ListNamespace(ctx context.Context, opts ...client.ListOption) (*v1.NamespaceList, error)
}

// NamespaceTransitionFunction instructs the NamespaceWriter how to transition between an existing
// Namespace object and a desired on an Upsert
type NamespaceTransitionFunction func(existing, desired *v1.Namespace) error

// Writer knows how to create, delete, and update Namespaces.
type NamespaceWriter interface {
	// Create saves the Namespace object.
	CreateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.CreateOption) error

	// Delete deletes the Namespace object.
	DeleteNamespace(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given Namespace object.
	UpdateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error

	// Patch patches the given Namespace object.
	PatchNamespace(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Namespace objects matching the given options.
	DeleteAllOfNamespace(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Namespace object.
	UpsertNamespace(ctx context.Context, obj *v1.Namespace, transitionFuncs ...NamespaceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Namespace object.
type NamespaceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Namespace object.
	UpdateNamespaceStatus(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error

	// Patch patches the given Namespace object's subresource.
	PatchNamespaceStatus(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Namespaces.
type NamespaceClient interface {
	NamespaceReader
	NamespaceWriter
	NamespaceStatusWriter
}

type namespaceClient struct {
	client client.Client
}

func NewNamespaceClient(client client.Client) *namespaceClient {
	return &namespaceClient{client: client}
}

func (c *namespaceClient) GetNamespace(ctx context.Context, name string) (*v1.Namespace, error) {
	obj := &v1.Namespace{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *namespaceClient) ListNamespace(ctx context.Context, opts ...client.ListOption) (*v1.NamespaceList, error) {
	list := &v1.NamespaceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *namespaceClient) CreateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *namespaceClient) DeleteNamespace(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &v1.Namespace{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *namespaceClient) UpdateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *namespaceClient) PatchNamespace(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *namespaceClient) DeleteAllOfNamespace(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &v1.Namespace{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *namespaceClient) UpsertNamespace(ctx context.Context, obj *v1.Namespace, transitionFuncs ...NamespaceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Namespace), desired.(*v1.Namespace)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *namespaceClient) UpdateNamespaceStatus(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *namespaceClient) PatchNamespaceStatus(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides NamespaceClients for multiple clusters.
type MulticlusterNamespaceClient interface {
	// Cluster returns a NamespaceClient for the given cluster
	Cluster(cluster string) (NamespaceClient, error)
}

type multiclusterNamespaceClient struct {
	client multicluster.Client
}

func NewMulticlusterNamespaceClient(client multicluster.Client) MulticlusterNamespaceClient {
	return &multiclusterNamespaceClient{client: client}
}

func (m *multiclusterNamespaceClient) Cluster(cluster string) (NamespaceClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewNamespaceClient(client), nil
}
