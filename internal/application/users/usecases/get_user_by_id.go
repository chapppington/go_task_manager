package users

import (
	"context"

	"crud/internal/domain/users"

	"github.com/google/uuid"
)

// GetUserByIDUseCase use case для получения пользователя по ID
type GetUserByIDUseCase struct {
	repo users.BaseUsersRepository
}

// NewGetUserByIDUseCase создает новый use case
func NewGetUserByIDUseCase(repo users.BaseUsersRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		repo: repo,
	}
}

// Execute выполняет получение пользователя
func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id uuid.UUID) (*users.User, error) {
	return uc.repo.GetByID(ctx, id)
}
