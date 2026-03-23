package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/models"
)

type PhotoRepository interface {
	Create(ctx context.Context, photo *models.Photo) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Photo, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, offset int, limit int) ([]*models.Photo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
