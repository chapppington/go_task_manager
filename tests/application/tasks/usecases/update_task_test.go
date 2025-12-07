package application

import (
	"context"
	"testing"

	tasks "crud/internal/application/tasks/usecases"
	tasks_domain "crud/internal/domain/tasks"
	vo "crud/internal/domain/tasks/value_objects"
	"crud/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTaskUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveFromContainer[*tasks.CreateTaskUseCase](container)
	require.NoError(t, err)

	updateUseCase, err := tests.ResolveFromContainer[*tasks.UpdateTaskUseCase](container)
	require.NoError(t, err)

	getUseCase, err := tests.ResolveFromContainer[*tasks.GetTaskByIDUseCase](container)
	require.NoError(t, err)

	userID := uuid.New()

	t.Run("update title", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Original Title", "Description", "todo")
		require.NoError(t, err)

		newTitle := "Updated Title"
		updatedTask, err := updateUseCase.Execute(ctx, task.ID, &newTitle, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "Updated Title", updatedTask.Title.Value())
		assert.Equal(t, "Description", updatedTask.Description)
		assert.Equal(t, "todo", updatedTask.Status.Value())
	})

	t.Run("update description", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Title", "Original Description", "todo")
		require.NoError(t, err)

		newDescription := "Updated Description"
		updatedTask, err := updateUseCase.Execute(ctx, task.ID, nil, &newDescription, nil)
		require.NoError(t, err)
		assert.Equal(t, "Title", updatedTask.Title.Value())
		assert.Equal(t, "Updated Description", updatedTask.Description)
		assert.Equal(t, "todo", updatedTask.Status.Value())
	})

	t.Run("update status", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Title", "Description", "todo")
		require.NoError(t, err)

		newStatus := "in_progress"
		updatedTask, err := updateUseCase.Execute(ctx, task.ID, nil, nil, &newStatus)
		require.NoError(t, err)
		assert.Equal(t, "Title", updatedTask.Title.Value())
		assert.Equal(t, "Description", updatedTask.Description)
		assert.Equal(t, "in_progress", updatedTask.Status.Value())
	})

	t.Run("update all fields", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Title", "Description", "todo")
		require.NoError(t, err)

		newTitle := "New Title"
		newDescription := "New Description"
		newStatus := "done"
		updatedTask, err := updateUseCase.Execute(ctx, task.ID, &newTitle, &newDescription, &newStatus)
		require.NoError(t, err)
		assert.Equal(t, "New Title", updatedTask.Title.Value())
		assert.Equal(t, "New Description", updatedTask.Description)
		assert.Equal(t, "done", updatedTask.Status.Value())

		// Проверяем, что изменения сохранились
		retrievedTask, err := getUseCase.Execute(ctx, task.ID)
		require.NoError(t, err)
		assert.Equal(t, "New Title", retrievedTask.Title.Value())
		assert.Equal(t, "New Description", retrievedTask.Description)
		assert.Equal(t, "done", retrievedTask.Status.Value())
	})

	t.Run("invalid title", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Title", "Description", "todo")
		require.NoError(t, err)

		invalidTitle := ""
		_, err = updateUseCase.Execute(ctx, task.ID, &invalidTitle, nil, nil)
		assert.True(t, vo.IsInvalidTitle(err))
	})

	t.Run("invalid status", func(t *testing.T) {
		task, err := createUseCase.Execute(ctx, userID, "Title", "Description", "todo")
		require.NoError(t, err)

		invalidStatus := "invalid_status"
		_, err = updateUseCase.Execute(ctx, task.ID, nil, nil, &invalidStatus)
		assert.True(t, vo.IsInvalidStatus(err))
	})

	t.Run("task not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		newTitle := "New Title"
		_, err := updateUseCase.Execute(ctx, nonExistentID, &newTitle, nil, nil)
		assert.True(t, tasks_domain.IsTaskNotFound(err))
	})
}
