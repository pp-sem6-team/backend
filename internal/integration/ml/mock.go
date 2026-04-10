package ml

import (
	"context"
	"encoding/json"

	"github.com/pp-sem6-team/backend/internal/domain"
	"gorm.io/datatypes"
)

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) Analyze(ctx context.Context, objectKey string) (*Result, error) {
	stats := map[string]any{
		"accuracy":        0.82,
		"processing_time": 1.37,
	}

	bytes, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}

	return &Result{
		SkinType: domain.Oily,
		Data:     datatypes.JSON(bytes),
	}, nil
}
