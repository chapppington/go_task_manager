package value_objects

import (
	"testing"

	vo "crud/internal/domain/users/value_objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserNameValueObject(t *testing.T) {
	// Тест валидного имени
	name, err := vo.NewUserNameValueObject("Test User")
	require.NoError(t, err)
	assert.Equal(t, "Test User", name.Value())

	// Тест пустого имени
	_, err = vo.NewUserNameValueObject("")
	assert.True(t, vo.IsInvalidName(err))

	// Тест слишком короткого имени
	_, err = vo.NewUserNameValueObject("A")
	assert.True(t, vo.IsInvalidName(err))

	// Тест слишком длинного имени
	longName := make([]byte, 101)
	for i := range longName {
		longName[i] = 'A'
	}
	_, err = vo.NewUserNameValueObject(string(longName))
	assert.True(t, vo.IsInvalidName(err))

	// Тест сравнения имен
	name1, _ := vo.NewUserNameValueObject("Test User")
	name2, _ := vo.NewUserNameValueObject("Test User")
	name3, _ := vo.NewUserNameValueObject("Other User")

	assert.True(t, name1.Equals(name2), "Expected name1 and name2 to be equal")
	assert.False(t, name1.Equals(name3), "Expected name1 and name3 to be different")
}
