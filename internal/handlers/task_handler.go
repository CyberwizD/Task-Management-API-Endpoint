package handlers

import (
	"errors"
	"net/http"

	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/models"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/services"

	"github.com/gin-gonic/gin"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service services.TaskService
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(
			"Invalid request body",
			err.Error(),
		))
		return
	}

	// Create task through service
	task, err := h.service.CreateTask(&req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := "Internal server error"

		// Handle specific errors
		if errors.Is(err, models.ErrInvalidTitle) {
			statusCode = http.StatusBadRequest
			errorMsg = "Validation failed"
		}

		c.JSON(statusCode, models.NewErrorResponse(errorMsg, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks handles GET /tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	response, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(
			"Failed to retrieve tasks",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetTaskByID handles GET /tasks/:id
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := "Internal server error"

		// Handle specific errors
		if errors.Is(err, models.ErrTaskNotFound) {
			statusCode = http.StatusNotFound
			errorMsg = "Task not found"
		} else if errors.Is(err, models.ErrInvalidTaskID) {
			statusCode = http.StatusBadRequest
			errorMsg = "Invalid request"
		}

		c.JSON(statusCode, models.NewErrorResponse(errorMsg, err.Error()))
		return
	}

	c.JSON(http.StatusOK, task)
}
