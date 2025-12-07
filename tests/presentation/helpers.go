package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1_tasks "crud/internal/presentation/api/v1/tasks"
	v1_users "crud/internal/presentation/api/v1/users"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// CreateUserViaHTTP создает пользователя через HTTP запрос и возвращает ответ
func CreateUserViaHTTP(t *testing.T, router chi.Router, email, name string) *v1_users.UserResponse {
	reqBody := v1_users.CreateUserRequest{
		Email: email,
		Name:  name,
	}
	jsonBody, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusCreated, responseRecorder.Code)

	var response v1_users.UserResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)
	require.NotEmpty(t, response.ID)

	return &response
}

// CreateTaskViaHTTP создает задачу через HTTP запрос и возвращает ответ
func CreateTaskViaHTTP(t *testing.T, router chi.Router, userID, title, description, status string) *v1_tasks.TaskResponse {
	reqBody := v1_tasks.CreateTaskRequest{
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
	}
	jsonBody, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusCreated, responseRecorder.Code)

	var response v1_tasks.TaskResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)
	require.NotEmpty(t, response.ID)

	return &response
}

// ExecuteRequest выполняет HTTP запрос и возвращает recorder
func ExecuteRequest(router chi.Router, method, path string, body interface{}) *httptest.ResponseRecorder {
	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	var req *http.Request
	if jsonBody != nil {
		req = httptest.NewRequest(method, path, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

// DecodeJSONResponse декодирует JSON ответ и возвращает декодированную структуру
func DecodeJSONResponse[T any](t *testing.T, response *httptest.ResponseRecorder) T {
	var result T
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err)
	return result
}

// DecodeJSONListResponse декодирует JSON ответ со списком (с полями data и total)
func DecodeJSONListResponse(t *testing.T, response *httptest.ResponseRecorder) ([]interface{}, int64) {
	var listResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&listResponse)
	require.NoError(t, err)

	data, ok := listResponse["data"].([]interface{})
	require.True(t, ok)

	total, ok := listResponse["total"].(float64)
	require.True(t, ok)

	return data, int64(total)
}
