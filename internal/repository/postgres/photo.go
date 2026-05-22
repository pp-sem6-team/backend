package postgres

import (
	"context"

	"github.com/google/uuid"
	dbutil "github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) Create(ctx context.Context, photo *model.Photo) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).Create(photo).Error,
	)
}

func (r *photoRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Photo, error) {
	var photo model.Photo

	if err := dbutil.MapError(
		r.db.WithContext(ctx).First(&photo, "id = ?", id).Error,
	); err != nil {
		return nil, err
	}

	return &photo, nil
}

func (r *photoRepository) ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*model.Photo, int64, error) {
	var photos []*model.Photo
	var total int64

	db := r.db.WithContext(ctx).
		Model(&model.Photo{}).
		Where("user_id = ?", userID)

	if err := dbutil.MapError(db.Count(&total).Error); err != nil {
		return nil, 0, err
	}

	if err := dbutil.MapError(
		db.Order("uploaded_at ASC, id ASC").
			Offset(offset).
			Limit(limit).
			Find(&photos).Error,
	); err != nil {
		return nil, 0, err
	}

	return photos, total, nil
}

func (r *photoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Delete(&model.Photo{}, "id = ?", id).Error,
	)
}
