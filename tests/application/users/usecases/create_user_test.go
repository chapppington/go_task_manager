package application

import (
	"context"
	"testing"

	users "crud/internal/application/users/usecases"
	users_domain "crud/internal/domain/users"
	vo "crud/internal/domain/users/value_objects"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	// Получаем use case из контейнера
	useCase, err := tests.ResolveFromContainer[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	t.Run("successful creation", func(t *testing.T) {
		user, err := useCase.Execute(ctx, "test@example.com", "Test User")
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test@example.com", user.Email.Value())
		assert.Equal(t, "Test User", user.Name.Value())
		assert.NotEqual(t, uuid.Nil, user.ID)
		assert.False(t, user.CreatedAt.IsZero())
		assert.False(t, user.UpdatedAt.IsZero())
	})

	t.Run("invalid email", func(t *testing.T) {
		user, err := useCase.Execute(ctx, "invalid-email", "Test User")
		assert.Nil(t, user)
		assert.True(t, vo.IsInvalidEmail(err))
	})

	t.Run("empty email", func(t *testing.T) {
		user, err := useCase.Execute(ctx, "", "Test User")
		assert.Nil(t, user)
		assert.True(t, vo.IsInvalidEmail(err))
	})

	t.Run("empty name", func(t *testing.T) {
		user, err := useCase.Execute(ctx, "test@example.com", "")
		assert.Nil(t, user)
		assert.True(t, vo.IsInvalidName(err))
	})

	t.Run("duplicate email", func(t *testing.T) {
		email := "duplicate@example.com"
		user1, err := useCase.Execute(ctx, email, "User 1")
		require.NoError(t, err)
		require.NotNil(t, user1)

		user2, err := useCase.Execute(ctx, email, "User 2")
		assert.Nil(t, user2)
		assert.True(t, users_domain.IsUserAlreadyExists(err))
	})
}
