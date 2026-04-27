package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type CreateAnalysisResponse struct {
	ID         uuid.UUID             `json:"id"`
	PhotoID    uuid.UUID             `json:"photo_id"`
	FileURL    string                `json:"file_url"`
	Status     domain.AnalysisStatus `json:"status"`
	UploadedAt time.Time             `json:"uploaded_at"`
}

type AnalysisListItemResponse struct {
	ID        uuid.UUID             `json:"id"`
	PhotoID   uuid.UUID             `json:"photo_id"`
	FileURL   string                `json:"file_url"`
	Status    domain.AnalysisStatus `json:"status"`
	SkinType  *domain.SkinType      `json:"skin_type"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type AnalysisDetailResponse struct {
	ID              uuid.UUID               `json:"id"`
	PhotoID         uuid.UUID               `json:"photo_id"`
	FileURL         string                  `json:"file_url"`
	Status          domain.AnalysisStatus   `json:"status"`
	SkinType        *domain.SkinType        `json:"skin_type"`
	AnalysisData    map[string]any          `json:"analysis_data"`
	Recommendations []domain.Recommendation `json:"recommendations,omitempty"`
	Ingredients     []domain.Ingredient     `json:"ingredients,omitempty"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}
