package models

import (
	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type Recommendation struct {
	ID          uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SkinType    domain.SkinType `gorm:"type:skin_type;not null"`
	Title       string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text;not null"`
}
