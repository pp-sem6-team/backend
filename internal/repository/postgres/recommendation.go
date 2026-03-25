package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
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

func (r *recommendationRepository) Create(ctx context.Context, recommendation *models.Recommendation) error {
	return r.db.WithContext(ctx).Create(recommendation).Error
}

func (r *recommendationRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Recommendation, error) {
	var recommendation models.Recommendation
	if err := r.db.WithContext(ctx).First(&recommendation, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &recommendation, nil
}

func (r *recommendationRepository) ListBySkinType(
	ctx context.Context,
	skinType domain.SkinType,
	offset int,
	limit int,
) ([]*models.Recommendation, int64, error) {
	var recommendations []*models.Recommendation
	var total int64
	db := r.db.WithContext(ctx).Model(&models.Recommendation{}).Where("skin_type = ?", skinType)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("title ASC, id ASC").Offset(offset).Limit(limit).Find(&recommendations).Error; err != nil {
		return nil, 0, err
	}
	return recommendations, total, nil
}

func (r *recommendationRepository) List(
	ctx context.Context,
	offset int,
	limit int,
) ([]*models.Recommendation, int64, error) {
	var recommendations []*models.Recommendation
	var total int64
	db := r.db.WithContext(ctx).Model(&models.Recommendation{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("title ASC, id ASC").Offset(offset).Limit(limit).Find(&recommendations).Error; err != nil {
		return nil, 0, err
	}
	return recommendations, total, nil
}

func (r *recommendationRepository) Update(ctx context.Context, recommendation *models.Recommendation) error {
	return r.db.WithContext(ctx).Model(&models.Recommendation{}).
		Where("id = ?", recommendation.ID).Updates(map[string]any{
		"skin_type":   recommendation.SkinType,
		"title":       recommendation.Title,
		"description": recommendation.Description,
	}).Error
}

func (r *recommendationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Recommendation{}, "id = ?", id).Error
}
