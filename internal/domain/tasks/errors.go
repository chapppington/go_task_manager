package tasks

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// TaskNotFoundError представляет ошибку, когда задача не найдена
type TaskNotFoundError struct {
	TaskID uuid.UUID
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("task with ID %s not found", e.TaskID)
}

// TaskAlreadyExistsError представляет ошибку, когда задача уже существует
type TaskAlreadyExistsError struct {
	TaskID uuid.UUID
}

func (e *TaskAlreadyExistsError) Error() string {
	return fmt.Sprintf("task with ID %s already exists", e.TaskID)
}

// InvalidTaskDataError представляет ошибку валидации данных задачи
type InvalidTaskDataError struct {
	Field   string
	Message string
}

func (e *InvalidTaskDataError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("invalid task data: field '%s' - %s", e.Field, e.Message)
	}
	return fmt.Sprintf("invalid task data: %s", e.Message)
}

// TaskOperationFailedError представляет ошибку при выполнении операции с задачей
type TaskOperationFailedError struct {
	Operation string
	Reason    string
}

func (e *TaskOperationFailedError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("task operation '%s' failed: %s", e.Operation, e.Reason)
	}
	return fmt.Sprintf("task operation '%s' failed", e.Operation)
}

// IsTaskNotFound проверяет, является ли ошибка ошибкой "задача не найдена"
func IsTaskNotFound(err error) bool {
	var taskNotFoundErr *TaskNotFoundError
	return errors.As(err, &taskNotFoundErr)
}

// IsTaskAlreadyExists проверяет, является ли ошибка ошибкой "задача уже существует"
func IsTaskAlreadyExists(err error) bool {
	var taskExistsErr *TaskAlreadyExistsError
	return errors.As(err, &taskExistsErr)
}

// IsInvalidTaskData проверяет, является ли ошибка ошибкой валидации данных задачи
func IsInvalidTaskData(err error) bool {
	var invalidDataErr *InvalidTaskDataError
	return errors.As(err, &invalidDataErr)
}
