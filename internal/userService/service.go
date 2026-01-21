package userservice

import (
	"errors"
	taskservice "gomeWork/internal/taskService"
	"strings"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(email string, password string) (User, error)
	GetAllUser() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, email string, password string) (User, error)
	DeleteUser(id string) error
	GetTasksForUser(userID string) ([]taskservice.Task, error)
}

type userService struct {
	repo        UserRepository
	taskservice taskservice.TaskService
}

func NewUserService(r UserRepository, ts taskservice.TaskService) UserService {
	return &userService{
		repo:        r,
		taskservice: ts,
	}
}

func (s *userService) CreateUser(email string, password string) (User, error) {
	if email == "" {
		return User{}, errors.New("email can not be empty")
	}
	if password == "" {
		return User{}, errors.New("password can not be empty")
	}

	email = strings.ToLower(strings.TrimSpace(email))
	if !strings.Contains(email, "@") {
		return User{}, errors.New("invalid email format")
	}

	// Проверка на уникальность email
	allUser, err := s.repo.GetAllUser()
	if err == nil { // Измените это условие
		for _, user := range allUser {
			if strings.EqualFold(user.Email, email) {
				return User{}, errors.New("user with this email already exists")
			}
		}
	}

	user := User{
		ID:       uuid.New().String(), // Генерируем UUID
		Email:    email,
		Password: password,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUser() ([]User, error) {
	return s.repo.GetAllUser()
}

func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, email string, password string) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, errors.New("user not found")
	}

	if email != "" {
		user.Email = strings.ToLower(strings.TrimSpace(email))
	}
	if password != "" {
		if len(password) < 6 {
			return User{}, errors.New("password must be at least 6 characters")
		}
		user.Password = password
	}
	err = s.repo.UpdateUser(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

// GetTasksForUser implements UserService.
func (s *userService) GetTasksForUser(userID string) ([]taskservice.Task, error) {
	if userID == "" {
		return nil, errors.New("userID is required")
	}
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return s.taskservice.GetTasksByUserID(userID)
}
