package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task модель для базы данных
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(20);not null;default:'todo';index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// TableName указывает имя таблицы для GORM
func (Task) TableName() string {
	return "tasks"
}
