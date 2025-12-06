package tasks

import (
	"context"

	"crud/internal/domain/tasks"

	"github.com/google/uuid"
)

// ListTasksUseCase use case для получения списка задач
type ListTasksUseCase struct {
	repo tasks.BaseTasksRepository
}

// NewListTasksUseCase создает новый use case
func NewListTasksUseCase(repo tasks.BaseTasksRepository) *ListTasksUseCase {
	return &ListTasksUseCase{
		repo: repo,
	}
}

// Execute выполняет получение списка задач
func (uc *ListTasksUseCase) Execute(
	ctx context.Context,
	userID *uuid.UUID,
	status *string,
	page, pageSize int,
) ([]*tasks.Task, int64, error) {
	return uc.repo.List(ctx, userID, status, page, pageSize)
}
