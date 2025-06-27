package middleware

import (
	"log"
	"net/http"

	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/models"

	"github.com/gin-gonic/gin"
)

// ErrorHandler provides centralized error handling
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				log.Printf("Panic recovered: %v", err)

				// Return internal server error
				c.JSON(http.StatusInternalServerError, models.NewErrorResponse(
					"Internal server error",
					"An unexpected error occurred",
				))
				c.Abort()
			}
		}()

		c.Next()

		// Handle any errors that occurred during request processing
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Log the error
			log.Printf("Request error: %v", err)

			// If response hasn't been written yet, return error response
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, models.NewErrorResponse(
					"Internal server error",
					"An error occurred while processing the request",
				))
			}
		}
	}
}
