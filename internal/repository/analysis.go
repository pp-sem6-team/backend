package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
)

type AnalysisRepository interface {
	Create(ctx context.Context, analysis *model.Analysis) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Analysis, error)
	GetByPhotoID(ctx context.Context, photoID uuid.UUID) (*model.Analysis, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*model.Analysis, int64, error)
	Update(ctx context.Context, analysis *model.Analysis) error
	Delete(ctx context.Context, id uuid.UUID) error
}
