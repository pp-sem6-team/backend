package domain

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserInput struct {
	Email     *string
	Name      *string
	BirthDate *time.Time
	Gender    *Gender
}

type User struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birth_date"`
	Gender    *Gender    `json:"gender"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
