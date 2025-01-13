// Code generated by MockGen. DO NOT EDIT.
// Source: ./common/flagdinjector/flagdinjector.go

// Package commonmock is a generated GoMock package.
package commonmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/open-feature/open-feature-operator/apis/core/v1beta1"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockIFlagdContainerInjector is a mock of IFlagdContainerInjector interface.
type MockIFlagdContainerInjector struct {
	ctrl     *gomock.Controller
	recorder *MockIFlagdContainerInjectorMockRecorder
}

// MockIFlagdContainerInjectorMockRecorder is the mock recorder for MockIFlagdContainerInjector.
type MockIFlagdContainerInjectorMockRecorder struct {
	mock *MockIFlagdContainerInjector
}

// NewMockIFlagdContainerInjector creates a new mock instance.
func NewMockIFlagdContainerInjector(ctrl *gomock.Controller) *MockIFlagdContainerInjector {
	mock := &MockIFlagdContainerInjector{ctrl: ctrl}
	mock.recorder = &MockIFlagdContainerInjectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFlagdContainerInjector) EXPECT() *MockIFlagdContainerInjectorMockRecorder {
	return m.recorder
}

// EnableClusterRoleBinding mocks base method.
func (m *MockIFlagdContainerInjector) EnableClusterRoleBinding(ctx context.Context, namespace, serviceAccountName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableClusterRoleBinding", ctx, namespace, serviceAccountName)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableClusterRoleBinding indicates an expected call of EnableClusterRoleBinding.
func (mr *MockIFlagdContainerInjectorMockRecorder) EnableClusterRoleBinding(ctx, namespace, serviceAccountName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableClusterRoleBinding", reflect.TypeOf((*MockIFlagdContainerInjector)(nil).EnableClusterRoleBinding), ctx, namespace, serviceAccountName)
}

// InjectFlagd mocks base method.
func (m *MockIFlagdContainerInjector) InjectFlagd(ctx context.Context, objectMeta *v10.ObjectMeta, podSpec *v1.PodSpec, flagSourceConfig *v1beta1.FeatureFlagSourceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectFlagd", ctx, objectMeta, podSpec, flagSourceConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectFlagd indicates an expected call of InjectFlagd.
func (mr *MockIFlagdContainerInjectorMockRecorder) InjectFlagd(ctx, objectMeta, podSpec, flagSourceConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectFlagd", reflect.TypeOf((*MockIFlagdContainerInjector)(nil).InjectFlagd), ctx, objectMeta, podSpec, flagSourceConfig)
}
