package application

import (
	"context"
	"testing"

	tasks "crud/internal/application/tasks/usecases"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTaskByIDUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	// Получаем use cases и репозиторий из контейнера
	createUseCase, err := tests.ResolveTest[*tasks.CreateTaskUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveTest[*tasks.GetTaskByIDUseCase](container)
	require.NoError(t, err)

	userID := uuid.New()

	t.Run("successful retrieval", func(t *testing.T) {
		// Создаем задачу
		createdTask, err := createUseCase.Execute(ctx, userID, "Test Task", "Test Description", "todo")
		require.NoError(t, err)
		require.NotNil(t, createdTask)

		// Получаем задачу по ID
		task, err := getUseCase.Execute(ctx, createdTask.ID)
		require.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, createdTask.ID, task.ID)
		assert.Equal(t, createdTask.UserID, task.UserID)
		assert.Equal(t, createdTask.Title.Value(), task.Title.Value())
		assert.Equal(t, createdTask.Description, task.Description)
		assert.Equal(t, createdTask.Status.Value(), task.Status.Value())
	})

	t.Run("task not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		task, err := getUseCase.Execute(ctx, nonExistentID)
		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Contains(t, err.Error(), "not found")
	})
}
