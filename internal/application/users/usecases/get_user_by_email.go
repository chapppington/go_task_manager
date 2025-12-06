package users

import (
	"context"

	"crud/internal/domain/users"
)

// GetUserByEmailUseCase use case для получения пользователя по email
type GetUserByEmailUseCase struct {
	repo users.BaseUsersRepository
}

// NewGetUserByEmailUseCase создает новый use case
func NewGetUserByEmailUseCase(repo users.BaseUsersRepository) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		repo: repo,
	}
}

// Execute выполняет получение пользователя по email
func (uc *GetUserByEmailUseCase) Execute(ctx context.Context, email string) (*users.User, error) {
	return uc.repo.GetByEmail(ctx, email)
}
