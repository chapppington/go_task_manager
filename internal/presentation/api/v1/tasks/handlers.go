package tasks

import (
	"crud/internal/application"
	tasks_usecases "crud/internal/application/tasks/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

// Handler обработчик для задач
type Handler struct {
	container *dig.Container
}

// NewHandler создает новый обработчик задач
func NewHandler(container *dig.Container) *Handler {
	return &Handler{
		container: container,
	}
}

// CreateTask создает новую задачу
// POST /api/v1/tasks
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	useCase, err := application.ResolveFromContainer[*tasks_usecases.CreateTaskUseCase](h.container)
	if err != nil {
		http.Error(w, "Failed to resolve use case", http.StatusInternalServerError)
		return
	}

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	task, err := useCase.Execute(r.Context(), userID, req.Title, req.Description, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TaskDTOFromEntity(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetTaskByID получает задачу по ID
// GET /api/v1/tasks/{id}
func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	useCase, err := application.ResolveFromContainer[*tasks_usecases.GetTaskByIDUseCase](h.container)
	if err != nil {
		http.Error(w, "Failed to resolve use case", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := useCase.Execute(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := TaskDTOFromEntity(task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ListTasks получает список задач
// GET /api/v1/tasks
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	useCase, err := application.ResolveFromContainer[*tasks_usecases.ListTasksUseCase](h.container)
	if err != nil {
		http.Error(w, "Failed to resolve use case", http.StatusInternalServerError)
		return
	}

	page := 1
	pageSize := 10
	var userID *uuid.UUID
	var status *string

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pageSize = ps
		}
	}

	if userIDStr := r.URL.Query().Get("user_id"); userIDStr != "" {
		if id, err := uuid.Parse(userIDStr); err == nil {
			userID = &id
		}
	}

	if statusStr := r.URL.Query().Get("status"); statusStr != "" {
		status = &statusStr
	}

	tasks, total, err := useCase.Execute(r.Context(), userID, status, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		response[i] = TaskDTOFromEntity(task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":      response,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// UpdateTask обновляет задачу
// PUT /api/v1/tasks/{id}
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	useCase, err := application.ResolveFromContainer[*tasks_usecases.UpdateTaskUseCase](h.container)
	if err != nil {
		http.Error(w, "Failed to resolve use case", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := useCase.Execute(r.Context(), id, req.Title, req.Description, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TaskDTOFromEntity(task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteTask удаляет задачу
// DELETE /api/v1/tasks/{id}
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	useCase, err := application.ResolveFromContainer[*tasks_usecases.DeleteTaskUseCase](h.container)
	if err != nil {
		http.Error(w, "Failed to resolve use case", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := useCase.Execute(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
