package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

type Task struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    			id INTEGER PRIMARY KEY CHECK (id = 1),
				username TEXT NOT NULL
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS tasks (
	 			id INTEGER PRIMARY KEY,
				text TEXT NOT NULL,
				status TEXT CHECK(status IN ('todo', 'in_progress', 'done')) NOT NULL DEFAULT 'todo'
	);`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("База данных инициализирована")
}

func AddTask(text string) {
	_, err := db.Exec("INSERT INTO tasks (text, status) VALUES(?, 'todo')", text)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTasks() []Task {
	rows, err := db.Query("SELECT id, text, status FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Text, &task.Status)
		tasks = append(tasks, task)
	}

	return tasks
}

func RemoveTask(id int) {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsername() (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = 1").Scan(&username)
	if err != nil {
		return "", nil
	}
	return username, nil
}

func SetUsername(username string) error {
	_, err := db.Exec(`
		INSERT INTO users VALUES(id, username) VALUES (1, ?)
		ON CONFLICT(id) DO UPDATE SET username = ?`, username, username)
	return err
}
