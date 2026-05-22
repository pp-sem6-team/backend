package postgres

import (
	"context"

	"github.com/google/uuid"
	dbutil "github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type recommendationRepository struct {
	db *gorm.DB
}

func NewRecommendationRepository(db *gorm.DB) repository.RecommendationRepository {
	return &recommendationRepository{db: db}
}

func (r *recommendationRepository) Create(ctx context.Context, recommendation *model.Recommendation) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).Create(recommendation).Error,
	)
}

func (r *recommendationRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Recommendation, error) {
	var recommendation model.Recommendation

	if err := dbutil.MapError(
		r.db.WithContext(ctx).First(&recommendation, "id = ?", id).Error,
	); err != nil {
		return nil, err
	}

	return &recommendation, nil
}

func (r *recommendationRepository) ListBySkinType(
	ctx context.Context,
	skinType domain.SkinType,
	offset int,
	limit int,
) ([]*model.Recommendation, int64, error) {
	var recommendations []*model.Recommendation
	var total int64

	db := r.db.WithContext(ctx).
		Model(&model.Recommendation{}).
		Where("skin_type = ?", skinType)

	if err := dbutil.MapError(db.Count(&total).Error); err != nil {
		return nil, 0, err
	}

	if err := dbutil.MapError(
		db.Order("title ASC, id ASC").
			Offset(offset).
			Limit(limit).
			Find(&recommendations).Error,
	); err != nil {
		return nil, 0, err
	}

	return recommendations, total, nil
}

func (r *recommendationRepository) List(
	ctx context.Context,
	offset int,
	limit int,
) ([]*model.Recommendation, int64, error) {
	var recommendations []*model.Recommendation
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Recommendation{})

	if err := dbutil.MapError(db.Count(&total).Error); err != nil {
		return nil, 0, err
	}

	if err := dbutil.MapError(
		db.Order("title ASC, id ASC").
			Offset(offset).
			Limit(limit).
			Find(&recommendations).Error,
	); err != nil {
		return nil, 0, err
	}

	return recommendations, total, nil
}

func (r *recommendationRepository) Update(ctx context.Context, recommendation *model.Recommendation) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Model(&model.Recommendation{}).
			Where("id = ?", recommendation.ID).
			Updates(map[string]any{
				"skin_type":   recommendation.SkinType,
				"title":       recommendation.Title,
				"description": recommendation.Description,
			}).Error,
	)
}

func (r *recommendationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Delete(&model.Recommendation{}, "id = ?", id).Error,
	)
}
