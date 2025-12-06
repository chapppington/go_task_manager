package tasks

import (
	"time"

	tasks_domain "crud/internal/domain/tasks"
)

// CreateTaskRequest запрос на создание задачи
type CreateTaskRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// UpdateTaskRequest запрос на обновление задачи
type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

// TaskResponse ответ с данными задачи
type TaskResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TaskDTOFromEntity создает TaskResponse из сущности задачи
func TaskDTOFromEntity(task *tasks_domain.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID.String(),
		UserID:      task.UserID.String(),
		Title:       task.Title.Value(),
		Description: task.Description,
		Status:      task.Status.Value(),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}
}
