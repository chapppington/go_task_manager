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
	title2, _ := vo.NewTaskTitleValueObject("Master Go")

	status, _ := vo.NewTaskStatusValueObject("todo")
	userID := uuid.New()

	task1 := tasks.NewTask(userID, title1, "Description 1", status)
	task2 := tasks.NewTask(userID, title2, "Description 2", status)

	assert.False(t, task1.Equals(task2), "Expected task1 and task2 to be different")
}
