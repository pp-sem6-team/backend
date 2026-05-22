package domain

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type AnalysisStatus string

const (
	Processing AnalysisStatus = "processing"
	Completed  AnalysisStatus = "completed"
	Failed     AnalysisStatus = "failed"
)

type SkinType string

const (
	Oily        SkinType = "oily"
	Dry         SkinType = "dry"
	Normal      SkinType = "normal"
	Combination SkinType = "combination"
)
