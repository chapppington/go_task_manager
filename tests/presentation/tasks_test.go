package presentation

import (
	"net/http"
	"testing"

	v1_tasks "crud/internal/presentation/api/v1/tasks"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	userResponse := CreateUserViaHTTP(t, router, "taskuser@example.com", "Task User")

	// Создаем задачу
	response := CreateTaskViaHTTP(t, router, userResponse.ID, "Test Task", "Test Description", "todo")

	assert.Equal(t, "Test Task", response.Title)
	assert.Equal(t, "Test Description", response.Description)
	assert.Equal(t, "todo", response.Status)
	assert.Equal(t, userResponse.ID, response.UserID)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.CreatedAt)
	assert.NotEmpty(t, response.UpdatedAt)
}

func TestGetTaskByID(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя и задачу
	userResponse := CreateUserViaHTTP(t, router, "gettaskuser@example.com", "Get Task User")
	createTaskResponse := CreateTaskViaHTTP(t, router, userResponse.ID, "Get Task", "Get Task Description", "in_progress")

	// Получаем задачу по ID
	response := ExecuteRequest(router, http.MethodGet, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	assert.Equal(t, http.StatusOK, response.Code)

	var getResponse v1_tasks.TaskResponse
	DecodeJSONResponse(t, response, &getResponse)

	assert.Equal(t, createTaskResponse.ID, getResponse.ID)
	assert.Equal(t, createTaskResponse.Title, getResponse.Title)
	assert.Equal(t, createTaskResponse.Description, getResponse.Description)
	assert.Equal(t, createTaskResponse.Status, getResponse.Status)
}

func TestListTasks(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	userResponse := CreateUserViaHTTP(t, router, "listtasksuser@example.com", "List Tasks User")

	// Создаем несколько задач
	tasks := []struct {
		title, description, status string
	}{
		{"Task 1", "Description 1", "todo"},
		{"Task 2", "Description 2", "in_progress"},
		{"Task 3", "Description 3", "done"},
	}

	for _, task := range tasks {
		CreateTaskViaHTTP(t, router, userResponse.ID, task.title, task.description, task.status)
	}

	// Получаем список задач
	response := ExecuteRequest(router, http.MethodGet, "/api/v1/tasks", nil)
	assert.Equal(t, http.StatusOK, response.Code)

	data, total := DecodeJSONListResponse(t, response)
	assert.GreaterOrEqual(t, len(data), len(tasks))
	assert.GreaterOrEqual(t, int(total), len(tasks))
}

func TestListTasksWithFilters(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя
	userResponse := CreateUserViaHTTP(t, router, "filtertasksuser@example.com", "Filter Tasks User")

	// Создаем задачи с разными статусами
	CreateTaskViaHTTP(t, router, userResponse.ID, "Todo Task", "Description", "todo")
	CreateTaskViaHTTP(t, router, userResponse.ID, "Done Task", "Description", "done")

	// Получаем список задач с фильтром по статусу
	response := ExecuteRequest(router, http.MethodGet, "/api/v1/tasks?status=todo&user_id="+userResponse.ID, nil)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateTask(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя и задачу
	userResponse := CreateUserViaHTTP(t, router, "updatetaskuser@example.com", "Update Task User")
	createTaskResponse := CreateTaskViaHTTP(t, router, userResponse.ID, "Update Task", "Original Description", "todo")

	// Обновляем задачу
	updatedTitle := "Updated Task Title"
	updatedDescription := "Updated Description"
	updatedStatus := "in_progress"
	updateReqBody := v1_tasks.UpdateTaskRequest{
		Title:       &updatedTitle,
		Description: &updatedDescription,
		Status:      &updatedStatus,
	}

	response := ExecuteRequest(router, http.MethodPut, "/api/v1/tasks/"+createTaskResponse.ID, updateReqBody)
	assert.Equal(t, http.StatusOK, response.Code)

	var updateResponse v1_tasks.TaskResponse
	DecodeJSONResponse(t, response, &updateResponse)

	assert.Equal(t, createTaskResponse.ID, updateResponse.ID)
	assert.Equal(t, "Updated Task Title", updateResponse.Title)
	assert.Equal(t, "Updated Description", updateResponse.Description)
	assert.Equal(t, "in_progress", updateResponse.Status)
}

func TestDeleteTask(t *testing.T) {
	router := NewTestRouterWithContainer()

	// Создаем пользователя и задачу
	userResponse := CreateUserViaHTTP(t, router, "deletetaskuser@example.com", "Delete Task User")
	createTaskResponse := CreateTaskViaHTTP(t, router, userResponse.ID, "Delete Task", "Delete Description", "todo")

	// Удаляем задачу
	response := ExecuteRequest(router, http.MethodDelete, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	assert.Equal(t, http.StatusNoContent, response.Code)

	// Проверяем, что задача действительно удалена
	response = ExecuteRequest(router, http.MethodGet, "/api/v1/tasks/"+createTaskResponse.ID, nil)
	assert.Equal(t, http.StatusNotFound, response.Code)
}
