package ml

import (
	"github.com/pp-sem6-team/backend/internal/domain"
	"gorm.io/datatypes"
)

type Result struct {
	SkinType domain.SkinType `json:"skin_type"`
	Data     datatypes.JSON  `json:"analysis_data"`
}
