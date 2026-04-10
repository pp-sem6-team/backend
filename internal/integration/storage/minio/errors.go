package minio

import "errors"

var (
	ErrUploadFailed = errors.New("upload failed")
	ErrDeleteFailed = errors.New("delete failed")
	ErrGetFailed    = errors.New("get failed")
)
