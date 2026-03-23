package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type RecommendationRepository interface {
	Create(ctx context.Context, recommendation *models.Recommendation) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Recommendation, error)
	ListBySkinType(ctx context.Context, skinType domain.SkinType, offset int, limit int) ([]*models.Recommendation, error)
	List(ctx context.Context, offset int, limit int) ([]*models.Recommendation, error)
	Update(ctx context.Context, recommendation *models.Recommendation) error
	Delete(ctx context.Context, id uuid.UUID) error
}
