package users

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для пользователей
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Создаем handler с контейнером
	handler := NewHandler(container)

	// Настраиваем маршруты
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)
		r.Get("/", handler.ListUsers)
		r.Get("/{id}", handler.GetUserByID)
		r.Get("/email/{email}", handler.GetUserByEmail)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})

	return nil
}
