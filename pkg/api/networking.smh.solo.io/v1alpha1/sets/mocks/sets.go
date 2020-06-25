// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/sets"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockTrafficPolicySet is a mock of TrafficPolicySet interface.
type MockTrafficPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicySetMockRecorder
}

// MockTrafficPolicySetMockRecorder is the mock recorder for MockTrafficPolicySet.
type MockTrafficPolicySetMockRecorder struct {
	mock *MockTrafficPolicySet
}

// NewMockTrafficPolicySet creates a new mock instance.
func NewMockTrafficPolicySet(ctrl *gomock.Controller) *MockTrafficPolicySet {
	mock := &MockTrafficPolicySet{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicySet) EXPECT() *MockTrafficPolicySetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockTrafficPolicySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockTrafficPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockTrafficPolicySet)(nil).Keys))
}

// List mocks base method.
func (m *MockTrafficPolicySet) List() []*v1alpha1.TrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.TrafficPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockTrafficPolicySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrafficPolicySet)(nil).List))
}

// Map mocks base method.
func (m *MockTrafficPolicySet) Map() map[string]*v1alpha1.TrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.TrafficPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockTrafficPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockTrafficPolicySet)(nil).Map))
}

// Insert mocks base method.
func (m *MockTrafficPolicySet) Insert(trafficPolicy ...*v1alpha1.TrafficPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockTrafficPolicySetMockRecorder) Insert(trafficPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrafficPolicySet)(nil).Insert), trafficPolicy...)
}

// Equal mocks base method.
func (m *MockTrafficPolicySet) Equal(trafficPolicySet v1alpha1sets.TrafficPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", trafficPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockTrafficPolicySetMockRecorder) Equal(trafficPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTrafficPolicySet)(nil).Equal), trafficPolicySet)
}

// Has mocks base method.
func (m *MockTrafficPolicySet) Has(trafficPolicy *v1alpha1.TrafficPolicy) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", trafficPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockTrafficPolicySetMockRecorder) Has(trafficPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTrafficPolicySet)(nil).Has), trafficPolicy)
}

// Delete mocks base method.
func (m *MockTrafficPolicySet) Delete(trafficPolicy *v1alpha1.TrafficPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", trafficPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockTrafficPolicySetMockRecorder) Delete(trafficPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrafficPolicySet)(nil).Delete), trafficPolicy)
}

// Union mocks base method.
func (m *MockTrafficPolicySet) Union(set v1alpha1sets.TrafficPolicySet) v1alpha1sets.TrafficPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.TrafficPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockTrafficPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockTrafficPolicySet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockTrafficPolicySet) Difference(set v1alpha1sets.TrafficPolicySet) v1alpha1sets.TrafficPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.TrafficPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockTrafficPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockTrafficPolicySet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockTrafficPolicySet) Intersection(set v1alpha1sets.TrafficPolicySet) v1alpha1sets.TrafficPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.TrafficPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockTrafficPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockTrafficPolicySet)(nil).Intersection), set)
}

// MockAccessControlPolicySet is a mock of AccessControlPolicySet interface.
type MockAccessControlPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicySetMockRecorder
}

// MockAccessControlPolicySetMockRecorder is the mock recorder for MockAccessControlPolicySet.
type MockAccessControlPolicySetMockRecorder struct {
	mock *MockAccessControlPolicySet
}

// NewMockAccessControlPolicySet creates a new mock instance.
func NewMockAccessControlPolicySet(ctrl *gomock.Controller) *MockAccessControlPolicySet {
	mock := &MockAccessControlPolicySet{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicySet) EXPECT() *MockAccessControlPolicySetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockAccessControlPolicySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockAccessControlPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Keys))
}

// List mocks base method.
func (m *MockAccessControlPolicySet) List() []*v1alpha1.AccessControlPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.AccessControlPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockAccessControlPolicySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccessControlPolicySet)(nil).List))
}

