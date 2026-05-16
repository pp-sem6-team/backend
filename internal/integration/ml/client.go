package ml

import (
	"context"
)

type Client interface {
	Analyze(ctx context.Context, fileName string, data []byte) (*Result, error)
}
