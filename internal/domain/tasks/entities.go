package tasks

import (
	"time"

	"crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
)

// Task представляет сущность задачи
type Task struct {
	ID          int64
	OID         uuid.UUID // Object ID для сравнения
	UserID      int64
	Title       value_objects.TaskTitleValueObject
	Description string
	Status      value_objects.TaskStatusValueObject
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask создает новую задачу
func NewTask(
	userID int64,
	title value_objects.TaskTitleValueObject,
	description string,
	status value_objects.TaskStatusValueObject,
) *Task {
	now := time.Now()
	return &Task{
		OID:         uuid.New(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Equals проверяет равенство двух задач по хешу
func (t *Task) Equals(other *Task) bool {
	if t == nil || other == nil {
		return t == other
	}
	return t.Hash() == other.Hash()
}

// Hash возвращает хеш задачи (для использования в map)
func (t *Task) Hash() string {
	return t.OID.String()
}
