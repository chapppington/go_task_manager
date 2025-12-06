package value_objects

import (
	"fmt"
	"slices"
	"strings"
)

var validStatuses = []string{
	"todo",
	"in_progress",
	"done",
}

// TaskStatusValueObject представляет статус задачи с валидацией
type TaskStatusValueObject struct {
	value string
}

// NewTaskStatusValueObject создает новый TaskStatusValueObject с валидацией
func NewTaskStatusValueObject(status string) (TaskStatusValueObject, error) {
	status = strings.TrimSpace(status)
	if status == "" {
		return TaskStatusValueObject{}, fmt.Errorf("status cannot be empty")
	}
	if !slices.Contains(validStatuses, status) {
		return TaskStatusValueObject{}, fmt.Errorf("invalid status: %s. Valid statuses: %s", status, strings.Join(validStatuses, ", "))
	}
	return TaskStatusValueObject{value: status}, nil
}

// Value возвращает строковое значение статуса
func (s TaskStatusValueObject) Value() string {
	return s.value
}

// Equals проверяет равенство двух статусов
func (s TaskStatusValueObject) Equals(other TaskStatusValueObject) bool {
	return s.value == other.value
}

// IsValid проверяет валидность статуса
func (s TaskStatusValueObject) IsValid() bool {
	return slices.Contains(validStatuses, s.value)
}
