package application

import (
	"context"
	"testing"

	users "crud/internal/application/users/usecases"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByIDUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	// Получаем use cases из контейнера
	createUseCase, err := tests.ResolveTest[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveTest[*users.GetUserByIDUseCase](container)
	require.NoError(t, err)

	t.Run("successful retrieval", func(t *testing.T) {
		// Создаем пользователя с уникальным email
		email := "getbyid-test@example.com"
		createdUser, err := createUseCase.Execute(ctx, email, "Test User")
		require.NoError(t, err)
		require.NotNil(t, createdUser)

		// Получаем пользователя по ID
		user, err := getUseCase.Execute(ctx, createdUser.ID)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.Equal(t, createdUser.Email.Value(), user.Email.Value())
		assert.Equal(t, createdUser.Name.Value(), user.Name.Value())
	})

	t.Run("user not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		user, err := getUseCase.Execute(ctx, nonExistentID)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "not found")
	})
}
