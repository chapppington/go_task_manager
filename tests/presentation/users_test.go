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
