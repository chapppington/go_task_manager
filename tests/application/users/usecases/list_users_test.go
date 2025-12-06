package application

import (
	"context"
	"testing"

	users "crud/internal/application/users/usecases"
	"crud/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListUsersUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveTest[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	listUseCase, err := tests.ResolveTest[*users.ListUsersUseCase](container)
	require.NoError(t, err)

	// Создаем несколько пользователей
	_, err = createUseCase.Execute(ctx, "user1@example.com", "User 1")
	require.NoError(t, err)

	_, err = createUseCase.Execute(ctx, "user2@example.com", "User 2")
	require.NoError(t, err)

	_, err = createUseCase.Execute(ctx, "user3@example.com", "User 3")
	require.NoError(t, err)

	t.Run("list all users", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, 1, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 3)
	})

	t.Run("pagination first page", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, 1, 2)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 2)
	})

	t.Run("pagination second page", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, 2, 2)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 1)
	})

	t.Run("empty page", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, 10, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 0)
	})
}
