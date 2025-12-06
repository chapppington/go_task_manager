package users

import (
	"crud/internal/application"
	users_usecases "crud/internal/application/users/usecases"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для пользователей
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Получаем use cases из контейнера
	createUserUseCase, err := application.ResolveFromContainer[*users_usecases.CreateUserUseCase](container)
	if err != nil {
		return err
	}

	getUserByIDUseCase, err := application.ResolveFromContainer[*users_usecases.GetUserByIDUseCase](container)
	if err != nil {
		return err
	}

	getUserByEmailUseCase, err := application.ResolveFromContainer[*users_usecases.GetUserByEmailUseCase](container)
	if err != nil {
		return err
	}

	listUsersUseCase, err := application.ResolveFromContainer[*users_usecases.ListUsersUseCase](container)
	if err != nil {
		return err
	}

	updateUserUseCase, err := application.ResolveFromContainer[*users_usecases.UpdateUserUseCase](container)
	if err != nil {
		return err
	}

	deleteUserUseCase, err := application.ResolveFromContainer[*users_usecases.DeleteUserUseCase](container)
	if err != nil {
		return err
	}

	// Создаем handler
	handler := NewHandler(
		createUserUseCase,
		getUserByIDUseCase,
		getUserByEmailUseCase,
		listUsersUseCase,
		updateUserUseCase,
		deleteUserUseCase,
	)

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
