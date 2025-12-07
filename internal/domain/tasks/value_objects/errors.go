package value_objects

import (
	"errors"
	"fmt"
)

// InvalidTitleError представляет ошибку валидации заголовка задачи
type InvalidTitleError struct {
	Value   string
	Message string
}

func (e *InvalidTitleError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("invalid task title: %s", e.Value)
}

// InvalidStatusError представляет ошибку валидации статуса задачи
type InvalidStatusError struct {
	Value       string
	ValidValues []string
	Message     string
}

func (e *InvalidStatusError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if len(e.ValidValues) > 0 {
		return fmt.Sprintf("invalid task status '%s'. Valid statuses: %v", e.Value, e.ValidValues)
	}
	return fmt.Sprintf("invalid task status: %s", e.Value)
}

// IsInvalidTitle проверяет, является ли ошибка ошибкой валидации заголовка
func IsInvalidTitle(err error) bool {
	var invalidTitleErr *InvalidTitleError
	return errors.As(err, &invalidTitleErr)
}

// IsInvalidStatus проверяет, является ли ошибка ошибкой валидации статуса
func IsInvalidStatus(err error) bool {
	var invalidStatusErr *InvalidStatusError
	return errors.As(err, &invalidStatusErr)
}
