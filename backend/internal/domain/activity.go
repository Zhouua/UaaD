package domain

import (
	"time"

	"gorm.io/gorm"
)

// Status values: DRAFT, PREHEAT, PUBLISHED, SELLING_OUT, SOLD_OUT, OFFLINE, CANCELLED
// Category values: CONCERT, CONFERENCE, EXPO, ESPORTS, EXHIBITION, OTHER
type Activity struct {
	ID            uint64         `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"type:varchar(200);not null" json:"title"`
	Description   string         `gorm:"type:text;not null" json:"description"`
	CoverURL      *string        `gorm:"type:varchar(500)" json:"cover_url"`
	Location      string         `gorm:"type:varchar(200);not null" json:"location"`
	Latitude      *float64       `gorm:"type:decimal(10,7)" json:"latitude"`
	Longitude     *float64       `gorm:"type:decimal(10,7)" json:"longitude"`
	Category      string         `gorm:"type:varchar(20);not null;default:'OTHER'" json:"category"`
	Tags          *string        `gorm:"type:json" json:"tags"`
	MaxCapacity   int            `gorm:"not null;default:0" json:"max_capacity"`
	EnrollOpenAt  time.Time      `gorm:"not null" json:"enroll_open_at"`
	EnrollCloseAt time.Time      `gorm:"not null" json:"enroll_close_at"`
	ActivityAt    time.Time      `gorm:"not null" json:"activity_at"`
	Price         float64        `gorm:"type:decimal(10,2);not null;default:0.00" json:"price"`
	Status        string         `gorm:"type:varchar(20);not null;default:'DRAFT'" json:"status"`
	CreatedBy     uint64         `gorm:"not null" json:"created_by"`
	ViewCount     int64          `gorm:"not null;default:0" json:"view_count"`
	EnrollCount   int64          `gorm:"not null;default:0" json:"enroll_count"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
