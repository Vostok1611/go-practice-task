package taskservice

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(newTask Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(newTask Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(newTask Task) error {
	return r.db.Create(&newTask).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id string) (Task, error) {
	var newTask Task
	err := r.db.First(&newTask, "id = ?", id).Error
	return newTask, err
}

func (r *taskRepository) UpdateTask(newTask Task) error {
	return r.db.Save(&newTask).Error
}

func (r *taskRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
