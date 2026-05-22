package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string         `gorm:"type:varchar(255);unique;not null"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	Name         string         `gorm:"type:varchar(255);not null"`
	BirthDate    *time.Time     `gorm:"type:date"`
	Gender       *domain.Gender `gorm:"type:gender"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;not null;autoUpdateTime"`

	Photos []Photo `gorm:"foreignKey:UserID"`
}
