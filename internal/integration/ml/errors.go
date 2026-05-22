package ml

import "errors"

var (
	ErrRequestFailed   = errors.New("ml request failed")
	ErrInvalidResponse = errors.New("invalid ml response")
)
