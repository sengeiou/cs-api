// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/interface/script.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockILusScript is a mock of ILusScript interface.
type MockILusScript struct {
	ctrl     *gomock.Controller
	recorder *MockILusScriptMockRecorder
}

// MockILusScriptMockRecorder is the mock recorder for MockILusScript.
type MockILusScriptMockRecorder struct {
	mock *MockILusScript
}

// NewMockILusScript creates a new mock instance.
func NewMockILusScript(ctrl *gomock.Controller) *MockILusScript {
	mock := &MockILusScript{ctrl: ctrl}
	mock.recorder = &MockILusScriptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILusScript) EXPECT() *MockILusScriptMockRecorder {
	return m.recorder
}

// RemoveToken mocks base method.
func (m *MockILusScript) RemoveToken(ctx context.Context, clientType, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveToken", ctx, clientType, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveToken indicates an expected call of RemoveToken.
func (mr *MockILusScriptMockRecorder) RemoveToken(ctx, clientType, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveToken", reflect.TypeOf((*MockILusScript)(nil).RemoveToken), ctx, clientType, name)
}

// SetToken mocks base method.
func (m *MockILusScript) SetToken(ctx context.Context, clientType, name, token string, value interface{}, duration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetToken", ctx, clientType, name, token, value, duration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetToken indicates an expected call of SetToken.
func (mr *MockILusScriptMockRecorder) SetToken(ctx, clientType, name, token, value, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockILusScript)(nil).SetToken), ctx, clientType, name, token, value, duration)
}
