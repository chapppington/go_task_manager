package tasks

import (
	"context"

	"github.com/google/uuid"
)

// Repository определяет интерфейс для работы с задачами
type Repository interface {
	// Create создает новую задачу
	Create(ctx context.Context, task *Task) (*Task, error)

	// GetByID возвращает задачу по ID
	GetByID(ctx context.Context, id uuid.UUID) (*Task, error)

	// List возвращает список задач с фильтрацией и пагинацией
	List(ctx context.Context, userID *uuid.UUID, status *string, page, pageSize int) ([]*Task, int64, error)

	// Update обновляет данные задачи
	Update(ctx context.Context, task *Task) (*Task, error)

	// Delete удаляет задачу по ID
	Delete(ctx context.Context, id uuid.UUID) error
}
