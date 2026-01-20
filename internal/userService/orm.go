package userservice

import (
	taskservice "gomeWork/internal/taskService"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string             `gorm:"primaryKey;type:uuid" json:"id"`
	Email     string             `gorm:"uniqueIndex" json:"email"`
	Password  string             `gorm:"not null" json:"-"`
	CreatedAt time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"-" gorm:"index"`
	Tasks     []taskservice.Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"` // ← Связь "один-ко-многим"

}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
