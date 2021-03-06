// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the discovery.smh.solo.io/v1alpha1 APIs
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

// clienset for the discovery.smh.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
	KubernetesClusters() KubernetesClusterClient
	// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
	MeshServices() MeshServiceClient
	// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
	MeshWorkloads() MeshWorkloadClient
	// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
	Meshes() MeshClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
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

// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) KubernetesClusters() KubernetesClusterClient {
	return NewKubernetesClusterClient(c.client)
}

// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) MeshServices() MeshServiceClient {
	return NewMeshServiceClient(c.client)
}

// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) MeshWorkloads() MeshWorkloadClient {
	return NewMeshWorkloadClient(c.client)
}

// clienset for the discovery.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) Meshes() MeshClient {
	return NewMeshClient(c.client)
}

// Reader knows how to read and list KubernetesClusters.
type KubernetesClusterReader interface {
	// Get retrieves a KubernetesCluster for the given object key
	GetKubernetesCluster(ctx context.Context, key client.ObjectKey) (*KubernetesCluster, error)

	// List retrieves list of KubernetesClusters for a given namespace and list options.
	ListKubernetesCluster(ctx context.Context, opts ...client.ListOption) (*KubernetesClusterList, error)
}

// KubernetesClusterTransitionFunction instructs the KubernetesClusterWriter how to transition between an existing
// KubernetesCluster object and a desired on an Upsert
type KubernetesClusterTransitionFunction func(existing, desired *KubernetesCluster) error

