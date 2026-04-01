package domain

import (
	"time"
)

// Type values: ENROLL_SUCCESS, ENROLL_FAIL, ORDER_EXPIRE, ACTIVITY_REMINDER
type Notification struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(200);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Type      string    `gorm:"type:varchar(30);not null" json:"type"`
	RelatedID *uint64   `json:"related_id"`
	IsRead    bool      `gorm:"not null;default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
