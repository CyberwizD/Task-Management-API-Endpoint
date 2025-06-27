package repository

import (
	"sync"
	"github.com/CyberwizD/Task-Management-API-Endpoint/internal/models"
)

// TaskRepository interface defines the contract for task data operations
type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}

// InMemoryTaskRepository implements TaskRepository using in-memory storage
type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository creates a new in-memory task repository
func NewInMemoryTaskRepository() TaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*models.Task),
	}
}

// Create adds a new task to the repository
func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return nil
}

// GetAll retrieves all tasks from the repository
func (r *InMemoryTaskRepository) GetAll() ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

// GetByID retrieves a task by its ID
func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, models.ErrTaskNotFound
	}
	
	// Return a copy to prevent external modification
	taskCopy := *task
	return &taskCopy, nil
}

// Update updates an existing task
func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return models.ErrTaskNotFound
	}

	r.tasks[task.ID] = task
	return nil
}

// Delete removes a task from the repository
func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return models.ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}
