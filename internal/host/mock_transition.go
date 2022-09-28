// Code generated by MockGen. DO NOT EDIT.
// Source: transition.go

// Package host is a generated GoMock package.
package host

import (
	reflect "reflect"
	time "time"

	stateswitch "github.com/filanov/stateswitch"
	gomock "github.com/golang/mock/gomock"
)

// MockTransitionHandler is a mock of TransitionHandler interface.
type MockTransitionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTransitionHandlerMockRecorder
}

// MockTransitionHandlerMockRecorder is the mock recorder for MockTransitionHandler.
type MockTransitionHandlerMockRecorder struct {
	mock *MockTransitionHandler
}

// NewMockTransitionHandler creates a new mock instance.
func NewMockTransitionHandler(ctrl *gomock.Controller) *MockTransitionHandler {
	mock := &MockTransitionHandler{ctrl: ctrl}
	mock.recorder = &MockTransitionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransitionHandler) EXPECT() *MockTransitionHandlerMockRecorder {
	return m.recorder
}

// HasClusterError mocks base method.
func (m *MockTransitionHandler) HasClusterError(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasClusterError", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasClusterError indicates an expected call of HasClusterError.
func (mr *MockTransitionHandlerMockRecorder) HasClusterError(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasClusterError", reflect.TypeOf((*MockTransitionHandler)(nil).HasClusterError), sw, args)
}

// HasInstallationInProgressTimedOut mocks base method.
func (m *MockTransitionHandler) HasInstallationInProgressTimedOut(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasInstallationInProgressTimedOut", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasInstallationInProgressTimedOut indicates an expected call of HasInstallationInProgressTimedOut.
func (mr *MockTransitionHandlerMockRecorder) HasInstallationInProgressTimedOut(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasInstallationInProgressTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).HasInstallationInProgressTimedOut), sw, arg1)
}

// HasStatusTimedOut mocks base method.
func (m *MockTransitionHandler) HasStatusTimedOut(timeout time.Duration) stateswitch.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasStatusTimedOut", timeout)
	ret0, _ := ret[0].(stateswitch.Condition)
	return ret0
}

// HasStatusTimedOut indicates an expected call of HasStatusTimedOut.
func (mr *MockTransitionHandlerMockRecorder) HasStatusTimedOut(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasStatusTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).HasStatusTimedOut), timeout)
}

// HostNotResponsiveWhileInstallation mocks base method.
func (m *MockTransitionHandler) HostNotResponsiveWhileInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostNotResponsiveWhileInstallation", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HostNotResponsiveWhileInstallation indicates an expected call of HostNotResponsiveWhileInstallation.
func (mr *MockTransitionHandlerMockRecorder) HostNotResponsiveWhileInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostNotResponsiveWhileInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).HostNotResponsiveWhileInstallation), sw, args)
}

// HostNotResponsiveWhilePreparingInstallation mocks base method.
func (m *MockTransitionHandler) HostNotResponsiveWhilePreparingInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostNotResponsiveWhilePreparingInstallation", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HostNotResponsiveWhilePreparingInstallation indicates an expected call of HostNotResponsiveWhilePreparingInstallation.
func (mr *MockTransitionHandlerMockRecorder) HostNotResponsiveWhilePreparingInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostNotResponsiveWhilePreparingInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).HostNotResponsiveWhilePreparingInstallation), sw, args)
}

