package application

import (
	"context"
	"testing"

	tasks "crud/internal/application/tasks/usecases"
	vo "crud/internal/domain/tasks/value_objects"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	// Получаем use case из контейнера
	useCase, err := tests.ResolveFromContainer[*tasks.CreateTaskUseCase](container)
	require.NoError(t, err)

	// Создаем пользователя для теста
	userID := uuid.New()

	t.Run("successful creation", func(t *testing.T) {
		task, err := useCase.Execute(ctx, userID, "Test Task", "Test Description", "todo")
		require.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, userID, task.UserID)
		assert.Equal(t, "Test Task", task.Title.Value())
		assert.Equal(t, "Test Description", task.Description)
		assert.Equal(t, "todo", task.Status.Value())
		assert.NotEqual(t, uuid.Nil, task.ID)
		assert.False(t, task.CreatedAt.IsZero())
		assert.False(t, task.UpdatedAt.IsZero())
	})

	t.Run("invalid title", func(t *testing.T) {
		task, err := useCase.Execute(ctx, userID, "", "Test Description", "todo")
		assert.Nil(t, task)
		assert.True(t, vo.IsInvalidTitle(err))
	})

	t.Run("invalid status", func(t *testing.T) {
		task, err := useCase.Execute(ctx, userID, "Test Task", "Test Description", "invalid_status")
		assert.Nil(t, task)
		assert.True(t, vo.IsInvalidStatus(err))
	})

	t.Run("valid statuses", func(t *testing.T) {
		validStatuses := []string{"todo", "in_progress", "done"}
		for _, status := range validStatuses {
			task, err := useCase.Execute(ctx, userID, "Test Task", "Test Description", status)
			require.NoError(t, err, "status: %s", status)
			assert.Equal(t, status, task.Status.Value())
		}
	})
}
