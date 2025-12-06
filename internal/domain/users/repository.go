package users

import (
	"context"

	"github.com/google/uuid"
)

// BaseUsersRepository определяет интерфейс для работы с пользователями
type BaseUsersRepository interface {
	// Create создает нового пользователя
	Create(ctx context.Context, user *User) (*User, error)

	// GetByID возвращает пользователя по ID
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)

	// GetByEmail возвращает пользователя по email
	GetByEmail(ctx context.Context, email string) (*User, error)

	// List возвращает список пользователей с пагинацией
	List(ctx context.Context, page, pageSize int) ([]*User, int64, error)

	// Update обновляет данные пользователя
	Update(ctx context.Context, user *User) (*User, error)

	// Delete удаляет пользователя по ID
	Delete(ctx context.Context, id uuid.UUID) error
}
