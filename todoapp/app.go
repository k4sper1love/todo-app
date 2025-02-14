package main

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todoapp/repository"
)

// App struct
type App struct {
	ctx      context.Context
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

// NewApp creates a new App application struct
func NewApp(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *App {
	return &App{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	var msg strings.Builder
	currentHour := time.Now().Hour()

	switch {
	case currentHour >= 6 && currentHour < 12:
		msg.WriteString("☀️ Доброе утро")
	case currentHour >= 12 && currentHour < 18:
		msg.WriteString("🌤️ Добрый день")
	case currentHour >= 18 && currentHour < 21:
		msg.WriteString("🌆 Добрый вечер")
	default: // Остальное — это ночь
		msg.WriteString("🌙 Доброй ночи")
	}

	msg.WriteString(fmt.Sprintf(", %s!", name))
	return msg.String()
}

func (a *App) AddTask(text string, deadline *string, hasPriority bool) error {
	return a.taskRepo.AddTask(text, deadline, hasPriority)
}

func (a *App) GetActiveTasks() ([]repository.Task, error) {
	return a.taskRepo.GetActiveTasks()
}

func (a *App) GetCompletedTasks() ([]repository.Task, error) {
	return a.taskRepo.GetCompletedTasks()
}

func (a *App) UpdateTask(id int, text string, dueAt *string) error {
	return a.taskRepo.UpdateTask(id, text, dueAt)
}

func (a *App) UpdateTaskStatus(id int, status string) error {
	return a.taskRepo.UpdateTaskStatus(id, status)
}

func (a *App) UpdateTaskPriority(id int, hasPriority bool) error {
	return a.taskRepo.UpdateTaskPriority(id, hasPriority)
}

func (a *App) DeleteTask(id int) error {
	return a.taskRepo.DeleteTask(id)
}

func (a *App) GetUsername() (string, error) {
	return a.userRepo.GetUsername()
}

func (a *App) SetUsername(username string) error {
	return a.userRepo.SetUsername(username)
}
