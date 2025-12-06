package tasks

import (
	"context"

	"crud/internal/domain/tasks"

	"github.com/google/uuid"
)

// DeleteTaskUseCase use case для удаления задачи
type DeleteTaskUseCase struct {
	repo tasks.BaseTasksRepository
}

// NewDeleteTaskUseCase создает новый use case
func NewDeleteTaskUseCase(repo tasks.BaseTasksRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		repo: repo,
	}
}

// Execute выполняет удаление задачи
func (uc *DeleteTaskUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
