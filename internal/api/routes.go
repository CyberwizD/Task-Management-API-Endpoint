package api

import (
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.GetAllTasks)
			tasks.GET("/:id", taskHandler.GetTaskByID)
		}
	}
}
