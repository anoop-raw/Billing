package models

import "time"

type Loan struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Amount       float64   `json:"amount" gorm:"not null"`
	InterestRate float64   `json:"interest_rate" gorm:"not null"`
	TermWeeks    int       `json:"term_weeks" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

type Payment struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	LoanID    uint       `json:"loan_id" gorm:"not null;index"`
	Week      int        `json:"week" gorm:"not null"`
	Amount    float64    `json:"amount" gorm:"not null"`
	PaidAt    *time.Time `json:"paid_at"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
