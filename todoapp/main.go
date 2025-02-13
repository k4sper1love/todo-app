package main

import (
	"database/sql"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	_ "modernc.org/sqlite"
	"todoapp/repository"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewSQLiteRepository(db)
	if err = repo.InitDB(); err != nil {
		log.Fatal(err)
	}

	// Create an instance of the app structure
	app := NewApp(repo, repo)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "todoapp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}
