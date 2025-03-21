package main

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todoapp/repository"
)

// App struct represents the main application logic.
type App struct {
	ctx         context.Context
	taskRepo    repository.TaskRepository
	profileRepo repository.ProfileRepository
}

// NewApp creates a new instance of the App.
func NewApp(taskRepo repository.TaskRepository, profileRepo repository.ProfileRepository) *App {
	return &App{
		taskRepo:    taskRepo,
		profileRepo: profileRepo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting message based on the current time of day.
func (a *App) Greet(name string) string {
	var msg strings.Builder
	currentHour := time.Now().Hour()

	// Determine the appropriate greeting
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

// AddTask adds a new task with an optional deadline and priority flag.
func (a *App) AddTask(profileID int, text string, deadline *string, hasPriority bool) error {
	return a.taskRepo.AddTask(profileID, text, deadline, hasPriority)
}

// GetActiveTasks retrieves a list of all active (incomplete) tasks.
func (a *App) GetActiveTasks(profileID int) ([]repository.Task, error) {
	return a.taskRepo.GetActiveTasks(profileID)
}

// GetCompletedTasks retrieves a list of all completed tasks.
func (a *App) GetCompletedTasks(profileID int) ([]repository.Task, error) {
	return a.taskRepo.GetCompletedTasks(profileID)
}

// UpdateTask modifies the text and/or due date of an existing task.
func (a *App) UpdateTask(id int, text string, dueAt *string) error {
	return a.taskRepo.UpdateTask(id, text, dueAt)
}

// UpdateTaskStatus changes the status of a task (e.g., "todo", "done").
func (a *App) UpdateTaskStatus(id int, status string) error {
	return a.taskRepo.UpdateTaskStatus(id, status)
}

// UpdateTaskPriority sets or removes the priority flag for a task.
func (a *App) UpdateTaskPriority(id int, hasPriority bool) error {
	return a.taskRepo.UpdateTaskPriority(id, hasPriority)
}

// DeleteTask removes a task by its ID.
func (a *App) DeleteTask(id int) error {
	return a.taskRepo.DeleteTask(id)
}

func (a *App) AddProfile(name string) (int, error) {
	return a.profileRepo.AddProfile(name)
}

func (a *App) GetProfile(id int) (repository.Profile, error) {
	return a.profileRepo.GetProfile(id)
}

func (a *App) GetProfiles() ([]repository.Profile, error) {
	return a.profileRepo.GetProfiles()
}

func (a *App) UpdateProfile(id int, name string) error {
	return a.profileRepo.UpdateProfile(id, name)
}

func (a *App) DeleteProfile(id int) error {
	return a.profileRepo.DeleteProfile(id)
}
