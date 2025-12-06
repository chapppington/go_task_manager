package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	users_usecases "crud/internal/application/users/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Handler обработчик для пользователей
type Handler struct {
	createUseCase     *users_usecases.CreateUserUseCase
	getByIDUseCase    *users_usecases.GetUserByIDUseCase
	getByEmailUseCase *users_usecases.GetUserByEmailUseCase
	listUseCase       *users_usecases.ListUsersUseCase
	updateUseCase     *users_usecases.UpdateUserUseCase
	deleteUseCase     *users_usecases.DeleteUserUseCase
}

// NewHandler создает новый обработчик пользователей
func NewHandler(
	createUseCase *users_usecases.CreateUserUseCase,
	getByIDUseCase *users_usecases.GetUserByIDUseCase,
	getByEmailUseCase *users_usecases.GetUserByEmailUseCase,
	listUseCase *users_usecases.ListUsersUseCase,
	updateUseCase *users_usecases.UpdateUserUseCase,
	deleteUseCase *users_usecases.DeleteUserUseCase,
) *Handler {
	return &Handler{
		createUseCase:     createUseCase,
		getByIDUseCase:    getByIDUseCase,
		getByEmailUseCase: getByEmailUseCase,
		listUseCase:       listUseCase,
		updateUseCase:     updateUseCase,
		deleteUseCase:     deleteUseCase,
	}
}

// CreateUser создает нового пользователя
// POST /api/v1/users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.createUseCase.Execute(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := UserDTOFromEntity(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetUserByID получает пользователя по ID
// GET /api/v1/users/{id}
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.getByIDUseCase.Execute(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := UserDTOFromEntity(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserByEmail получает пользователя по email
// GET /api/v1/users/email/{email}
func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		http.Error(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	user, err := h.getByEmailUseCase.Execute(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := UserDTOFromEntity(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ListUsers получает список пользователей
// GET /api/v1/users
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page := 1
	pageSize := 10

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

	users, total, err := h.listUseCase.Execute(r.Context(), page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]UserResponse, len(users))
	for i, user := range users {
		response[i] = UserDTOFromEntity(user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":      response,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// UpdateUser обновляет пользователя
// PUT /api/v1/users/{id}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.updateUseCase.Execute(r.Context(), id, req.Email, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := UserDTOFromEntity(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteUser удаляет пользователя
// DELETE /api/v1/users/{id}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.deleteUseCase.Execute(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
