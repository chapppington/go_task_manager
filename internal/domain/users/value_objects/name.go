package value_objects

import (
	"fmt"
	"strings"
)

// UserNameValueObject представляет имя пользователя с валидацией
type UserNameValueObject struct {
	value string
}

// NewUserNameValueObject создает новый UserNameValueObject с валидацией
func NewUserNameValueObject(name string) (UserNameValueObject, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return UserNameValueObject{}, fmt.Errorf("name cannot be empty")
	}
	if len(name) < 2 {
		return UserNameValueObject{}, fmt.Errorf("name must be at least 2 characters long")
	}
	if len(name) > 100 {
		return UserNameValueObject{}, fmt.Errorf("name must be at most 100 characters long")
	}
	return UserNameValueObject{value: name}, nil
}

// Value возвращает строковое значение имени
func (n UserNameValueObject) Value() string {
	return n.value
}

// String реализует интерфейс fmt.Stringer
func (n UserNameValueObject) String() string {
	return n.value
}

// Equals проверяет равенство двух имен
func (n UserNameValueObject) Equals(other UserNameValueObject) bool {
	return n.value == other.value
}
