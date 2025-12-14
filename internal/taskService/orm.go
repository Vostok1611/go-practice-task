package taskservice

import "gorm.io/gorm"

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
