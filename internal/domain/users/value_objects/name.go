package value_objects

import (
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
		return UserNameValueObject{}, &InvalidNameError{Message: "name cannot be empty"}
	}
	if len(name) < 2 {
		return UserNameValueObject{}, &InvalidNameError{Value: name, Message: "name must be at least 2 characters long"}
	}
	if len(name) > 100 {
		return UserNameValueObject{}, &InvalidNameError{Value: name, Message: "name must be at most 100 characters long"}
	}
	return UserNameValueObject{value: name}, nil
}

// Value возвращает строковое значение имени
func (n UserNameValueObject) Value() string {
	return n.value
}

// Equals проверяет равенство двух имен
func (n UserNameValueObject) Equals(other UserNameValueObject) bool {
	return n.value == other.value
}
