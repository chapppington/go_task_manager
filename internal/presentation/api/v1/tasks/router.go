package tasks

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для задач
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Создаем handler с контейнером
	handler := NewHandler(container)

	// Настраиваем маршруты
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.CreateTask)
		r.Get("/", handler.ListTasks)
		r.Get("/{id}", handler.GetTaskByID)
		r.Put("/{id}", handler.UpdateTask)
		r.Delete("/{id}", handler.DeleteTask)
	})

	return nil
}
