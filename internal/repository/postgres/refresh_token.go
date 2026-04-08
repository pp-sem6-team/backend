package postgres

import (
	"context"

	"github.com/google/uuid"
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
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *refreshTokenRepository) GetByTokenHash(ctx context.Context, tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	if err := r.db.WithContext(ctx).First(&token, "token_hash = ?", tokenHash).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *refreshTokenRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.RefreshToken{}, "id = ?", id).Error
}

func (r *refreshTokenRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.RefreshToken{}, "user_id = ?", userID).Error
}

func (r *refreshTokenRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < NOW()").
		Delete(&model.RefreshToken{}).Error
}
