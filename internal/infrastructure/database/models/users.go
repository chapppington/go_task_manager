package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User модель для базы данных
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName указывает имя таблицы для GORM
func (User) TableName() string {
	return "users"
}
