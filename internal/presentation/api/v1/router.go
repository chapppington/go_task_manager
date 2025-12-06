package v1

import (
	users_usecases "crud/internal/application/users/usecases"
	"crud/internal/presentation/api/v1/handlers"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты API v1
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

	// Создаем handlers
	usersHandler := handlers.NewUsersHandler(createUserUseCase, getUserByIDUseCase)

	// Настраиваем маршруты
	r.Route("/users", func(r chi.Router) {
		r.Post("/", usersHandler.CreateUser)
		r.Get("/{id}", usersHandler.GetUserByID)
	})

	return nil
}
