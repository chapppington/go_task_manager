package tasks

import (
	tasks_usecases "crud/internal/application/tasks/usecases"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для задач
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Получаем use cases из контейнера
	var createTaskUseCase *tasks_usecases.CreateTaskUseCase
	if err := container.Invoke(func(uc *tasks_usecases.CreateTaskUseCase) {
		createTaskUseCase = uc
	}); err != nil {
		return err
	}

	var getTaskByIDUseCase *tasks_usecases.GetTaskByIDUseCase
	if err := container.Invoke(func(uc *tasks_usecases.GetTaskByIDUseCase) {
		getTaskByIDUseCase = uc
	}); err != nil {
		return err
	}

	var listTasksUseCase *tasks_usecases.ListTasksUseCase
	if err := container.Invoke(func(uc *tasks_usecases.ListTasksUseCase) {
		listTasksUseCase = uc
	}); err != nil {
		return err
	}

	var updateTaskUseCase *tasks_usecases.UpdateTaskUseCase
	if err := container.Invoke(func(uc *tasks_usecases.UpdateTaskUseCase) {
		updateTaskUseCase = uc
	}); err != nil {
		return err
	}

	var deleteTaskUseCase *tasks_usecases.DeleteTaskUseCase
	if err := container.Invoke(func(uc *tasks_usecases.DeleteTaskUseCase) {
		deleteTaskUseCase = uc
	}); err != nil {
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
