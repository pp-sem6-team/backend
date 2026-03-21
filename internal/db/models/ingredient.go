package models

import (
	"github.com/google/uuid"
)

type Ingredient struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(255);unique;not null"`
	Description string    `gorm:"type:text"`

	SkinTypeIngredients []SkinTypeIngredient `gorm:"foreignKey:IngredientID"`
}
