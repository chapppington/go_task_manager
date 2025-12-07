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

func TestListTasksUseCase_Execute(t *testing.T) {
	ctx := context.Background()

	// Создаем новый контейнер для теста
	container := tests.NewTestContainer()

	createUseCase, err := tests.ResolveFromContainer[*tasks.CreateTaskUseCase](container)
	require.NoError(t, err)

	listUseCase, err := tests.ResolveFromContainer[*tasks.ListTasksUseCase](container)
	require.NoError(t, err)

	userID1 := uuid.New()
	userID2 := uuid.New()

	// Создаем несколько задач
	task1, err := createUseCase.Execute(ctx, userID1, "Task 1", "Description 1", "todo")
	require.NoError(t, err)

	task2, err := createUseCase.Execute(ctx, userID1, "Task 2", "Description 2", "in_progress")
	require.NoError(t, err)

	_, err = createUseCase.Execute(ctx, userID2, "Task 3", "Description 3", "done")
	require.NoError(t, err)

	t.Run("list all tasks", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, nil, nil, 1, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 3)
	})

	t.Run("filter by userID", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, &userID1, nil, 1, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, result, 2)
		assert.Contains(t, []uuid.UUID{result[0].ID, result[1].ID}, task1.ID)
		assert.Contains(t, []uuid.UUID{result[0].ID, result[1].ID}, task2.ID)
	})

	t.Run("filter by status", func(t *testing.T) {
		status := "todo"
		result, total, err := listUseCase.Execute(ctx, nil, &status, 1, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Len(t, result, 1)
		assert.Equal(t, task1.ID, result[0].ID)
	})

	t.Run("filter by userID and status", func(t *testing.T) {
		status := "in_progress"
		result, total, err := listUseCase.Execute(ctx, &userID1, &status, 1, 10)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Len(t, result, 1)
		assert.Equal(t, task2.ID, result[0].ID)
	})

	t.Run("pagination", func(t *testing.T) {
		result, total, err := listUseCase.Execute(ctx, nil, nil, 1, 2)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Len(t, result, 2)

		result2, total2, err := listUseCase.Execute(ctx, nil, nil, 2, 2)
		require.NoError(t, err)
		assert.Equal(t, int64(3), total2)
		assert.Len(t, result2, 1)
	})
}
