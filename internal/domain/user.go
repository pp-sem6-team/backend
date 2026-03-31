package domain

import (
	"time"
)

type UpdateUserInput struct {
	Email     *string
	Name      *string
	BirthDate *time.Time
	Gender    *Gender
}
