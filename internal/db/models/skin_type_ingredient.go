package models

import (
	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
)

type SkinTypeIngredient struct {
	SkinType     domain.SkinType `gorm:"type:skin_type;not null;primaryKey"`
	IngredientID uuid.UUID       `gorm:"type:uuid;not null;primaryKey"`

	Ingredient *Ingredient `gorm:"foreignKey:IngredientID"`
}
