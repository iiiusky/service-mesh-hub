// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TrafficPolicySet interface {
	Keys() sets.String
	List() []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	Map() map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy
	Insert(trafficPolicy ...*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	Equal(trafficPolicySet TrafficPolicySet) bool
	Has(trafficPolicy *networking_smh_solo_io_v1alpha1.TrafficPolicy) bool
	Delete(trafficPolicy *networking_smh_solo_io_v1alpha1.TrafficPolicy)
	Union(set TrafficPolicySet) TrafficPolicySet
	Difference(set TrafficPolicySet) TrafficPolicySet
	Intersection(set TrafficPolicySet) TrafficPolicySet
}

func makeGenericTrafficPolicySet(trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range trafficPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type trafficPolicySet struct {
	set sksets.ResourceSet
}

func NewTrafficPolicySet(trafficPolicyList ...*networking_smh_solo_io_v1alpha1.TrafficPolicy) TrafficPolicySet {
	return &trafficPolicySet{set: makeGenericTrafficPolicySet(trafficPolicyList)}
}

func (s trafficPolicySet) Keys() sets.String {
	return s.set.Keys()
}

func (s trafficPolicySet) List() []*networking_smh_solo_io_v1alpha1.TrafficPolicy {
	var trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	for _, obj := range s.set.List() {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy))
	}
	return trafficPolicyList
}

func (s trafficPolicySet) Map() map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	}
	return newMap
}

func (s trafficPolicySet) Insert(
	trafficPolicyList ...*networking_smh_solo_io_v1alpha1.TrafficPolicy,
) {
	for _, obj := range trafficPolicyList {
		s.set.Insert(obj)
	}
}

func (s trafficPolicySet) Has(trafficPolicy *networking_smh_solo_io_v1alpha1.TrafficPolicy) bool {
	return s.set.Has(trafficPolicy)
}

func (s trafficPolicySet) Equal(
	trafficPolicySet TrafficPolicySet,
) bool {
	return s.set.Equal(makeGenericTrafficPolicySet(trafficPolicySet.List()))
}

func (s trafficPolicySet) Delete(TrafficPolicy *networking_smh_solo_io_v1alpha1.TrafficPolicy) {
	s.set.Delete(TrafficPolicy)
}

func (s trafficPolicySet) Union(set TrafficPolicySet) TrafficPolicySet {
	return NewTrafficPolicySet(append(s.List(), set.List()...)...)
}

func (s trafficPolicySet) Difference(set TrafficPolicySet) TrafficPolicySet {
	newSet := s.set.Difference(makeGenericTrafficPolicySet(set.List()))
	return trafficPolicySet{set: newSet}
}

func (s trafficPolicySet) Intersection(set TrafficPolicySet) TrafficPolicySet {
	newSet := s.set.Intersection(makeGenericTrafficPolicySet(set.List()))
	var trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	for _, obj := range newSet.List() {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy))
	}
	return NewTrafficPolicySet(trafficPolicyList...)
}

