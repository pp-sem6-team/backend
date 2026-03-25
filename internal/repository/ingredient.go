package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
)

type IngredientRepository interface {
	Create(ctx context.Context, ingredient *models.Ingredient) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Ingredient, error)
	GetByName(ctx context.Context, name string) (*models.Ingredient, error)
	List(ctx context.Context, offset int, limit int) ([]*models.Ingredient, int64, error)
	Update(ctx context.Context, ingredient *models.Ingredient) error
	Delete(ctx context.Context, id uuid.UUID) error
}
