package repository

import (
	"database/sql"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

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

func (r *SQLiteRepository) AddTask(text string, deadline *string, hasPriority bool) error {
	_, err := r.db.Exec(""+
		"INSERT INTO tasks (text, status, due_at, has_priority) "+
		"VALUES(?, 'todo', ?, ?)", text, deadline, hasPriority)
	return err
}

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

func (r *SQLiteRepository) UpdateTask(id int, text string, dueAt *string) error {
	_, err := r.db.Exec(`UPDATE tasks SET text = ?, due_at = ? WHERE id = ?`, text, dueAt, id)
	return err
}

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

func (r *SQLiteRepository) UpdateTaskPriority(id int, hasPriority bool) error {
	_, err := r.db.Exec(`UPDATE tasks SET has_priority = ? WHERE id = ?`, hasPriority, id)
	return err
}

func (r *SQLiteRepository) GetUsername() (string, error) {
	var username string
	err := r.db.QueryRow("SELECT username FROM users WHERE id = 1").Scan(&username)
	return username, err
}

func (r *SQLiteRepository) SetUsername(username string) error {
	_, err := r.db.Exec(`
		INSERT INTO users (id, username) VALUES (1, ?)
		ON CONFLICT(id) DO UPDATE SET username = EXCLUDED.username`, username)
	return err
}

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

func (r *SQLiteRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
