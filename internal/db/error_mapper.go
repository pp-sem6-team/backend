package db

import (
	"errors"
	"strings"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

const pgUniqueViolation = "23505"

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgUniqueViolation {
			return ErrAlreadyExists
		}
	}

	if strings.Contains(err.Error(), "duplicate key") {
		return ErrAlreadyExists
	}

	return ErrInternal
}
