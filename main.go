package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

type TaskRequest struct {
	Task string `json:"task"`
}

var tasks = []Task{}

func postTasks(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task cannot be empty"})
	}

	newTask := Task{
		ID:   uuid.NewString(),
		Task: req.Task,
	}

	tasks = append(tasks, newTask)

	return c.JSON(http.StatusOK, newTask)
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func patсhTasks(c echo.Context) error {
	id := c.Param("id")

	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task cannot be empty"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Task = req.Task
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func deleteTasks(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func main() {
	e := echo.New()

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTasks)
	e.PATCH("/tasks/:id", patсhTasks)
	e.DELETE("/tasks/:id", deleteTasks)

	e.Start("localhost:8080")
}
