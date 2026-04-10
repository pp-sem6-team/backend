package jwt

import "errors"

var (
	ErrTokenGenerationFailed = errors.New("failed to generate token")
	ErrInvalidToken          = errors.New("invalid token")
	ErrExpiredToken          = errors.New("token is expired")
)
