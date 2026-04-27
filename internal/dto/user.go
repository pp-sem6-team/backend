package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type UserRequest struct {
	Email     *string        `json:"email,omitempty" binding:"omitempty,email" example:"user@example.com"`
	Name      *string        `json:"name,omitempty"`
	BirthDate *time.Time     `json:"birth_date,omitempty" example:"2000-01-01T00:00:00Z"`
	Gender    *domain.Gender `json:"gender,omitempty" binding:"omitempty,oneof=male female"`
}

type UserPasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	BirthDate *time.Time     `json:"birth_date"`
	Gender    *domain.Gender `json:"gender"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
