package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Task represents a task entity
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// TasksResponse represents the response for getting all tasks
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
	Count int    `json:"count"`
}

// NewTask creates a new task instance with generated ID and timestamps
func NewTask(title, description string) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.New().String(),
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Validate validates the task data
func (t *Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return ErrInvalidTitle
	}
	return nil
}
