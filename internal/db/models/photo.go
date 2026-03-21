package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	ObjectKey  string    `gorm:"type:varchar(1024);not null"`
	UploadedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`

	User     User      `gorm:"foreignKey:UserID"`
	Analysis *Analysis `gorm:"foreignKey:PhotoID"`
}
