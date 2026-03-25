package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) repository.IngredientRepository {
	return &ingredientRepository{db: db}
}

func (r *ingredientRepository) Create(ctx context.Context, ingredient *models.Ingredient) error {
	return r.db.WithContext(ctx).Create(ingredient).Error
}

func (r *ingredientRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	if err := r.db.WithContext(ctx).First(&ingredient, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *ingredientRepository) GetByName(ctx context.Context, name string) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	if err := r.db.WithContext(ctx).First(&ingredient, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *ingredientRepository) List(ctx context.Context, offset int, limit int) ([]*models.Ingredient, int64, error) {
	var ingredients []*models.Ingredient
	var total int64
	db := r.db.WithContext(ctx).Model(&models.Ingredient{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("name ASC, id ASC").Offset(offset).Limit(limit).Find(&ingredients).Error; err != nil {
		return nil, 0, err
	}
	return ingredients, total, nil
}

func (r *ingredientRepository) Update(ctx context.Context, ingredient *models.Ingredient) error {
	return r.db.WithContext(ctx).Model(&models.Ingredient{}).
		Where("id = ?", ingredient.ID).Updates(map[string]any{
		"name":        ingredient.Name,
		"description": ingredient.Description,
	}).Error
}

func (r *ingredientRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Ingredient{}, "id = ?", id).Error
}
