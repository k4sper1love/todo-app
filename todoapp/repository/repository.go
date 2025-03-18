// Package repository provides interfaces and structures for interacting with the database.
// It implements the Repository Pattern to separate data access logic from business logic.
package repository

import (
	"database/sql"
)

// TaskRepository defines the interface for working with tasks in the database.
type TaskRepository interface {
	// AddTask inserts a new task with optional deadline and priority flag.
	AddTask(profileID int, text string, deadline *string, hasPriority bool) error

	// GetActiveTasks retrieves all active (incomplete) tasks.
	GetActiveTasks(profileID int) ([]Task, error)

	// GetCompletedTasks retrieves all completed tasks.
	GetCompletedTasks(profileID int) ([]Task, error)

	// UpdateTask modifies the text and/or due date of an existing task.
	UpdateTask(id int, text string, dueAt *string) error

	// UpdateTaskStatus changes the status of a task (e.g., "todo", "done").
	UpdateTaskStatus(id int, status string) error

	// UpdateTaskPriority updates the priority flag for a task.
	UpdateTaskPriority(id int, hasPriority bool) error

	// DeleteTask removes a task by its ID.
	DeleteTask(id int) error
}

type ProfileRepository interface {
	AddProfile(name string) (int, error)

	GetProfile(id int) (Profile, error)

	GetProfiles() ([]Profile, error)

	UpdateProfile(id int, name string) error

	DeleteProfile(id int) error
}

// Task represents a task entity in the database.
type Task struct {
	ID          int            `json:"id"`
	ProfileID   int            `json:"profile_id"`
	Text        string         `json:"text"`
	Status      string         `json:"status"`
	HasPriority bool           `json:"has_priority"`
	CreatedAt   sql.NullString `json:"created_at"`
	DueAt       sql.NullString `json:"due_at"`
	CompletedAt sql.NullString `json:"completed_at"`
}

type Profile struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt sql.NullString `json:"created_at"`
}
