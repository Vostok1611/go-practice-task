package main

import (
	"gomeWork/internal/db"
	"gomeWork/internal/handlers"
	taskservice "gomeWork/internal/taskService"
	"gomeWork/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	newTaskRepo := taskservice.NewTaskRepository(database)
	newTaskService := taskservice.NewTaskService(newTaskRepo)
	newTaskHandler := handlers.NewTaskHandler(newTaskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(newTaskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start wirh err: %v", err)
	}

}
