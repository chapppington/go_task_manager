package users

import (
	"context"

	"crud/internal/domain/users"
	vo "crud/internal/domain/users/value_objects"
)

// CreateUserUseCase use case для создания пользователя
type CreateUserUseCase struct {
	repo users.BaseUsersRepository
}

// NewCreateUserUseCase создает новый use case
func NewCreateUserUseCase(repo users.BaseUsersRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

// Execute выполняет создание пользователя
func (uc *CreateUserUseCase) Execute(
	ctx context.Context,
	email string,
	name string,
) (*users.User, error) {
	emailVO, err := vo.NewEmailValueObject(email)
	if err != nil {
		return nil, err
	}

	nameVO, err := vo.NewUserNameValueObject(name)
	if err != nil {
		return nil, err
	}

	user := users.NewUser(emailVO, nameVO)
	return uc.repo.Create(ctx, user)
}
