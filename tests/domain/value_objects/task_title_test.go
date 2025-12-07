package value_objects

import (
	"testing"

	vo "crud/internal/domain/tasks/value_objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskTitleValueObject(t *testing.T) {
	// Тест валидного заголовка
	title, err := vo.NewTaskTitleValueObject("Learn gRPC")
	require.NoError(t, err)
	assert.Equal(t, "Learn gRPC", title.Value())

	// Тест пустого заголовка
	_, err = vo.NewTaskTitleValueObject("")
	assert.True(t, vo.IsInvalidTitle(err))

	// Тест слишком длинного заголовка
	longTitle := make([]byte, 201)
	for i := range longTitle {
		longTitle[i] = 'A'
	}
	_, err = vo.NewTaskTitleValueObject(string(longTitle))
	assert.True(t, vo.IsInvalidTitle(err))

	// Тест сравнения заголовков
	title1, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title2, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title3, _ := vo.NewTaskTitleValueObject("Master Go")

	assert.True(t, title1.Equals(title2), "Expected title1 and title2 to be equal")
	assert.False(t, title1.Equals(title3), "Expected title1 and title3 to be different")
}
