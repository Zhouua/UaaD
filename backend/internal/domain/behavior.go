package domain

import (
	"time"
)

// BehaviorType values: VIEW, COLLECT, SHARE, CLICK, SEARCH
type UserBehavior struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	UserID       uint64    `gorm:"not null" json:"user_id"`
	ActivityID   uint64    `gorm:"not null" json:"activity_id"`
	BehaviorType string    `gorm:"type:varchar(20);not null" json:"behavior_type"`
	Detail       *string   `gorm:"type:json" json:"detail"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
}
