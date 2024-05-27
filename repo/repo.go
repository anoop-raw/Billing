package repo

import (
	"math"

	"github.com/anoop-raw/Billing/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	GetLoanByID(loanID uint) (*models.Loan, error)
	UpdatePayment(payment *models.Payment) error
	GetPaymentsByLoanID(loanID uint) ([]models.Payment, error)
	CreateLoan(loan *models.Loan) error
}

type SQLRepository struct {
	db *gorm.DB
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreateLoan(loan *models.Loan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(loan).Error; err != nil {
			return err
		}

		payments := make([]models.Payment, loan.TermWeeks)
		// Convert weeks to years
		years := float64(loan.TermWeeks) / 52

		// Calculate the total amount with simple interest
		totalAmount := loan.Amount + (loan.Amount*loan.InterestRate*years)/100

		// Calculate the weekly amount
		weeklyAmount := totalAmount / float64(loan.TermWeeks)
		weeklyAmount = math.Round(weeklyAmount*100) / 100
		for i := 0; i < loan.TermWeeks; i++ {
			payments[i] = models.Payment{
				LoanID: loan.ID,
				Week:   i + 1,
				Amount: weeklyAmount,
			}
		}

		if err := tx.Create(&payments).Error; err != nil {
			return err
		}

		return nil
	})
}

// Payment repository methods
func (r *SQLRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *SQLRepository) GetLoanByID(loanID uint) (*models.Loan, error) {
	var loan models.Loan
	if err := r.db.First(&loan, loanID).Error; err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *SQLRepository) GetPaymentsByLoanID(loanID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Where("loan_id = ?", loanID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
