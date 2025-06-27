package main

import (
	"log"

	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/api"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/config"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/handlers"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/middleware"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/repository"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Initialize repository
	taskRepo := repository.NewInMemoryTaskRepository()

	// Initialize services
	taskService := services.NewTaskService(taskRepo)

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.Logger())

	// Setup routes
	api.SetupRoutes(router, taskHandler)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
