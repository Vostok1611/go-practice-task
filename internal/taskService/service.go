package taskservice

import (
	"errors"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task string, is_done string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, task string, is_done string) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task string, is_done string) (Task, error) {
	if task == "" {
		return Task{}, errors.New("task cannot be empty")
	}

	if is_done == "" {
		return Task{}, errors.New("is_done cannot be empty")
	}

	newTask := Task{
		ID:     uuid.NewString(),
		Task:   task,
		IsDone: is_done,
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

func (s *taskService) UpdateTask(id string, task string, is_done string) (Task, error) {
	if task == "" {
		return Task{}, errors.New("task cannot be empty")
	}

	if is_done == "" {
		return Task{}, errors.New("is_done cannot be empty")
	}

	newTask, err := s.repo.GetTaskByID(id)

	if err != nil {
		return Task{}, err
	}

	newTask.Task = task
	newTask.IsDone = is_done

	if err := s.repo.UpdateTask(newTask); err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
