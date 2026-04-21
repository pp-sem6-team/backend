package dto

import (
	"time"

	"github.com/pp-sem6-team/backend/internal/domain"
)

type RegisterRequest struct {
	Email     string         `json:"email" binding:"required,email" example:"user@example.com"`
	Password  string         `json:"password" binding:"required"`
	Name      string         `json:"name" binding:"required"`
	BirthDate *time.Time     `json:"birth_date,omitempty" example:"2000-01-01T00:00:00Z"`
	Gender    *domain.Gender `json:"gender,omitempty" binding:"omitempty,oneof=male female"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LogoutRequest = RefreshRequest

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}
