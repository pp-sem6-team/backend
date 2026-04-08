package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *model.RefreshToken) error
	GetByTokenHash(ctx context.Context, tokenHash string) (*model.RefreshToken, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteExpired(ctx context.Context) error
}
