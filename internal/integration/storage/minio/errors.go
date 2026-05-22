package minio

import "errors"

var (
	ErrStorageUnavailable = errors.New("storage unavailable")
	ErrUploadFailed       = errors.New("upload failed")
	ErrDeleteFailed       = errors.New("delete failed")
	ErrGetFailed          = errors.New("get failed")
)
