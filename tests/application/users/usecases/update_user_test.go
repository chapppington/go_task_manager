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

func TestUpdateUserUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveFromContainer[*users.CreateUserUseCase](container)
	require.NoError(t, err)

	updateUseCase, err := tests.ResolveFromContainer[*users.UpdateUserUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveFromContainer[*users.GetUserByIDUseCase](container)
	require.NoError(t, err)

	t.Run("update email", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "original@example.com", "Original Name")
		require.NoError(t, err)

		newEmail := "updated@example.com"
		updatedUser, err := updateUseCase.Execute(ctx, user.ID, &newEmail, nil)
		require.NoError(t, err)
		assert.Equal(t, "updated@example.com", updatedUser.Email.Value())
		assert.Equal(t, "Original Name", updatedUser.Name.Value())
	})

	t.Run("update name", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "name-test@example.com", "Original Name")
		require.NoError(t, err)

		newName := "Updated Name"
		updatedUser, err := updateUseCase.Execute(ctx, user.ID, nil, &newName)
		require.NoError(t, err)
		assert.Equal(t, "name-test@example.com", updatedUser.Email.Value())
		assert.Equal(t, "Updated Name", updatedUser.Name.Value())
	})

	t.Run("update all fields", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "all-test@example.com", "Original Name")
		require.NoError(t, err)

		newEmail := "all-updated@example.com"
		newName := "Updated Name"
		updatedUser, err := updateUseCase.Execute(ctx, user.ID, &newEmail, &newName)
		require.NoError(t, err)
		assert.Equal(t, "all-updated@example.com", updatedUser.Email.Value())
		assert.Equal(t, "Updated Name", updatedUser.Name.Value())

		// Проверяем, что изменения сохранились
		retrievedUser, err := getUseCase.Execute(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, "all-updated@example.com", retrievedUser.Email.Value())
		assert.Equal(t, "Updated Name", retrievedUser.Name.Value())
	})

	t.Run("invalid email", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "invalid-test@example.com", "Name")
		require.NoError(t, err)

		invalidEmail := "invalid-email"
		_, err = updateUseCase.Execute(ctx, user.ID, &invalidEmail, nil)
		assert.True(t, vo.IsInvalidEmail(err))
	})

	t.Run("invalid name", func(t *testing.T) {
		user, err := createUseCase.Execute(ctx, "invalid-name-test@example.com", "Name")
		require.NoError(t, err)

		invalidName := ""
		_, err = updateUseCase.Execute(ctx, user.ID, nil, &invalidName)
		assert.True(t, vo.IsInvalidName(err))
	})

	t.Run("user not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		newEmail := "new@example.com"
		_, err := updateUseCase.Execute(ctx, nonExistentID, &newEmail, nil)
		assert.True(t, users_domain.IsUserNotFound(err))
	})
}
