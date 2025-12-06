package users

import (
	"context"

	"crud/internal/domain/users"
)

// ListUsersUseCase use case для получения списка пользователей
type ListUsersUseCase struct {
	repo users.BaseUsersRepository
}

// NewListUsersUseCase создает новый use case
func NewListUsersUseCase(repo users.BaseUsersRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		repo: repo,
	}
}

// Execute выполняет получение списка пользователей
func (uc *ListUsersUseCase) Execute(ctx context.Context, page, pageSize int) ([]*users.User, int64, error) {
	return uc.repo.List(ctx, page, pageSize)
}
