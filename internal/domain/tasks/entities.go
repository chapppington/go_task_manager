package tasks

import (
	"time"

	"crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
)

// Task представляет сущность задачи
type Task struct {
	ID          uuid.UUID // Object ID для сравнения
	UserID      uuid.UUID
	Title       value_objects.TaskTitleValueObject
	Description string
	Status      value_objects.TaskStatusValueObject
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask создает новую задачу
func NewTask(
	userID uuid.UUID,
	title value_objects.TaskTitleValueObject,
	description string,
	status value_objects.TaskStatusValueObject,
) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Equals проверяет равенство двух задач по ID
func (t *Task) Equals(other *Task) bool {
	if t == nil || other == nil {
		return t == other
	}
	return t.ID == other.ID
}
