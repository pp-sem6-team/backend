package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) Create(ctx context.Context, photo *models.Photo) error {
	return r.db.WithContext(ctx).Create(photo).Error
}

func (r *photoRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Photo, error) {
	var photo models.Photo
	if err := r.db.WithContext(ctx).First(&photo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}

func (r *photoRepository) ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*models.Photo, int64, error) {
	var photos []*models.Photo
	var total int64
	db := r.db.WithContext(ctx).Model(&models.Photo{}).Where("user_id = ?", userID)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("uploaded_at ASC, id ASC").Offset(offset).Limit(limit).Find(&photos).Error; err != nil {
		return nil, 0, err
	}
	return photos, total, nil
}

func (r *photoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Photo{}, "id = ?", id).Error
}
