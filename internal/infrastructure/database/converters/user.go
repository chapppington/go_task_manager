package converters

import (
	"crud/internal/domain/users"
	"crud/internal/domain/users/value_objects"
	"crud/internal/infrastructure/database/models"
)

// UserModelToEntity конвертирует GORM модель в domain entity
func UserModelToEntity(model *models.User) (*users.User, error) {
	if model == nil {
		return nil, nil
	}

	email, err := value_objects.NewEmailValueObject(model.Email)
	if err != nil {
		return nil, err
	}

	name, err := value_objects.NewUserNameValueObject(model.Name)
	if err != nil {
		return nil, err
	}

	return &users.User{
		ID:        model.ID,
		Email:     email,
		Name:      name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

// UserEntityToModel конвертирует domain entity в GORM модель
func UserEntityToModel(user *users.User) *models.User {
	if user == nil {
		return nil
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email.Value(),
		Name:      user.Name.Value(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
