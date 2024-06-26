// Code generated by MockGen. DO NOT EDIT.
// Source: ./services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBillingServiceInterface is a mock of BillingServiceInterface interface.
type MockBillingServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBillingServiceInterfaceMockRecorder
}

// MockBillingServiceInterfaceMockRecorder is the mock recorder for MockBillingServiceInterface.
type MockBillingServiceInterfaceMockRecorder struct {
	mock *MockBillingServiceInterface
}

// NewMockBillingServiceInterface creates a new mock instance.
func NewMockBillingServiceInterface(ctrl *gomock.Controller) *MockBillingServiceInterface {
	mock := &MockBillingServiceInterface{ctrl: ctrl}
	mock.recorder = &MockBillingServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBillingServiceInterface) EXPECT() *MockBillingServiceInterfaceMockRecorder {
	return m.recorder
}

// GetOutstanding mocks base method.
func (m *MockBillingServiceInterface) GetOutstanding(loanID uint) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutstanding", loanID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutstanding indicates an expected call of GetOutstanding.
func (mr *MockBillingServiceInterfaceMockRecorder) GetOutstanding(loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutstanding", reflect.TypeOf((*MockBillingServiceInterface)(nil).GetOutstanding), loanID)
}

// IsDelinquent mocks base method.
func (m *MockBillingServiceInterface) IsDelinquent(loanID uint, weekNumber int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDelinquent", loanID, weekNumber)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDelinquent indicates an expected call of IsDelinquent.
func (mr *MockBillingServiceInterfaceMockRecorder) IsDelinquent(loanID, weekNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDelinquent", reflect.TypeOf((*MockBillingServiceInterface)(nil).IsDelinquent), loanID, weekNumber)
}

// MakePayment mocks base method.
func (m *MockBillingServiceInterface) MakePayment(loanID uint, amount float64, week int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakePayment", loanID, amount, week)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakePayment indicates an expected call of MakePayment.
func (mr *MockBillingServiceInterfaceMockRecorder) MakePayment(loanID, amount, week interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakePayment", reflect.TypeOf((*MockBillingServiceInterface)(nil).MakePayment), loanID, amount, week)
}
