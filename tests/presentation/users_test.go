package presentation

import (
	"net/http"
	"testing"

	v1_users "crud/internal/presentation/api/v1/users"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	router := NewTestRouterWithContainer()

	response := CreateUserViaHTTP(t, router, "test@example.com", "Test User")

	assert.Equal(t, "test@example.com", response.Email)
	assert.Equal(t, "Test User", response.Name)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.CreatedAt)
	assert.NotEmpty(t, response.UpdatedAt)
}

func TestGetUserByID(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	createResponse := CreateUserViaHTTP(t, router, "get@example.com", "Get User")

	// Получаем пользователя по ID
	getRR := ExecuteRequest(router, http.MethodGet, "/api/v1/users/"+createResponse.ID, nil)
	assert.Equal(t, http.StatusOK, getRR.Code)

	var getResponse v1_users.UserResponse
	DecodeJSONResponse(t, getRR, &getResponse)

	assert.Equal(t, createResponse.ID, getResponse.ID)
	assert.Equal(t, createResponse.Email, getResponse.Email)
	assert.Equal(t, createResponse.Name, getResponse.Name)
}

func TestGetUserByEmail(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	createResponse := CreateUserViaHTTP(t, router, "email@example.com", "Email User")

	// Получаем пользователя по email
	getRR := ExecuteRequest(router, http.MethodGet, "/api/v1/users/email/"+createResponse.Email, nil)
	assert.Equal(t, http.StatusOK, getRR.Code)

	var getResponse v1_users.UserResponse
	DecodeJSONResponse(t, getRR, &getResponse)

	assert.Equal(t, createResponse.ID, getResponse.ID)
	assert.Equal(t, createResponse.Email, getResponse.Email)
	assert.Equal(t, createResponse.Name, getResponse.Name)
}

func TestListUsers(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем несколько пользователей
	emails := []string{"list1@example.com", "list2@example.com", "list3@example.com"}
	var createdUsers []v1_users.UserResponse

	for i, email := range emails {
		userResponse := CreateUserViaHTTP(t, router, email, "List User "+string(rune('1'+i)))
		createdUsers = append(createdUsers, *userResponse)
	}

	// Получаем список пользователей
	listRR := ExecuteRequest(router, http.MethodGet, "/api/v1/users", nil)
	assert.Equal(t, http.StatusOK, listRR.Code)

	data, total := DecodeJSONListResponse(t, listRR)
	assert.GreaterOrEqual(t, len(data), len(createdUsers))
	assert.GreaterOrEqual(t, int(total), len(createdUsers))
}

func TestUpdateUser(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	createResponse := CreateUserViaHTTP(t, router, "update@example.com", "Update User")

	// Обновляем пользователя
	updatedName := "Updated Name"
	updateReqBody := v1_users.UpdateUserRequest{
		Name: &updatedName,
	}

	updateRR := ExecuteRequest(router, http.MethodPut, "/api/v1/users/"+createResponse.ID, updateReqBody)
	assert.Equal(t, http.StatusOK, updateRR.Code)

	var updateResponse v1_users.UserResponse
	DecodeJSONResponse(t, updateRR, &updateResponse)

	assert.Equal(t, createResponse.ID, updateResponse.ID)
	assert.Equal(t, createResponse.Email, updateResponse.Email)
	assert.Equal(t, "Updated Name", updateResponse.Name)
}

func TestDeleteUser(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	createResponse := CreateUserViaHTTP(t, router, "delete@example.com", "Delete User")

	// Удаляем пользователя
	deleteRR := ExecuteRequest(router, http.MethodDelete, "/api/v1/users/"+createResponse.ID, nil)
	assert.Equal(t, http.StatusNoContent, deleteRR.Code)

	// Проверяем, что пользователь действительно удален
	getRR := ExecuteRequest(router, http.MethodGet, "/api/v1/users/"+createResponse.ID, nil)
	assert.Equal(t, http.StatusNotFound, getRR.Code)
}
