package dummy

import (
	"context"
	"fmt"
	"sync"

	"crud/internal/domain/users"

	"github.com/google/uuid"
)

// UsersRepository in-memory реализация репозитория пользователей
type UsersRepository struct {
	mu    sync.RWMutex
	users []*users.User
}

// NewUsersRepository создает новый in-memory репозиторий пользователей
func NewUsersRepository() *UsersRepository {
	return &UsersRepository{
		users: make([]*users.User, 0),
	}
}

// Create создает нового пользователя
func (r *UsersRepository) Create(ctx context.Context, user *users.User) (*users.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	email := user.Email.Value()
	for _, u := range r.users {
		if u.Email.Value() == email {
			return nil, fmt.Errorf("user with email %s already exists", email)
		}
	}

	r.users = append(r.users, user)
	return user, nil
}

// GetByID возвращает пользователя по ID
func (r *UsersRepository) GetByID(ctx context.Context, id uuid.UUID) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found: %s", id)
}

// GetByEmail возвращает пользователя по email
func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email.Value() == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with email %s not found", email)
}

// List возвращает список пользователей с пагинацией
func (r *UsersRepository) List(ctx context.Context, page, pageSize int) ([]*users.User, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := int64(len(r.users))

	// Пагинация
	start := (page - 1) * pageSize
	if start < 0 {
		start = 0
	}
	end := start + pageSize
	if end > len(r.users) {
		end = len(r.users)
	}

	if start >= len(r.users) {
		return []*users.User{}, total, nil
	}

	return r.users[start:end], total, nil
}

// Update обновляет данные пользователя
func (r *UsersRepository) Update(ctx context.Context, user *users.User) (*users.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found: %s", user.ID)
}

// Delete удаляет пользователя по ID
func (r *UsersRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user not found: %s", id)
}
