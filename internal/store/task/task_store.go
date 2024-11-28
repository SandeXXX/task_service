package task

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/SandeXXX/task_service/internal/models"
	"github.com/SandeXXX/task_service/internal/store"
)

type TaskStore struct {
	Store *store.Store
}

func (t *TaskStore) CreateTask(task *models.CreateTaskParams) error {
	queryBuilder := sq.Insert(tableName).
		Columns(
			"title",
			"body",
			"complited").
		Values(
			task.Title,
			task.Body,
			task.Completed,
		)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building insert query: %v", err)
	}
	_, err = t.Store.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error executing insert query: %v", err)
	}
	return nil
}
func (t *TaskStore) GetTask(id string) (*models.Task, error) {
	queryBuilder := sq.Select(
		"title",
		"body",
		"complited",
		"createdat",
		"updatedat",
	).
		From(tableName).
		Where(sq.Eq{"task_id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %v", err)
	}
	var task models.Task
	err = t.Store.DB.QueryRow(query, args...).Scan(
		&task.Id,
		&task.Title,
		&task.Body,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task not found")
		}
	}

	return &task, nil
}
func (t *TaskStore) UpdateTask(id string, updatedTask *models.Task) error {
	queryBuilder := sq.Update(tableName).
		Set("title", updatedTask.Title).
		Set("body", updatedTask.Body).
		Set("completed", updatedTask.Completed).
		Set("updated_at", sq.Expr("CURRENT_TIMESTAMP")).
		Where(sq.Eq{"task_id": id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building update query: %v", err)
	}
	_, err = t.Store.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error executing update query: %v", err)
	}
	return nil
}
func (t *TaskStore) DeleteTask(id string) error {
	queryBuilder := sq.Delete(tableName).
		Where(sq.Eq{"task_id": id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	_, execErr := t.Store.DB.Exec(query, args...)
	if execErr != nil {
		return execErr
	}

	return nil
}
