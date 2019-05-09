// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/mesh_ingress_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// MockMeshIngressWatcher is a mock of MeshIngressWatcher interface
type MockMeshIngressWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshIngressWatcherMockRecorder
}

// MockMeshIngressWatcherMockRecorder is the mock recorder for MockMeshIngressWatcher
type MockMeshIngressWatcherMockRecorder struct {
	mock *MockMeshIngressWatcher
}

// NewMockMeshIngressWatcher creates a new mock instance
func NewMockMeshIngressWatcher(ctrl *gomock.Controller) *MockMeshIngressWatcher {
	mock := &MockMeshIngressWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshIngressWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshIngressWatcher) EXPECT() *MockMeshIngressWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockMeshIngressWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan v1.MeshIngressList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1.MeshIngressList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockMeshIngressWatcherMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMeshIngressWatcher)(nil).Watch), namespace, opts)
}

// MockMeshIngressClient is a mock of MeshIngressClient interface
type MockMeshIngressClient struct {
	ctrl     *gomock.Controller
	recorder *MockMeshIngressClientMockRecorder
}

// MockMeshIngressClientMockRecorder is the mock recorder for MockMeshIngressClient
type MockMeshIngressClientMockRecorder struct {
	mock *MockMeshIngressClient
}

// NewMockMeshIngressClient creates a new mock instance
func NewMockMeshIngressClient(ctrl *gomock.Controller) *MockMeshIngressClient {
	mock := &MockMeshIngressClient{ctrl: ctrl}
	mock.recorder = &MockMeshIngressClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshIngressClient) EXPECT() *MockMeshIngressClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockMeshIngressClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockMeshIngressClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockMeshIngressClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockMeshIngressClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockMeshIngressClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockMeshIngressClient)(nil).Register))
}

// Read mocks base method
func (m *MockMeshIngressClient) Read(namespace, name string, opts clients.ReadOpts) (*v1.MeshIngress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1.MeshIngress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockMeshIngressClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMeshIngressClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockMeshIngressClient) Write(resource *v1.MeshIngress, opts clients.WriteOpts) (*v1.MeshIngress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1.MeshIngress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockMeshIngressClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockMeshIngressClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockMeshIngressClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMeshIngressClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshIngressClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockMeshIngressClient) List(namespace string, opts clients.ListOpts) (v1.MeshIngressList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1.MeshIngressList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockMeshIngressClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshIngressClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockMeshIngressClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1.MeshIngressList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1.MeshIngressList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockMeshIngressClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMeshIngressClient)(nil).Watch), namespace, opts)
}
