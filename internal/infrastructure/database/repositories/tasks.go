package repositories

import (
	"context"
	"errors"

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
		return nil, &tasks.InvalidTaskDataError{Field: "task", Message: "task cannot be nil"}
	}

	model := converters.TaskEntityToModel(task)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		// Проверяем, не является ли это ошибкой дубликата (если есть уникальный индекс)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &tasks.TaskAlreadyExistsError{TaskID: task.ID}
		}
		return nil, &tasks.TaskOperationFailedError{Operation: "create", Reason: err.Error()}
	}

	return converters.TaskModelToEntity(model)
}

// GetByID возвращает задачу по ID
func (r *TasksRepository) GetByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	var model models.Task
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &tasks.TaskNotFoundError{TaskID: id}
		}
		return nil, &tasks.TaskOperationFailedError{Operation: "get_by_id", Reason: err.Error()}
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
		return nil, 0, &tasks.TaskOperationFailedError{Operation: "list_count", Reason: err.Error()}
	}

	// Получение данных с пагинацией
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&taskModels).Error; err != nil {
		return nil, 0, &tasks.TaskOperationFailedError{Operation: "list", Reason: err.Error()}
	}

	domainTasks := make([]*tasks.Task, 0, len(taskModels))
	for _, model := range taskModels {
		task, err := converters.TaskModelToEntity(model)
		if err != nil {
			return nil, 0, &tasks.TaskOperationFailedError{Operation: "list_convert", Reason: err.Error()}
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
		return nil, &tasks.InvalidTaskDataError{Field: "task", Message: "task cannot be nil"}
	}

	model := converters.TaskEntityToModel(task)
	model.UpdatedAt = task.UpdatedAt

	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &tasks.TaskNotFoundError{TaskID: task.ID}
		}
		return nil, &tasks.TaskOperationFailedError{Operation: "update", Reason: err.Error()}
	}

	return converters.TaskModelToEntity(model)
}

// Delete удаляет задачу по ID
func (r *TasksRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", id)
	if result.Error != nil {
		return &tasks.TaskOperationFailedError{Operation: "delete", Reason: result.Error.Error()}
	}
	if result.RowsAffected == 0 {
		return &tasks.TaskNotFoundError{TaskID: id}
	}
	return nil
}
