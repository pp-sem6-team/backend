package ml

import (
	"context"
	"encoding/json"

	"github.com/pp-sem6-team/backend/internal/domain"
	"gorm.io/datatypes"
)

type MockClient struct{}

func NewMockClient() Client {
	return &MockClient{}
}

func (c *MockClient) Health(ctx context.Context) error {
	return nil
}

func (c *MockClient) Analyze(
	ctx context.Context,
	fileName string,
	data []byte,
) (*Result, error) {
	mockResponse := map[string]any{
		"model_version": "mock-1.0.0",
		"confidence":    0.82,
		"probabilities": map[string]float64{
			"combination": 0.05,
			"dry":         0.03,
			"normal":      0.10,
			"oily":        0.82,
		},
		"warnings": []string{},
	}

	rawJSON, err := json.Marshal(mockResponse)
	if err != nil {
		return nil, err
	}

	return &Result{
		SkinType: domain.Oily,
		Data:     datatypes.JSON(rawJSON),
	}, nil
}
