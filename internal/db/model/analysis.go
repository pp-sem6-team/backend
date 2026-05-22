package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
	"gorm.io/datatypes"
)

type Analysis struct {
	ID           uuid.UUID             `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PhotoID      uuid.UUID             `gorm:"type:uuid;not null"`
	Status       domain.AnalysisStatus `gorm:"type:analysis_status;not null;default:'processing'"`
	SkinType     *domain.SkinType      `gorm:"type:skin_type"`
	AnalysisData datatypes.JSON        `gorm:"type:jsonb;not null;default:'{}'"`
	CreatedAt    time.Time             `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt    time.Time             `gorm:"type:timestamptz;not null;autoUpdateTime"`

	Photo *Photo `gorm:"foreignKey:PhotoID"`
}
