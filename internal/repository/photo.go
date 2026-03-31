package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
)

type PhotoRepository interface {
	Create(ctx context.Context, photo *model.Photo) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Photo, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*model.Photo, int64, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
