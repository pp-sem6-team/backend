package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/security/jwt"
)

func AuthMiddleware(jwtManager *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "missing authorization header",
				Code:    "MISSING_AUTH_HEADER",
			})
			return
		}

		tokenStr, ok := extractToken(authHeader)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "invalid authorization header",
				Code:    "INVALID_AUTH_HEADER",
			})
			return
		}

		userID, tokenType, err := jwtManager.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "invalid token",
				Code:    "INVALID_TOKEN",
			})
			return
		}

		if tokenType != jwt.TokenTypeAccess {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "access token required",
				Code:    "ACCESS_TOKEN_REQUIRED",
			})
			return
		}

		c.Set(string(UserIDKey), userID)

		c.Next()
	}
}

func extractToken(authHeader string) (string, bool) {
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		return "", false
	}

	if strings.ToLower(parts[0]) != "bearer" {
		return "", false
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", false
	}

	return token, true
}
