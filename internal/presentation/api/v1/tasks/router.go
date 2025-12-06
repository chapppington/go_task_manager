package tasks

import (
	"crud/internal/application"
	tasks_usecases "crud/internal/application/tasks/usecases"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для задач
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Получаем use cases из контейнера
	createTaskUseCase, err := application.ResolveFromContainer[*tasks_usecases.CreateTaskUseCase](container)
	if err != nil {
		return err
	}

	getTaskByIDUseCase, err := application.ResolveFromContainer[*tasks_usecases.GetTaskByIDUseCase](container)
	if err != nil {
		return err
	}

	listTasksUseCase, err := application.ResolveFromContainer[*tasks_usecases.ListTasksUseCase](container)
	if err != nil {
		return err
	}

	updateTaskUseCase, err := application.ResolveFromContainer[*tasks_usecases.UpdateTaskUseCase](container)
	if err != nil {
		return err
	}

	deleteTaskUseCase, err := application.ResolveFromContainer[*tasks_usecases.DeleteTaskUseCase](container)
	if err != nil {
		return err
	}

	// Создаем handler
	handler := NewHandler(
		createTaskUseCase,
		getTaskByIDUseCase,
		listTasksUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
	)

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
