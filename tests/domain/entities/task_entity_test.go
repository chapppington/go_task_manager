package entities

import (
	"testing"

	"crud/internal/domain/tasks"
	vo "crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskEntity_Creation(t *testing.T) {
	title, err := vo.NewTaskTitleValueObject("Learn gRPC")
	require.NoError(t, err)

	status, err := vo.NewTaskStatusValueObject("todo")
	require.NoError(t, err)

	userID := uuid.New()
	task := tasks.NewTask(userID, title, "Study gRPC basics", status)

	assert.Equal(t, "Learn gRPC", task.Title.Value())
	assert.Equal(t, "Study gRPC basics", task.Description)
	assert.Equal(t, "todo", task.Status.Value())
	assert.Equal(t, userID, task.UserID)
	assert.NotEqual(t, uuid.Nil, task.ID)
	assert.False(t, task.CreatedAt.IsZero())
	assert.False(t, task.UpdatedAt.IsZero())
}

func TestTaskEntity_Equality(t *testing.T) {
	title1, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title2, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title3, _ := vo.NewTaskTitleValueObject("Master Go")

	status, _ := vo.NewTaskStatusValueObject("todo")
	userID := uuid.New()

	// Создаем задачи с одинаковым ID
	task1 := tasks.NewTask(userID, title1, "Description 1", status)
	task2 := tasks.NewTask(userID, title2, "Description 2", status)
	task2.ID = task1.ID // Устанавливаем одинаковый ID

	// Создаем задачу с другим ID
	task3 := tasks.NewTask(userID, title3, "Description 3", status)

	// Задачи с одинаковым ID должны быть равны
	assert.True(t, task1.Equals(task2), "Expected task1 and task2 to be equal (same ID)")

	// Задачи с разным ID не должны быть равны
	assert.False(t, task1.Equals(task3), "Expected task1 and task3 to be different (different ID)")
}
