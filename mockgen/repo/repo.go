// Code generated by MockGen. DO NOT EDIT.
// Source: ./repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	reflect "reflect"

	models "github.com/anoop-raw/Billing/models"
	gomock "github.com/golang/mock/gomock"
)

// MockLoanRepository is a mock of LoanRepository interface.
type MockLoanRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoanRepositoryMockRecorder
}

// MockLoanRepositoryMockRecorder is the mock recorder for MockLoanRepository.
type MockLoanRepositoryMockRecorder struct {
	mock *MockLoanRepository
}

// NewMockLoanRepository creates a new mock instance.
func NewMockLoanRepository(ctrl *gomock.Controller) *MockLoanRepository {
	mock := &MockLoanRepository{ctrl: ctrl}
	mock.recorder = &MockLoanRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoanRepository) EXPECT() *MockLoanRepositoryMockRecorder {
	return m.recorder
}

// CreateLoan mocks base method.
func (m *MockLoanRepository) CreateLoan(loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLoan indicates an expected call of CreateLoan.
func (mr *MockLoanRepositoryMockRecorder) CreateLoan(loan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockLoanRepository)(nil).CreateLoan), loan)
}

// GetLoanByID mocks base method.
func (m *MockLoanRepository) GetLoanByID(loanID uint) (*models.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByID", loanID)
	ret0, _ := ret[0].(*models.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByID indicates an expected call of GetLoanByID.
func (mr *MockLoanRepositoryMockRecorder) GetLoanByID(loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByID", reflect.TypeOf((*MockLoanRepository)(nil).GetLoanByID), loanID)
}

// GetPaymentsByLoanID mocks base method.
func (m *MockLoanRepository) GetPaymentsByLoanID(loanID uint) ([]models.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentsByLoanID", loanID)
	ret0, _ := ret[0].([]models.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentsByLoanID indicates an expected call of GetPaymentsByLoanID.
func (mr *MockLoanRepositoryMockRecorder) GetPaymentsByLoanID(loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentsByLoanID", reflect.TypeOf((*MockLoanRepository)(nil).GetPaymentsByLoanID), loanID)
}

// UpdatePayment mocks base method.
func (m *MockLoanRepository) UpdatePayment(payment *models.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", payment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePayment indicates an expected call of UpdatePayment.
func (mr *MockLoanRepositoryMockRecorder) UpdatePayment(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockLoanRepository)(nil).UpdatePayment), payment)
}