// IsDay2Host mocks base method.
func (m *MockTransitionHandler) IsDay2Host(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDay2Host", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDay2Host indicates an expected call of IsDay2Host.
func (mr *MockTransitionHandlerMockRecorder) IsDay2Host(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDay2Host", reflect.TypeOf((*MockTransitionHandler)(nil).IsDay2Host), sw, args)
}

// IsHostInDone mocks base method.
func (m *MockTransitionHandler) IsHostInDone(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHostInDone", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHostInDone indicates an expected call of IsHostInDone.
func (mr *MockTransitionHandlerMockRecorder) IsHostInDone(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHostInDone", reflect.TypeOf((*MockTransitionHandler)(nil).IsHostInDone), sw, arg1)
}

// IsHostInReboot mocks base method.
func (m *MockTransitionHandler) IsHostInReboot(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHostInReboot", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHostInReboot indicates an expected call of IsHostInReboot.
func (mr *MockTransitionHandlerMockRecorder) IsHostInReboot(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHostInReboot", reflect.TypeOf((*MockTransitionHandler)(nil).IsHostInReboot), sw, arg1)
}

// IsLogCollectionTimedOut mocks base method.
func (m *MockTransitionHandler) IsLogCollectionTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLogCollectionTimedOut", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLogCollectionTimedOut indicates an expected call of IsLogCollectionTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsLogCollectionTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLogCollectionTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsLogCollectionTimedOut), sw, args)
}

// IsUnboundHost mocks base method.
func (m *MockTransitionHandler) IsUnboundHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnboundHost", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUnboundHost indicates an expected call of IsUnboundHost.
func (mr *MockTransitionHandlerMockRecorder) IsUnboundHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnboundHost", reflect.TypeOf((*MockTransitionHandler)(nil).IsUnboundHost), sw, args)
}

// IsValidRoleForInstallation mocks base method.
func (m *MockTransitionHandler) IsValidRoleForInstallation(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidRoleForInstallation", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidRoleForInstallation indicates an expected call of IsValidRoleForInstallation.
func (mr *MockTransitionHandlerMockRecorder) IsValidRoleForInstallation(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidRoleForInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).IsValidRoleForInstallation), sw, arg1)
}

// PostBindHost mocks base method.
func (m *MockTransitionHandler) PostBindHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBindHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostBindHost indicates an expected call of PostBindHost.
func (mr *MockTransitionHandlerMockRecorder) PostBindHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBindHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostBindHost), sw, args)
}

// PostCancelInstallation mocks base method.
func (m *MockTransitionHandler) PostCancelInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCancelInstallation", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCancelInstallation indicates an expected call of PostCancelInstallation.
func (mr *MockTransitionHandlerMockRecorder) PostCancelInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCancelInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).PostCancelInstallation), sw, args)
}

// PostHostInstallationFailed mocks base method.
func (m *MockTransitionHandler) PostHostInstallationFailed(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostHostInstallationFailed", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostHostInstallationFailed indicates an expected call of PostHostInstallationFailed.
func (mr *MockTransitionHandlerMockRecorder) PostHostInstallationFailed(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHostInstallationFailed", reflect.TypeOf((*MockTransitionHandler)(nil).PostHostInstallationFailed), sw, args)
}

// PostHostMediaDisconnected mocks base method.
func (m *MockTransitionHandler) PostHostMediaDisconnected(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostHostMediaDisconnected", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostHostMediaDisconnected indicates an expected call of PostHostMediaDisconnected.
func (mr *MockTransitionHandlerMockRecorder) PostHostMediaDisconnected(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHostMediaDisconnected", reflect.TypeOf((*MockTransitionHandler)(nil).PostHostMediaDisconnected), sw, args)
}

// PostInstallHost mocks base method.
func (m *MockTransitionHandler) PostInstallHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInstallHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostInstallHost indicates an expected call of PostInstallHost.
func (mr *MockTransitionHandlerMockRecorder) PostInstallHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInstallHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostInstallHost), sw, args)
}

// PostPreparingForInstallationHost mocks base method.
func (m *MockTransitionHandler) PostPreparingForInstallationHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostPreparingForInstallationHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostPreparingForInstallationHost indicates an expected call of PostPreparingForInstallationHost.
func (mr *MockTransitionHandlerMockRecorder) PostPreparingForInstallationHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostPreparingForInstallationHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostPreparingForInstallationHost), sw, args)
}

