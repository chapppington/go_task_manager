package users

import (
	users_usecases "crud/internal/application/users/usecases"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты для пользователей
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Получаем use cases из контейнера
	var createUserUseCase *users_usecases.CreateUserUseCase
	if err := container.Invoke(func(uc *users_usecases.CreateUserUseCase) {
		createUserUseCase = uc
	}); err != nil {
		return err
	}

	var getUserByIDUseCase *users_usecases.GetUserByIDUseCase
	if err := container.Invoke(func(uc *users_usecases.GetUserByIDUseCase) {
		getUserByIDUseCase = uc
	}); err != nil {
		return err
	}

	var getUserByEmailUseCase *users_usecases.GetUserByEmailUseCase
	if err := container.Invoke(func(uc *users_usecases.GetUserByEmailUseCase) {
		getUserByEmailUseCase = uc
	}); err != nil {
		return err
	}

	var listUsersUseCase *users_usecases.ListUsersUseCase
	if err := container.Invoke(func(uc *users_usecases.ListUsersUseCase) {
		listUsersUseCase = uc
	}); err != nil {
		return err
	}

	var updateUserUseCase *users_usecases.UpdateUserUseCase
	if err := container.Invoke(func(uc *users_usecases.UpdateUserUseCase) {
		updateUserUseCase = uc
	}); err != nil {
		return err
	}

	var deleteUserUseCase *users_usecases.DeleteUserUseCase
	if err := container.Invoke(func(uc *users_usecases.DeleteUserUseCase) {
		deleteUserUseCase = uc
	}); err != nil {
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