type AccessControlPolicySet interface {
	Keys() sets.String
	List() []*networking_smh_solo_io_v1alpha1.AccessControlPolicy
	Map() map[string]*networking_smh_solo_io_v1alpha1.AccessControlPolicy
	Insert(accessControlPolicy ...*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	Equal(accessControlPolicySet AccessControlPolicySet) bool
	Has(accessControlPolicy *networking_smh_solo_io_v1alpha1.AccessControlPolicy) bool
	Delete(accessControlPolicy *networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	Union(set AccessControlPolicySet) AccessControlPolicySet
	Difference(set AccessControlPolicySet) AccessControlPolicySet
	Intersection(set AccessControlPolicySet) AccessControlPolicySet
}

func makeGenericAccessControlPolicySet(accessControlPolicyList []*networking_smh_solo_io_v1alpha1.AccessControlPolicy) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range accessControlPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type accessControlPolicySet struct {
	set sksets.ResourceSet
}

func NewAccessControlPolicySet(accessControlPolicyList ...*networking_smh_solo_io_v1alpha1.AccessControlPolicy) AccessControlPolicySet {
	return &accessControlPolicySet{set: makeGenericAccessControlPolicySet(accessControlPolicyList)}
}

func (s accessControlPolicySet) Keys() sets.String {
	return s.set.Keys()
}

func (s accessControlPolicySet) List() []*networking_smh_solo_io_v1alpha1.AccessControlPolicy {
	var accessControlPolicyList []*networking_smh_solo_io_v1alpha1.AccessControlPolicy
	for _, obj := range s.set.List() {
		accessControlPolicyList = append(accessControlPolicyList, obj.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy))
	}
	return accessControlPolicyList
}

func (s accessControlPolicySet) Map() map[string]*networking_smh_solo_io_v1alpha1.AccessControlPolicy {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.AccessControlPolicy{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	}
	return newMap
}

func (s accessControlPolicySet) Insert(
	accessControlPolicyList ...*networking_smh_solo_io_v1alpha1.AccessControlPolicy,
) {
	for _, obj := range accessControlPolicyList {
		s.set.Insert(obj)
	}
}

func (s accessControlPolicySet) Has(accessControlPolicy *networking_smh_solo_io_v1alpha1.AccessControlPolicy) bool {
	return s.set.Has(accessControlPolicy)
}

func (s accessControlPolicySet) Equal(
	accessControlPolicySet AccessControlPolicySet,
) bool {
	return s.set.Equal(makeGenericAccessControlPolicySet(accessControlPolicySet.List()))
}

func (s accessControlPolicySet) Delete(AccessControlPolicy *networking_smh_solo_io_v1alpha1.AccessControlPolicy) {
	s.set.Delete(AccessControlPolicy)
}

func (s accessControlPolicySet) Union(set AccessControlPolicySet) AccessControlPolicySet {
	return NewAccessControlPolicySet(append(s.List(), set.List()...)...)
}

func (s accessControlPolicySet) Difference(set AccessControlPolicySet) AccessControlPolicySet {
	newSet := s.set.Difference(makeGenericAccessControlPolicySet(set.List()))
	return accessControlPolicySet{set: newSet}
}

func (s accessControlPolicySet) Intersection(set AccessControlPolicySet) AccessControlPolicySet {
	newSet := s.set.Intersection(makeGenericAccessControlPolicySet(set.List()))
	var accessControlPolicyList []*networking_smh_solo_io_v1alpha1.AccessControlPolicy
	for _, obj := range newSet.List() {
		accessControlPolicyList = append(accessControlPolicyList, obj.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy))
	}
	return NewAccessControlPolicySet(accessControlPolicyList...)
}

type VirtualMeshSet interface {
	Keys() sets.String
	List() []*networking_smh_solo_io_v1alpha1.VirtualMesh
	Map() map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh
	Insert(virtualMesh ...*networking_smh_solo_io_v1alpha1.VirtualMesh)
	Equal(virtualMeshSet VirtualMeshSet) bool
	Has(virtualMesh *networking_smh_solo_io_v1alpha1.VirtualMesh) bool
	Delete(virtualMesh *networking_smh_solo_io_v1alpha1.VirtualMesh)
	Union(set VirtualMeshSet) VirtualMeshSet
	Difference(set VirtualMeshSet) VirtualMeshSet
	Intersection(set VirtualMeshSet) VirtualMeshSet
}

func makeGenericVirtualMeshSet(virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range virtualMeshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualMeshSet struct {
	set sksets.ResourceSet
}

func NewVirtualMeshSet(virtualMeshList ...*networking_smh_solo_io_v1alpha1.VirtualMesh) VirtualMeshSet {
	return &virtualMeshSet{set: makeGenericVirtualMeshSet(virtualMeshList)}
}

func (s virtualMeshSet) Keys() sets.String {
	return s.set.Keys()
}

func (s virtualMeshSet) List() []*networking_smh_solo_io_v1alpha1.VirtualMesh {
	var virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh
	for _, obj := range s.set.List() {
		virtualMeshList = append(virtualMeshList, obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh))
	}
	return virtualMeshList
}

func (s virtualMeshSet) Map() map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	}
	return newMap
}

func (s virtualMeshSet) Insert(
	virtualMeshList ...*networking_smh_solo_io_v1alpha1.VirtualMesh,
) {
	for _, obj := range virtualMeshList {
		s.set.Insert(obj)
	}
}

func (s virtualMeshSet) Has(virtualMesh *networking_smh_solo_io_v1alpha1.VirtualMesh) bool {
	return s.set.Has(virtualMesh)
}

func (s virtualMeshSet) Equal(
	virtualMeshSet VirtualMeshSet,
) bool {
	return s.set.Equal(makeGenericVirtualMeshSet(virtualMeshSet.List()))
}

