package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	users_usecases "crud/internal/application/users/usecases"
	"crud/internal/presentation/api/v1/dto"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// UsersHandler обработчик для пользователей
type UsersHandler struct {
	createUseCase  *users_usecases.CreateUserUseCase
	getByIDUseCase *users_usecases.GetUserByIDUseCase
}

// NewUsersHandler создает новый обработчик пользователей
func NewUsersHandler(
	createUseCase *users_usecases.CreateUserUseCase,
	getByIDUseCase *users_usecases.GetUserByIDUseCase,
) *UsersHandler {
	return &UsersHandler{
		createUseCase:  createUseCase,
		getByIDUseCase: getByIDUseCase,
	}
}

// CreateUser создает нового пользователя
// POST /api/v1/users
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.createUseCase.Execute(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := dto.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email.Value(),
		Name:      user.Name.Value(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetUserByID получает пользователя по ID
// GET /api/v1/users/{id}
func (h *UsersHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	response := dto.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email.Value(),
		Name:      user.Name.Value(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
