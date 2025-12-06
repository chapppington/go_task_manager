package users

import (
	"context"

	"crud/internal/domain/users"

	"github.com/google/uuid"
)

// DeleteUserUseCase use case для удаления пользователя
type DeleteUserUseCase struct {
	repo users.BaseUsersRepository
}

// NewDeleteUserUseCase создает новый use case
func NewDeleteUserUseCase(repo users.BaseUsersRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repo,
	}
}

// Execute выполняет удаление пользователя
func (uc *DeleteUserUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
