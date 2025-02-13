package repository

import "database/sql"

type TaskRepository interface {
	AddTask(text string, deadline *string, hasPriority bool) error
	GetTasks() ([]Task, error)
	DeleteTask(id int) error
	UpdateTaskStatus(id int, status string) error
	UpdateTaskPriority(id int, hasPriority bool) error
}

type UserRepository interface {
	GetUsername() (string, error)
	SetUsername(username string) error
}

type Task struct {
	ID          int            `json:"id"`
	Text        string         `json:"text"`
	Status      string         `json:"status"`
	Deadline    sql.NullString `json:"deadline"`
	HasPriority bool           `json:"has_priority"`
}
