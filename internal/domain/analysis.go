package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AnalysisCreated struct {
	ID         uuid.UUID      `json:"id"`
	PhotoID    uuid.UUID      `json:"photo_id"`
	FileURL    string         `json:"file_url"`
	Status     AnalysisStatus `json:"status"`
	UploadedAt time.Time      `json:"uploaded_at"`
}

type AnalysisListItem struct {
	ID        uuid.UUID      `json:"id"`
	PhotoID   uuid.UUID      `json:"photo_id"`
	FileURL   string         `json:"file_url"`
	Status    AnalysisStatus `json:"status"`
	SkinType  *SkinType      `json:"skin_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type AnalysisDetail struct {
	ID              uuid.UUID        `json:"id"`
	PhotoID         uuid.UUID        `json:"photo_id"`
	FileURL         string           `json:"file_url"`
	Status          AnalysisStatus   `json:"status"`
	SkinType        *SkinType        `json:"skin_type"`
	AnalysisData    datatypes.JSON   `json:"analysis_data"`
	Recommendations []Recommendation `json:"recommendations,omitempty"`
	Ingredients     []Ingredient     `json:"ingredients,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}