// Writer knows how to create, delete, and update KubernetesClusters.
type KubernetesClusterWriter interface {
	// Create saves the KubernetesCluster object.
	CreateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.CreateOption) error

	// Delete deletes the KubernetesCluster object.
	DeleteKubernetesCluster(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given KubernetesCluster object.
	UpdateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error

	// Patch patches the given KubernetesCluster object.
	PatchKubernetesCluster(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all KubernetesCluster objects matching the given options.
	DeleteAllOfKubernetesCluster(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the KubernetesCluster object.
	UpsertKubernetesCluster(ctx context.Context, obj *KubernetesCluster, transitionFuncs ...KubernetesClusterTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a KubernetesCluster object.
type KubernetesClusterStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given KubernetesCluster object.
	UpdateKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error

	// Patch patches the given KubernetesCluster object's subresource.
	PatchKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on KubernetesClusters.
type KubernetesClusterClient interface {
	KubernetesClusterReader
	KubernetesClusterWriter
	KubernetesClusterStatusWriter
}

type kubernetesClusterClient struct {
	client client.Client
}

func NewKubernetesClusterClient(client client.Client) *kubernetesClusterClient {
	return &kubernetesClusterClient{client: client}
}

func (c *kubernetesClusterClient) GetKubernetesCluster(ctx context.Context, key client.ObjectKey) (*KubernetesCluster, error) {
	obj := &KubernetesCluster{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *kubernetesClusterClient) ListKubernetesCluster(ctx context.Context, opts ...client.ListOption) (*KubernetesClusterList, error) {
	list := &KubernetesClusterList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *kubernetesClusterClient) CreateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) DeleteKubernetesCluster(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &KubernetesCluster{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) UpdateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) PatchKubernetesCluster(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *kubernetesClusterClient) DeleteAllOfKubernetesCluster(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &KubernetesCluster{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) UpsertKubernetesCluster(ctx context.Context, obj *KubernetesCluster, transitionFuncs ...KubernetesClusterTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*KubernetesCluster), desired.(*KubernetesCluster)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *kubernetesClusterClient) UpdateKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) PatchKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list MeshServices.
type MeshServiceReader interface {
	// Get retrieves a MeshService for the given object key
	GetMeshService(ctx context.Context, key client.ObjectKey) (*MeshService, error)

	// List retrieves list of MeshServices for a given namespace and list options.
	ListMeshService(ctx context.Context, opts ...client.ListOption) (*MeshServiceList, error)
}

// MeshServiceTransitionFunction instructs the MeshServiceWriter how to transition between an existing
// MeshService object and a desired on an Upsert
type MeshServiceTransitionFunction func(existing, desired *MeshService) error

// Writer knows how to create, delete, and update MeshServices.
type MeshServiceWriter interface {
	// Create saves the MeshService object.
	CreateMeshService(ctx context.Context, obj *MeshService, opts ...client.CreateOption) error

	// Delete deletes the MeshService object.
	DeleteMeshService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given MeshService object.
	UpdateMeshService(ctx context.Context, obj *MeshService, opts ...client.UpdateOption) error

	// Patch patches the given MeshService object.
	PatchMeshService(ctx context.Context, obj *MeshService, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all MeshService objects matching the given options.
	DeleteAllOfMeshService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the MeshService object.
	UpsertMeshService(ctx context.Context, obj *MeshService, transitionFuncs ...MeshServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a MeshService object.
type MeshServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given MeshService object.
	UpdateMeshServiceStatus(ctx context.Context, obj *MeshService, opts ...client.UpdateOption) error

	// Patch patches the given MeshService object's subresource.
	PatchMeshServiceStatus(ctx context.Context, obj *MeshService, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on MeshServices.
type MeshServiceClient interface {
	MeshServiceReader
	MeshServiceWriter
	MeshServiceStatusWriter
}

type meshServiceClient struct {
	client client.Client
}

func NewMeshServiceClient(client client.Client) *meshServiceClient {
	return &meshServiceClient{client: client}
}

func (c *meshServiceClient) GetMeshService(ctx context.Context, key client.ObjectKey) (*MeshService, error) {
	obj := &MeshService{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *meshServiceClient) ListMeshService(ctx context.Context, opts ...client.ListOption) (*MeshServiceList, error) {
	list := &MeshServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *meshServiceClient) CreateMeshService(ctx context.Context, obj *MeshService, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *meshServiceClient) DeleteMeshService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &MeshService{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *meshServiceClient) UpdateMeshService(ctx context.Context, obj *MeshService, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *meshServiceClient) PatchMeshService(ctx context.Context, obj *MeshService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *meshServiceClient) DeleteAllOfMeshService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &MeshService{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *meshServiceClient) UpsertMeshService(ctx context.Context, obj *MeshService, transitionFuncs ...MeshServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*MeshService), desired.(*MeshService)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *meshServiceClient) UpdateMeshServiceStatus(ctx context.Context, obj *MeshService, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *meshServiceClient) PatchMeshServiceStatus(ctx context.Context, obj *MeshService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list MeshWorkloads.
type MeshWorkloadReader interface {
	// Get retrieves a MeshWorkload for the given object key
	GetMeshWorkload(ctx context.Context, key client.ObjectKey) (*MeshWorkload, error)

	// List retrieves list of MeshWorkloads for a given namespace and list options.
	ListMeshWorkload(ctx context.Context, opts ...client.ListOption) (*MeshWorkloadList, error)
}

// MeshWorkloadTransitionFunction instructs the MeshWorkloadWriter how to transition between an existing
// MeshWorkload object and a desired on an Upsert
type MeshWorkloadTransitionFunction func(existing, desired *MeshWorkload) error

// Writer knows how to create, delete, and update MeshWorkloads.
type MeshWorkloadWriter interface {
	// Create saves the MeshWorkload object.
	CreateMeshWorkload(ctx context.Context, obj *MeshWorkload, opts ...client.CreateOption) error

	// Delete deletes the MeshWorkload object.
	DeleteMeshWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given MeshWorkload object.
	UpdateMeshWorkload(ctx context.Context, obj *MeshWorkload, opts ...client.UpdateOption) error

	// Patch patches the given MeshWorkload object.
	PatchMeshWorkload(ctx context.Context, obj *MeshWorkload, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all MeshWorkload objects matching the given options.
	DeleteAllOfMeshWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the MeshWorkload object.
	UpsertMeshWorkload(ctx context.Context, obj *MeshWorkload, transitionFuncs ...MeshWorkloadTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a MeshWorkload object.
type MeshWorkloadStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given MeshWorkload object.
	UpdateMeshWorkloadStatus(ctx context.Context, obj *MeshWorkload, opts ...client.UpdateOption) error

	// Patch patches the given MeshWorkload object's subresource.
	PatchMeshWorkloadStatus(ctx context.Context, obj *MeshWorkload, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on MeshWorkloads.
type MeshWorkloadClient interface {
	MeshWorkloadReader
	MeshWorkloadWriter
	MeshWorkloadStatusWriter
}

type meshWorkloadClient struct {
	client client.Client
}

func NewMeshWorkloadClient(client client.Client) *meshWorkloadClient {
	return &meshWorkloadClient{client: client}
}

func (c *meshWorkloadClient) GetMeshWorkload(ctx context.Context, key client.ObjectKey) (*MeshWorkload, error) {
	obj := &MeshWorkload{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *meshWorkloadClient) ListMeshWorkload(ctx context.Context, opts ...client.ListOption) (*MeshWorkloadList, error) {
	list := &MeshWorkloadList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *meshWorkloadClient) CreateMeshWorkload(ctx context.Context, obj *MeshWorkload, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *meshWorkloadClient) DeleteMeshWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &MeshWorkload{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *meshWorkloadClient) UpdateMeshWorkload(ctx context.Context, obj *MeshWorkload, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *meshWorkloadClient) PatchMeshWorkload(ctx context.Context, obj *MeshWorkload, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *meshWorkloadClient) DeleteAllOfMeshWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &MeshWorkload{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *meshWorkloadClient) UpsertMeshWorkload(ctx context.Context, obj *MeshWorkload, transitionFuncs ...MeshWorkloadTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*MeshWorkload), desired.(*MeshWorkload)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *meshWorkloadClient) UpdateMeshWorkloadStatus(ctx context.Context, obj *MeshWorkload, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *meshWorkloadClient) PatchMeshWorkloadStatus(ctx context.Context, obj *MeshWorkload, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Meshs.
type MeshReader interface {
	// Get retrieves a Mesh for the given object key
	GetMesh(ctx context.Context, key client.ObjectKey) (*Mesh, error)

	// List retrieves list of Meshs for a given namespace and list options.
	ListMesh(ctx context.Context, opts ...client.ListOption) (*MeshList, error)
}

// MeshTransitionFunction instructs the MeshWriter how to transition between an existing
// Mesh object and a desired on an Upsert
type MeshTransitionFunction func(existing, desired *Mesh) error

// Writer knows how to create, delete, and update Meshs.
type MeshWriter interface {
	// Create saves the Mesh object.
	CreateMesh(ctx context.Context, obj *Mesh, opts ...client.CreateOption) error

	// Delete deletes the Mesh object.
	DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Mesh object.
	UpdateMesh(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object.
	PatchMesh(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Mesh objects matching the given options.
	DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Mesh object.
	UpsertMesh(ctx context.Context, obj *Mesh, transitionFuncs ...MeshTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Mesh object.
type MeshStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Mesh object.
	UpdateMeshStatus(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object's subresource.
	PatchMeshStatus(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Meshs.
type MeshClient interface {
	MeshReader
	MeshWriter
	MeshStatusWriter
}

type meshClient struct {
	client client.Client
}

func NewMeshClient(client client.Client) *meshClient {
	return &meshClient{client: client}
}

func (c *meshClient) GetMesh(ctx context.Context, key client.ObjectKey) (*Mesh, error) {
	obj := &Mesh{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *meshClient) ListMesh(ctx context.Context, opts ...client.ListOption) (*MeshList, error) {
	list := &MeshList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *meshClient) CreateMesh(ctx context.Context, obj *Mesh, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *meshClient) DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Mesh{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *meshClient) UpdateMesh(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMesh(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *meshClient) DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Mesh{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *meshClient) UpsertMesh(ctx context.Context, obj *Mesh, transitionFuncs ...MeshTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Mesh), desired.(*Mesh)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *meshClient) UpdateMeshStatus(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMeshStatus(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
