package taskservice

import (
	"errors"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task string, is_done string, userID string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, task string, is_done string, userID string) (Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task string, is_done string, userID string) (Task, error) {
	if task == "" {
		return Task{}, errors.New("task cannot be empty")
	}

	if is_done == "" {
		return Task{}, errors.New("is_done cannot be empty")
	}
	if userID == "" {
		return Task{}, errors.New("user_id is required")
	}

	newTask := Task{
		ID:     uuid.NewString(),
		Task:   task,
		IsDone: is_done,
		UserID: userID,
	}

	err := s.repo.CreateTask(newTask)
	if err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, task string, is_done string, userID string) (Task, error) {
	existingTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, errors.New("task not found")
	}

	//  ОБНОВЛЯЕМ ТОЛЬКО ЕСЛИ ПЕРЕДАНО НЕ ПУСТОЕ ЗНАЧЕНИЕ
	// Пустая строка "" означает "не менять это поле"
	if task != "" {
		existingTask.Task = task
	}
	if is_done != "" {
		existingTask.IsDone = is_done
	}
	if userID != "" {
		existingTask.UserID = userID
	}

	if task == "" && is_done == "" && userID == "" {
		return Task{}, errors.New("nothing to update")
	}

	// Сохраняем в БД
	if err := s.repo.UpdateTask(existingTask); err != nil {
		return Task{}, err
	}

	return existingTask, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

// GetTasksByUserID implements TaskService.
func (s *taskService) GetTasksByUserID(userID string) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
