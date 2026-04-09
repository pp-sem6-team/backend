package ml

import "github.com/pp-sem6-team/backend/internal/domain"

type Result struct {
	SkinType domain.SkinType `json:"skin_type"`
	Data     map[string]any  `json:"analysis_data"`
}
