package repository

import (
	"database/sql"
	"log"
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
		`CREATE TABLE IF NOT EXISTS profiles (
    			id INTEGER PRIMARY KEY,
				name TEXT NOT NULL,
    			created_at TEXT DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`CREATE TABLE IF NOT EXISTS tasks (
	 			id INTEGER PRIMARY KEY,
	 			profile_id INTEGER,
				text TEXT NOT NULL,
				status TEXT CHECK(status IN ('todo', 'in_progress', 'done')) NOT NULL DEFAULT 'todo',
                has_priority boolean NOT NULL DEFAULT false,
    			created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    			due_at TEXT,
    			completed_at TEXT,
    			FOREIGN KEY (profile_id) REFERENCES profiles(id)
	);`)

	return err
}

// AddTask inserts a new task into the database.
func (r *SQLiteRepository) AddTask(profileID int, text string, deadline *string, hasPriority bool) error {
	_, err := r.db.Exec(`
		INSERT INTO tasks (profile_id, text, status, due_at, has_priority)
		VALUES(?, ?, 'todo', ?, ?)`, profileID, text, deadline, hasPriority)
	log.Println(err)
	return err
}

// GetActiveTasks retrieves all tasks that are not marked as "done".
func (r *SQLiteRepository) GetActiveTasks(profileID int) ([]Task, error) {
	rows, err := r.db.Query(`
		SELECT * FROM tasks 
		WHERE status IN ('todo', 'in_progress')
			AND profile_id = ?
		ORDER BY 
		    has_priority DESC, 
		    CASE 
		        WHEN due_at IS NOT NULL THEN 0
		        ELSE 1
		    END,
		    due_at ASC,
		    created_at DESC`, profileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	return scanTasks(rows)
}

// GetCompletedTasks retrieves all completed tasks, ordered by completion date.
func (r *SQLiteRepository) GetCompletedTasks(profileID int) ([]Task, error) {
	rows, err := r.db.Query(`
		SELECT * FROM tasks 
		WHERE status = 'done'
			AND profile_id = ?
		ORDER BY completed_at DESC`, profileID)
	if err != nil {
		log.Println(err)
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

func (r *SQLiteRepository) AddProfile(name string) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO profiles (name) 
		VALUES (?)
		RETURNING id`, name).Scan(&id)
	return id, err
}

func (r *SQLiteRepository) GetProfile(id int) (Profile, error) {
	var profile Profile

	row := r.db.QueryRow(`
		SELECT * FROM profiles WHERE id = ?`, id)
	err := row.Scan(&profile.ID, &profile.Name, &profile.CreatedAt)

	return profile, err
}

func (r *SQLiteRepository) GetProfiles() ([]Profile, error) {
	rows, err := r.db.Query(`
		SELECT * FROM profiles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanProfiles(rows)
}

func (r *SQLiteRepository) UpdateProfile(id int, name string) error {
	_, err := r.db.Exec(`
		UPDATE profiles
		SET name = ?
		WHERE id = ?`, name, id)

	return err
}

func (r *SQLiteRepository) DeleteProfile(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM profiles WHERE id = ?`, id)
	return err
}

// scanTasks scans rows into a slice of Task structs.
func scanTasks(rows *sql.Rows) ([]Task, error) {
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.ProfileID, &task.Text, &task.Status, &task.HasPriority, &task.CreatedAt, &task.DueAt, &task.CompletedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func scanProfiles(rows *sql.Rows) ([]Profile, error) {
	var profiles []Profile
	for rows.Next() {
		var profile Profile
		err := rows.Scan(&profile.ID, &profile.Name, &profile.CreatedAt)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
