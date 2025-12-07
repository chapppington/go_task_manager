package value_objects

import (
	"testing"

	vo "crud/internal/domain/users/value_objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmailValueObject(t *testing.T) {
	// Тест валидного email
	email, err := vo.NewEmailValueObject("test@example.com")
	require.NoError(t, err)
	assert.Equal(t, "test@example.com", email.Value())

	// Тест пустого email
	_, err = vo.NewEmailValueObject("")
	assert.True(t, vo.IsInvalidEmail(err))

	// Тест невалидного email
	_, err = vo.NewEmailValueObject("invalid-email")
	assert.True(t, vo.IsInvalidEmail(err))

	// Тест сравнения email
	email1, _ := vo.NewEmailValueObject("test@example.com")
	email2, _ := vo.NewEmailValueObject("test@example.com")
	email3, _ := vo.NewEmailValueObject("other@example.com")

	assert.True(t, email1.Equals(email2), "Expected email1 and email2 to be equal")
	assert.False(t, email1.Equals(email3), "Expected email1 and email3 to be different")
}
