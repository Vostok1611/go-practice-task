package main

import (
	"gomeWork/internal/db"
	"gomeWork/internal/handlers"
	taskservice "gomeWork/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	newTaskRepo := taskservice.NewTaskRepository(database)
	newTaskService := taskservice.NewTaskService(newTaskRepo)
	newTaskHandler := handlers.NewTaskHandler(newTaskService)

	e.GET("/tasks", newTaskHandler.GetTasks)
	e.POST("/tasks", newTaskHandler.PostTasks)
	e.PATCH("/tasks/:id", newTaskHandler.PatсhTasks)
	e.DELETE("/tasks/:id", newTaskHandler.DeleteTasks)

	e.Start("localhost:8080")
}
