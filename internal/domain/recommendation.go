package domain

import "github.com/google/uuid"

type Recommendation struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
