package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"github.com/pp-sem6-team/backend/internal/security"
	"github.com/pp-sem6-team/backend/internal/security/jwt"
)

type AuthService struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	jwtManager       *jwt.Manager
}

func NewAuthService(
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	jwtManager *jwt.Manager,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtManager:       jwtManager,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	email, password, name string,
	birth_date *time.Time,
	gender *domain.Gender,
) (string, string, error) {
	hash, err := security.HashPassword(password)
	if err != nil {
		return "", "", err
	}

	newUser := &model.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hash),
		Name:         name,
		BirthDate:    birth_date,
		Gender:       gender,
	}

	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return "", "", err
	}

	return s.generateTokens(ctx, newUser.ID)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if err := security.CheckPassword(password, user.PasswordHash); err != nil {
		return "", "", domain.ErrInvalidCreds
	}

	return s.generateTokens(ctx, user.ID)
}

func (s *AuthService) Logout(ctx context.Context, userID uuid.UUID) error {
	return s.refreshTokenRepo.DeleteByUserID(ctx, userID)
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	hash := security.HashToken(refreshToken)

	stored, err := s.refreshTokenRepo.GetByTokenHash(ctx, hash)
	if err != nil {
		return "", "", err
	}

	if time.Now().After(stored.ExpiresAt) {
		_ = s.refreshTokenRepo.Delete(ctx, stored.ID)
		return "", "", domain.ErrTokenExpired
	}

	if err := s.refreshTokenRepo.Delete(ctx, stored.ID); err != nil {
		return "", "", err
	}

	return s.generateTokens(ctx, stored.UserID)
}

func (s *AuthService) generateTokens(ctx context.Context, userID uuid.UUID) (string, string, error) {
	accessToken, _, err := s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, expiresAt, err := s.jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	hash := security.HashToken(refreshToken)

	token := &model.RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		TokenHash: hash,
		ExpiresAt: expiresAt,
	}

	if err := s.refreshTokenRepo.Create(ctx, token); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
