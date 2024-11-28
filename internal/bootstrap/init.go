package bootstrap

import (
	"github.com/SandeXXX/task_service/internal/config"
	task "github.com/SandeXXX/task_service/internal/service"
	"github.com/SandeXXX/task_service/internal/store"
	task_store "github.com/SandeXXX/task_service/internal/store/task"
	"github.com/pkg/errors"
)

func CreateTaskService() (*task.TaskService, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "CreateConfig")
	}
	dbStore, err := store.New(cfg.Database)
	if err != nil {
		return nil, err
	}
	taskStore := &task_store.TaskStore{Store: dbStore}
	return task.New(taskStore)
}
