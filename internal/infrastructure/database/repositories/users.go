package repositories

import (
	"context"
	"fmt"

	"crud/internal/domain/users"
	"crud/internal/infrastructure/database/converters"
	"crud/internal/infrastructure/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UsersRepository GORM реализация репозитория пользователей
type UsersRepository struct {
	db *gorm.DB
}

// NewUsersRepository создает новый GORM репозиторий пользователей
func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create создает нового пользователя
func (r *UsersRepository) Create(ctx context.Context, user *users.User) (*users.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	model := converters.UserEntityToModel(user)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return converters.UserModelToEntity(model)
}

// GetByID возвращает пользователя по ID
func (r *UsersRepository) GetByID(ctx context.Context, id uuid.UUID) (*users.User, error) {
	var model models.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %s", id)
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return converters.UserModelToEntity(&model)
}

// GetByEmail возвращает пользователя по email
func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	var model models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return converters.UserModelToEntity(&model)
}

// List возвращает список пользователей с пагинацией
func (r *UsersRepository) List(ctx context.Context, page, pageSize int) ([]*users.User, int64, error) {
	var userModels []*models.User
	var total int64

	// Подсчет общего количества
	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	// Получение данных с пагинацией
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&userModels).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}

	domainUsers := make([]*users.User, 0, len(userModels))
	for _, model := range userModels {
		user, err := converters.UserModelToEntity(model)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to convert user: %w", err)
		}
		if user != nil {
			domainUsers = append(domainUsers, user)
		}
	}

	return domainUsers, total, nil
}

// Update обновляет данные пользователя
func (r *UsersRepository) Update(ctx context.Context, user *users.User) (*users.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	model := converters.UserEntityToModel(user)
	model.UpdatedAt = user.UpdatedAt

	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %s", user.ID)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return converters.UserModelToEntity(model)
}

// Delete удаляет пользователя по ID
func (r *UsersRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found: %s", id)
	}
	return nil
}
