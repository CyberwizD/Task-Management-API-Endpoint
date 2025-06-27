package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger returns a custom logging middleware
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s %s %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.ErrorMessage,
		)
	})
}

// CustomLogger provides structured logging
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status
		status := c.Writer.Status()

		// Build log message
		logData := map[string]interface{}{
			"timestamp":  start.Format(time.RFC3339),
			"method":     c.Request.Method,
			"path":       path,
			"query":      raw,
			"status":     status,
			"latency":    latency.String(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		// Log errors if any
		if len(c.Errors) > 0 {
			logData["errors"] = c.Errors.String()
		}

		// Simple log output (in production, use structured logger like logrus or zap)
		log.Printf("API Request: %+v", logData)
	}
}
