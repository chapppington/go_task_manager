package application

import (
	"context"
	"testing"

	users "crud/internal/application/users/usecases"
	users_domain "crud/internal/domain/users"
	"crud/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByEmailUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveTest[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	getByEmailUseCase, err := tests.ResolveTest[*users.GetUserByEmailUseCase](container)
	require.NoError(t, err)

	t.Run("successful retrieval", func(t *testing.T) {
		email := "getbyemail-test@example.com"
		createdUser, err := createUseCase.Execute(ctx, email, "Test User")
		require.NoError(t, err)
		require.NotNil(t, createdUser)

		user, err := getByEmailUseCase.Execute(ctx, email)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.Equal(t, createdUser.Email.Value(), user.Email.Value())
		assert.Equal(t, createdUser.Name.Value(), user.Name.Value())
	})

	t.Run("user not found", func(t *testing.T) {
		_, err := getByEmailUseCase.Execute(ctx, "nonexistent@example.com")
		assert.True(t, users_domain.IsUserNotFound(err))
	})
}
