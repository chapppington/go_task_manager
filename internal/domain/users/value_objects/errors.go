package value_objects

import (
	"errors"
	"fmt"
)

// InvalidNameError представляет ошибку валидации имени пользователя
type InvalidNameError struct {
	Value   string
	Message string
}

func (e *InvalidNameError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("invalid user name: %s", e.Value)
}

// InvalidEmailError представляет ошибку валидации email пользователя
type InvalidEmailError struct {
	Value   string
	Message string
}

func (e *InvalidEmailError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("invalid email: %s", e.Value)
}

// IsInvalidName проверяет, является ли ошибка ошибкой валидации имени
func IsInvalidName(err error) bool {
	var invalidNameErr *InvalidNameError
	return errors.As(err, &invalidNameErr)
}

// IsInvalidEmail проверяет, является ли ошибка ошибкой валидации email
func IsInvalidEmail(err error) bool {
	var invalidEmailErr *InvalidEmailError
	return errors.As(err, &invalidEmailErr)
}
