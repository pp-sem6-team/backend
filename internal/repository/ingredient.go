package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
)

type IngredientRepository interface {
	Create(ctx context.Context, ingredient *model.Ingredient) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Ingredient, error)
	GetByName(ctx context.Context, name string) (*model.Ingredient, error)
	List(ctx context.Context, offset int, limit int) ([]*model.Ingredient, int64, error)
	Update(ctx context.Context, ingredient *model.Ingredient) error
	Delete(ctx context.Context, id uuid.UUID) error
}
