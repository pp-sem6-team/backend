package ml

import (
	"context"
)

type Client interface {
	Analyze(ctx context.Context, objectKey string) (*Result, error)
}
