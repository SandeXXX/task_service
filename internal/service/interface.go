package task

import "github.com/SandeXXX/task_service/internal/models"

type TaskRepository interface {
	CreateTask(task *models.CreateTaskParams) error
	GetTask(id string) (*models.Task, error)
	UpdateTask(id string, updatedTask *models.Task) error
	DeleteTask(id string) error
}
