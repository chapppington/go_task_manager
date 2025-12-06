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

func TestDeleteTaskUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveTest[*tasks.CreateTaskUseCase](container)
	require.NoError(t, err)

	deleteUseCase, err := tests.ResolveTest[*tasks.DeleteTaskUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveTest[*tasks.GetTaskByIDUseCase](container)
	require.NoError(t, err)

	userID := uuid.New()

	t.Run("successful deletion", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Task to Delete", "Description", "todo")
		require.NoError(t, err)

		err = deleteUseCase.Execute(ctx, task.ID)
		require.NoError(t, err)

		// Проверяем, что задача удалена
		_, err = getUseCase.Execute(ctx, task.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("task not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		err := deleteUseCase.Execute(ctx, nonExistentID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}
