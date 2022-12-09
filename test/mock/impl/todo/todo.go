// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dytlan/moonlay-todo-list/impl/todo (interfaces: Service)

// Package mockTodo is a generated GoMock package.
package mockTodo

import (
	gomock "github.com/golang/mock/gomock"
	v4 "github.com/labstack/echo/v4"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockService) Delete(arg0 v4.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0)
}

// GetByIDs mocks base method
func (m *MockService) GetByIDs(arg0 v4.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockServiceMockRecorder) GetByIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockService)(nil).GetByIDs), arg0)
}

// List mocks base method
func (m *MockService) List(arg0 v4.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), arg0)
}

// Upsert mocks base method
func (m *MockService) Upsert(arg0 v4.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockServiceMockRecorder) Upsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockService)(nil).Upsert), arg0)
}
