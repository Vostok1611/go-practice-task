package taskservice

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository - поддельный репозиторий для тестов
type MockTaskRepository struct {
	mock.Mock
}

// CreateTask соответствует интерфейсу TaskRepository
func (m *MockTaskRepository) CreateTask(newTask Task) error {
	args := m.Called(newTask)
	return args.Error(0) // Только ошибка!
}

// GetAllTasks соответствует интерфейсу TaskRepository
func (m *MockTaskRepository) GetAllTasks() ([]Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Task), args.Error(1)
}

// GetTaskByID соответствует интерфейсу TaskRepository
func (m *MockTaskRepository) GetTaskByID(id string) (Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return Task{}, args.Error(1)
	}
	return args.Get(0).(Task), args.Error(1)
}

// UpdateTask соответствует интерфейсу TaskRepository
func (m *MockTaskRepository) UpdateTask(newTask Task) error {
	args := m.Called(newTask)
	return args.Error(0)
}

// DeleteTask соответствует интерфейсу TaskRepository
func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
