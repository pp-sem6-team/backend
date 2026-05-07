package db

import (
	"errors"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/pp-sem6-team/backend/internal/apperrors"
	"gorm.io/gorm"
)

const pgUniqueViolation = "23505"

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgUniqueViolation {
			return apperrors.ErrAlreadyExists
		}
	}

	if strings.Contains(err.Error(), "duplicate key") {
		return apperrors.ErrAlreadyExists
	}

	return apperrors.ErrInternal
}
