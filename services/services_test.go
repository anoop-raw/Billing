package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/anoop-raw/Billing/mockgen/repo"
	"github.com/anoop-raw/Billing/models"
	"github.com/golang/mock/gomock"
)

func TestBillingService_CreateLoan(t *testing.T) {
	tests := []struct {
		name                    string
		amount                  float64
		interestRate            float64
		weeks                   int
		mockCreateLoanOutput    error
		mockCreatePaymentOutput error
		want                    *models.Loan
		wantErr                 bool
	}{
		{
			name:                    "Success",
			amount:                  5000000,
			interestRate:            0.1,
			weeks:                   50,
			mockCreateLoanOutput:    nil,
			mockCreatePaymentOutput: nil,
			want: &models.Loan{
				Amount:       5000000,
				InterestRate: 0.1,
				TermWeeks:    50,
			},
			wantErr: false,
		},
		{
			name:                    "Failure",
			amount:                  5000000,
			interestRate:            0.1,
			weeks:                   50,
			mockCreateLoanOutput:    errors.New("error"),
			mockCreatePaymentOutput: nil,
			want:                    nil,
			wantErr:                 true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock_repo.NewMockLoanRepository(ctrl)
			mockRepo.EXPECT().CreateLoan(gomock.Any()).Return(tt.mockCreateLoanOutput).Times(1)

			svc := NewBillingService(mockRepo)

			got, err := svc.CreateLoan(tt.amount, tt.interestRate, tt.weeks)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				tt.want.ID = got.ID
				tt.want.CreatedAt = got.CreatedAt
				tt.want.UpdatedAt = got.UpdatedAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLoan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBillingService_GetLoan(t *testing.T) {
	tests := []struct {
		name               string
		loanID             uint
		mockLoanOutput     *models.Loan
		mockPaymentsOutput []models.Payment
		mockLoanErr        error
		mockPaymentsErr    error
		want               *models.Loan
		wantPayments       []models.Payment
		wantErr            bool
	}{
		{
			name:   "Success",
			loanID: 1,
			mockLoanOutput: &models.Loan{
				ID:           1,
				Amount:       5000000,
				InterestRate: 0.1,
				TermWeeks:    50,
			},
			mockPaymentsOutput: []models.Payment{
				{ID: 1, LoanID: 1, Week: 1, Amount: 110000},
				{ID: 2, LoanID: 1, Week: 2, Amount: 110000},
			},
			mockLoanErr:     nil,
			mockPaymentsErr: nil,
			want: &models.Loan{
				ID:           1,
				Amount:       5000000,
				InterestRate: 0.1,
				TermWeeks:    50,
			},
			wantPayments: []models.Payment{
				{ID: 1, LoanID: 1, Week: 1, Amount: 110000},
				{ID: 2, LoanID: 1, Week: 2, Amount: 110000},
			},
			wantErr: false,
		},
		{
			name:               "Loan Fetch Failure",
			loanID:             2,
			mockLoanOutput:     nil,
			mockPaymentsOutput: nil,
			mockLoanErr:        errors.New("error"),
			mockPaymentsErr:    nil,
			want:               nil,
			wantPayments:       nil,
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock_repo.NewMockLoanRepository(ctrl)
			mockRepo.EXPECT().GetLoanByID(tt.loanID).Return(tt.mockLoanOutput, tt.mockLoanErr).Times(1)

			if tt.mockLoanErr == nil {
				mockRepo.EXPECT().GetPaymentsByLoanID(tt.loanID).Return(tt.mockPaymentsOutput, tt.mockPaymentsErr).Times(1)
			}

			svc := NewBillingService(mockRepo)

			got, gotPayments, err := svc.GetLoan(tt.loanID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoan() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(gotPayments, tt.wantPayments) {
				t.Errorf("GetLoan() gotPayments = %v, want %v", gotPayments, tt.wantPayments)
			}
		})
	}
}
