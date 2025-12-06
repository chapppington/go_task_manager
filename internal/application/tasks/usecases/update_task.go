package tasks

import (
	"context"

	"crud/internal/domain/tasks"
	vo "crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
)

// UpdateTaskUseCase use case для обновления задачи
type UpdateTaskUseCase struct {
	repo tasks.BaseTasksRepository
}

// NewUpdateTaskUseCase создает новый use case
func NewUpdateTaskUseCase(repo tasks.BaseTasksRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		repo: repo,
	}
}

// Execute выполняет обновление задачи
func (uc *UpdateTaskUseCase) Execute(
	ctx context.Context,
	id uuid.UUID,
	titleStr *string,
	description *string,
	statusStr *string,
) (*tasks.Task, error) {
	// Получаем существующую задачу
	task, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем поля, если они переданы
	if titleStr != nil {
		title, err := vo.NewTaskTitleValueObject(*titleStr)
		if err != nil {
			return nil, err
		}
		task.Title = title
	}

	if description != nil {
		task.Description = *description
	}

	if statusStr != nil {
		status, err := vo.NewTaskStatusValueObject(*statusStr)
		if err != nil {
			return nil, err
		}
		task.Status = status
	}

	return uc.repo.Update(ctx, task)
}
