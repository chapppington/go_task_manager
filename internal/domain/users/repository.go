package users

import "context"

// Repository определяет интерфейс для работы с пользователями
type Repository interface {
	// Create создает нового пользователя
	Create(ctx context.Context, user *User) (*User, error)

	// GetByID возвращает пользователя по ID
	GetByID(ctx context.Context, id int64) (*User, error)

	// GetByEmail возвращает пользователя по email
	GetByEmail(ctx context.Context, email string) (*User, error)

	// List возвращает список пользователей с пагинацией
	List(ctx context.Context, page, pageSize int) ([]*User, int64, error)

	// Update обновляет данные пользователя
	Update(ctx context.Context, user *User) (*User, error)

	// Delete удаляет пользователя по ID
	Delete(ctx context.Context, id int64) error
}
