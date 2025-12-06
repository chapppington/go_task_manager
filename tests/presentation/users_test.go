package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1_users "crud/internal/presentation/api/v1/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Подготавливаем запрос
	reqBody := v1_users.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	jsonBody, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Проверяем ответ
	var response v1_users.UserResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, reqBody.Email, response.Email)
	assert.Equal(t, reqBody.Name, response.Name)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.CreatedAt)
	assert.NotEmpty(t, response.UpdatedAt)
}

func TestGetUserByID(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createReqBody := v1_users.CreateUserRequest{
		Email: "get@example.com",
		Name:  "Get User",
	}
	jsonBody, err := json.Marshal(createReqBody)
	require.NoError(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	router.ServeHTTP(createRR, createReq)

	assert.Equal(t, http.StatusCreated, createRR.Code)

	var createResponse v1_users.UserResponse
	err = json.NewDecoder(createRR.Body).Decode(&createResponse)
	require.NoError(t, err)
	require.NotEmpty(t, createResponse.ID)

	// Теперь получаем пользователя по ID
	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+createResponse.ID, nil)
	getRR := httptest.NewRecorder()

	router.ServeHTTP(getRR, getReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, getRR.Code)

	// Проверяем ответ
	var getResponse v1_users.UserResponse
	err = json.NewDecoder(getRR.Body).Decode(&getResponse)
	require.NoError(t, err)

	assert.Equal(t, createResponse.ID, getResponse.ID)
	assert.Equal(t, createResponse.Email, getResponse.Email)
	assert.Equal(t, createResponse.Name, getResponse.Name)
}

func TestGetUserByEmail(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createReqBody := v1_users.CreateUserRequest{
		Email: "email@example.com",
		Name:  "Email User",
	}
	jsonBody, err := json.Marshal(createReqBody)
	require.NoError(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	router.ServeHTTP(createRR, createReq)

	assert.Equal(t, http.StatusCreated, createRR.Code)

	var createResponse v1_users.UserResponse
	err = json.NewDecoder(createRR.Body).Decode(&createResponse)
	require.NoError(t, err)
	require.NotEmpty(t, createResponse.Email)

	// Теперь получаем пользователя по email
	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/users/email/"+createResponse.Email, nil)
	getRR := httptest.NewRecorder()

	router.ServeHTTP(getRR, getReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, getRR.Code)

	// Проверяем ответ
	var getResponse v1_users.UserResponse
	err = json.NewDecoder(getRR.Body).Decode(&getResponse)
	require.NoError(t, err)

	assert.Equal(t, createResponse.ID, getResponse.ID)
	assert.Equal(t, createResponse.Email, getResponse.Email)
	assert.Equal(t, createResponse.Name, getResponse.Name)
}

func TestListUsers(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Создаем несколько пользователей
	emails := []string{"list1@example.com", "list2@example.com", "list3@example.com"}
	var createdUsers []v1_users.UserResponse

	for i, email := range emails {
		createReqBody := v1_users.CreateUserRequest{
			Email: email,
			Name:  "List User " + string(rune('1'+i)),
		}
		jsonBody, err := json.Marshal(createReqBody)
		require.NoError(t, err)

		createReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
		createReq.Header.Set("Content-Type", "application/json")
		createRR := httptest.NewRecorder()
		router.ServeHTTP(createRR, createReq)

		assert.Equal(t, http.StatusCreated, createRR.Code)

		var userResponse v1_users.UserResponse
		err = json.NewDecoder(createRR.Body).Decode(&userResponse)
		require.NoError(t, err)
		createdUsers = append(createdUsers, userResponse)
	}

	// Получаем список пользователей
	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	listRR := httptest.NewRecorder()

	router.ServeHTTP(listRR, listReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, listRR.Code)

	// Проверяем ответ
	var listResponse map[string]interface{}
	err := json.NewDecoder(listRR.Body).Decode(&listResponse)
	require.NoError(t, err)

	data, ok := listResponse["data"].([]interface{})
	require.True(t, ok)
	assert.GreaterOrEqual(t, len(data), len(createdUsers))

	total, ok := listResponse["total"].(float64)
	require.True(t, ok)
	assert.GreaterOrEqual(t, int(total), len(createdUsers))
}

func TestUpdateUser(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createReqBody := v1_users.CreateUserRequest{
		Email: "update@example.com",
		Name:  "Update User",
	}
	jsonBody, err := json.Marshal(createReqBody)
	require.NoError(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	router.ServeHTTP(createRR, createReq)

	assert.Equal(t, http.StatusCreated, createRR.Code)

	var createResponse v1_users.UserResponse
	err = json.NewDecoder(createRR.Body).Decode(&createResponse)
	require.NoError(t, err)
	require.NotEmpty(t, createResponse.ID)

	// Обновляем пользователя
	updatedName := "Updated Name"
	updateReqBody := v1_users.UpdateUserRequest{
		Name: &updatedName,
	}
	updateJsonBody, err := json.Marshal(updateReqBody)
	require.NoError(t, err)

	updateReq := httptest.NewRequest(http.MethodPut, "/api/v1/users/"+createResponse.ID, bytes.NewBuffer(updateJsonBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateRR := httptest.NewRecorder()

	router.ServeHTTP(updateRR, updateReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, updateRR.Code)

	// Проверяем ответ
	var updateResponse v1_users.UserResponse
	err = json.NewDecoder(updateRR.Body).Decode(&updateResponse)
	require.NoError(t, err)

	assert.Equal(t, createResponse.ID, updateResponse.ID)
	assert.Equal(t, createResponse.Email, updateResponse.Email)
	assert.Equal(t, "Updated Name", updateResponse.Name)
}

func TestDeleteUser(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createReqBody := v1_users.CreateUserRequest{
		Email: "delete@example.com",
		Name:  "Delete User",
	}
	jsonBody, err := json.Marshal(createReqBody)
	require.NoError(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	router.ServeHTTP(createRR, createReq)

	assert.Equal(t, http.StatusCreated, createRR.Code)

	var createResponse v1_users.UserResponse
	err = json.NewDecoder(createRR.Body).Decode(&createResponse)
	require.NoError(t, err)
	require.NotEmpty(t, createResponse.ID)

	// Удаляем пользователя
	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+createResponse.ID, nil)
	deleteRR := httptest.NewRecorder()

	router.ServeHTTP(deleteRR, deleteReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusNoContent, deleteRR.Code)

	// Проверяем, что пользователь действительно удален
	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+createResponse.ID, nil)
	getRR := httptest.NewRecorder()

	router.ServeHTTP(getRR, getReq)

	// Должен вернуть 404
	assert.Equal(t, http.StatusNotFound, getRR.Code)
}
