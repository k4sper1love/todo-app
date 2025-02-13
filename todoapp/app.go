package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
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
	case currentHour < 6 || currentHour > 21:
		msg.WriteString("🌙 Доброй ночи")
	case currentHour < 12:
		msg.WriteString("☀️ Доброе утро")
	case currentHour < 18:
		msg.WriteString("🌤️ Добрый день")
	case currentHour < 21:
		msg.WriteString("🌆 Добрый вечер")
	default:
		msg.WriteString("😊 Доброго времени суток")
	}

	msg.WriteString(fmt.Sprintf(", %s!", name))

	return msg.String()
}

func (a *App) AddTask(text string, deadline *string, hasPriority bool) {
	AddTask(text, deadline, hasPriority)
}

func (a *App) GetTasks() []Task {
	return GetTasks()
}

func (a *App) RemoveTask(id int) {
	RemoveTask(id)
}

func (a *App) GetUsername() (string, error) {
	return GetUsername()
}

func (a *App) SetUsername(username string) error {
	return SetUsername(username)
}

func (a *App) UpdateTaskStatus(id int, status string) error {
	return UpdateTaskStatus(id, status)
}

func (a *App) UpdateTaskPriority(id int, hasPriority bool) error {
	return UpdateTaskPriority(id, hasPriority)
}