// Map mocks base method.
func (m *MockAccessControlPolicySet) Map() map[string]*v1alpha1.AccessControlPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.AccessControlPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockAccessControlPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Map))
}

// Insert mocks base method.
func (m *MockAccessControlPolicySet) Insert(accessControlPolicy ...*v1alpha1.AccessControlPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range accessControlPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockAccessControlPolicySetMockRecorder) Insert(accessControlPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Insert), accessControlPolicy...)
}

// Equal mocks base method.
func (m *MockAccessControlPolicySet) Equal(accessControlPolicySet v1alpha1sets.AccessControlPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", accessControlPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockAccessControlPolicySetMockRecorder) Equal(accessControlPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Equal), accessControlPolicySet)
}

// Has mocks base method.
func (m *MockAccessControlPolicySet) Has(accessControlPolicy *v1alpha1.AccessControlPolicy) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", accessControlPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockAccessControlPolicySetMockRecorder) Has(accessControlPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Has), accessControlPolicy)
}

// Delete mocks base method.
func (m *MockAccessControlPolicySet) Delete(accessControlPolicy *v1alpha1.AccessControlPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", accessControlPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockAccessControlPolicySetMockRecorder) Delete(accessControlPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Delete), accessControlPolicy)
}

// Union mocks base method.
func (m *MockAccessControlPolicySet) Union(set v1alpha1sets.AccessControlPolicySet) v1alpha1sets.AccessControlPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.AccessControlPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockAccessControlPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockAccessControlPolicySet) Difference(set v1alpha1sets.AccessControlPolicySet) v1alpha1sets.AccessControlPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.AccessControlPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockAccessControlPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockAccessControlPolicySet) Intersection(set v1alpha1sets.AccessControlPolicySet) v1alpha1sets.AccessControlPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.AccessControlPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockAccessControlPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockAccessControlPolicySet)(nil).Intersection), set)
}

// MockVirtualMeshSet is a mock of VirtualMeshSet interface.
type MockVirtualMeshSet struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshSetMockRecorder
}

// MockVirtualMeshSetMockRecorder is the mock recorder for MockVirtualMeshSet.
type MockVirtualMeshSetMockRecorder struct {
	mock *MockVirtualMeshSet
}

// NewMockVirtualMeshSet creates a new mock instance.
func NewMockVirtualMeshSet(ctrl *gomock.Controller) *MockVirtualMeshSet {
	mock := &MockVirtualMeshSet{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshSet) EXPECT() *MockVirtualMeshSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockVirtualMeshSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockVirtualMeshSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockVirtualMeshSet)(nil).Keys))
}

// List mocks base method.
func (m *MockVirtualMeshSet) List() []*v1alpha1.VirtualMesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.VirtualMesh)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockVirtualMeshSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMeshSet)(nil).List))
}

// Map mocks base method.
func (m *MockVirtualMeshSet) Map() map[string]*v1alpha1.VirtualMesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.VirtualMesh)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockVirtualMeshSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockVirtualMeshSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockVirtualMeshSet) Insert(virtualMesh ...*v1alpha1.VirtualMesh) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualMesh {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockVirtualMeshSetMockRecorder) Insert(virtualMesh ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockVirtualMeshSet)(nil).Insert), virtualMesh...)
}

// Equal mocks base method.
func (m *MockVirtualMeshSet) Equal(virtualMeshSet v1alpha1sets.VirtualMeshSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", virtualMeshSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockVirtualMeshSetMockRecorder) Equal(virtualMeshSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVirtualMeshSet)(nil).Equal), virtualMeshSet)
}

// Has mocks base method.
func (m *MockVirtualMeshSet) Has(virtualMesh *v1alpha1.VirtualMesh) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", virtualMesh)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockVirtualMeshSetMockRecorder) Has(virtualMesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockVirtualMeshSet)(nil).Has), virtualMesh)
}

