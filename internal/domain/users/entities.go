package users

import (
	"time"

	"crud/internal/domain/users/value_objects"

	"github.com/google/uuid"
)

// User представляет сущность пользователя
type User struct {
	ID        uuid.UUID // Object ID для сравнения
	Email     value_objects.EmailValueObject
	Name      value_objects.UserNameValueObject
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser создает нового пользователя
func NewUser(email value_objects.EmailValueObject, name value_objects.UserNameValueObject) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Equals проверяет равенство двух пользователей по ID
func (u *User) Equals(other *User) bool {
	if u == nil || other == nil {
		return u == other
	}
	return u.ID == other.ID
}
