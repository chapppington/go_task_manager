package dummy

import (
	"context"
	"sync"

	"crud/internal/domain/tasks"

	"github.com/google/uuid"
)

// TasksRepository in-memory реализация репозитория задач
type TasksRepository struct {
	mu    sync.RWMutex
	tasks []*tasks.Task
}

// NewTasksRepository создает новый in-memory репозиторий задач
func NewTasksRepository() *TasksRepository {
	return &TasksRepository{
		tasks: make([]*tasks.Task, 0),
	}
}

// Create создает новую задачу
func (r *TasksRepository) Create(ctx context.Context, task *tasks.Task) (*tasks.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task == nil {
		return nil, &tasks.InvalidTaskDataError{Field: "task", Message: "task cannot be nil"}
	}

	// Проверяем, не существует ли уже задача с таким ID
	for _, t := range r.tasks {
		if t.ID == task.ID {
			return nil, &tasks.TaskAlreadyExistsError{TaskID: task.ID}
		}
	}

	r.tasks = append(r.tasks, task)
	return task, nil
}

// GetByID возвращает задачу по ID
func (r *TasksRepository) GetByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, &tasks.TaskNotFoundError{TaskID: id}
}

// List возвращает список задач с фильтрацией и пагинацией
func (r *TasksRepository) List(
	ctx context.Context,
	userID *uuid.UUID,
	status *string,
	page, pageSize int,
) ([]*tasks.Task, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []*tasks.Task
	for _, task := range r.tasks {
		if userID != nil && task.UserID != *userID {
			continue
		}
		if status != nil && task.Status.Value() != *status {
			continue
		}
		filtered = append(filtered, task)
	}

	total := int64(len(filtered))

	// Пагинация
	start := (page - 1) * pageSize
	if start < 0 {
		start = 0
	}
	end := start + pageSize
	if end > len(filtered) {
		end = len(filtered)
	}

	if start >= len(filtered) {
		return []*tasks.Task{}, total, nil
	}

	return filtered[start:end], total, nil
}

// Update обновляет данные задачи
func (r *TasksRepository) Update(ctx context.Context, task *tasks.Task) (*tasks.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task == nil {
		return nil, &tasks.InvalidTaskDataError{Field: "task", Message: "task cannot be nil"}
	}

	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			return task, nil
		}
	}

	return nil, &tasks.TaskNotFoundError{TaskID: task.ID}
}

// Delete удаляет задачу по ID
func (r *TasksRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}

	return &tasks.TaskNotFoundError{TaskID: id}
}
