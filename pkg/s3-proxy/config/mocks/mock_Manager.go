// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/config (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/config"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddOnChangeHook mocks base method
func (m *MockManager) AddOnChangeHook(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddOnChangeHook", arg0)
}

// AddOnChangeHook indicates an expected call of AddOnChangeHook
func (mr *MockManagerMockRecorder) AddOnChangeHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOnChangeHook", reflect.TypeOf((*MockManager)(nil).AddOnChangeHook), arg0)
}

// GetConfig mocks base method
func (m *MockManager) GetConfig() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockManagerMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockManager)(nil).GetConfig))
}

// Load mocks base method
func (m *MockManager) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockManagerMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockManager)(nil).Load))
}
