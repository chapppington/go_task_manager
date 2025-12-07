package application

import (
	"context"
	"testing"

	users "crud/internal/application/users/usecases"
	users_domain "crud/internal/domain/users"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteUserUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveFromContainer[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	deleteUseCase, err := tests.ResolveFromContainer[*users.DeleteUserUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveFromContainer[*users.GetUserByIDUseCase](container)
	require.NoError(t, err)

	t.Run("successful deletion", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "delete-test@example.com", "User to Delete")
		require.NoError(t, err)

		err = deleteUseCase.Execute(ctx, user.ID)
		require.NoError(t, err)

		// Проверяем, что пользователь удален
		_, err = getUseCase.Execute(ctx, user.ID)
		assert.True(t, users_domain.IsUserNotFound(err))
	})

	t.Run("user not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		err := deleteUseCase.Execute(ctx, nonExistentID)
		assert.True(t, users_domain.IsUserNotFound(err))
	})
}