func (s virtualMeshSet) Delete(VirtualMesh *networking_smh_solo_io_v1alpha1.VirtualMesh) {
	s.set.Delete(VirtualMesh)
}

func (s virtualMeshSet) Union(set VirtualMeshSet) VirtualMeshSet {
	return NewVirtualMeshSet(append(s.List(), set.List()...)...)
}

func (s virtualMeshSet) Difference(set VirtualMeshSet) VirtualMeshSet {
	newSet := s.set.Difference(makeGenericVirtualMeshSet(set.List()))
	return virtualMeshSet{set: newSet}
}

func (s virtualMeshSet) Intersection(set VirtualMeshSet) VirtualMeshSet {
	newSet := s.set.Intersection(makeGenericVirtualMeshSet(set.List()))
	var virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh
	for _, obj := range newSet.List() {
		virtualMeshList = append(virtualMeshList, obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh))
	}
	return NewVirtualMeshSet(virtualMeshList...)
}

type FailoverServiceSet interface {
	Keys() sets.String
	List() []*networking_smh_solo_io_v1alpha1.FailoverService
	Map() map[string]*networking_smh_solo_io_v1alpha1.FailoverService
	Insert(failoverService ...*networking_smh_solo_io_v1alpha1.FailoverService)
	Equal(failoverServiceSet FailoverServiceSet) bool
	Has(failoverService *networking_smh_solo_io_v1alpha1.FailoverService) bool
	Delete(failoverService *networking_smh_solo_io_v1alpha1.FailoverService)
	Union(set FailoverServiceSet) FailoverServiceSet
	Difference(set FailoverServiceSet) FailoverServiceSet
	Intersection(set FailoverServiceSet) FailoverServiceSet
}

func makeGenericFailoverServiceSet(failoverServiceList []*networking_smh_solo_io_v1alpha1.FailoverService) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range failoverServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type failoverServiceSet struct {
	set sksets.ResourceSet
}

func NewFailoverServiceSet(failoverServiceList ...*networking_smh_solo_io_v1alpha1.FailoverService) FailoverServiceSet {
	return &failoverServiceSet{set: makeGenericFailoverServiceSet(failoverServiceList)}
}

func (s failoverServiceSet) Keys() sets.String {
	return s.set.Keys()
}

func (s failoverServiceSet) List() []*networking_smh_solo_io_v1alpha1.FailoverService {
	var failoverServiceList []*networking_smh_solo_io_v1alpha1.FailoverService
	for _, obj := range s.set.List() {
		failoverServiceList = append(failoverServiceList, obj.(*networking_smh_solo_io_v1alpha1.FailoverService))
	}
	return failoverServiceList
}

func (s failoverServiceSet) Map() map[string]*networking_smh_solo_io_v1alpha1.FailoverService {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.FailoverService{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.FailoverService)
	}
	return newMap
}

func (s failoverServiceSet) Insert(
	failoverServiceList ...*networking_smh_solo_io_v1alpha1.FailoverService,
) {
	for _, obj := range failoverServiceList {
		s.set.Insert(obj)
	}
}

func (s failoverServiceSet) Has(failoverService *networking_smh_solo_io_v1alpha1.FailoverService) bool {
	return s.set.Has(failoverService)
}

func (s failoverServiceSet) Equal(
	failoverServiceSet FailoverServiceSet,
) bool {
	return s.set.Equal(makeGenericFailoverServiceSet(failoverServiceSet.List()))
}

func (s failoverServiceSet) Delete(FailoverService *networking_smh_solo_io_v1alpha1.FailoverService) {
	s.set.Delete(FailoverService)
}

func (s failoverServiceSet) Union(set FailoverServiceSet) FailoverServiceSet {
	return NewFailoverServiceSet(append(s.List(), set.List()...)...)
}

func (s failoverServiceSet) Difference(set FailoverServiceSet) FailoverServiceSet {
	newSet := s.set.Difference(makeGenericFailoverServiceSet(set.List()))
	return failoverServiceSet{set: newSet}
}

func (s failoverServiceSet) Intersection(set FailoverServiceSet) FailoverServiceSet {
	newSet := s.set.Intersection(makeGenericFailoverServiceSet(set.List()))
	var failoverServiceList []*networking_smh_solo_io_v1alpha1.FailoverService
	for _, obj := range newSet.List() {
		failoverServiceList = append(failoverServiceList, obj.(*networking_smh_solo_io_v1alpha1.FailoverService))
	}
	return NewFailoverServiceSet(failoverServiceList...)
}
