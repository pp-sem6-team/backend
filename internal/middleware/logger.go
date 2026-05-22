package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/logger"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		status := c.Writer.Status()

		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("duration", duration),
			zap.String("ip", c.ClientIP()),
		}

		if len(c.Errors) > 0 || status >= 500 {
			logger.Log.Error("http request", fields...)
		} else {
			logger.Log.Info("http request", fields...)
		}
	}
}