// PostReclaim mocks base method.
func (m *MockTransitionHandler) PostReclaim(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostReclaim", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostReclaim indicates an expected call of PostReclaim.
func (mr *MockTransitionHandlerMockRecorder) PostReclaim(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostReclaim", reflect.TypeOf((*MockTransitionHandler)(nil).PostReclaim), sw, args)
}

// PostRefreshHost mocks base method.
func (m *MockTransitionHandler) PostRefreshHost(reason string) stateswitch.PostTransition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshHost", reason)
	ret0, _ := ret[0].(stateswitch.PostTransition)
	return ret0
}

// PostRefreshHost indicates an expected call of PostRefreshHost.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshHost(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshHost), reason)
}

// PostRefreshHostRefreshStageUpdateTime mocks base method.
func (m *MockTransitionHandler) PostRefreshHostRefreshStageUpdateTime(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshHostRefreshStageUpdateTime", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRefreshHostRefreshStageUpdateTime indicates an expected call of PostRefreshHostRefreshStageUpdateTime.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshHostRefreshStageUpdateTime(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshHostRefreshStageUpdateTime", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshHostRefreshStageUpdateTime), sw, args)
}

// PostRefreshLogsProgress mocks base method.
func (m *MockTransitionHandler) PostRefreshLogsProgress(progress string) stateswitch.PostTransition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshLogsProgress", progress)
	ret0, _ := ret[0].(stateswitch.PostTransition)
	return ret0
}

// PostRefreshLogsProgress indicates an expected call of PostRefreshLogsProgress.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshLogsProgress(progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshLogsProgress", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshLogsProgress), progress)
}

// PostRefreshReclaimTimeout mocks base method.
func (m *MockTransitionHandler) PostRefreshReclaimTimeout(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshReclaimTimeout", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRefreshReclaimTimeout indicates an expected call of PostRefreshReclaimTimeout.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshReclaimTimeout(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshReclaimTimeout", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshReclaimTimeout), sw, args)
}

// PostRegisterDuringInstallation mocks base method.
func (m *MockTransitionHandler) PostRegisterDuringInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRegisterDuringInstallation", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRegisterDuringInstallation indicates an expected call of PostRegisterDuringInstallation.
func (mr *MockTransitionHandlerMockRecorder) PostRegisterDuringInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRegisterDuringInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).PostRegisterDuringInstallation), sw, args)
}

// PostRegisterDuringReboot mocks base method.
func (m *MockTransitionHandler) PostRegisterDuringReboot(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRegisterDuringReboot", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRegisterDuringReboot indicates an expected call of PostRegisterDuringReboot.
func (mr *MockTransitionHandlerMockRecorder) PostRegisterDuringReboot(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRegisterDuringReboot", reflect.TypeOf((*MockTransitionHandler)(nil).PostRegisterDuringReboot), sw, args)
}

// PostRegisterHost mocks base method.
func (m *MockTransitionHandler) PostRegisterHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRegisterHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRegisterHost indicates an expected call of PostRegisterHost.
func (mr *MockTransitionHandlerMockRecorder) PostRegisterHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRegisterHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostRegisterHost), sw, args)
}

// PostResetHost mocks base method.
func (m *MockTransitionHandler) PostResetHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostResetHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostResetHost indicates an expected call of PostResetHost.
func (mr *MockTransitionHandlerMockRecorder) PostResetHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostResetHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostResetHost), sw, args)
}

// PostResettingPendingUserAction mocks base method.
func (m *MockTransitionHandler) PostResettingPendingUserAction(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostResettingPendingUserAction", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostResettingPendingUserAction indicates an expected call of PostResettingPendingUserAction.
func (mr *MockTransitionHandlerMockRecorder) PostResettingPendingUserAction(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostResettingPendingUserAction", reflect.TypeOf((*MockTransitionHandler)(nil).PostResettingPendingUserAction), sw, args)
}

// PostUnbindHost mocks base method.
func (m *MockTransitionHandler) PostUnbindHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUnbindHost", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostUnbindHost indicates an expected call of PostUnbindHost.
func (mr *MockTransitionHandlerMockRecorder) PostUnbindHost(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUnbindHost", reflect.TypeOf((*MockTransitionHandler)(nil).PostUnbindHost), sw, args)
}
