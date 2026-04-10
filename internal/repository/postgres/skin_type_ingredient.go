package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type skinTypeIngredientRepository struct {
	db *gorm.DB
}

func NewSkinTypeIngredientRepository(db *gorm.DB) repository.SkinTypeIngredientRepository {
	return &skinTypeIngredientRepository{db: db}
}

func (r *skinTypeIngredientRepository) Add(
	ctx context.Context,
	skinType domain.SkinType,
	ingredientID uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Create(&model.SkinTypeIngredient{SkinType: skinType, IngredientID: ingredientID}).Error
}

func (r *skinTypeIngredientRepository) Delete(
	ctx context.Context,
	skinType domain.SkinType,
	ingredientID uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Delete(&model.SkinTypeIngredient{}, "skin_type = ? AND ingredient_id = ?", skinType, ingredientID).Error
}

func (r *skinTypeIngredientRepository) ListBySkinType(
	ctx context.Context,
	skinType domain.SkinType,
	offset int,
	limit int,
) ([]*model.SkinTypeIngredient, int64, error) {
	var skinTypeIngredients []*model.SkinTypeIngredient
	var total int64
	db := r.db.WithContext(ctx).Model(&model.SkinTypeIngredient{}).Where("skin_type = ?", skinType)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("ingredient_id ASC").Offset(offset).Limit(limit).Find(&skinTypeIngredients).Error; err != nil {
		return nil, 0, err
	}
	return skinTypeIngredients, total, nil
}
