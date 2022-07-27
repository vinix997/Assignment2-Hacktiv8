// Code generated by MockGen. DO NOT EDIT.
// Source: handler/handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderHanlderInterface is a mock of OrderHanlderInterface interface.
type MockOrderHanlderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderHanlderInterfaceMockRecorder
}

// MockOrderHanlderInterfaceMockRecorder is the mock recorder for MockOrderHanlderInterface.
type MockOrderHanlderInterfaceMockRecorder struct {
	mock *MockOrderHanlderInterface
}

// NewMockOrderHanlderInterface creates a new mock instance.
func NewMockOrderHanlderInterface(ctrl *gomock.Controller) *MockOrderHanlderInterface {
	mock := &MockOrderHanlderInterface{ctrl: ctrl}
	mock.recorder = &MockOrderHanlderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderHanlderInterface) EXPECT() *MockOrderHanlderInterfaceMockRecorder {
	return m.recorder
}

// OrderHandler mocks base method.
func (m *MockOrderHanlderInterface) OrderHandler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OrderHandler", w, r)
}

// OrderHandler indicates an expected call of OrderHandler.
func (mr *MockOrderHanlderInterfaceMockRecorder) OrderHandler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderHandler", reflect.TypeOf((*MockOrderHanlderInterface)(nil).OrderHandler), w, r)
}