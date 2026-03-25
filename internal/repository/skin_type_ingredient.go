package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type SkinTypeIngredientRepository interface {
	Add(ctx context.Context, skinType domain.SkinType, ingredientID uuid.UUID) error
	Delete(ctx context.Context, skinType domain.SkinType, ingredientID uuid.UUID) error
	ListBySkinType(ctx context.Context, skinType domain.SkinType, offset int, limit int) ([]*models.SkinTypeIngredient, int64, error)
}
