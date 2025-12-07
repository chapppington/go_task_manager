package users

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// UserNotFoundError представляет ошибку, когда пользователь не найден
type UserNotFoundError struct {
	UserID uuid.UUID
	Email  string
}

func (e *UserNotFoundError) Error() string {
	if e.Email != "" {
		return fmt.Sprintf("user with email %s not found", e.Email)
	}
	return fmt.Sprintf("user with ID %s not found", e.UserID)
}

// UserAlreadyExistsError представляет ошибку, когда пользователь уже существует
type UserAlreadyExistsError struct {
	Email string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with email %s already exists", e.Email)
}

// InvalidUserDataError представляет ошибку валидации данных пользователя
type InvalidUserDataError struct {
	Field   string
	Message string
}

func (e *InvalidUserDataError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("invalid user data: field '%s' - %s", e.Field, e.Message)
	}
	return fmt.Sprintf("invalid user data: %s", e.Message)
}

// UserOperationFailedError представляет ошибку при выполнении операции с пользователем
type UserOperationFailedError struct {
	Operation string
	Reason    string
}

func (e *UserOperationFailedError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("user operation '%s' failed: %s", e.Operation, e.Reason)
	}
	return fmt.Sprintf("user operation '%s' failed", e.Operation)
}

// IsUserNotFound проверяет, является ли ошибка ошибкой "пользователь не найден"
func IsUserNotFound(err error) bool {
	var userNotFoundErr *UserNotFoundError
	return errors.As(err, &userNotFoundErr)
}

// IsUserAlreadyExists проверяет, является ли ошибка ошибкой "пользователь уже существует"
func IsUserAlreadyExists(err error) bool {
	var userExistsErr *UserAlreadyExistsError
	return errors.As(err, &userExistsErr)
}

// IsInvalidUserData проверяет, является ли ошибка ошибкой валидации данных пользователя
func IsInvalidUserData(err error) bool {
	var invalidDataErr *InvalidUserDataError
	return errors.As(err, &invalidDataErr)
}
