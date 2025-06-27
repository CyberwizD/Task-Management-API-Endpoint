package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS returns CORS middleware with development-friendly configuration
// NOTE: In production, this should be restricted to specific origins
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	// Development configuration - allows all origins
	// PRODUCTION WARNING: This should be restricted to specific trusted domains
	config.AllowAllOrigins = true

	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}

	config.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "Authorization",
		"X-Requested-With", "Accept", "Cache-Control",
	}

	config.ExposeHeaders = []string{
		"Content-Length", "X-Request-ID",
	}

	config.AllowCredentials = true
	config.MaxAge = 86400 // 24 hours

	return cors.New(config)
}
