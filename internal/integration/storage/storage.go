package storage

import (
	"context"
	"io"
	"time"
)

type Storage interface {
	Upload(ctx context.Context, objectKey string, r io.Reader, size int64, contentType string) error
	GetPresignedURL(ctx context.Context, objectKey string, expires time.Duration) (string, error)
	Delete(ctx context.Context, objectKey string) error
}
