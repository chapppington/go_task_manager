package value_objects

import (
	"testing"

	vo "crud/internal/domain/tasks/value_objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskStatusValueObject(t *testing.T) {
	// Тест валидных статусов
	validStatuses := []string{"todo", "in_progress", "done"}
	for _, statusStr := range validStatuses {
		status, err := vo.NewTaskStatusValueObject(statusStr)
		require.NoError(t, err, "Expected no error for valid status '%s'", statusStr)
		assert.Equal(t, statusStr, status.Value())
		assert.True(t, status.IsValid(), "Expected status '%s' to be valid", statusStr)
	}

	// Тест пустого статуса
	_, err := vo.NewTaskStatusValueObject("")
	assert.True(t, vo.IsInvalidStatus(err))

	// Тест невалидного статуса
	_, err = vo.NewTaskStatusValueObject("invalid_status")
	assert.True(t, vo.IsInvalidStatus(err))

	// Тест сравнения статусов
	status1, _ := vo.NewTaskStatusValueObject("todo")
	status2, _ := vo.NewTaskStatusValueObject("todo")
	status3, _ := vo.NewTaskStatusValueObject("done")

	assert.True(t, status1.Equals(status2), "Expected status1 and status2 to be equal")
	assert.False(t, status1.Equals(status3), "Expected status1 and status3 to be different")
}
