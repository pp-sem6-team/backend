package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/logger"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error(
					"panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.Int("status", 500),
					zap.Stack("stack"),
				)

				c.AbortWithStatus(500)
			}
		}()

		c.Next()
	}
}
