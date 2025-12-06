package tasks

import (
	"context"

	"crud/internal/domain/tasks"
	vo "crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
)

// CreateTaskUseCase use case для создания задачи
type CreateTaskUseCase struct {
	repo tasks.BaseTasksRepository
}

// NewCreateTaskUseCase создает новый use case
func NewCreateTaskUseCase(repo tasks.BaseTasksRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		repo: repo,
	}
}

// Execute выполняет создание задачи
func (uc *CreateTaskUseCase) Execute(
	ctx context.Context,
	userID uuid.UUID,
	title string,
	description string,
	status string,
) (*tasks.Task, error) {
	titleVO, err := vo.NewTaskTitleValueObject(title)
	if err != nil {
		return nil, err
	}

	statusVO, err := vo.NewTaskStatusValueObject(status)
	if err != nil {
		return nil, err
	}

	task := tasks.NewTask(userID, titleVO, description, statusVO)
	return uc.repo.Create(ctx, task)
}
