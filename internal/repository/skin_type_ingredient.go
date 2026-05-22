package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type SkinTypeIngredientRepository interface {
	Add(ctx context.Context, skinType domain.SkinType, ingredientID uuid.UUID) error
	Delete(ctx context.Context, skinType domain.SkinType, ingredientID uuid.UUID) error
	ListBySkinType(ctx context.Context, skinType domain.SkinType, offset int, limit int) ([]*model.SkinTypeIngredient, int64, error)
}
