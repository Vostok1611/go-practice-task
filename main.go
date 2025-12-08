package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Task string

type TaskRequest struct {
	Task string `json:"task"`
}

type TaskResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Task    string `json:"task"`
}

func postTask(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	Task = req.Task

	return c.JSON(http.StatusOK, TaskResponse{
		Status:  "success",
		Message: "Task saved successfully",
		Task:    Task,
	})
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, task")
}

func main() {
	e := echo.New()

	e.GET("/hello,task", getTask)
	e.POST("/Task", postTask)

	e.Start("localhost:8080")
}
