package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1_tasks "crud/internal/presentation/api/v1/tasks"
	v1_users "crud/internal/presentation/api/v1/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "taskuser@example.com",
		Name:  "Task User",
	}
	userJsonBody, err := json.Marshal(createUserReq)
	require.NoError(t, err)

	createUserHTTPReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(userJsonBody))
	createUserHTTPReq.Header.Set("Content-Type", "application/json")
	createUserRR := httptest.NewRecorder()
	router.ServeHTTP(createUserRR, createUserHTTPReq)

	assert.Equal(t, http.StatusCreated, createUserRR.Code)

	var userResponse v1_users.UserResponse
	err = json.NewDecoder(createUserRR.Body).Decode(&userResponse)
	require.NoError(t, err)
	require.NotEmpty(t, userResponse.ID)

	// Теперь создаем задачу
	reqBody := v1_tasks.CreateTaskRequest{
		UserID:      userResponse.ID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "todo",
	}
	jsonBody, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Проверяем ответ
	var response v1_tasks.TaskResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, reqBody.Title, response.Title)
	assert.Equal(t, reqBody.Description, response.Description)
	assert.Equal(t, reqBody.Status, response.Status)
	assert.Equal(t, userResponse.ID, response.UserID)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.CreatedAt)
	assert.NotEmpty(t, response.UpdatedAt)
}

func TestGetTaskByID(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "gettaskuser@example.com",
		Name:  "Get Task User",
	}
	userJsonBody, err := json.Marshal(createUserReq)
	require.NoError(t, err)

	createUserHTTPReq := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(userJsonBody))
	createUserHTTPReq.Header.Set("Content-Type", "application/json")
	createUserRR := httptest.NewRecorder()
	router.ServeHTTP(createUserRR, createUserHTTPReq)

	assert.Equal(t, http.StatusCreated, createUserRR.Code)

	var userResponse v1_users.UserResponse
	err = json.NewDecoder(createUserRR.Body).Decode(&userResponse)
	require.NoError(t, err)
	require.NotEmpty(t, userResponse.ID)

	// Создаем задачу
	createTaskReq := v1_tasks.CreateTaskRequest{
		UserID:      userResponse.ID,
		Title:       "Get Task",
		Description: "Get Task Description",
		Status:      "in_progress",
	}
	taskJsonBody, err := json.Marshal(createTaskReq)
	require.NoError(t, err)

	createTaskHTTPReq := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(taskJsonBody))
	createTaskHTTPReq.Header.Set("Content-Type", "application/json")
	createTaskRR := httptest.NewRecorder()
	router.ServeHTTP(createTaskRR, createTaskHTTPReq)

	assert.Equal(t, http.StatusCreated, createTaskRR.Code)

	var createTaskResponse v1_tasks.TaskResponse
	err = json.NewDecoder(createTaskRR.Body).Decode(&createTaskResponse)
	require.NoError(t, err)
	require.NotEmpty(t, createTaskResponse.ID)

	// Теперь получаем задачу по ID
	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	getRR := httptest.NewRecorder()

	router.ServeHTTP(getRR, getReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, getRR.Code)

	// Проверяем ответ
	var getResponse v1_tasks.TaskResponse
	err = json.NewDecoder(getRR.Body).Decode(&getResponse)
	require.NoError(t, err)

	assert.Equal(t, createTaskResponse.ID, getResponse.ID)
	assert.Equal(t, createTaskResponse.Title, getResponse.Title)
	assert.Equal(t, createTaskResponse.Description, getResponse.Description)
	assert.Equal(t, createTaskResponse.Status, getResponse.Status)
}
