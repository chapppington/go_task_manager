package tasks

import (
	"context"

	"crud/internal/domain/tasks"

	"github.com/google/uuid"
)

// GetTaskByIDUseCase use case для получения задачи по ID
type GetTaskByIDUseCase struct {
	repo tasks.BaseTasksRepository
}

// NewGetTaskByIDUseCase создает новый use case
func NewGetTaskByIDUseCase(repo tasks.BaseTasksRepository) *GetTaskByIDUseCase {
	return &GetTaskByIDUseCase{
		repo: repo,
	}
}

// Execute выполняет получение задачи
func (uc *GetTaskByIDUseCase) Execute(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	return uc.repo.GetByID(ctx, id)
}
