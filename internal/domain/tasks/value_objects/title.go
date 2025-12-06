package value_objects

import (
	"fmt"
	"strings"
)

// TaskTitleValueObject представляет заголовок задачи с валидацией
type TaskTitleValueObject struct {
	value string
}

// NewTaskTitleValueObject создает новый TaskTitleValueObject с валидацией
func NewTaskTitleValueObject(title string) (TaskTitleValueObject, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return TaskTitleValueObject{}, fmt.Errorf("title cannot be empty")
	}
	if len(title) < 1 {
		return TaskTitleValueObject{}, fmt.Errorf("title must be at least 1 character long")
	}
	if len(title) > 200 {
		return TaskTitleValueObject{}, fmt.Errorf("title must be at most 200 characters long")
	}
	return TaskTitleValueObject{value: title}, nil
}

// Value возвращает строковое значение заголовка
func (t TaskTitleValueObject) Value() string {
	return t.value
}

// Equals проверяет равенство двух заголовков
func (t TaskTitleValueObject) Equals(other TaskTitleValueObject) bool {
	return t.value == other.value
}
