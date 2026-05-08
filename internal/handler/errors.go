package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/integration/storage/minio"
	"github.com/pp-sem6-team/backend/internal/security/jwt"
)

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	var (
		status  int
		message string
		code    string
	)

	switch {
	case errors.Is(err, domain.ErrInvalidCreds):
		status, message, code = http.StatusUnauthorized, "invalid credentials", "INVALID_CREDENTIALS"

	case errors.Is(err, domain.ErrTokenExpired):
		status, message, code = http.StatusUnauthorized, "token expired", "TOKEN_EXPIRED"

	case errors.Is(err, domain.ErrForbidden):
		status, message, code = http.StatusForbidden, "forbidden", "FORBIDDEN"

	case errors.Is(err, domain.ErrInvalidFileType):
		status, message, code = http.StatusBadRequest, "invalid file type", "INVALID_FILE_TYPE"

	case errors.Is(err, db.ErrNotFound):
		status, message, code = http.StatusNotFound, "resource not found", "NOT_FOUND"

	case errors.Is(err, db.ErrAlreadyExists):
		status, message, code = http.StatusConflict, "resource already exists", "ALREADY_EXISTS"

	case errors.Is(err, db.ErrInternal):
		status, message, code = http.StatusInternalServerError, "database error", "DB_ERROR"

	case errors.Is(err, minio.ErrStorageUnavailable):
		status, message, code = http.StatusServiceUnavailable, "storage unavailable", "STORAGE_UNAVAILABLE"

	case errors.Is(err, minio.ErrUploadFailed):
		status, message, code = http.StatusInternalServerError, "upload failed", "UPLOAD_FAILED"

	case errors.Is(err, minio.ErrDeleteFailed):
		status, message, code = http.StatusInternalServerError, "delete failed", "DELETE_FAILED"

	case errors.Is(err, minio.ErrGetFailed):
		status, message, code = http.StatusInternalServerError, "failed to get file", "GET_FAILED"

	case errors.Is(err, jwt.ErrTokenGenerationFailed):
		status, message, code = http.StatusInternalServerError, "token generation failed", "TOKEN_GENERATION_FAILED"

	case errors.Is(err, jwt.ErrInvalidToken):
		status, message, code = http.StatusUnauthorized, "invalid token", "INVALID_TOKEN"

	case errors.Is(err, jwt.ErrExpiredToken):
		status, message, code = http.StatusUnauthorized, "token expired", "TOKEN_EXPIRED"

	default:
		status, message, code = http.StatusInternalServerError, "internal server error", "INTERNAL_ERROR"
	}

	c.JSON(status, dto.ErrorResponse{
		Message: message,
		Code:    code,
	})

	return true
}
