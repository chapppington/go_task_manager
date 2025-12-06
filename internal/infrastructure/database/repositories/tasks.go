package repositories

import (
	"context"
	"fmt"

	"crud/internal/domain/tasks"
	"crud/internal/infrastructure/database/converters"
	"crud/internal/infrastructure/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TasksRepository GORM реализация репозитория задач
type TasksRepository struct {
	db *gorm.DB
}

// NewTasksRepository создает новый GORM репозиторий задач
func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

// Create создает новую задачу
func (r *TasksRepository) Create(ctx context.Context, task *tasks.Task) (*tasks.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task cannot be nil")
	}

	model := converters.TaskEntityToModel(task)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return converters.TaskModelToEntity(model)
}

// GetByID возвращает задачу по ID
func (r *TasksRepository) GetByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	var model models.Task
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task not found: %s", id)
		}
		return nil, fmt.Errorf("failed to get task by id: %w", err)
	}

	return converters.TaskModelToEntity(&model)
}

// List возвращает список задач с фильтрацией и пагинацией
func (r *TasksRepository) List(
	ctx context.Context,
	userID *uuid.UUID,
	status *string,
	page, pageSize int,
) ([]*tasks.Task, int64, error) {
	var taskModels []*models.Task
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Task{})

	// Применение фильтров
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// Подсчет общего количества с учетом фильтров
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count tasks: %w", err)
	}

	// Получение данных с пагинацией
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&taskModels).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list tasks: %w", err)
	}

	domainTasks := make([]*tasks.Task, 0, len(taskModels))
	for _, model := range taskModels {
		task, err := converters.TaskModelToEntity(model)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to convert task: %w", err)
		}
		if task != nil {
			domainTasks = append(domainTasks, task)
		}
	}

	return domainTasks, total, nil
}

// Update обновляет данные задачи
func (r *TasksRepository) Update(ctx context.Context, task *tasks.Task) (*tasks.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task cannot be nil")
	}

	model := converters.TaskEntityToModel(task)
	model.UpdatedAt = task.UpdatedAt

	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task not found: %s", task.ID)
		}
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return converters.TaskModelToEntity(model)
}

// Delete удаляет задачу по ID
func (r *TasksRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete task: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task not found: %s", id)
	}
	return nil
}