// Delete mocks base method.
func (m *MockVirtualMeshSet) Delete(virtualMesh *v1alpha1.VirtualMesh) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", virtualMesh)
}

// Delete indicates an expected call of Delete.
func (mr *MockVirtualMeshSetMockRecorder) Delete(virtualMesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMeshSet)(nil).Delete), virtualMesh)
}

// Union mocks base method.
func (m *MockVirtualMeshSet) Union(set v1alpha1sets.VirtualMeshSet) v1alpha1sets.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockVirtualMeshSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockVirtualMeshSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockVirtualMeshSet) Difference(set v1alpha1sets.VirtualMeshSet) v1alpha1sets.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockVirtualMeshSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockVirtualMeshSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockVirtualMeshSet) Intersection(set v1alpha1sets.VirtualMeshSet) v1alpha1sets.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualMeshSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockVirtualMeshSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockVirtualMeshSet)(nil).Intersection), set)
}

// MockFailoverServiceSet is a mock of FailoverServiceSet interface.
type MockFailoverServiceSet struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverServiceSetMockRecorder
}

// MockFailoverServiceSetMockRecorder is the mock recorder for MockFailoverServiceSet.
type MockFailoverServiceSetMockRecorder struct {
	mock *MockFailoverServiceSet
}

// NewMockFailoverServiceSet creates a new mock instance.
func NewMockFailoverServiceSet(ctrl *gomock.Controller) *MockFailoverServiceSet {
	mock := &MockFailoverServiceSet{ctrl: ctrl}
	mock.recorder = &MockFailoverServiceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFailoverServiceSet) EXPECT() *MockFailoverServiceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockFailoverServiceSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockFailoverServiceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFailoverServiceSet)(nil).Keys))
}

// List mocks base method.
func (m *MockFailoverServiceSet) List() []*v1alpha1.FailoverService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.FailoverService)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFailoverServiceSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFailoverServiceSet)(nil).List))
}

// Map mocks base method.
func (m *MockFailoverServiceSet) Map() map[string]*v1alpha1.FailoverService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.FailoverService)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockFailoverServiceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFailoverServiceSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockFailoverServiceSet) Insert(failoverService ...*v1alpha1.FailoverService) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range failoverService {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockFailoverServiceSetMockRecorder) Insert(failoverService ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFailoverServiceSet)(nil).Insert), failoverService...)
}

// Equal mocks base method.
func (m *MockFailoverServiceSet) Equal(failoverServiceSet v1alpha1sets.FailoverServiceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", failoverServiceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockFailoverServiceSetMockRecorder) Equal(failoverServiceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFailoverServiceSet)(nil).Equal), failoverServiceSet)
}

// Has mocks base method.
func (m *MockFailoverServiceSet) Has(failoverService *v1alpha1.FailoverService) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", failoverService)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockFailoverServiceSetMockRecorder) Has(failoverService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFailoverServiceSet)(nil).Has), failoverService)
}

// Delete mocks base method.
func (m *MockFailoverServiceSet) Delete(failoverService *v1alpha1.FailoverService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", failoverService)
}

// Delete indicates an expected call of Delete.
func (mr *MockFailoverServiceSetMockRecorder) Delete(failoverService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFailoverServiceSet)(nil).Delete), failoverService)
}

// Union mocks base method.
func (m *MockFailoverServiceSet) Union(set v1alpha1sets.FailoverServiceSet) v1alpha1sets.FailoverServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.FailoverServiceSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockFailoverServiceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFailoverServiceSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockFailoverServiceSet) Difference(set v1alpha1sets.FailoverServiceSet) v1alpha1sets.FailoverServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.FailoverServiceSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockFailoverServiceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFailoverServiceSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockFailoverServiceSet) Intersection(set v1alpha1sets.FailoverServiceSet) v1alpha1sets.FailoverServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.FailoverServiceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockFailoverServiceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFailoverServiceSet)(nil).Intersection), set)
}
