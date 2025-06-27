package services

import (
	"strings"

	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/models"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/repository"
)

// TaskService interface defines the business logic contract
type TaskService interface {
	CreateTask(req *models.CreateTaskRequest) (*models.Task, error)
	GetAllTasks() (*models.TasksResponse, error)
	GetTaskByID(id string) (*models.Task, error)
}

// taskService implements TaskService interface
type taskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new task service
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

// CreateTask creates a new task with business logic validation
func (s *taskService) CreateTask(req *models.CreateTaskRequest) (*models.Task, error) {
	// Additional business logic validation
	if strings.TrimSpace(req.Title) == "" {
		return nil, models.ErrInvalidTitle
	}

	// Create new task
	task := models.NewTask(req.Title, req.Description)

	// Validate task
	if err := task.Validate(); err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTasks retrieves all tasks
func (s *taskService) GetAllTasks() (*models.TasksResponse, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &models.TasksResponse{
		Tasks: tasks,
		Count: len(tasks),
	}, nil
}

// GetTaskByID retrieves a task by its ID
func (s *taskService) GetTaskByID(id string) (*models.Task, error) {
	// Validate ID format
	if strings.TrimSpace(id) == "" {
		return nil, models.ErrInvalidTaskID
	}

	// Get from repository
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
