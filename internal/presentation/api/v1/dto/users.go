package dto

// CreateUserRequest запрос на создание пользователя
type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UserResponse ответ с данными пользователя
type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
