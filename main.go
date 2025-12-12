package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type Task struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Task      string         `json:"task"`
	IsDone    string         `json:"is_done"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	IsDone string `json:"is_done"`
}

func postTasks(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task cannot be empty"})
	}

	newTask := Task{
		ID:     uuid.NewString(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	if err := db.Create(&newTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add task"})
	}

	return c.JSON(http.StatusCreated, newTask)
}

func getTasks(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
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

	var newTask Task
	if err := db.First(&newTask, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find task"})
	}
	newTask.Task = req.Task
	newTask.IsDone = req.IsDone

	if err := db.Save(&newTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, newTask)
}

func deleteTasks(c echo.Context) error {
	id := c.Param("id")

	var task Task

	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	if err := db.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	e := echo.New()

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTasks)
	e.PATCH("/tasks/:id", patсhTasks)
	e.DELETE("/tasks/:id", deleteTasks)

	e.Start("localhost:8080")
}
