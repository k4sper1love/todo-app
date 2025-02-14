package repository

import (
	"database/sql"
)

// SQLiteRepository is a repository implementation using SQLite as the database.
type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository creates a new instance of SQLiteRepository.
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// InitDB initializes the database schema, creating necessary tables if they don't exist.
func (r *SQLiteRepository) InitDB() error {
	_, err := r.db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    			id INTEGER PRIMARY KEY CHECK (id = 1),
				username TEXT NOT NULL
		);`)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`CREATE TABLE IF NOT EXISTS tasks (
	 			id INTEGER PRIMARY KEY,
				text TEXT NOT NULL,
				status TEXT CHECK(status IN ('todo', 'in_progress', 'done')) NOT NULL DEFAULT 'todo',
                has_priority boolean NOT NULL DEFAULT false,
    			created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    			due_at TEXT,
    			completed_at TEXT
	);`)

	return nil
}

// AddTask inserts a new task into the database.
func (r *SQLiteRepository) AddTask(text string, deadline *string, hasPriority bool) error {
	_, err := r.db.Exec(""+
		"INSERT INTO tasks (text, status, due_at, has_priority) "+
		"VALUES(?, 'todo', ?, ?)", text, deadline, hasPriority)
	return err
}

// GetActiveTasks retrieves all tasks that are not marked as "done".
func (r *SQLiteRepository) GetActiveTasks() ([]Task, error) {
	rows, err := r.db.Query(`
		SELECT * FROM tasks WHERE status IN ('todo', 'in_progress')
		ORDER BY 
		    has_priority DESC, 
		    CASE 
		        WHEN due_at IS NOT NULL THEN 0
		        ELSE 1
		    END,
		    due_at ASC,
		    created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTasks(rows)
}

// GetCompletedTasks retrieves all completed tasks, ordered by completion date.
func (r *SQLiteRepository) GetCompletedTasks() ([]Task, error) {
	rows, err := r.db.Query(`
		SELECT * FROM tasks WHERE status = 'done'
		ORDER BY completed_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTasks(rows)
}

// UpdateTask updates the text and due date of a specific task.
func (r *SQLiteRepository) UpdateTask(id int, text string, dueAt *string) error {
	_, err := r.db.Exec(`UPDATE tasks SET text = ?, due_at = ? WHERE id = ?`, text, dueAt, id)
	return err
}

// UpdateTaskStatus updates the status of a task and sets the completed_at timestamp if necessary.
func (r *SQLiteRepository) UpdateTaskStatus(id int, status string) error {
	_, err := r.db.Exec(`
			UPDATE tasks 
			SET status = ?,
			    completed_at = CASE
			    	WHEN ? = 'done' THEN CURRENT_TIMESTAMP
					WHEN ? = 'todo' THEN NULL
					ELSE completed_at
				END
			WHERE id = ?`,
		status, status, status, id)
	return err
}

// UpdateTaskPriority updates the priority flag of a specific task.
func (r *SQLiteRepository) UpdateTaskPriority(id int, hasPriority bool) error {
	_, err := r.db.Exec(`UPDATE tasks SET has_priority = ? WHERE id = ?`, hasPriority, id)
	return err
}

// DeleteTask removes a task from the database by its ID.
func (r *SQLiteRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

// GetUsername retrieves the stored username from the users table.
func (r *SQLiteRepository) GetUsername() (string, error) {
	var username string
	err := r.db.QueryRow("SELECT username FROM users WHERE id = 1").Scan(&username)
	return username, err
}

// SetUsername inserts or updates the username in the database.
func (r *SQLiteRepository) SetUsername(username string) error {
	_, err := r.db.Exec(`
		INSERT INTO users (id, username) VALUES (1, ?)
		ON CONFLICT(id) DO UPDATE SET username = EXCLUDED.username`, username)
	return err
}

// scanTasks scans rows into a slice of Task structs.
func scanTasks(rows *sql.Rows) ([]Task, error) {
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Text, &task.Status, &task.HasPriority, &task.CreatedAt, &task.DueAt, &task.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
