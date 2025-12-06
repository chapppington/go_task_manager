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

func TestListTasks(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "listtasksuser@example.com",
		Name:  "List Tasks User",
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

	// Создаем несколько задач
	tasks := []v1_tasks.CreateTaskRequest{
		{UserID: userResponse.ID, Title: "Task 1", Description: "Description 1", Status: "todo"},
		{UserID: userResponse.ID, Title: "Task 2", Description: "Description 2", Status: "in_progress"},
		{UserID: userResponse.ID, Title: "Task 3", Description: "Description 3", Status: "done"},
	}

	for _, taskReq := range tasks {
		taskJsonBody, err := json.Marshal(taskReq)
		require.NoError(t, err)

		createTaskHTTPReq := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(taskJsonBody))
		createTaskHTTPReq.Header.Set("Content-Type", "application/json")
		createTaskRR := httptest.NewRecorder()
		router.ServeHTTP(createTaskRR, createTaskHTTPReq)

		assert.Equal(t, http.StatusCreated, createTaskRR.Code)
	}

	// Получаем список задач
	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
	listRR := httptest.NewRecorder()

	router.ServeHTTP(listRR, listReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, listRR.Code)

	// Проверяем ответ
	var listResponse map[string]interface{}
	err = json.NewDecoder(listRR.Body).Decode(&listResponse)
	require.NoError(t, err)

	data, ok := listResponse["data"].([]interface{})
	require.True(t, ok)
	assert.GreaterOrEqual(t, len(data), len(tasks))

	total, ok := listResponse["total"].(float64)
	require.True(t, ok)
	assert.GreaterOrEqual(t, int(total), len(tasks))
}

func TestListTasksWithFilters(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "filtertasksuser@example.com",
		Name:  "Filter Tasks User",
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

	// Создаем задачи с разными статусами
	tasks := []v1_tasks.CreateTaskRequest{
		{UserID: userResponse.ID, Title: "Todo Task", Description: "Description", Status: "todo"},
		{UserID: userResponse.ID, Title: "Done Task", Description: "Description", Status: "done"},
	}

	for _, taskReq := range tasks {
		taskJsonBody, err := json.Marshal(taskReq)
		require.NoError(t, err)

		createTaskHTTPReq := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(taskJsonBody))
		createTaskHTTPReq.Header.Set("Content-Type", "application/json")
		createTaskRR := httptest.NewRecorder()
		router.ServeHTTP(createTaskRR, createTaskHTTPReq)

		assert.Equal(t, http.StatusCreated, createTaskRR.Code)
	}

	// Получаем список задач с фильтром по статусу
	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/tasks?status=todo&user_id="+userResponse.ID, nil)
	listRR := httptest.NewRecorder()

	router.ServeHTTP(listRR, listReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, listRR.Code)

	var listResponse map[string]interface{}
	err = json.NewDecoder(listRR.Body).Decode(&listResponse)
	require.NoError(t, err)
}

func TestUpdateTask(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "updatetaskuser@example.com",
		Name:  "Update Task User",
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
		Title:       "Update Task",
		Description: "Original Description",
		Status:      "todo",
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

	// Обновляем задачу
	updatedTitle := "Updated Task Title"
	updatedDescription := "Updated Description"
	updatedStatus := "in_progress"
	updateReqBody := v1_tasks.UpdateTaskRequest{
		Title:       &updatedTitle,
		Description: &updatedDescription,
		Status:      &updatedStatus,
	}
	updateJsonBody, err := json.Marshal(updateReqBody)
	require.NoError(t, err)

	updateReq := httptest.NewRequest(http.MethodPut, "/api/v1/tasks/"+createTaskResponse.ID, bytes.NewBuffer(updateJsonBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateRR := httptest.NewRecorder()

	router.ServeHTTP(updateRR, updateReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, updateRR.Code)

	// Проверяем ответ
	var updateResponse v1_tasks.TaskResponse
	err = json.NewDecoder(updateRR.Body).Decode(&updateResponse)
	require.NoError(t, err)

	assert.Equal(t, createTaskResponse.ID, updateResponse.ID)
	assert.Equal(t, "Updated Task Title", updateResponse.Title)
	assert.Equal(t, "Updated Description", updateResponse.Description)
	assert.Equal(t, "in_progress", updateResponse.Status)
}

func TestDeleteTask(t *testing.T) {
	// Создаем тестовый роутер
	router := NewTestRouterWithContainer()

	// Сначала создаем пользователя
	createUserReq := v1_users.CreateUserRequest{
		Email: "deletetaskuser@example.com",
		Name:  "Delete Task User",
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
		Title:       "Delete Task",
		Description: "Delete Description",
		Status:      "todo",
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

	// Удаляем задачу
	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	deleteRR := httptest.NewRecorder()

	router.ServeHTTP(deleteRR, deleteReq)

	// Проверяем статус код
	assert.Equal(t, http.StatusNoContent, deleteRR.Code)

	// Проверяем, что задача действительно удалена
	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	getRR := httptest.NewRecorder()

	router.ServeHTTP(getRR, getReq)

	// Должен вернуть 404
	assert.Equal(t, http.StatusNotFound, getRR.Code)
}
