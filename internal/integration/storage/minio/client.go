package minio

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pp-sem6-team/backend/internal/config"
	"github.com/pp-sem6-team/backend/internal/integration/storage"
)

type Client struct {
	client *minio.Client
	bucket string
}

var _ storage.Storage = (*Client)(nil)

func New(cfg *config.Config) (*Client, error) {
	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.RootUser, cfg.Minio.RootPassword, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	exists, err := mc.BucketExists(
		context.Background(),
		cfg.Minio.Bucket,
	)

	if err != nil {
		return nil, err
	}

	if !exists {
		err = mc.MakeBucket(
			context.Background(),
			cfg.Minio.Bucket,
			minio.MakeBucketOptions{},
		)
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		client: mc,
		bucket: cfg.Minio.Bucket,
	}, nil
}

func (c *Client) Health(ctx context.Context) error {
	_, err := c.client.ListBuckets(ctx)
	return err
}

func (c *Client) Upload(
	ctx context.Context,
	objectKey string,
	r io.Reader,
	size int64,
	contentType string,
) error {

	_, err := c.client.PutObject(
		ctx,
		c.bucket,
		objectKey,
		r,
		size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)

	if err != nil {
		return ErrUploadFailed
	}

	return nil
}

func (c *Client) GetPresignedURL(
	ctx context.Context,
	objectKey string,
	expires time.Duration,
) (string, error) {

	url, err := c.client.PresignedGetObject(
		ctx,
		c.bucket,
		objectKey,
		expires,
		nil,
	)

	if err != nil {
		return "", ErrGetFailed
	}

	return url.String(), nil
}

func (c *Client) Delete(
	ctx context.Context,
	objectKey string,
) error {

	err := c.client.RemoveObject(
		ctx,
		c.bucket,
		objectKey,
		minio.RemoveObjectOptions{},
	)

	if err != nil {
		return ErrDeleteFailed
	}

	return nil
}
