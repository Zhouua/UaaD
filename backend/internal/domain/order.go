package domain

import (
	"time"
)

// Status values: PENDING, PAID, CLOSED, REFUNDED
type Order struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	OrderNo      string     `gorm:"type:varchar(32);not null;uniqueIndex" json:"order_no"`
	EnrollmentID uint64     `gorm:"not null;uniqueIndex" json:"enrollment_id"`
	UserID       uint64     `gorm:"not null" json:"user_id"`
	ActivityID   uint64     `gorm:"not null" json:"activity_id"`
	Amount       float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status       string     `gorm:"type:varchar(20);not null;default:'PENDING'" json:"status"`
	PaidAt       *time.Time `json:"paid_at"`
	ExpiredAt    time.Time  `gorm:"not null" json:"expired_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
