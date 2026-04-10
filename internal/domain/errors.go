package domain

import "errors"

var (
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrTokenExpired = errors.New("token expired")
	ErrForbidden    = errors.New("forbidden")
)
