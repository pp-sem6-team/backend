package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type RecommendationRepository interface {
	Create(ctx context.Context, recommendation *model.Recommendation) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Recommendation, error)
	ListBySkinType(ctx context.Context, skinType domain.SkinType, offset int, limit int) ([]*model.Recommendation, int64, error)
	List(ctx context.Context, offset int, limit int) ([]*model.Recommendation, int64, error)
	Update(ctx context.Context, recommendation *model.Recommendation) error
	Delete(ctx context.Context, id uuid.UUID) error
}
