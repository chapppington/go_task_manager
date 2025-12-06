package users

import (
	"time"

	"crud/internal/domain/users/value_objects"

	"github.com/google/uuid"
)

// User представляет сущность пользователя
type User struct {
	ID        int64
	OID       uuid.UUID // Object ID для сравнения
	Email     value_objects.EmailValueObject
	Name      value_objects.UserNameValueObject
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser создает нового пользователя
func NewUser(email value_objects.EmailValueObject, name value_objects.UserNameValueObject) *User {
	now := time.Now()
	return &User{
		OID:       uuid.New(),
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Equals проверяет равенство двух пользователей по хешу
func (u *User) Equals(other *User) bool {
	if u == nil || other == nil {
		return u == other
	}
	return u.Hash() == other.Hash()
}

// Hash возвращает хеш пользователя (для использования в map)
func (u *User) Hash() string {
	return u.OID.String()
}
