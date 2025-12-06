package users

import (
	"context"

	"crud/internal/domain/users"
	vo "crud/internal/domain/users/value_objects"

	"github.com/google/uuid"
)

// UpdateUserUseCase use case для обновления пользователя
type UpdateUserUseCase struct {
	repo users.BaseUsersRepository
}

// NewUpdateUserUseCase создает новый use case
func NewUpdateUserUseCase(repo users.BaseUsersRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repo,
	}
}

// Execute выполняет обновление пользователя
func (uc *UpdateUserUseCase) Execute(
	ctx context.Context,
	id uuid.UUID,
	emailStr *string,
	nameStr *string,
) (*users.User, error) {
	// Получаем существующего пользователя
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем поля, если они переданы
	if emailStr != nil {
		email, err := vo.NewEmailValueObject(*emailStr)
		if err != nil {
			return nil, err
		}
		user.Email = email
	}

	if nameStr != nil {
		name, err := vo.NewUserNameValueObject(*nameStr)
		if err != nil {
			return nil, err
		}
		user.Name = name
	}

	return uc.repo.Update(ctx, user)
}
