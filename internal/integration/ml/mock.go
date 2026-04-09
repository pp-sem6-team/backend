package ml

import (
	"context"

	"github.com/pp-sem6-team/backend/internal/domain"
)

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) Analyze(ctx context.Context, objectKey string) (*Result, error) {
	return &Result{
		SkinType: domain.Oily,
		Data: map[string]any{
			"accuracy":        0.82,
			"processing_time": 1.37,
		},
	}, nil
}
