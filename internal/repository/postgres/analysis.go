package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type analysisRepository struct {
	db *gorm.DB
}

func NewAnalysisRepository(db *gorm.DB) repository.AnalysisRepository {
	return &analysisRepository{db: db}
}

func (r *analysisRepository) Create(ctx context.Context, analysis *model.Analysis) error {
	return r.db.WithContext(ctx).Create(analysis).Error
}

func (r *analysisRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Analysis, error) {
	var analysis model.Analysis
	if err := r.db.WithContext(ctx).First(&analysis, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *analysisRepository) GetByPhotoID(ctx context.Context, photoID uuid.UUID) (*model.Analysis, error) {
	var analysis model.Analysis
	if err := r.db.WithContext(ctx).First(&analysis, "photo_id = ?", photoID).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *analysisRepository) ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*model.Analysis, int64, error) {
	var analyses []*model.Analysis
	var total int64
	db := r.db.WithContext(ctx).
		Model(&model.Analysis{}).
		Joins("JOIN photos ON photos.id = analyses.photo_id").
		Where("photos.user_id = ?", userID)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("created_at ASC, id ASC").Offset(offset).Limit(limit).Find(&analyses).Error; err != nil {
		return nil, 0, err
	}
	return analyses, total, nil
}

func (r *analysisRepository) Update(ctx context.Context, analysis *model.Analysis) error {
	return r.db.WithContext(ctx).Model(&model.Analysis{}).
		Where("id = ?", analysis.ID).Updates(map[string]any{
		"photo_id":      analysis.PhotoID,
		"status":        analysis.Status,
		"skin_type":     analysis.SkinType,
		"analysis_data": analysis.AnalysisData,
	}).Error
}

func (r *analysisRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status domain.AnalysisStatus) error {
	return r.db.WithContext(ctx).Model(&model.Analysis{}).Where("id = ?", id).Update("status", status).Error
}

func (r *analysisRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Analysis{}, "id = ?", id).Error
}
