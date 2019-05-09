// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewMeshGroup(namespace, name string) *MeshGroup {
	meshgroup := &MeshGroup{}
	meshgroup.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return meshgroup
}

func (r *MeshGroup) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *MeshGroup) SetStatus(status core.Status) {
	r.Status = status
}

func (r *MeshGroup) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Meshes,
	)
}

type MeshGroupList []*MeshGroup

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list MeshGroupList) Find(namespace, name string) (*MeshGroup, error) {
	for _, meshGroup := range list {
		if meshGroup.GetMetadata().Name == name {
			if namespace == "" || meshGroup.GetMetadata().Namespace == namespace {
				return meshGroup, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find meshGroup %v.%v", namespace, name)
}

func (list MeshGroupList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, meshGroup := range list {
		ress = append(ress, meshGroup)
	}
	return ress
}

func (list MeshGroupList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, meshGroup := range list {
		ress = append(ress, meshGroup)
	}
	return ress
}

func (list MeshGroupList) Names() []string {
	var names []string
	for _, meshGroup := range list {
		names = append(names, meshGroup.GetMetadata().Name)
	}
	return names
}

func (list MeshGroupList) NamespacesDotNames() []string {
	var names []string
	for _, meshGroup := range list {
		names = append(names, meshGroup.GetMetadata().Namespace+"."+meshGroup.GetMetadata().Name)
	}
	return names
}

func (list MeshGroupList) Sort() MeshGroupList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list MeshGroupList) Clone() MeshGroupList {
	var meshGroupList MeshGroupList
	for _, meshGroup := range list {
		meshGroupList = append(meshGroupList, resources.Clone(meshGroup).(*MeshGroup))
	}
	return meshGroupList
}

func (list MeshGroupList) Each(f func(element *MeshGroup)) {
	for _, meshGroup := range list {
		f(meshGroup)
	}
}

func (list MeshGroupList) EachResource(f func(element resources.Resource)) {
	for _, meshGroup := range list {
		f(meshGroup)
	}
}

func (list MeshGroupList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *MeshGroup) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var _ resources.Resource = &MeshGroup{}

// Kubernetes Adapter for MeshGroup

func (o *MeshGroup) GetObjectKind() schema.ObjectKind {
	t := MeshGroupCrd.TypeMeta()
	return &t
}

func (o *MeshGroup) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*MeshGroup)
}

var MeshGroupCrd = crd.NewCrd("supergloo.solo.io",
	"meshgroups",
	"supergloo.solo.io",
	"v1",
	"MeshGroup",
	"mg",
	false,
	&MeshGroup{})
