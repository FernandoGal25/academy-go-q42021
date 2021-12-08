// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure/datastore/csv.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileManager is a mock of FileManager interface.
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager.
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance.
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// BuildHandler mocks base method.
func (m *MockFileManager) BuildHandler() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildHandler")
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildHandler indicates an expected call of BuildHandler.
func (mr *MockFileManagerMockRecorder) BuildHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildHandler", reflect.TypeOf((*MockFileManager)(nil).BuildHandler))
}

// Close mocks base method.
func (m *MockFileManager) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockFileManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileManager)(nil).Close))
}

// GetHeader mocks base method.
func (m *MockFileManager) GetHeader() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockFileManagerMockRecorder) GetHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockFileManager)(nil).GetHeader))
}

// Read mocks base method.
func (m *MockFileManager) Read() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileManagerMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileManager)(nil).Read))
}

// ReadAll mocks base method.
func (m *MockFileManager) ReadAll() ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockFileManagerMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockFileManager)(nil).ReadAll))
}

// Write mocks base method.
func (m *MockFileManager) Write(r []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockFileManagerMockRecorder) Write(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileManager)(nil).Write), r)
}