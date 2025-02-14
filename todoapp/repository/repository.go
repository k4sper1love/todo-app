package repository

import (
	"database/sql"
)

type TaskRepository interface {
	AddTask(text string, deadline *string, hasPriority bool) error
	GetActiveTasks() ([]Task, error)
	GetCompletedTasks() ([]Task, error)
	UpdateTask(id int, text string, dueAt *string) error
	UpdateTaskStatus(id int, status string) error
	UpdateTaskPriority(id int, hasPriority bool) error
	DeleteTask(id int) error
}

type UserRepository interface {
	GetUsername() (string, error)
	SetUsername(username string) error
}

type Task struct {
	ID          int            `json:"id"`
	Text        string         `json:"text"`
	Status      string         `json:"status"`
	HasPriority bool           `json:"has_priority"`
	CreatedAt   sql.NullString `json:"created_at"`
	DueAt       sql.NullString `json:"due_at"`
	CompletedAt sql.NullString `json:"completed_at"`
}
