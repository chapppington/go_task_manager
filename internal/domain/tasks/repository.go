package tasks

import "context"

// Repository определяет интерфейс для работы с задачами
type Repository interface {
	// Create создает новую задачу
	Create(ctx context.Context, task *Task) (*Task, error)

	// GetByID возвращает задачу по ID
	GetByID(ctx context.Context, id int64) (*Task, error)

	// List возвращает список задач с фильтрацией и пагинацией
	List(ctx context.Context, userID *int64, status *string, page, pageSize int) ([]*Task, int64, error)

	// Update обновляет данные задачи
	Update(ctx context.Context, task *Task) (*Task, error)

	// Delete удаляет задачу по ID
	Delete(ctx context.Context, id int64) error
}
