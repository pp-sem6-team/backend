package postgres

import (
	"context"

	"github.com/google/uuid"
	dbutil "github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) repository.RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) Create(ctx context.Context, token *model.RefreshToken) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).Create(token).Error,
	)
}

func (r *refreshTokenRepository) GetByTokenHash(ctx context.Context, tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken

	if err := dbutil.MapError(
		r.db.WithContext(ctx).First(&token, "token_hash = ?", tokenHash).Error,
	); err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *refreshTokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Delete(&model.RefreshToken{}, "id = ?", id).Error,
	)
}

func (r *refreshTokenRepository) DeleteByUserIDAndToken(ctx context.Context, userID uuid.UUID, tokenHash string) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Where("user_id = ? AND token_hash = ?", userID, tokenHash).
			Delete(&model.RefreshToken{}).Error,
	)
}

func (r *refreshTokenRepository) DeleteExpired(ctx context.Context) error {
	return dbutil.MapError(
		r.db.WithContext(ctx).
			Where("expires_at < NOW()").
			Delete(&model.RefreshToken{}).Error,
	)
}
