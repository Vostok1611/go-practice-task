package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var task string = "World"

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

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"eror": "Task cannot be empty"})
	}

	task = req.Task

	return c.JSON(http.StatusOK, TaskResponse{
		Status:  "success",
		Message: "Task saved successfully",
		Task:    task,
	})
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "hello, " + task})
}

func main() {
	e := echo.New()

	e.GET("/task", getTask)
	e.POST("/task", postTask)

	e.Start("localhost:8080")
}
