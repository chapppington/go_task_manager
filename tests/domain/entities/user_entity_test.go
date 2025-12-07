package entities

import (
	"testing"

	"crud/internal/domain/users"
	vo "crud/internal/domain/users/value_objects"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserEntity_Creation(t *testing.T) {
	email, err := vo.NewEmailValueObject("test@example.com")
	require.NoError(t, err)

	name, err := vo.NewUserNameValueObject("Test User")
	require.NoError(t, err)

	user := users.NewUser(email, name)

	assert.Equal(t, "test@example.com", user.Email.Value())
	assert.Equal(t, "Test User", user.Name.Value())
	assert.NotEqual(t, uuid.Nil, user.ID)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestUserEntity_Equality(t *testing.T) {
	email1, _ := vo.NewEmailValueObject("test@example.com")
	email2, _ := vo.NewEmailValueObject("other@example.com")

	name, _ := vo.NewUserNameValueObject("Test User")

	user1 := users.NewUser(email1, name)
	user2 := users.NewUser(email2, name)

	assert.False(t, user1.Equals(user2), "Expected user1 and user2 to be different")
}
