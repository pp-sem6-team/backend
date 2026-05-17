package ml

import (
	"context"
)

type Client interface {
	Health(ctx context.Context) error
	Analyze(ctx context.Context, fileName string, data []byte) (*Result, error)
}
