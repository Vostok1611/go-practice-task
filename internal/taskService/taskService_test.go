package taskservice

/*
import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskCreate(t *testing.T) {
	tests := []struct {
		name        string
		taskText    string
		isDone      string
		mockSetup   func(m *MockTaskRepository)
		wantErr     bool
		checkResult func(t *testing.T, result Task)
	}{
		{
			name:     "Первая проверка на успех",
			taskText: "Гибон опасен для живых чуществ",
			isDone:   "Ку-ку-ку дза-дза",
			mockSetup: func(m *MockTaskRepository) {
				// Только ошибка, без задачи!
				m.On("CreateTask", mock.AnythingOfType("Task")).Return(nil)
			},
			wantErr: false,
			checkResult: func(t *testing.T, result Task) {
				assert.NotEmpty(t, result.ID)
				assert.Equal(t, "Гибон опасен для живых чуществ", result.Task)
				assert.Equal(t, "Ку-ку-ку дза-дза", result.IsDone)
			},
		},
		{
			name:     "Вторая проверка провал",
			taskText: "WTFA",
			isDone:   "PEPE",
			mockSetup: func(m *MockTaskRepository) {
				// Только ошибка!
				m.On("CreateTask", mock.AnythingOfType("Task")).
					Return(errors.New("db error"))
			},
			wantErr:     true,
			checkResult: nil,
		},
	}

	// ... остальной код теста без изменений

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}
			service := NewTaskService(mockRepo)
			result, err := service.CreateTask(tt.taskText, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.checkResult != nil {
					tt.checkResult(t, result)
				}
			}
			mockRepo.AssertExpectations(t)
		})
	}

}

func TestGetAllTasks(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(m *MockTaskRepository)
		wantErr     bool
		checkResult func(t *testing.T, result []Task)
	}{
		{
			name: "Успех получении задачи",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]Task{
					{ID: "1", Task: "OKKO", IsDone: "Sirega"},
					{ID: "2", Task: "true", IsDone: "false"},
				}, nil)
			},
			wantErr: false,
			checkResult: func(t *testing.T, result []Task) {
				assert.Len(t, result, 2)
				assert.Equal(t, "OKKO", result[0].Task)
				assert.Equal(t, "Sirega", result[0].IsDone)
				assert.Equal(t, "true", result[1].Task)
				assert.Equal(t, "false", result[1].IsDone)
			},
		},
		{
			name: "Нет успеха, при получение задачи, то есть ошибка получается",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]Task{}, errors.New("База данных не доступна"))
			},
			wantErr:     true,
			checkResult: nil,
		},
		{
			name: "Делаем пустой стписок задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]Task{}, nil)
			},
			wantErr: false,
			checkResult: func(t *testing.T, result []Task) {
				assert.Empty(t, result)
				assert.Len(t, result, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)
			service := NewTaskService(mockRepo)
			result, err := service.GetAllTasks()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.checkResult != nil {
					tt.checkResult(t, result)
				}
			}
			mockRepo.AssertExpectations(t)
		})

	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name        string
		taskID      string
		taskText    string
		isDone      string
		mockSetup   func(m *MockTaskRepository)
		wantErr     bool
		checkResult func(t *testing.T, result Task)
	}{
		{
			name:     "Успешное обновление задачи",
			taskID:   "123456",
			taskText: "google",
			isDone:   "true",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskByID", "123456").Return(Task{
					ID:     "123456",
					Task:   "not google",
					IsDone: "false",
				}, nil)
				m.On("UpdateTask", mock.AnythingOfType("Task")).Return(nil)
			},
			wantErr: false,
			checkResult: func(t *testing.T, result Task) {
				assert.Equal(t, "123456", result.ID)
				assert.Equal(t, "google", result.Task)
				assert.Equal(t, "true", result.IsDone)
			},
		},
		{
			name:     "Ошибка при обновлении БД",
			taskID:   "777",
			taskText: "Nike",
			isDone:   "false",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskByID", "777").Return(Task{
					ID:     "777",
					Task:   "Abibas",
					IsDone: "true",
				}, nil)
				m.On("UpdateTask", mock.AnythingOfType("Task")).Return(errors.New("Ошибка в БД"))
			},
			wantErr:     true,
			checkResult: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)
			service := NewTaskService(mockRepo)
			result, err := service.UpdateTask(tt.taskID, tt.taskText, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.checkResult != nil {
					tt.checkResult(t, result)
				}
			}
			mockRepo.AssertExpectations(t)
		})
	}

}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		taskID    string
		mockSetap func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name:   "Успешное удаление",
			taskID: "123",
			mockSetap: func(m *MockTaskRepository) {
				m.On("DeleteTask", "123").Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Ошибка при удалении",
			taskID: "456",
			mockSetap: func(m *MockTaskRepository) {
				m.On("DeleteTask", "456").Return(errors.New("Не удалось удалить задачу"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetap(mockRepo)
			service := NewTaskService(mockRepo)
			err := service.DeleteTask(tt.taskID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
*/
