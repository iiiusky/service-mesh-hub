// Code generated by MockGen. DO NOT EDIT.
// Source: clientset.go

// Package clients is a generated GoMock package.
package clients

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	v10 "k8s.io/api/core/v1"
)

// MockWebhookResourceClient is a mock of WebhookResourceClient interface
type MockWebhookResourceClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookResourceClientMockRecorder
}

// MockWebhookResourceClientMockRecorder is the mock recorder for MockWebhookResourceClient
type MockWebhookResourceClientMockRecorder struct {
	mock *MockWebhookResourceClient
}

// NewMockWebhookResourceClient creates a new mock instance
func NewMockWebhookResourceClient(ctrl *gomock.Controller) *MockWebhookResourceClient {
	mock := &MockWebhookResourceClient{ctrl: ctrl}
	mock.recorder = &MockWebhookResourceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebhookResourceClient) EXPECT() *MockWebhookResourceClientMockRecorder {
	return m.recorder
}

// ListMeshes mocks base method
func (m *MockWebhookResourceClient) ListMeshes(namespace string, opts clients.ListOpts) (v1.MeshList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMeshes", namespace, opts)
	ret0, _ := ret[0].(v1.MeshList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMeshes indicates an expected call of ListMeshes
func (mr *MockWebhookResourceClientMockRecorder) ListMeshes(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMeshes", reflect.TypeOf((*MockWebhookResourceClient)(nil).ListMeshes), namespace, opts)
}

// GetConfigMap mocks base method
func (m *MockWebhookResourceClient) GetConfigMap(namespace, name string) (*v10.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", namespace, name)
	ret0, _ := ret[0].(*v10.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap
func (mr *MockWebhookResourceClientMockRecorder) GetConfigMap(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockWebhookResourceClient)(nil).GetConfigMap), namespace, name)
}