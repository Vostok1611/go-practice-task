package main

import (
	"gomeWork/internal/db"
	"gomeWork/internal/handlers"
	taskservice "gomeWork/internal/taskService"
	userservice "gomeWork/internal/userService"
	"gomeWork/internal/web/api"
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

	newUserRepo := userservice.NewUserRepository(database)
	newUserService := userservice.NewUserService(newUserRepo, newTaskService)

	newTaskHandler := handlers.NewTaskHandler(newTaskService)
	newUserHandler := handlers.NewUserHandler(newUserService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	combinedHandler := struct {
		*handlers.TaskHandler
		*handlers.UserHandler
	}{
		TaskHandler: newTaskHandler,
		UserHandler: newUserHandler,
	}

	strictHandler := api.NewStrictHandler(combinedHandler, nil)

	api.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start wirh err: %v", err)
	}

}
