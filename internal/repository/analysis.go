package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
)

type AnalysisRepository interface {
	Create(ctx context.Context, analysis *models.Analysis) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Analysis, error)
	GetByPhotoID(ctx context.Context, photoID uuid.UUID) (*models.Analysis, error)
	ListByUserID(ctx context.Context, id uuid.UUID) ([]*models.Analysis, error)
	Update(ctx context.Context, analysis *models.Analysis) error
	Delete(ctx context.Context, id uuid.UUID) error
}
