package middleware

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogMiddleware() gin.HandlerFunc {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		level := slog.LevelInfo
		if c.Writer.Status() >= 500 {
			level = slog.LevelError // Server error
		} else if c.Writer.Status() >= 400 {
			level = slog.LevelWarn // Client error
		}

		logger.LogAttrs(
			context.Background(),
			level,
			"Request Handled",
			slog.Int("status", c.Writer.Status()),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("ip", c.ClientIP()),
			slog.String("latency", latency.String()),
			slog.String("user_agent", c.Request.UserAgent()),
		)
	}
}
