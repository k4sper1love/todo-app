package repository

import "database/sql"

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
                deadline TEXT,
                has_priority boolean NOT NULL DEFAULT false
	);`)

	return nil
}

func (r *SQLiteRepository) AddTask(text string, deadline *string, hasPriority bool) error {
	_, err := r.db.Exec("INSERT INTO tasks (text, status, deadline, has_priority) VALUES(?, 'todo', ?, ?)", text, deadline, hasPriority)
	return err
}

func (r *SQLiteRepository) GetTasks() ([]Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Text, &task.Status, &task.Deadline, &task.HasPriority)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *SQLiteRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (r *SQLiteRepository) UpdateTaskStatus(id int, status string) error {
	_, err := r.db.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, status, id)
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
