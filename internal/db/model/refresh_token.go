package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	TokenHash string    `gorm:"type:text;not null;unique"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
	ExpiresAt time.Time `gorm:"type:timestamptz;not null"`

	User *User `gorm:"foreignKey:UserID"`
}
