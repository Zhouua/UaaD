package repository

import (
	"github.com/uaad/backend/internal/domain"
	"gorm.io/gorm"
)

// BehaviorRepository defines data access methods for user behaviors.
type BehaviorRepository interface {
	// Create inserts a single behavior record.
	Create(behavior *domain.UserBehavior) error
	// BatchCreate inserts multiple behavior records in one transaction.
	BatchCreate(behaviors []*domain.UserBehavior) error
	// ListByUserID returns recent behaviors for a given user.
	ListByUserID(userID uint64, limit int) ([]domain.UserBehavior, error)
	// CountByActivityAndType counts behaviors matching activity and type.
	CountByActivityAndType(activityID uint64, behaviorType string) (int64, error)
}

type behaviorRepository struct {
	db *gorm.DB
}

func NewBehaviorRepository(db *gorm.DB) BehaviorRepository {
	return &behaviorRepository{db: db}
}

func (r *behaviorRepository) Create(behavior *domain.UserBehavior) error {
	return r.db.Create(behavior).Error
}

func (r *behaviorRepository) BatchCreate(behaviors []*domain.UserBehavior) error {
	if len(behaviors) == 0 {
		return nil
	}
	return r.db.CreateInBatches(behaviors, len(behaviors)).Error
}

func (r *behaviorRepository) ListByUserID(userID uint64, limit int) ([]domain.UserBehavior, error) {
	var behaviors []domain.UserBehavior
	if err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&behaviors).Error; err != nil {
		return nil, err
	}
	return behaviors, nil
}

func (r *behaviorRepository) CountByActivityAndType(activityID uint64, behaviorType string) (int64, error) {
	var count int64
	if err := r.db.Model(&domain.UserBehavior{}).
		Where("activity_id = ? AND behavior_type = ?", activityID, behaviorType).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
