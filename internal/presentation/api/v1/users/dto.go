package users

import (
	"time"

	users_domain "crud/internal/domain/users"
)

// CreateUserRequest запрос на создание пользователя
type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UpdateUserRequest запрос на обновление пользователя
type UpdateUserRequest struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// UserResponse ответ с данными пользователя
type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ResponseFromEntity создает UserResponse из сущности пользователя
func ResponseFromEntity(user *users_domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email.Value(),
		Name:      user.Name.Value(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
