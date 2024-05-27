package services

import (
	"errors"
	"time"

	"github.com/anoop-raw/Billing/models"
	"github.com/anoop-raw/Billing/repo"
)

type BillingServiceInterface interface {
	GetOutstanding(loanID uint) (float64, error)
	IsDelinquent(loanID uint, weekNumber int64) (bool, error)
	MakePayment(loanID uint, amount float64, week int) error
}

type BillingService struct {
	loanRepo repo.LoanRepository
}

func NewBillingService(loanRepo repo.LoanRepository) *BillingService {
	return &BillingService{loanRepo: loanRepo}
}

func (s *BillingService) CreateLoan(amount float64, interestRate float64, weeks int) (*models.Loan, error) {
	loan := &models.Loan{
		Amount:       amount,
		InterestRate: interestRate,
		TermWeeks:    weeks,
	}
	if err := s.loanRepo.CreateLoan(loan); err != nil {
		return nil, err
	}
	return loan, nil
}

func (s *BillingService) GetOutstanding(loanID uint) (float64, error) {
	payments, err := s.loanRepo.GetPaymentsByLoanID(loanID)
	if err != nil {
		return 0, err
	}

	totalPaid := 0.0
	totalAmount := 0.0
	for _, payment := range payments {
		totalAmount += payment.Amount
		if payment.PaidAt != nil {
			totalPaid += payment.Amount
		}
	}

	return totalAmount - totalPaid, nil
}

// weekNumber is current week till customer has to pay the installment.
func (s *BillingService) IsDelinquent(loanID uint, weekNumber int64) (bool, error) {
	payments, err := s.loanRepo.GetPaymentsByLoanID(loanID)
	if err != nil {
		return false, err
	}

	missedPayments := 0
	lastIndx := -1
	for i := weekNumber - 1; i >= 0; i-- {
		if payments[i].PaidAt == nil {
			missedPayments++
			if lastIndx == -1 {
				lastIndx = int(i)
			} else if lastIndx-int(i) == 1 {
				return true, nil
			} else {
				lastIndx = int(i)
			}
		} else {
			break
		}
	}

	return false, nil
}

func (s *BillingService) MakePayment(loanID uint, amount float64, week int) error {
	payments, err := s.loanRepo.GetPaymentsByLoanID(loanID)
	if err != nil {
		return err
	}

	for i, payment := range payments {
		if payment.Week == week {
			if payment.PaidAt != nil {
				return errors.New("payment already made for this week")
			}

			if payment.Amount != amount {
				return errors.New("incorrect amount ")
			}

			now := time.Now()
			payments[i].PaidAt = &now
			return s.loanRepo.UpdatePayment(&payments[i])
		}
	}

	return errors.New("payment not found for the given week")
}
